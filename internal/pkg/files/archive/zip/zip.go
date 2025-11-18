package zip

import (
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/cuhsat/zip/pkg/zip"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x50, 0x4B, 0x03, 0x04,
	})
}

func Deflate(path, pass string) (e []files.Entry) {
	r, err := zip.OpenReader(path)

	if err != nil {
		log.Println(err)
		return
	}

	defer sys.Handle(r.Close)

	for _, f := range r.File {
		if strings.HasSuffix(f.Name, "/") {
			continue
		}

		if len(pass) > 0 {
			f.SetPassword(pass)
		}

		a, err := f.Open()

		if err != nil {
			log.Println(err)
			continue
		}

		buf, err := io.ReadAll(a)

		_ = a.Close()

		if err != nil {
			log.Println(err)
			continue
		}

		e = append(e, files.Entry{
			Name: filepath.Join(path, f.Name),
			Data: buf,
		})
	}

	return
}
