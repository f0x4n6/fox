package br

import (
	"io"
	"os"

	"github.com/andybalholm/brotli"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0xCE, 0xB2, 0xCF, 0x81,
	})
}

func Deflate(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r := brotli.NewReader(f)

	return io.ReadAll(r)
}
