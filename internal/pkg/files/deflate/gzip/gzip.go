package gzip

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(b []byte) bool {
	return files.HasMagic(b, 0, []byte{
		0x1F, 0x8B, 0x08,
	})
}

func Deflate(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	defer sys.Ignore(r.Close)

	return io.ReadAll(r)
}
