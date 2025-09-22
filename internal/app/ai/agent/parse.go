package agent

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/pkg/sys"
)

// agent commands
const (
	Dir = "list"
	Add = "set"
	Del = "del"
)

// agent targets
const (
	Model = "model"
	Embed = "embed"
)

func (a *Agent) parse(query string) bool {
	cmd, model, _ := strings.Cut(query, " ")

	if !slices.Contains([]string{Add, Del, Dir}, cmd) {
		return false
	}

	var err error

	switch cmd {
	case Dir:
		err = a.list()
	case Add:
		err = a.addModel(model)
	case Del:
		err = a.delModel(model)
	default:
		err = fmt.Errorf("unknown command")
	}

	if err != nil {
		sys.Error(err)
	}

	return true
}

func (a *Agent) addModel(model string) error {
	v, model, _ := strings.Cut(model, " ")

	a.busy.Store(true)

	err := a.llm.AddModel(model, func(res api.ProgressResponse) error {
		if res.Completed >= res.Total {
			if a.busy.Load() {
				_, _ = a.File.WriteString(fmt.Sprintf("Using %s %s\n", v, model))
				a.busy.Store(false)
			}
		} else {
			a.ch <- "."
		}
		return nil
	})

	if err != nil {
		a.busy.Store(false)
		return err
	}

	switch v {
	case Model:
		a.ctx.ChangeModel(model)
	case Embed:
		a.ctx.ChangeEmbed(model)
	default:
		err = fmt.Errorf("unknown target")
	}

	return err
}

func (a *Agent) delModel(model string) error {
	err := a.llm.DelModel(model)

	if err != nil {
		return err
	}

	// clear model, if current is deleted
	if a.ctx.Model() == model {
		a.ctx.ChangeModel("")
	}

	return nil
}

func (a *Agent) list() error {
	var ms []string

	res, err := a.llm.List()

	if err != nil {
		return err
	}

	for _, m := range res.Models {
		ms = append(ms, m.Name)
	}

	slices.Sort(ms)

	for _, m := range ms {
		a.ch <- fmt.Sprintf("\n- %s", m)
	}

	a.ch <- "\n\n"

	return nil
}
