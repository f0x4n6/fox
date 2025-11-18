package evtx

import (
	"bytes"
	"log"
	"os"

	"github.com/0xrawsec/golang-evtx/evtx"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x45, 0x6C, 0x66, 0x46, 0x69, 0x6C, 0x65, 0x00,
	})
}

func Convert(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r, err := evtx.New(f)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(r.Close)

	buf := bytes.NewBuffer(nil)

	for e := range r.Events() {
		_, err := buf.Write(evtx.ToJSON(e))

		if err != nil {
			log.Println(err)
			continue
		}

		_, err = buf.WriteRune('\n')

		if err != nil {
			log.Println(err)
		}
	}

	return buf.Bytes(), nil
}
