// Package fortinet is experimental and not yet documented.
package fortinet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/pierrec/lz4/v4"

	"github.com/cuhsat/fox/internal/pkg/files"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
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
	defer f.Close()

	t := fs.Create(path)
	defer t.Close()

	for {
		log, err := decode(f)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			sys.Error(err)
			break
		}

		str, err := deflate(log)

		if err != nil {
			sys.Error(err)
			break
		}

		_, _ = t.WriteString(str)

		err = forward(f)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			sys.Error(err)
			break
		}
	}

	return t.Name()
}

func decode(f fs.File) (llog, error) {
	var err error

	log := llog{}

	err = binary.Read(f, binary.LittleEndian, &log.Magic)

	if log.Magic != 0xCEEC && log.Magic != 0xCFEC {
		return log, errors.New("log format not supported")
	}

	err = binary.Read(f, binary.LittleEndian, &log.Flags)
	err = binary.Read(f, binary.LittleEndian, &log.Unused)
	err = binary.Read(f, binary.LittleEndian, &log.LDevId)
	err = binary.Read(f, binary.LittleEndian, &log.LDevName)
	err = binary.Read(f, binary.LittleEndian, &log.LVDom)
	err = binary.Read(f, binary.BigEndian, &log.Entries)
	err = binary.Read(f, binary.BigEndian, &log.LCompressed)
	err = binary.Read(f, binary.BigEndian, &log.LDecompressed)
	err = binary.Read(f, binary.BigEndian, &log.Timestamp)

	if err != nil {
		return log, err
	}

	log.LEntries = log.Entries * 2
	log.LAscii = uint16(log.LDevId + log.LDevName + log.LVDom)

	if log.Flags&4 == 1 {
		log.Padding = log.LEntries
	}

	log.Body = make([]byte, log.LAscii+log.LEntries+log.Padding+log.LCompressed)

	_, _ = io.ReadFull(f, log.Body)

	i, j := 0, int(log.LDevId)

	log.DevId = string(log.Body[i:j])

	i, j = j, j+int(log.LDevName)

	log.DevName = string(log.Body[i:j])

	i, j = j, j+int(log.LVDom)

	log.VDom = string(log.Body[i:j])

	return log, nil
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
