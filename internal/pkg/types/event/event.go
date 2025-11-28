package event

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/cuhsat/fox/v4/internal"
)

const CEF = "%s %s CEF:1|fox|hunt|%s|100|%s|%d|"

type Event struct {
	Time      time.Time
	Host      string
	Message   string
	Severity  int8
	Extension map[string]any
}

func (evt *Event) String() string {
	var sb strings.Builder

	msg := evt.Message
	msg = strings.ReplaceAll(msg, `\`, `\\`)
	msg = strings.ReplaceAll(msg, `|`, `\|`)

	if len(msg) > 512 {
		msg = msg[:512]
	}

	sb.WriteString(fmt.Sprintf(CEF,
		evt.Time.Format("Jan 02 2006 15:04:05.000"),
		evt.Host,
		app.Version[1:],
		msg,
		evt.Severity,
	))

	for _, k := range slices.Sorted(maps.Keys(evt.Extension)) {
		if v := evt.Extension[k]; v != nil {
			s := fmt.Sprintf("%v", v)

			k = strings.ReplaceAll(k, `=`, `\=`)
			s = strings.ReplaceAll(s, `=`, `\=`)

			sb.WriteString(fmt.Sprintf("%s=%s ", k, s))
		}
	}

	return strings.TrimSpace(sb.String())
}
