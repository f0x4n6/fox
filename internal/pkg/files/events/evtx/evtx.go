package evtx

import (
	"regexp"
	"time"

	"github.com/cuhsat/fox/v3/internal/pkg/files/events"
)

var re = regexp.MustCompile(`"Computer":\s*"(?P<host>.*?)".*"(EventID|Value)":\s*"(?P<data>.*?)".*"SystemTime":\s*"(?P<time>.+?)"`)

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

	ev := &events.Event{
		Time: ts.UTC(),
		Host: m[y],
		Data: "EventID " + m[z],
	}

	if e, ok := db[m[z]]; ok {
		ev.Data = e.msg
	}

	return ev
}
