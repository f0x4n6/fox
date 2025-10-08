package journal

import (
	"regexp"
	"time"

	"github.com/cuhsat/fox/internal/pkg/files/events"
)

var re = regexp.MustCompile(`"Timestamp":\s*"(?P<time>.+?)".*"_HOSTNAME":\s*"(?P<host>.+?)".*"MESSAGE":\s*"(?P<data>.*?)"`)

func Extract(line string) *events.Event {
	m := re.FindStringSubmatch(line)

	if len(m) < 4 {
		return nil
	}

	x := re.SubexpIndex("time")
	y := re.SubexpIndex("host")
	z := re.SubexpIndex("data")

	ts, err := time.Parse(time.RFC3339, m[x])

	if err != nil {
		return nil
	}

	return &events.Event{
		Time: ts.UTC(),
		Host: m[y],
		Data: m[z],
	}
}
