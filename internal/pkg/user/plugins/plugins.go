package plugins

import (
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/viper"

	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/user"
)

type Callback func(path, dir string)

type Plugins struct {
	plugins map[string]Plugin `mapstructure:"plugin"`
}

type Plugin struct {
	re *regexp.Regexp

	Name string
	Path string
	Exec []string
}

func New() *Plugins {
	ps := new(Plugins)

	cfg := viper.New()

	if !user.LoadConfig(cfg, "plugins") {
		return nil
	}

	err := cfg.Unmarshal(ps)

	if err != nil {
		log.Println(err)
		return nil
	}

	return ps
}

func (ps *Plugins) Plugins() []Plugin {
	rs := make([]Plugin, len(ps.plugins))

	for key := range ps.plugins {
		p := ps.plugins[key]
		p.re = regexp.MustCompile(p.Path)

		rs = append(rs, p)
	}

	return rs
}

func (p *Plugin) Match(path string) bool {
	if p.re != nil {
		return p.re.MatchString(path)
	} else {
		return false
	}
}

func (p *Plugin) Execute(file string, fn Callback) {
	var dir string

	// create temp dir if necessary
	if slices.ContainsFunc(p.Exec, func(s string) bool {
		return strings.Contains(s, "TEMP")
	}) {
		dir = user.TempDir("plugin")
	}

	// replace and persist
	rep := strings.NewReplacer(
		"FILE", user.Persist(file),
		"TEMP", dir,
	)

	cmds := make([]string, 0)

	for _, cmd := range p.Exec {
		cmds = append(cmds, rep.Replace(cmd))
	}

	fn(sys.Exec(cmds).Name(), dir)
}
