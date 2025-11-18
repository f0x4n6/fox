package journal

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Velocidex/go-journalctl/parser"

	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

func Detect(path string) bool {
	return filepath.Ext(path) == ".journal"
}

func Convert(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer sys.Handle(f.Close)

	j, err := parser.OpenFile(f)

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)

	for l := range j.GetLogs(context.Background()) {
		_, err := buf.WriteString(fmt.Sprintf("%v\n", l))

		if err != nil {
			log.Println(err)
		}
	}

	return buf.Bytes(), err
}
