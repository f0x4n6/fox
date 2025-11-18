package zstd

import (
	"io"
	"os"

	"github.com/klauspost/compress/zstd"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	for _, m := range [][]byte{
		{0x1E, 0xB5, 0x2F, 0xFD}, // v0.1
		{0x22, 0xB5, 0x2F, 0xFD}, // v0.2
		{0x23, 0xB5, 0x2F, 0xFD}, // v0.3
		{0x24, 0xB5, 0x2F, 0xFD}, // v0.4
		{0x25, 0xB5, 0x2F, 0xFD}, // v0.5
		{0x26, 0xB5, 0x2F, 0xFD}, // v0.6
		{0x27, 0xB5, 0x2F, 0xFD}, // v0.7
		{0x28, 0xB5, 0x2F, 0xFD}, // v0.8
	} {
		if files.HasMagic(path, 0, m) {
			return true
		}
	}

	return false
}

func Deflate(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	r, err := zstd.NewReader(f)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	return io.ReadAll(r)
}
