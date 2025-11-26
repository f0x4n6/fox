package journal

import (
	"bytes"
	"regexp"

	"github.com/Velocidex/go-journalctl/parser"
)

var Magic = regexp.MustCompile("LPKSHHRH")

func Parse(b []byte, off int64) (*parser.JournalFile, error) {
	return parser.OpenFile(bytes.NewReader(b[off:]))
}
