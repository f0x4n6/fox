package events

import (
	"time"
)

type Extract func(string) *Event

type Event struct {
	// Event timestamp
	Time time.Time
	// Event hostname
	Host string
	// Event message
	Data string
}
