package agent

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/pkg/sys"
)

const syntax = `^(stop|list|(get\s+(model|embed))|(set\s+(model|embed)\s+.+)|(del\s+.+))$`

func (a *Agent) parse(ctx context.Context, query string) bool {
	if !regexp.MustCompile(syntax).MatchString(query) {
		return false // no command
	}

	var err error

	cmd, param, _ := strings.Cut(query, " ")

	switch cmd {
	case "stop":
		a.stop(ctx)
	case "list":
		err = a.getModel(ctx, "")
	case "get":
		err = a.getModel(ctx, param)
	case "set":
		err = a.setModel(ctx, param)
	case "del":
		err = a.delModel(ctx, param)
	default:
		err = fmt.Errorf("unknown command")
	}

	if err != nil {
		sys.Error(err)
	}

	return true
}

func (a *Agent) stop(ctx context.Context) {
	ctx.Done()
	a.output("Stopped\n")
}

func (a *Agent) setModel(ctx context.Context, param string) error {
	v, model, _ := strings.Cut(param, " ")

	err := a.download(ctx, model)

	if err != nil {
		return err
	}

	switch v {
	case "model":
		a.ctx.ChangeModel(model)
	case "embed":
		a.ctx.ChangeEmbed(model)
	default:
		return fmt.Errorf("unknown target")
	}

	return nil
}

func (a *Agent) getModel(ctx context.Context, param string) error {
	switch param {
	case "":
		res, err := a.llm.Models(ctx)

		if err != nil {
			return err
		}

		for _, m := range res.Models {
			a.output(fmt.Sprintln(m.Name))
		}

	case "model":
		a.output(fmt.Sprintln(a.ctx.Model()))

	case "embed":
		a.output(fmt.Sprintln(a.ctx.Embed()))

	default:
		return fmt.Errorf("unknown target")
	}

	return nil
}

func (a *Agent) delModel(ctx context.Context, param string) error {
	err := a.llm.DelModel(ctx, param)

	if err != nil {
		return err
	}

	if a.ctx.Model() == param {
		a.ctx.ChangeModel("")
	}

	if a.ctx.Embed() == param {
		a.ctx.ChangeEmbed("")
	}

	a.output(fmt.Sprintf("Deleted model %s\n\n", param))

	return nil
}

func (a *Agent) download(ctx context.Context, model string) error {
	a.busy.Store(true)
	a.down.Store(0)

	err := a.llm.AddModel(ctx, model, func(res api.ProgressResponse) error {
		p := uint32((float32(res.Completed) / float32(res.Total)) * 100)

		if a.busy.Load() && (p > a.down.Load() && p < 100) {
			a.output(fmt.Sprintf("Downloading %s %2d%%\n", model, p))
			a.down.Store(p)
		}

		if p == 100 && a.busy.Load() {
			a.output(fmt.Sprintf("Using model %s\n", model))
			a.busy.Store(false)
		}

		return nil
	})

	if err != nil {
		a.busy.Store(false)
		a.down.Store(0)
	}

	return err
}
