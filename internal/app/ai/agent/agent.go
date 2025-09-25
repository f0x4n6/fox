package agent

import (
	"context"
	"fmt"
	"path"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/app"
	"github.com/cuhsat/fox/internal/app/ai/agent/llm"
	"github.com/cuhsat/fox/internal/app/ai/agent/rag"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
)

const banner = "Hello, how may I help you please?"

type Agent struct {
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

func New(ctx *app.Context, heap *heap.Heap) *Agent {
	a := &Agent{
		File: fs.Create(path.Join(heap.Path, "agent")),
		heap: heap,

		ctx: ctx,
		llm: llm.New(ctx.Model(), time.Minute*30),
		rag: rag.New(),

		ch: make(chan string, 64),
	}

	a.output(fmt.Sprintln(banner))

	a.busy.Store(false)

	go a.listen()

	return a
}

func (a *Agent) Prompt(query string) {
	a.output(fmt.Sprintf("\n%c %s\n", text.Icons().Ps1, query))
}

func (a *Agent) Process(query string) {
	var ctx context.Context

	if a.done != nil {
		a.done() // stop current activity
	}

	ctx, a.done = context.WithCancel(context.Background())

	query = strings.TrimSpace(query)

	if !a.parse(ctx, query) {
		a.query(ctx, query)
	}
}

func (a *Agent) Close() {
	if a.done != nil {
		a.done()
	}

	close(a.ch)
}

func (a *Agent) query(ctx context.Context, query string) {
	col := a.rag.Embed(ctx, a.ctx.Embed(), a.heap)

	if col == nil {
		return
	}

	lines := a.rag.Query(ctx, query, col)

	if len(lines) == 0 {
		return
	}

	a.busy.Store(true)

	err := a.llm.Query(ctx, a.ctx.Model(), query, lines, func(res api.ChatResponse) error {
		if len(res.Message.Content) == 0 {
			a.busy.Store(false)
			a.ch <- "\r\n"
		} else {
			a.ch <- res.Message.Content
		}

		return nil
	})

	if err != nil {
		a.busy.Store(false)
		sys.Error(err)
	}
}

func (a *Agent) listen() {
	var sb strings.Builder

	flg, end := flags.Get(), true

	for s := range a.ch {
		// response start
		if end {
			s = strings.TrimSpace(s)
		}

		s = strings.Replace(s, "  ", "", 1)

		// response chunk
		if !flg.Print {
			a.output(s)
		} else {
			_, _ = fmt.Print(s)
		}

		// response end
		end = s == "\r\n"

		sb.WriteString(s)

		if end {
			a.llm.AddAssistant(sb.String())
			sb.Reset()
		}
	}
}

func (a *Agent) output(s string) {
	_, _ = a.File.WriteString(s)
}
