package rar

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nwaples/rardecode"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x52, 0x61, 0x72, 0x21, 0x1A, 0x07,
	})
}

func Deflate(path, pass string) (e []files.Entry) {
	a, err := os.Open(path)

	if err != nil {
		return
	}

	defer sys.Handle(a.Close)

	r, err := rardecode.NewReader(a, pass)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		h, err := r.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
			break
		}

		if strings.HasSuffix(h.Name, "/") {
			continue
		}

		buf, err := io.ReadAll(r)

		if err != nil {
			log.Println(err)
			continue
		}

		e = append(e, files.Entry{
			Name: filepath.Join(path, h.Name),
			Data: buf,
		})
	}

	return
}
