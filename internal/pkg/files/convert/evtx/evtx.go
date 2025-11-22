package evtx

import (
	"bytes"
	"log"

	"github.com/0xrawsec/golang-evtx/evtx"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
)

func Detect(b []byte) bool {
	return files.HasMagic(b, 0, []byte{
		'E', 'l', 'f', 'F', 'i', 'l', 'e', 0,
	})
}

func Convert(b []byte) ([]byte, error) {
	r, err := evtx.New(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

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

	_ = r.Close()

	return buf.Bytes(), nil
}
