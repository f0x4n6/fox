package agent

import (
	"fmt"
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
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

const welcome = "How may I assist you today?"

type Agent struct {
	File fs.File
	busy atomic.Bool

	hs *heapset.HeapSet

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

	_, _ = a.File.WriteString(welcome + "\n\n")

	a.busy.Store(false)

	go a.gather()

	return a
}

func (a *Agent) IsBusy() bool {
	return a.busy.Load()
}

func (a *Agent) HeapSet(hs *heapset.HeapSet) {
	a.hs = hs
}

func (a *Agent) Process(query string) {
	query = strings.TrimSpace(query)

	if !a.parse(query) {
		a.query(query)
	}
}

func (a *Agent) Write(query string) {
	_, _ = a.File.WriteString(fmt.Sprintf("%c %s\n", text.Icons().Ps1, query))
}

func (a *Agent) Close() {
	close(a.ch)
}

func (a *Agent) query(query string) {
	col := a.rag.Embed("fox", a.ctx.Embed(), a.hs)

	if col == nil {
		return
	}

	ctx := a.rag.Query(query, col)

	if len(ctx) == 0 {
		return
	}

	a.busy.Store(true)

	err := a.llm.Query(a.ctx.Model(), query, ctx, func(res api.ChatResponse) error {
		if len(res.Message.Content) == 0 {
			a.busy.Store(false)
			a.ch <- "\n\n"
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
	flg, end := flags.Get(), true

	var sb strings.Builder

	for s := range a.ch {
		// response start
		if end {
			s = strings.TrimSpace(s)
		}

		s = strings.Replace(s, "  ", "", 1)

		// response chunk
		if !flg.Print {
			_, _ = a.File.WriteString(s)
		} else {
			_, _ = fmt.Print(s)
		}

		// response end
		end = s == "\n\n"

		sb.WriteString(s)

		if end {
			a.llm.AddAssistant(sb.String())
			sb.Reset()
		}
	}
}
