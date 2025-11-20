package cmd

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/0xrawsec/golang-utils/log"
	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/v4/internal/pkg/types/page"
)

type Hunt struct {
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Info struct {
	Min   float64  `default:"0.0"`
	Max   float64  `default:"1.0"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Text struct {
	Min   int      `default:"3"`
	Max   int      `default:"256"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Hash struct {
	Type  []string `sep:"," default:"SHA256"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Dump struct {
	Paths []string `arg:"" name:"path" type:"path"`
}

type Show struct {
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Globals struct {
	// Commands
	Hunt Hunt `cmd:"" aliases:"u"`
	Info Info `cmd:"" aliases:"i"`
	Text Text `cmd:"" aliases:"s,strings"`
	Hash Hash `cmd:"" aliases:"h"`
	Dump Dump `cmd:"" aliases:"x,hex"`
	Show Show `cmd:"" default:"withargs" aliases:"p,print"`

	// File limits
	Head  bool `short:"h" xor:"head,tail"`
	Tail  bool `short:"t" xor:"head,tail"`
	Lines int  `short:"n" default:"10"`
	Bytes int  `short:"c" default:"16"`

	// File loader
	Pass string `short:"p"`

	// Line filter
	Regex   string `short:"e"`
	Context int    `short:"C"`
	Before  int    `short:"B"`
	After   int    `short:"A"`

	// Evidence bag
	File string `short:"f"`
	Mode string `short:"m" enum:"none,text,json,jsonl,sqlite" default:"text"`

	// Evidence sign
	Sign string `short:"s"`

	// Evidence URL
	Url  string `short:"u"`
	Auth string `short:"a"`
	Ecs  bool   `xor:"ecs,hec"`
	Hec  bool   `xor:"ecs,hec" and:"hec,auth"`

	// Turn off
	Readonly  bool `short:"R"`
	Raw       bool `short:"r"`
	NoFile    bool `long:"no-file"`
	NoLine    bool `long:"no-line"`
	NoDeflate bool `long:"no-deflate"`
	NoConvert bool `long:"no-convert"`

	// Shortcuts
	Logstash bool `short:"L"`
	Splunk   bool `short:"S"`
	Sqlite   bool `short:"Q"`
	Jsonl    bool `short:"J"`
	Json     bool `short:"j"`
}

func (cmd *Hunt) Run(cli *Globals) error {
	return nil // TODO
}

func (cmd *Info) Run(cli *Globals) error {
	hs := load(cli.Info.Paths, cli)
	defer hs.ThrowAway()

	for _, h := range hs.Get() {
		if e := h.Entropy(
			cli.Info.Min,
			cli.Info.Max,
		); e != -1 {
			fmt.Printf("%8dL %8dB %.10f  %s\n", h.Len(), len(h.MMap()), e, h.String())
		}
	}

	return nil
}

func (cmd *Text) Run(cli *Globals) error {
	hs := load(cli.Text.Paths, cli)
	defer hs.ThrowAway()

	for _, h := range hs.Get() {
		for s := range h.Strings(
			cli.Text.Min,
			cli.Text.Max,
		) {
			if !cli.NoLine {
				fmt.Printf("%08x  %s\n", s.Off, strings.TrimSpace(s.Str))
			} else {
				fmt.Println(strings.TrimSpace(s.Str))
			}
		}
	}

	return nil
}

func (cmd *Hash) Run(cli *Globals) error {
	hs := load(cli.Hash.Paths, cli)
	defer hs.ThrowAway()

	for _, algo := range cli.Hash.Type {
		if len(cli.Hash.Type) > 1 {
			fmt.Println(text.Block(strings.ToUpper(algo), page.TermW))
		}

		for _, h := range hs.Get() {
			sum, err := h.HashSum(algo)

			if err != nil {
				log.Errorf("could not compute hash: %s", err.Error())
			}

			switch algo {
			case types.SDHASH:
				fmt.Printf("%s  %s\n", sum, h.String())
			default:
				fmt.Printf("%x  %s\n", sum, h.String())
			}
		}
	}

	return nil
}

func (cmd *Dump) Run(cli *Globals) error {
	hs := load(cli.Dump.Paths, cli)
	defer hs.ThrowAway()

	t := 0

	if cli.Tail {
		t = cli.Bytes
	}

	for _, h := range hs.Get() {
		if hs.Len() > 1 && !cli.NoFile {
			fmt.Println(text.Block(h.String(), page.TermW))
		}

		for l := range page.Hex(h, t).Lines {
			fmt.Println(l)
		}
	}

	return nil
}

func (cmd *Show) Run(cli *Globals) error {
	hs := load(cli.Show.Paths, cli)
	defer hs.ThrowAway()

	for _, h := range hs.Get() {
		if hs.Len() > 1 && !cli.NoFile {
			fmt.Println(text.Block(h.String(), page.TermW))
		}

		for l := range page.Text(h, 2).Lines {
			if !cli.NoLine && l.Nr == page.Sep {
				fmt.Println("--")
			} else if !cli.NoLine {
				fmt.Printf("%s %s\n", l.Nr, l)
			} else {
				fmt.Println(l)
			}
		}
	}

	return nil
}

func load(args []string, cli *Globals) *heapset.HeapSet {
	var re *regexp.Regexp

	if len(cli.Regex) > 0 {
		re = regexp.MustCompile(cli.Regex)
	}

	if cli.Context > 0 {
		cli.Before = cli.Context
		cli.After = cli.Context
	}

	if len(cli.File) == 0 {
		cli.File = time.Now().Format("2006-01-02")
	}

	if cli.Raw {
		cli.NoFile = true
		cli.NoLine = true
		cli.NoConvert = true
		cli.NoDeflate = true
	}

	if cli.Readonly {
		cli.Mode = types.NONE
	}

	if cli.Logstash {
		cli.Url = types.LOGSTASH
		cli.Ecs = true
	}

	if cli.Splunk {
		cli.Url = types.SPLUNK
		cli.Hec = true
	}

	if cli.Sqlite {
		cli.Mode = types.SQLITE
	}

	if cli.Jsonl {
		cli.Mode = types.JSONL
	}

	if cli.Json {
		cli.Mode = types.JSON
	}

	return heapset.New(args, &heapset.Options{
		Limit: &types.Limits{
			IsHead: cli.Head,
			IsTail: cli.Tail,
			Lines:  cli.Lines,
			Bytes:  cli.Bytes,
		},
		Filter: &types.Filters{
			Regex:  re,
			Before: cli.Before,
			After:  cli.After,
		},
		Password:  cli.Pass,
		NoDeflate: cli.NoDeflate,
		NoConvert: cli.NoConvert,
	})
}
