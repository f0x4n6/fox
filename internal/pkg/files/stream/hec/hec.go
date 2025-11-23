// Package hec specification: https://docs.splunk.com/Documentation/Splunk/latest/Data/FormateventsforHTTPEventCollector
package hec

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cuhsat/fox/v4/internal"
)

type Hec struct {
	token string

	Time   int64  `json:"time"`
	Source string `json:"source"`
	Event  string `json:"event"`
}

func New(token string) *Hec {
	return &Hec{
		token:  strings.ToLower(token),
		Source: fmt.Sprintf("%s %s", fox.Product, fox.Version),
	}
}

func (hec *Hec) Headers() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Splunk %s", hec.token),
	}
}

func (hec *Hec) String() string {
	buf, err := json.Marshal(hec)

	if err != nil {
		log.Print(err)
	}

	return string(buf)
}

func (hec *Hec) Write(s string) {
	hec.Time = time.Now().UTC().UnixMilli()
	hec.Event = strings.TrimRight(s, "\n")
}
