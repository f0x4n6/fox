package hunt

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cuhsat/fox/v4/internal"
)

const header = "%s %s CEF:1|fox|hunt|%s|100|%s|Unknown"

type Log struct {
	Ts      time.Time
	Id      int64
	User    string
	Host    string
	Channel string
	Message string
	// TODO: implement as map[string]string and join as extension
}

func (log *Log) String() string {
	var sb strings.Builder

	m := log.Message

	m = strings.ReplaceAll(m, `\`, `\\`)
	m = strings.ReplaceAll(m, `|`, `\|`)

	if len(m) > 512 {
		m = m[:512]
	}

	sb.WriteString(fmt.Sprintf(
		header,
		log.Ts.Format("Jan 02 2006 15:04:05.000"),
		log.Host,
		app.Version[1:],
		m,
	))

	if len(log.Channel) > 0 {
		sb.WriteString("|channel=")
		sb.WriteString(log.Channel)
	}

	if len(log.User) > 0 {
		sb.WriteString(" userid=")
		sb.WriteString(log.User)
	}

	if log.Id > 0 {
		sb.WriteString(" eventid=")
		sb.WriteString(strconv.Itoa(int(log.Id)))
	}

	return sb.String()
}

type Logs map[uint64]*Log
