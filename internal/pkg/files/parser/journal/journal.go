package journal

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/Velocidex/go-journalctl/parser"

	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
)

func Detect(path string) bool {
	return filepath.Ext(path) == ".journal"
}

func Parse(path string) string {
	f := fs.Open(path)
	defer f.Close()

	t := fs.Create(path)
	defer t.Close()

	j, err := parser.OpenFile(f)

	if err != nil {
		log.Println(err)
		return path
	}

	for l := range j.GetLogs(context.Background()) {
		_, err := t.WriteString(fmt.Sprintf("%v\n", l))

		if err != nil {
			log.Println(err)
		}
	}

	return t.Name()
}
