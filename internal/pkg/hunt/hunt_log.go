package hunt

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
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
			sb.WriteString(fmt.Sprintf("%s=%s ", k, v))
		}
	}

	return strings.TrimSpace(sb.String())
}

func FromEvtx(evt *evtx.GoEvtxMap) *Log {
	var ok bool

	log := Log{
		Time:      evt.TimeCreated().UTC(),
		Host:      evt.GetStringStrict(&hostPath),
		Extension: make(map[string]string),
	}

	if log.Message, ok = Events[evt.EventID()]; !ok {
		log.Message = fmt.Sprintf("Event ID %d (undescribed)", evt.EventID())
	}

	if log.Severity, ok = Levels[evt.EventID()]; !ok {
		log.Severity = 0 // unknown
	}

	log.Extension["channel"] = evt.Channel()
	log.Extension["eventid"] = strconv.Itoa(int(evt.EventID()))
	log.Extension["userid"], _ = evt.UserID()

	return &log
}

func FromJournal(od *ordereddict.Dict) *Log {
	sys, _ := od.Get("System")
	evt, _ := od.Get("EventData")

	ts, _ := sys.GetInt64("Timestamp")

	host, _ := sys.GetString("_HOSTNAME")

	log := Log{
		Time:      time.Unix(ts, 0),
		Host:      host,
		Extension: make(map[string]string),
	}

	log.Message, _ = evt.GetString("MESSAGE")

	//	if log.Message, ok = Events[evt.EventID()]; !ok {
	//		log.Message = fmt.Sprintf("Event ID %d (undescribed)", evt.EventID())
	//	}

	//	if log.Severity, ok = Levels[evt.EventID()]; !ok {
	//		log.Severity = 0 // unknown
	//	}

	//	log.Extension["channel"] = evt.Channel()
	//	log.Extension["eventid"] = strconv.Itoa(int(evt.EventID()))
	//	log.Extension["userid"], _ = evt.UserID()

	return &log
}
