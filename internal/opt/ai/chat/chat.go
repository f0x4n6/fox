package chat

import (
	"context"
	"fmt"
	"log"
	"path"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ai/chat/llm"
	"github.com/cuhsat/fox/internal/opt/ai/chat/rag"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/user/config"
)

const banner = "How may I help you today?"

type Chat struct {
	File fs.File

	state *opt.State
	heap  *heap.Heap

	llm *llm.LLM
	rag *rag.RAG

	resp chan string
	done context.CancelFunc
	down atomic.Uint32
	busy atomic.Bool
}

func New(state *opt.State, heap *heap.Heap) *Chat {
	c := &Chat{
		File: fs.Create(path.Join(heap.Path, "chat")),

		state: state,
		heap:  heap,

		llm: llm.New(state.Model(), time.Minute*30),
		rag: rag.New(),

		resp: make(chan string, 64),
	}

	c.stderr(fmt.Sprintln(banner))

	c.busy.Store(false)

	go c.listen()

	return c
}

func (c *Chat) Query(query string, echo bool) {
	var ctx context.Context

	if echo {
		c.stderr(fmt.Sprintf("\n%c %s\n", c.state.Icon.Ps1, query))
	}

	if c.done != nil {
		c.done() // stop current activity
	}

	ctx, c.done = context.WithCancel(context.Background())

	query = strings.TrimSpace(query)

	if !c.parse(ctx, query) {
		c.process(ctx, query)
	}
}

func (c *Chat) Close() {
	if c.done != nil {
		c.done()
	}

	close(c.resp)
}

func (c *Chat) process(ctx context.Context, query string) {
	col := c.rag.Embed(ctx, c.state.Embed(), c.heap)

	if col == nil {
		return
	}

	lines := c.rag.Query(ctx, query, col)

	if len(lines) == 0 {
		return
	}

	c.busy.Store(true)

	err := c.llm.Query(ctx, c.state.Model(), query, lines, func(res api.ChatResponse) error {
		if len(res.Message.Content) == 0 {
			c.busy.Store(false)
			c.resp <- "\r\n"
		} else {
			c.resp <- res.Message.Content
		}

		return nil
	})

	if err != nil {
		c.busy.Store(false)
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
	if flags.Get().Print {
		_, _ = fmt.Print(s)
	} else {
		c.stderr(s)
	}
}

func (c *Chat) stderr(s string) {
	_, _ = c.File.WriteString(s)
}

func (c *Chat) source() string {
	cfg := config.Get()

	return fmt.Sprintf(
		"\n%c Generated with %s:%d:%d:%.1f:%d:%.1f\n",
		c.state.Icon.HSep,
		c.state.Model(),
		cfg.GetInt("ai.num_ctx"),
		cfg.GetInt("ai.seed"),
		cfg.GetFloat64("ai.temp"),
		cfg.GetInt("ai.topk"),
		cfg.GetFloat64("ai.topp"),
	)
}
