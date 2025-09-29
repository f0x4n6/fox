package chat

import (
	"context"
	"fmt"
	"path"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/app"
	"github.com/cuhsat/fox/internal/app/ai/chat/llm"
	"github.com/cuhsat/fox/internal/app/ai/chat/rag"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
)

const banner = "How may I help you today?"

type Chat struct {
	File fs.File

	done context.CancelFunc
	down atomic.Uint32
	busy atomic.Bool

	heap *heap.Heap

	ctx *app.Context
	llm *llm.LLM
	rag *rag.RAG

	ch chan string
}

func New(ctx *app.Context, heap *heap.Heap) *Chat {
	a := &Chat{
		File: fs.Create(path.Join(heap.Path, "chat")),
		heap: heap,

		ctx: ctx,
		llm: llm.New(ctx.Model(), time.Minute*30),
		rag: rag.New(),

		ch: make(chan string, 64),
	}

	a.write(fmt.Sprintln(banner))

	a.busy.Store(false)

	go a.listen()

	return a
}

func (c *Chat) Query(query string, echo bool) {
	var ctx context.Context

	if echo {
		c.write(fmt.Sprintf("\n%c %s\n", c.ctx.Icon.Ps1, query))
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

	close(c.ch)
}

func (c *Chat) process(ctx context.Context, query string) {
	col := c.rag.Embed(ctx, c.ctx.Embed(), c.heap)

	if col == nil {
		return
	}

	lines := c.rag.Query(ctx, query, col)

	if len(lines) == 0 {
		return
	}

	c.busy.Store(true)

	err := c.llm.Query(ctx, c.ctx.Model(), query, lines, func(res api.ChatResponse) error {
		if len(res.Message.Content) == 0 {
			c.busy.Store(false)
			c.ch <- "\r\n"
		} else {
			c.ch <- res.Message.Content
		}

		return nil
	})

	if err != nil {
		c.busy.Store(false)
		sys.Error(err)
	}
}

func (c *Chat) listen() {
	var sb strings.Builder

	flg, end := flags.Get(), true

	for s := range c.ch {
		// response start
		if end {
			s = strings.TrimSpace(s)
		}

		s = strings.Replace(s, "  ", "", 1)

		// response chunk
		if !flg.Print {
			c.write(s)
		} else {
			_, _ = fmt.Print(s)
		}

		// response end
		end = s == "\r\n"

		sb.WriteString(s)

		if end {
			c.llm.AddAssistant(sb.String())
			sb.Reset()
		}
	}
}

func (c *Chat) write(s string) {
	_, _ = c.File.WriteString(s)
}
