package dissect

import (
	"regexp"
	"time"

	"github.com/cuhsat/fox/v3/internal/pkg/files/events"
)

var re = regexp.MustCompile(`"hostname":\s*"(?P<host>.+?)".*"ts":\s*"(?P<time>.+?)".*"EventID":\s*(?P<data>.+?),`)

func Extract(line string) *events.Event {
	m := re.FindStringSubmatch(line)

	if len(m) < 4 {
		return nil
	}

	x := re.SubexpIndex("time")
	y := re.SubexpIndex("host")
	z := re.SubexpIndex("data")

	ts, err := time.Parse(time.RFC3339Nano, m[x])

	if err != nil {
		return nil
	}

	ev := &events.Event{
		Time: ts.UTC(),
		Host: m[y],
		Data: "EventID " + m[z],
	}

	if e, ok := events.Database[m[z]]; ok {
		ev.Data = e.Msg
	}

	return ev
}
