package flags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cuhsat/fox/internal/pkg/types"
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
	case types.CRC32IEEE:
		fallthrough
	case types.CRC64ISO:
		fallthrough
	case types.CRC64ECMA:
		fallthrough
	case types.SDHASH:
		fallthrough
	case types.SSDEEP:
		fallthrough
	case types.TLSH:
		fallthrough
	case types.MD5:
		fallthrough
	case types.SHA1:
		fallthrough
	case types.SHA256:
		fallthrough
	case types.SHA3:
		fallthrough
	case types.SHA3224:
		fallthrough
	case types.SHA3256:
		fallthrough
	case types.SHA3384:
		fallthrough
	case types.SHA3512:
		a.Value = append(a.Value, v)
		return nil

	default:
		return errors.New("algorithm not recognized")
	}
}
