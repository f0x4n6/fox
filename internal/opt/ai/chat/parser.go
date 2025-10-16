package chat

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ollama/ollama/api"
)

const syntax = `^(stop|list|(get\s+(model|embed))|(set\s+(model|embed)\s+.+)|(del\s+.+))$`

func (c *Chat) parse(ctx context.Context, query string) bool {
	if !regexp.MustCompile(syntax).MatchString(query) {
		return false // no command
	}

	var err error

	cmd, param, _ := strings.Cut(query, " ")

	switch strings.ToLower(cmd) {
	case "stop":
		c.stop(ctx)
	case "list":
		err = c.getModel(ctx, "")
	case "get":
		err = c.getModel(ctx, param)
	case "set":
		err = c.setModel(ctx, param)
	case "del":
		err = c.delModel(ctx, param)
	default:
		err = fmt.Errorf("unknown command")
	}

	if err != nil {
		log.Println(err)
	}

	return true
}

func (c *Chat) stop(ctx context.Context) {
	ctx.Done()
	c.stderr("Stopped\n")
}

func (c *Chat) setModel(ctx context.Context, param string) error {
	v, model, _ := strings.Cut(param, " ")

	err := c.download(ctx, model)

	if err != nil {
		return err
	}

	switch strings.ToLower(v) {
	case "model":
		c.state.ChangeModel(model)
	case "embed":
		c.state.ChangeEmbed(model)
	default:
		return fmt.Errorf("unknown target")
	}

	return nil
}

func (c *Chat) getModel(ctx context.Context, param string) error {
	switch strings.ToLower(param) {
	case "":
		res, err := c.llm.Models(ctx)

		if err != nil {
			return err
		}

		for _, m := range res.Models {
			c.stderr(fmt.Sprintln(m.Name))
		}

	case "model":
		c.stderr(fmt.Sprintln(c.state.Model()))

	case "embed":
		c.stderr(fmt.Sprintln(c.state.Embed()))

	default:
		return fmt.Errorf("unknown target")
	}

	return nil
}

func (c *Chat) delModel(ctx context.Context, param string) error {
	err := c.llm.DelModel(ctx, param)

	if err != nil {
		return err
	}

	if c.state.Model() == param {
		c.state.ChangeModel("")
	}

	if c.state.Embed() == param {
		c.state.ChangeEmbed("")
	}

	c.stderr(fmt.Sprintf("Deleted model %s\n\n", param))

	return nil
}

func (c *Chat) download(ctx context.Context, model string) error {
	c.busy.Store(true)
	c.down.Store(0)

	err := c.llm.AddModel(ctx, model, func(res api.ProgressResponse) error {
		p := uint32((float32(res.Completed) / float32(res.Total)) * 100)

		if c.busy.Load() && (p > c.down.Load() && p < 100) {
			c.stderr(fmt.Sprintf("Downloading %s %2d%%\n", model, p))
			c.down.Store(p)
		}

		if p == 100 && c.busy.Load() {
			c.stderr(fmt.Sprintf("Using model %s\n", model))
			c.busy.Store(false)
		}

		return nil
	})

	if err != nil {
		c.busy.Store(false)
		c.down.Store(0)
	}

	return err
}
