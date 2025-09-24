package agent

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	fox "github.com/cuhsat/fox/internal"
	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/app"
	"github.com/cuhsat/fox/internal/app/ai/agent/llm"
	"github.com/cuhsat/fox/internal/app/ai/agent/rag"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

const welcome = "Please formulate your current work hypothesis"

type Agent struct {
	File fs.File

	done context.CancelFunc
	down atomic.Uint32
	busy atomic.Bool

	ctx *app.Context
	llm *llm.LLM
	rag *rag.RAG

	ch chan string
}

func New(ctx *app.Context) *Agent {
	a := &Agent{
		File: fs.Create("/fox/agent"),

		ctx: ctx,
		llm: llm.New(ctx.Model(), time.Minute*30),
		rag: rag.New(),

		ch: make(chan string, 64),
	}

	a.output(fmt.Sprintf("%s\n%s\n", fox.Ascii, welcome))

	a.busy.Store(false)

	go a.gather()

	return a
}

func (a *Agent) Process(query string, hs *heapset.HeapSet) {
	var ctx context.Context

	if a.done != nil {
		a.done() // stop current activity
	}

	ctx, a.done = context.WithCancel(context.Background())

	query = strings.TrimSpace(query)

	if !a.parse(ctx, query) {
		a.query(ctx, query, hs)
	}
}

func (a *Agent) Write(query string) {
	a.output(fmt.Sprintf("\n%c %s\n", text.Icons().Ps1, query))
}

func (a *Agent) Close() {
	if a.done != nil {
		a.done()
	}

	close(a.ch)
}

func (a *Agent) query(ctx context.Context, query string, hs *heapset.HeapSet) {
	name := flags.Get().Bag.Case

	col := a.rag.Embed(ctx, name, a.ctx.Embed(), hs)

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

func (a *Agent) gather() {
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
