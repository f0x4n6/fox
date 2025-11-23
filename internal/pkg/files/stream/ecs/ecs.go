// Package ecs specification: https://www.elastic.co/docs/reference/ecs/ecs-field-reference
package ecs

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/cuhsat/fox/v4/internal"
)

const version = "9.1.0"

type Ecs struct {
	Timestamp time.Time `json:"@timestamp"`
	Message   string    `json:"message"`
	Agent     struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"agent"`
	Ecs struct {
		Version string `json:"version"`
	} `json:"ecs"`
}

func New() *Ecs {
	ecs := new(Ecs)
	ecs.Ecs.Version = version
	ecs.Agent.Type = fox.Product
	ecs.Agent.Version = fox.Version[1:]
	return ecs
}

func (ecs *Ecs) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (ecs *Ecs) String() string {
	buf, err := json.Marshal(ecs)

	if err != nil {
		log.Print(err)
	}

	return string(buf)
}

func (ecs *Ecs) Write(s string) {
	ecs.Timestamp = time.Now().UTC()
	ecs.Message = strings.TrimRight(s, "\n")
}
