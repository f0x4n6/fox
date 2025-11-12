package zlib

import (
	"io"
	"log"

	"github.com/klauspost/compress/zlib"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	for _, m := range [][]byte{
		{0x78, 0x01}, // no compression
		{0x78, 0x5E}, // fast compression
		{0x78, 0x9C}, // default compression
		{0x78, 0xDA}, // best compression
	} {
		if files.HasMagic(path, 0, m) {
			return true
		}
	}

	return false
}

func Deflate(path string) string {
	a := fs.Open(path)
	defer sys.Handler(a.Close)

	r, err := zlib.NewReader(a)

	if err != nil {
		log.Println(err)
		return path
	}

	defer sys.Handler(r.Close)

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	_, err = io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
