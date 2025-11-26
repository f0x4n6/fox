package journal

import (
	"bytes"
	"regexp"

	"github.com/Velocidex/go-journalctl/parser"
)

var Magic = regexp.MustCompile("LPKSHHRH")

func Parse(b []byte, off int64) (*parser.JournalFile, error) {
	j, err := parser.OpenFile(bytes.NewReader(b[off:]))

	if err != nil {
		return nil, err
	}

	return j, nil
}
