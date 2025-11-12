package chat

import (
	"context"
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/v4/internal/opt/ai/chat/llm"
	"github.com/cuhsat/fox/v4/internal/opt/ai/chat/rag"
	"github.com/cuhsat/fox/v4/internal/pkg/flags"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v4/internal/pkg/user/config"
)

type Chat struct {
	File fs.File
	heap *heap.Heap

	llm *llm.LLM
	rag *rag.RAG

	resp chan string
}

func New(heap *heap.Heap) *Chat {
	flg := flags.Get()

	c := &Chat{
		File: fs.Create(path.Join(heap.Path, "chat")),

		heap: heap,

		llm: llm.New(flg.AI.Model, time.Minute*30),
		rag: rag.New(),

		resp: make(chan string, 64),
	}

	go c.listen()

	return c
}

func (c *Chat) Query(query string) {
	c.process(context.Background(), strings.TrimSpace(query))
}

func (c *Chat) Close() {
	close(c.resp)
}

func (c *Chat) process(ctx context.Context, query string) {
	col := c.rag.Embed(ctx, flags.Get().AI.Embed, c.heap)

	if col == nil {
		return
	}

	lines := c.rag.Query(ctx, query, col)

	if len(lines) == 0 {
		return
	}

	err := c.llm.Query(ctx, flags.Get().AI.Model, query, lines, func(res api.ChatResponse) error {
		if len(res.Message.Content) == 0 {
			c.resp <- "\r\n"
		} else {
			c.resp <- res.Message.Content
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

func (c *Chat) listen() {
	var sb strings.Builder
	var end = true

	for s := range c.resp {
		// response start
		if end {
			s = strings.TrimSpace(s)
		}

		s = strings.Replace(s, "  ", "", 1)

		end = s == "\r\n"

		// response chunk
		if end {
			c.stdout(c.source())
		} else {
			c.stdout(s)
		}

		// response end
		sb.WriteString(s)

		if end {
			c.llm.AddAssistant(sb.String())
			sb.Reset()
		}
	}
}

func (c *Chat) stdout(s string) {
	_, _ = fmt.Print(s)
}

func (c *Chat) stderr(s string) {
	_, _ = c.File.WriteString(s)
}

func (c *Chat) source() string {
	cfg := config.Get()
	flg := flags.Get()

	return fmt.Sprintf(
		"\n- Generated with %s:%d:%d:%.1f:%d:%.1f\n",
		flg.AI.Model,
		cfg.GetInt("ai.num_ctx"),
		cfg.GetInt("ai.seed"),
		cfg.GetFloat64("ai.temp"),
		cfg.GetInt("ai.topk"),
		cfg.GetFloat64("ai.topp"),
	)
}
