package blake3

import (
	"github.com/zeebo/blake3"
)

const (
	size256 = iota
	size512
)

type Blake3 struct {
	blake3.Hasher

	size int
}

func New256() *Blake3 {
	return &Blake3{*blake3.New(), size256}
}

func New512() *Blake3 {
	return &Blake3{*blake3.New(), size256}
}

func (b *Blake3) Sum(b []byte) []byte {
	switch b.size {
	case size256:
		return b.Sum512
	}
}
