package agent

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/ollama/ollama/api"

	"github.com/cuhsat/fox/internal/pkg/sys"
)

func (a *Agent) parse(query string) bool {
	var err error

	re := regexp.MustCompile(`^(set\s*(model|embed)\s*.+)|(del\s*.+)|list$`)

	if !re.MatchString(query) {
		return false
	}

	cmd, model, _ := strings.Cut(query, " ")

	switch cmd {
	case "set":
		err = a.addModel(model)
	case "del":
		err = a.delModel(model)
	case "list":
		err = a.models()
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
	a.down.Store(0)

	err := a.llm.AddModel(model, func(res api.ProgressResponse) error {
		p := uint32((float32(res.Completed) / float32(res.Total)) * 100)

		if a.busy.Load() && (p > a.down.Load() && p < 100) {
			a.output(fmt.Sprintf("Downloading %s ...% 2d%%\n", model, p))
			a.down.Store(p)
		}

		if p == 100 && a.busy.Load() {
			a.output(fmt.Sprintf("Using model %s\n\n", model))
			a.busy.Store(false)
		}

		return nil
	})

	if err != nil {
		a.busy.Store(false)
		return err
	}

	switch v {
	case "model":
		a.ctx.ChangeModel(model)
	case "embed":
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

	if a.ctx.Model() == model {
		a.ctx.ChangeModel("")
	}

	if a.ctx.Embed() == model {
		a.ctx.ChangeEmbed("")
	}

	a.output(fmt.Sprintf("Deleted model %s\n\n", model))

	return nil
}

func (a *Agent) models() error {
	var ms []string

	res, err := a.llm.Models()

	if err != nil {
		return err
	}

	for _, m := range res.Models {
		ms = append(ms, m.Name)
	}

	slices.Sort(ms)

	for _, m := range ms {
		a.output(fmt.Sprintln(m))
	}

	a.output("\n")

	return nil
}
