package bzip2

import (
	"compress/bzip2"
	"io"
	"log"

	"github.com/cuhsat/fox/internal/pkg/files"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x42, 0x5A, 0x68,
	})
}

func Deflate(path string) string {
	a := fs.Open(path)
	defer a.Close()

	r := bzip2.NewReader(a)

	t := fs.Create(path)
	defer t.Close()

	_, err := io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
