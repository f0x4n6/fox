package plain

import (
	"fmt"

	"github.com/cuhsat/fox/v3/internal/pkg/files/events"
)

func Format(e *events.Event) string {
	return fmt.Sprintf(
		"%s %s %s",
		e.Time.Format("2006-01-02T15:04:05.000Z"),
		e.Host,
		e.Data,
	)
}
