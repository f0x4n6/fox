package xz

import (
	"io"
	"log"

	"github.com/ulikunitz/xz"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00,
	})
}

func Deflate(path string) string {
	a := fs.Open(path)
	defer sys.Handler(a.Close)

	r, err := xz.NewReader(a)

	if err != nil {
		log.Println(err)
		return path
	}

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	_, err = io.Copy(t, r)

	if err != nil {
		log.Println(err)
		return path
	}

	return t.Name()
}
