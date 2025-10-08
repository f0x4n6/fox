package cef

import (
	"fmt"
	"strings"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/pkg/files/extract"
)

func Format(e *extract.Event) string {
	e.Data = strings.ReplaceAll(e.Data, `\`, `\\`)
	e.Data = strings.ReplaceAll(e.Data, `|`, `\|`)

	if len(e.Data) > 512 {
		e.Data = e.Data[:512]
	}

	return fmt.Sprintf(
		"%s %s CEF:1|%s|fox|%s|100|%s|Unknown",
		e.Time.Format("Jan 02 2006 15:04:05.000"),
		e.Host,
		fox.Product,
		fox.Version,
		e.Data,
	)
}
