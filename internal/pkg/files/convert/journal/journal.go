package journal

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/Velocidex/go-journalctl/parser"
	"github.com/cuhsat/fox/v4/internal/pkg/files"
)

func Detect(b []byte) bool {
	return files.HasMagic(b, 0, []byte{
		'L', 'P', 'K', 'S', 'H', 'H', 'R', 'H',
	})
}

func Convert(b []byte) ([]byte, error) {
	j, err := parser.OpenFile(bytes.NewReader(b))

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
