package files

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

type Entry struct {
	Name string
	Data []byte
}

type Deflate func(string, string) []Entry

func HasMagic(path string, off int, mask []byte) bool {
	buf := make([]byte, off+len(mask))

	f, err := os.Open(path)

	if err != nil {
		log.Println(err)
		return false
	}

	defer sys.Handle(f.Close)

	n, err := io.ReadFull(f, buf)

	if n < len(buf) {
		return false
	}

	if err != nil {
		log.Println(err)
		return false
	}

	return bytes.Equal(buf[off:], mask)
}
