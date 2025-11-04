// Package fortinet is experimental and not yet documented.
package fortinet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/pierrec/lz4/v4"

	"github.com/cuhsat/fox/v3/internal/pkg/files"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
)

type llog struct {
	// raw
	Magic         uint16
	Flags         uint8
	Unused        uint16
	LDevId        uint8
	LDevName      uint8
	LVDom         uint8
	Entries       uint16
	LCompressed   uint16
	LDecompressed uint16
	Timestamp     uint32

	// parsed
	LEntries uint16
	LAscii   uint16
	Padding  uint16
	DevId    string
	DevName  string
	VDom     string
	Body     []byte
}

func Detect(path string) bool {
	for _, m := range [][]byte{
		{0xEC, 0xCE}, // llog v5
		{0xEC, 0xCF}, // llog v5
	} {
		if files.HasMagic(path, 0, m) {
			return true
		}
	}

	return false
}

func Parse(path string) string {
	f := fs.Open(path)
	defer sys.Handler(f.Close)

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	for {
		l, err := decode(f)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		str, err := deflate(l)

		if err != nil {
			log.Println(err)
			break
		}

		_, _ = t.WriteString(str)

		err = forward(f)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Println(err)
			break
		}
	}

	return t.Name()
}

func decode(f fs.File) (llog, error) {
	var err error

	l := llog{}

	err = binary.Read(f, binary.LittleEndian, &l.Magic)

	if l.Magic != 0xCEEC && l.Magic != 0xCFEC {
		return l, errors.New("log format not supported")
	}

	err = binary.Read(f, binary.LittleEndian, &l.Flags)
	err = binary.Read(f, binary.LittleEndian, &l.Unused)
	err = binary.Read(f, binary.LittleEndian, &l.LDevId)
	err = binary.Read(f, binary.LittleEndian, &l.LDevName)
	err = binary.Read(f, binary.LittleEndian, &l.LVDom)
	err = binary.Read(f, binary.BigEndian, &l.Entries)
	err = binary.Read(f, binary.BigEndian, &l.LCompressed)
	err = binary.Read(f, binary.BigEndian, &l.LDecompressed)
	err = binary.Read(f, binary.BigEndian, &l.Timestamp)

	if err != nil {
		return l, err
	}

	l.LEntries = l.Entries * 2
	l.LAscii = uint16(l.LDevId + l.LDevName + l.LVDom)

	if l.Flags&4 == 1 {
		l.Padding = l.LEntries
	}

	l.Body = make([]byte, l.LAscii+l.LEntries+l.Padding+l.LCompressed)

	_, _ = io.ReadFull(f, l.Body)

	i, j := 0, int(l.LDevId)

	l.DevId = string(l.Body[i:j])

	i, j = j, j+int(l.LDevName)

	l.DevName = string(l.Body[i:j])

	i, j = j, j+int(l.LVDom)

	l.VDom = string(l.Body[i:j])

	return l, nil
}

func deflate(log llog) (string, error) {
	var sb strings.Builder

	i := log.LAscii + log.LEntries + log.Padding
	j := i + log.LCompressed

	data := make([]byte, log.LDecompressed+1)

	n, err := lz4.UncompressBlock(log.Body[i:j], data)

	if err != nil {
		return "", err
	}

	if uint16(n) != log.LDecompressed {
		return "", errors.New("invalid block length")
	}

	if log.Entries == 1 {
		_, _ = sb.WriteString(format(log, data))
	} else {
		p, l, b := 0, 0, log.Body[log.LAscii:log.LAscii+log.LEntries]

		for i := 0; i < int(log.LEntries); i += 2 {
			l = int(binary.BigEndian.Uint16(b[i : i+2]))
			_, _ = sb.WriteString(format(log, data[p:p+l]))
			p += l
		}
	}

	return sb.String(), nil
}

func format(log llog, b []byte) string {
	return fmt.Sprintf("devid=\"%s\" devname=\"%s\" vdom=\"%s\" %s\n", log.DevId, log.DevName, log.VDom, string(b))
}

func forward(f fs.File) error {
	var n uint16

	err := binary.Read(f, binary.LittleEndian, &n)

	if err != nil {
		return err
	}

	_, err = f.Seek(int64(n), io.SeekCurrent)

	return err
}
