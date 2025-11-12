package br

import (
	"io"
	"log"

	"github.com/andybalholm/brotli"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0xCE, 0xB2, 0xCF, 0x81,
	})
}

func Deflate(path string) string {
	a := fs.Open(path)
	defer sys.Handler(a.Close)

	r := brotli.NewReader(a)

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	_, err := io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
