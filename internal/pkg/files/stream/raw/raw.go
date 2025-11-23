package raw

import "strings"

type Raw struct {
	body strings.Builder
}

func New() *Raw {
	return new(Raw)
}

func (raw *Raw) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "text/plain",
	}
}

func (raw *Raw) String() string {
	return raw.body.String()
}

func (raw *Raw) Write(s string) {
	raw.body.WriteString(s)
}
