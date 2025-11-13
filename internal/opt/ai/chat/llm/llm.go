package llm

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/v4/internal"
	"github.com/cuhsat/fox/v4/internal/pkg/user/config"
)

type LLM struct {
	sync.RWMutex

	client  *api.Client   // chat client
	alive   *api.Duration // chat alive
	history []api.Message // chat history
}

func New(model string, keep time.Duration) *LLM {
	client, err := api.ClientFromEnvironment()

	if err != nil {
		log.Panicln(err)
	}

	alive := &api.Duration{Duration: keep}

	// preload model
	go func() {
		_ = client.Chat(context.Background(), &api.ChatRequest{
			Model:     model,
			KeepAlive: alive,
		}, func(cr api.ChatResponse) error {
			return nil // preloaded model
		})
	}()

	return &LLM{
		client:  client,
		alive:   alive,
		history: make([]api.Message, 0),
	}
}

func (llm *LLM) Query(ctx context.Context, model, query, lines string, fn api.ChatResponseFunc) error {
	llm.AddUser(fmt.Sprintf(fox.Prompt, query, lines))

	llm.RLock()

	cfg := config.Get()
	req := &api.ChatRequest{
		Model:     model,
		KeepAlive: llm.alive,
		Messages:  llm.history,
		Options: map[string]any{
			"num_ctx":     cfg.GetInt("ai.num_ctx"),
			"temperature": cfg.GetFloat64("ai.temp"),
			"seed":        cfg.GetInt("ai.seed"),
			"top_k":       cfg.GetInt("ai.top_k"),
			"top_p":       cfg.GetFloat64("ai.top_p"),
		},
	}

	llm.RUnlock()

	return llm.client.Chat(ctx, req, fn)
}

func (llm *LLM) AddUser(content string) {
	llm.addMessage("user", content)
}

func (llm *LLM) AddAssistant(content string) {
	llm.addMessage("assistant", content)
}

func (llm *LLM) addMessage(role, content string) {
	llm.Lock()
	llm.history = append(llm.history, api.Message{
		Role:    role,
		Content: content,
	})
	llm.Unlock()
}
