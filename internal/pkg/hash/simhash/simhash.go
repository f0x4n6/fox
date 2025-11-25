package simhash

import (
	"encoding/hex"

	"github.com/erfanmomeniii/simhash"
)

type SimHash struct {
	s *simhash.Simhash
}

func New() *SimHash {
	return &SimHash{simhash.NewSimhash()}
}

func (sh *SimHash) Reset() {
	sh.s = simhash.NewSimhash()
}

func (sh *SimHash) BlockSize() int {
	return 1
}

func (sh *SimHash) Size() int {
	return 8
}

func (sh *SimHash) Sum(_ []byte) []byte {
	b, _ := hex.DecodeString(sh.s.GenerateToken())

	return b
}

func (sh *SimHash) Write(b []byte) (int, error) {
	return len(b), sh.s.AddFeature(b, 1)
}
