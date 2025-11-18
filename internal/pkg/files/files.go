package files

import (
	"bytes"
)

type Entry struct {
	Name string
	Data []byte
}

type Detect func([]byte) bool

type Deflate func([]byte) ([]byte, error)

type Extract func([]byte, string, string) []Entry

func HasMagic(b []byte, o int, m []byte) bool {
	if len(b) < o+len(m) {
		return false
	}

	return bytes.Equal(b[o:o+len(m)], m)
}
