package evtx

import (
	"log"

	"github.com/0xrawsec/golang-evtx/evtx"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

const (
	lf = 0xa
)

func Detect(path string) bool {
	return files.HasMagic(path, 0, []byte{
		0x45, 0x6C, 0x66, 0x46, 0x69, 0x6C, 0x65, 0x00,
	})
}

func Parse(path string) string {
	f := fs.Open(path)
	defer sys.Handler(f.Close)

	t := fs.Create(path)
	defer sys.Handler(t.Close)

	r, err := evtx.New(f)

	if err != nil {
		log.Println(err)
		return path
	}

	defer sys.Handler(r.Close)

	for e := range r.Events() {
		_, err := t.Write(evtx.ToJSON(e))

		if err != nil {
			log.Println(err)
		}

		_, err = t.Write([]byte{lf})

		if err != nil {
			log.Println(err)
		}
	}

	return t.Name()
}
