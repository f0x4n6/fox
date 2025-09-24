package handler

import (
	"sync"

	"github.com/cuhsat/fox/internal/app"
	"github.com/cuhsat/fox/internal/app/ai/agent"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
)

type Handler struct {
	ctx    *app.Context
	agents sync.Map
}

func New(ctx *app.Context) *Handler {
	return &Handler{
		ctx: ctx,
	}
}

func (h *Handler) NewAgent(name string, heap *heap.Heap) *agent.Agent {
	if a, ok := h.agents.Load(name); ok {
		return a.(*agent.Agent)
	}

	a := agent.New(h.ctx, heap)
	h.agents.Store(name, a)
	return a
}

func (h *Handler) GetAgent(name string) *agent.Agent {
	if a, ok := h.agents.Load(name); ok {
		return a.(*agent.Agent)
	}

	sys.Error("agent not found")
	return nil
}

func (h *Handler) Close() {
	h.agents.Range(func(name, a any) bool {
		a.(*agent.Agent).Close()
		return true
	})
}
