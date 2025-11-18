package lz4

import (
	"io"
	"os"

	"github.com/pierrec/lz4"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x04, 0x22, 0x4D, 0x18,
	})
}

func Deflate(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r := lz4.NewReader(f)

	return io.ReadAll(r)
}
