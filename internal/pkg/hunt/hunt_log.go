package hunt

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/0xrawsec/golang-evtx/evtx"
	"github.com/Velocidex/ordereddict"
	"github.com/cuhsat/fox/v4/internal"
)

const cef = "%s %s CEF:1|fox|hunt|%s|%d|%s|%d|"

var (
	hostPath = evtx.Path("/Event/System/Computer")
)

type Log struct {
	Time      time.Time
	Host      string
	Message   string
	Severity  int8
	Extension map[string]string
}

func (log *Log) String() string {
	var sb strings.Builder

	msg := log.Message
	msg = strings.ReplaceAll(msg, `\`, `\\`)
	msg = strings.ReplaceAll(msg, `|`, `\|`)

	if len(msg) > 512 {
		msg = msg[:512]
	}

	sb.WriteString(fmt.Sprintf(cef,
		log.Time.Format("Jan 02 2006 15:04:05.000"),
		log.Host,
		app.Version[1:],
		100,
		msg,
		log.Severity,
	))

	for _, k := range slices.Sorted(maps.Keys(log.Extension)) {
		if v := log.Extension[k]; len(v) > 0 {
			k = strings.ReplaceAll(k, `=`, `\=`)
			v = strings.ReplaceAll(v, `=`, `\=`)
			sb.WriteString(fmt.Sprintf("%s=%s ", k, v))
		}
	}

	return strings.TrimSpace(sb.String())
}

func FromEvtx(evt *evtx.GoEvtxMap, ext bool) *Log {
	var ok bool

	l := Log{
		Time:      evt.TimeCreated().UTC(),
		Host:      evt.GetStringStrict(&hostPath),
		Extension: make(map[string]string),
	}

	if l.Message, ok = Events[evt.EventID()]; !ok {
		l.Message = fmt.Sprintf("Undescribed event: Event ID %d", evt.EventID())
	}

	if l.Severity, ok = Levels[evt.EventID()]; !ok {
		l.Severity = 0 // unknown
	}

	if ext {
		for k, v := range maps.All(*evt) {
			if !slices.Contains(Ignore, k) {
				l.Extension[k] = fmt.Sprintf("%v", v)
			}
		}
	}

	return &l
}

func FromJournal(od *ordereddict.Dict, ext bool) *Log {
	var sys, evt *ordereddict.Dict

	l := Log{Extension: make(map[string]string)}

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
			if !slices.Contains(Ignore, i.Key) && !strings.HasPrefix(i.Key, "(") {
				l.Extension[i.Key] = fmt.Sprintf("%v", i.Value)
			}
		}
	}

	return &l
}
