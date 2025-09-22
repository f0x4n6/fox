//go:build !minimal

package ai

import (
	"context"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/app"
	"github.com/cuhsat/fox/internal/app/ai/agent"
)

func NewAgent(ctx *app.Context) *agent.Agent {
	client, err := api.ClientFromEnvironment()

	if err != nil {
		return nil // no client found
	}

	cto, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = client.Heartbeat(cto)

	if err != nil {
		return nil // no server found
	}

	return agent.New(ctx)
}
