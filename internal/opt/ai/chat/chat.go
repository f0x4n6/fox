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
	a := &Chat{
		File: fs.Create(path.Join(heap.Path, "chat")),

		state: state,
		heap:  heap,

		llm: llm.New(state.Model(), time.Minute*30),
		rag: rag.New(),

		resp: make(chan string, 64),
	}

	a.write(fmt.Sprintln(banner))

	a.busy.Store(false)

	go a.listen()

	return a
}

func (c *Chat) Query(query string, echo bool) {
	var ctx context.Context

	if echo {
		c.write(fmt.Sprintf("\n%c %s\n", c.state.Icon.Ps1, query))
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

	flg, end := flags.Get(), true

	for s := range c.resp {
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
