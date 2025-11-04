package zstd

import (
	"io"
	"log"

	"github.com/klauspost/compress/zstd"

	"github.com/cuhsat/fox/v3/internal/pkg/files"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
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

func Deflate(path string) string {
	a := fs.Open(path)
	defer a.Close()

	r, err := zstd.NewReader(a)

	if err != nil {
		log.Println(err)
		return path
	}

	defer r.Close()

	t := fs.Create(path)
	defer t.Close()

	_, err = io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
