package extract

import (
	"time"
)

type Event struct {
	// Event timestamp
	Time time.Time
	// Event hostname
	Host string
	// Event message
	Data string
}
