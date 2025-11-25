package hunt

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/0xrawsec/golang-evtx/evtx"

	"github.com/cuhsat/fox/v4/internal"
)

const cef = "%s %s CEF:1|fox|hunt|%s|100|%s|%d|"

var hostPath = evtx.Path("/Event/System/Computer")

var Nil any

type Set map[uint64]any

type Log struct {
	Time      time.Time
	Host      string
	Message   string
	Severity  int8
	Extension map[string]string
}

func Transform(evt *evtx.GoEvtxMap) *Log {
	var ok bool

	log := Log{
		Time:      evt.TimeCreated().UTC(),
		Host:      evt.GetStringStrict(&hostPath),
		Extension: make(map[string]string),
	}

	if log.Message, ok = Database[evt.EventID()]; !ok {
		log.Message = "The event is undescribed"
	}

	if log.Severity, ok = Levels[evt.EventID()]; !ok {
		log.Severity = 0 // unknown
	}

	log.Extension["channel"] = evt.Channel()
	log.Extension["eventid"] = strconv.Itoa(int(evt.EventID()))
	log.Extension["userid"], _ = evt.UserID()

	return &log
}

func (log *Log) String() string {
	var sb strings.Builder

	m := log.Message

	m = strings.ReplaceAll(m, `\`, `\\`)
	m = strings.ReplaceAll(m, `|`, `\|`)

	if len(m) > 512 {
		m = m[:512]
	}

	sb.WriteString(fmt.Sprintf(cef,
		log.Time.Format("Jan 02 2006 15:04:05.000"),
		log.Host,
		app.Version[1:],
		m,
		log.Severity,
	))

	for k, v := range log.Extension {
		if len(v) > 0 {
			sb.WriteString(fmt.Sprintf("%s=%s ", k, v))
		}
	}

	return strings.TrimSpace(sb.String())
}
