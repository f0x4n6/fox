package gzip

import (
	"compress/gzip"
	"io"
	"os"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x1F, 0x8B, 0x08,
	})
}

func Deflate(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r, err := gzip.NewReader(f)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(r.Close)

	return io.ReadAll(r)
}
