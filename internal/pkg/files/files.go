package files

import (
	"bytes"
	"io"
	"log"

	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

type Item struct {
	Path string
	Name string
}

type Deflate func(string, string) []*Item

func HasMagic(p string, o int, m []byte) bool {
	buf := make([]byte, o+len(m))

	f := fs.Open(p)
	defer sys.Handler(f.Close)

	fi, err := f.Stat()

	if err != nil {
		log.Println(err)
		return false
	}

	if fi.Size() < int64(o+len(m)) {
		return false
	}

	_, err = io.ReadFull(f, buf)

	if err != nil {
		log.Println(err)
		return false
	}

	return bytes.Equal(buf[o:], m)
}
