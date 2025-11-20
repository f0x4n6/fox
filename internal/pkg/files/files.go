package files

import "bytes"

type Entry struct {
	Path string // Entry path
	Data []byte // Entry data
}

type Convert func([]byte) ([]byte, error)

type Deflate func([]byte) ([]byte, error)

type Extract func([]byte, string, string) []Entry

func HasMagic(b []byte, o int, m []byte) bool {
	if len(b) < o+len(m) {
		return false
	}

	return bytes.Equal(b[o:o+len(m)], m)
}
