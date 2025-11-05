package flags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cuhsat/fox/v3/internal/pkg/types"
)

type Algorithms struct {
	Value []string // patterns
}

func (a *Algorithms) String() string {
	return fmt.Sprintf("%v", a.Value)
}

func (a *Algorithms) Type() string {
	return "strings"
}

func (a *Algorithms) Set(v string) error {
	switch strings.ToLower(v) {

	// MD5
	case types.MD5:
		fallthrough

	// SHA family
	case types.SHA1, types.SHA256, types.SHA3:
		fallthrough

	// SHA3 family
	case types.SHA3224, types.SHA3256, types.SHA3384, types.SHA3512:
		fallthrough

	// BLAKE family
	case types.BLAKE3256, types.BLAKE3512:
		fallthrough

	// Similarity hashes
	case types.SDHASH, types.SSDEEP, types.TLSH:
		fallthrough

	// FNV-1 family
	case types.FNV1, types.FNV1A:
		fallthrough

	// XXH family
	case types.XXH64, types.XXH3:
		fallthrough

	// CRC32 family
	case types.CRC32IEEE, types.CRC64ISO, types.CRC64ECMA:
		a.Value = append(a.Value, v)
		return nil

	default:
		return errors.New("algorithm not recognized")
	}
}
