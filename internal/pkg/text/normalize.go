package text

import (
	"github.com/cuhsat/fox/internal/pkg/files/extract"
	"github.com/cuhsat/fox/internal/pkg/files/extract/evtx"
	"github.com/cuhsat/fox/internal/pkg/files/extract/journal"
	"github.com/cuhsat/fox/internal/pkg/files/format/cef"
	"github.com/cuhsat/fox/internal/pkg/files/format/log"
)

type extractor func(string) *extract.Event

type formator func(e *extract.Event) string

func Normalize(s string, c bool) string {
	var fmt formator

	if c {
		fmt = cef.Format
	} else {
		fmt = log.Format
	}

	for _, fn := range []extractor{
		evtx.Extract,
		journal.Extract,
	} {
		if e := fn(s); e != nil {
			return fmt(e)
		}
	}

	return ""
}
