package journal

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/Velocidex/go-journalctl/parser"
	"github.com/Velocidex/ordereddict"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/types/event"
)

const Magic = "LPKSHHRH"

func Detect(b []byte) bool {
	return files.HasMagic(b, 0, []byte(Magic))
}

func Format(b []byte) ([]byte, error) {
	j, err := parser.OpenFile(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)

	for l := range j.GetLogs(context.Background()) {
		_, err := buf.WriteString(fmt.Sprintf("%v\n", l))

		if err != nil {
			log.Println(err)
		}
	}

	return buf.Bytes(), err
}

func Decode(b []byte, off int64) (*parser.JournalFile, error) {
	return parser.OpenFile(bytes.NewReader(b[off:]))
}

func ToEvent(od *ordereddict.Dict, ext bool) *event.Event {
	var sys, evt *ordereddict.Dict

	l := event.Event{Extension: make(map[string]string)}

	if v, ok := od.Get("System"); ok {
		sys = v.(*ordereddict.Dict)
	} else {
		log.Print("System not found")
		return &l
	}

	if v, ok := od.Get("EventData"); ok {
		evt = v.(*ordereddict.Dict)
	} else {
		log.Print("EventData not found")
		return &l
	}

	for _, k := range []string{
		"_SOURCE_REALTIME_TIMESTAMP",
		"SYSLOG_TIMESTAMP",
		"Timestamp",
	} {
		if v, ok := sys.Get(k); ok {
			l.Time = v.(time.Time).UTC()
			break
		}
	}

	l.Host, _ = sys.GetString("_HOSTNAME")
	l.Message, _ = evt.GetString("MESSAGE")

	if len(l.Message) == 0 {
		l.Message = "Undescribed event"
	}

	if v, ok := evt.GetInt64("PRIORITY"); !ok {
		l.Severity = 10 - int8(v) // minimum 3
	}

	if ext {
		for _, i := range append(sys.Items(), evt.Items()...) {
			if !slices.Contains(ignore, i.Key) && !strings.HasPrefix(i.Key, "(") {
				l.Extension[i.Key] = fmt.Sprintf("%v", i.Value)
			}
		}
	}

	return &l
}

var ignore = []string{
	"Timestamp",
	"MESSAGE",
	"PRIORITY",
	"SYSLOG_TIMESTAMP",
	"_SOURCE_REALTIME_TIMESTAMP",
	"_HOSTNAME",
}
