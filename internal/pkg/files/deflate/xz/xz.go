package xz

import (
	"io"
	"os"

	"github.com/ulikunitz/xz"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00,
	})
}

func Deflate(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r, err := xz.NewReader(f)

	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}
