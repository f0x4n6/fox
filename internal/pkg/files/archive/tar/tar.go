package tar

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return files.HasMagic(path, 257, []byte{
		0x75, 0x73, 0x74, 0x61, 0x72,
	})
}

func Deflate(path, _ string) (e []files.Entry) {
	a, err := os.Open(path)

	if err != nil {
		return
	}

	defer sys.Handle(a.Close)

	r := tar.NewReader(a)

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
