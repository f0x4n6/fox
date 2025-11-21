package schema

import (
	"github.com/cuhsat/fox/v4/internal/pkg/files/evidence"
)

type Schema interface {
	String() string
	Headers() map[string]string
	SetMeta(meta evidence.Meta)
	AddLine(nr, grp uint, str string)
}
