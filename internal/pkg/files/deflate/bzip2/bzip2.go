package bzip2

import (
	"compress/bzip2"
	"io"
	"os"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x42, 0x5A, 0x68,
	})
}

func Deflate(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r := bzip2.NewReader(f)

	return io.ReadAll(r)
}
