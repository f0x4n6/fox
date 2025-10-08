package log

import (
	"fmt"

	"github.com/cuhsat/fox/internal/pkg/files/extract"
)

func Format(e *extract.Event) string {
	return fmt.Sprintf(
		"%s %s %s",
		e.Time.Format("2006-01-02T15:04:05.000Z"),
		e.Host,
		e.Data,
	)
}
