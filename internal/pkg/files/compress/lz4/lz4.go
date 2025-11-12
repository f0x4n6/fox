package lz4

import (
	"io"
	"log"

	"github.com/pierrec/lz4"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x04, 0x22, 0x4D, 0x18,
	})
}

func Deflate(path string) string {
	a := fs.Open(path)
	defer sys.Handler(a.Close)

	r := lz4.NewReader(a)

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	_, err := io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
