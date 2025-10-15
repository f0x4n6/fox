package heap

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"errors"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/eciavatta/sdhash"
	"github.com/glaslos/ssdeep"
	"github.com/glaslos/tlsh"
	"github.com/zeebo/xxh3"
)

type Hash map[string][]byte

func (h *Heap) HashSum(algo string) ([]byte, error) {
	algo = strings.ToLower(algo)

	h.RLock()
	sum, ok := h.hash[algo]
	h.RUnlock()

	if ok {
		return sum, nil
	}

	var imp hash.Hash

	switch algo {
	case types.MD5:
		imp = md5.New()
	case types.SHA1:
		imp = sha1.New()
	case types.SHA256:
		imp = sha256.New()
	case types.SHA3, types.SHA3224:
		imp = sha3.New224()
	case types.SHA3256:
		imp = sha3.New256()
	case types.SHA3384:
		imp = sha3.New384()
	case types.SHA3512:
		imp = sha3.New512()
	case types.FNV1:
		imp = fnv.New64()
	case types.FNV1A:
		imp = fnv.New64a()
	case types.XXH64:
		imp = xxhash.New()
	case types.XXH3:
		imp = xxh3.New()
	case types.SDHASH:
		imp = new(SDHash)
	case types.SSDEEP:
		imp = ssdeep.New()
	case types.TLSH:
		imp = tlsh.New()
	case types.CRC32IEEE:
		imp = crc32.NewIEEE()
	case types.CRC64ISO:
		imp = crc64.New(crc64.MakeTable(crc64.ISO))
	case types.CRC64ECMA:
		imp = crc64.New(crc64.MakeTable(crc64.ECMA))
	default:
		return nil, errors.New("algorithm not recognized")
	}

	imp.Reset()

	_, err := imp.Write(*h.MMap())

	if err != nil {
		return nil, err
	}

	sum = imp.Sum(nil)

	h.Lock()
	h.hash[algo] = sum
	h.Unlock()

	return sum, nil
}

type SDHash struct {
	f sdhash.SdbfFactory
	s sdhash.Sdbf
}

func (sd *SDHash) Reset() {
	sd.f = nil
	sd.s = nil
}

func (sd *SDHash) BlockSize() int {
	return sdhash.BlockSize
}

func (sd *SDHash) Size() int {
	return int(sd.s.Size())
}

func (sd *SDHash) Sum(_ []byte) []byte {
	sd.s = sd.f.Compute()

	return []byte(strings.TrimRight(sd.s.String(), "\n"))
}

func (sd *SDHash) Write(b []byte) (int, error) {
	var err error

	sd.f, err = sdhash.CreateSdbfFromBytes(b)

	return len(b), err
}
