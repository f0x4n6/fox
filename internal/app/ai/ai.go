//go:build !minimal

package ai

import (
	"context"
	"time"

	"github.com/ollama/ollama/api"
)

func Check() bool {
	client, err := api.ClientFromEnvironment()

	if err != nil {
		return false // no client found
	}

	cto, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = client.Heartbeat(cto)

	if err != nil {
		return false // no server found
	}

	return true
}
