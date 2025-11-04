package gzip

import (
	"compress/gzip"
	"io"
	"log"

	"github.com/cuhsat/fox/v3/internal/pkg/files"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x1F, 0x8B, 0x08,
	})
}

func Deflate(path string) string {
	a := fs.Open(path)
	defer sys.Handler(a.Close)

	r, err := gzip.NewReader(a)

	if err != nil {
		log.Println(err)
		return path
	}

	defer sys.Handler(r.Close)

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	_, err = io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
