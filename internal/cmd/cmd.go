package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/cuhsat/fox/v4/internal/pkg/files/stream"
	"github.com/cuhsat/fox/v4/internal/pkg/files/stream/ecs"
	"github.com/cuhsat/fox/v4/internal/pkg/files/stream/hec"
	"github.com/cuhsat/fox/v4/internal/pkg/files/stream/raw"
	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/buffer"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heapset"
)

type Hunt struct {
	All   bool     `short:"a"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Info struct {
	Min   float64  `default:"0.0"`
	Max   float64  `default:"1.0"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Text struct {
	Min   uint     `default:"3"`
	Max   uint     `default:"256"`
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

type Cli struct {
	// Commands
	Hunt Hunt `cmd:"" aliases:"u"`
	Info Info `cmd:"" aliases:"i,wc"`
	Text Text `cmd:"" aliases:"s,strings"`
	Hash Hash `cmd:"" aliases:"h"`
	Dump Dump `cmd:"" aliases:"x,hex"`
	Show Show `cmd:"" default:"withargs" aliases:"p,print,less,cat"`

	// File limits
	Head  bool `short:"h" xor:"head,tail"`
	Tail  bool `short:"t" xor:"head,tail"`
	Lines uint `short:"n" xor:"lines,bytes"`
	Bytes uint `short:"c" xor:"lines,bytes"`

	// File loader
	Pass string `short:"p"`

	// Line filter
	Regex   string `short:"e"`
	Context uint   `short:"C"`
	Before  uint   `short:"B"`
	After   uint   `short:"A"`

	// Data stream
	File string `short:"f"`
	Url  string `short:"u"`
	Auth string `short:"a"`
	Ecs  bool   `short:"E" xor:"ecs,hec"`
	Hec  bool   `short:"H" xor:"ecs,hec" and:"hec,auth"`

	// Turn off
	Readonly  bool `short:"R"`
	Raw       bool `short:"r"`
	NoFile    bool `long:"no-file"`
	NoLine    bool `long:"no-line"`
	NoDeflate bool `long:"no-deflate"`
	NoConvert bool `long:"no-convert"`

	// Localhost
	Logstash bool `short:"L" xor:"logstash,splunk"`
	Splunk   bool `short:"S" xor:"logstash,splunk"`

	// Standard
	Verbose int `short:"v" type:"counter"`

	// Internal
	w  io.WriteCloser   `kong:"-"`
	hs *heapset.HeapSet `kong:"-"`
}

func (cli *Cli) Bootstrap(args []string) *heapset.HeapSet {
	var re *regexp.Regexp
	var sw io.Writer

	if cli.Readonly {
		cli.File = ""
		log.Println("File output deactivated")
	}

	if len(cli.Regex) > 0 {
		re = regexp.MustCompile(cli.Regex)
	}

	if len(cli.Url) > 0 {
		switch {
		case cli.Hec:
			sw = hec.New(cli.Url, cli.Auth)
		case cli.Ecs:
			sw = ecs.New(cli.Url)
		default:
			sw = raw.New(cli.Url)
		}
	}

	if cli.Info.Min > cli.Info.Max {
		log.Fatal("invalid range")
	}

	if cli.Text.Min > cli.Text.Max {
		log.Fatal("invalid range")
	}

	if cli.Context > 0 {
		cli.Before = cli.Context
		cli.After = cli.Context
	}

	if cli.Raw {
		cli.NoFile = true
		cli.NoLine = true
		cli.NoConvert = true
		cli.NoDeflate = true
	}

	if cli.Logstash {
		cli.Url = types.LOGSTASH
		cli.Ecs = true
	}

	if cli.Splunk {
		cli.Url = types.SPLUNK
		cli.Hec = true
	}

	if len(cli.File)+len(cli.Url) > 0 {
		cli.w = stream.New(cli.File, sw)
	} else {
		cli.w = os.Stdout
	}

	cli.hs = heapset.New(args, &heapset.Options{
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
		Verbose:   cli.Verbose,
	})

	return cli.hs
}

func (cli *Cli) ThrowAway() {
	if len(cli.File) > 0 {
		_ = cli.w.Close()
	}

	cli.hs.ThrowAway()
}

func (cmd *Hunt) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Hunt.Paths)
	defer cli.ThrowAway()

	for _, h := range hs.Get() {
		h = h
	}

	// TODO:
	// 1. carve evtx (and journals?) from files or load files
	// 2. extract important events by event id or user name
	// 3. translate event id to meaningful description
	// 4. format event to CEF format, flag important events
	// 5. order by timestamp for super timeline (unique per XXH3?)
	// 6. output CEF log

	return nil
}

func (cmd *Info) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Info.Paths)
	defer cli.ThrowAway()

	for _, h := range hs.Get() {
		if e, ok := h.Entropy(
			cli.Info.Min,
			cli.Info.Max,
		); ok {
			_, _ = fmt.Fprintf(cli.w, "%10dL %10dB  %.10f  %s\n", h.Len(), len(h.MMap()), e, h.String())
		}
	}

	return nil
}

func (cmd *Text) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Text.Paths)
	defer cli.ThrowAway()

	for _, h := range hs.Get() {
		for s := range h.Strings(
			cli.Text.Min,
			cli.Text.Max,
		) {
			if !cli.NoLine {
				_, _ = fmt.Fprintf(cli.w, "%08x  %s\n", s.Off, s.Str)
			} else {
				_, _ = fmt.Fprintf(cli.w, "%s\n", s.Str)
			}
		}
	}

	return nil
}

func (cmd *Hash) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Hash.Paths)
	defer cli.ThrowAway()

	for _, algo := range cli.Hash.Type {
		if len(cli.Hash.Type) > 1 {
			_, _ = fmt.Fprintf(cli.w, "%s\n", text.Header(strings.ToUpper(algo)))
		}

		for _, h := range hs.Get() {
			sum, err := h.HashSum(algo)

			if err != nil {
				log.Printf("could not compute hash: %s", err.Error())
				continue
			}

			switch algo {
			case types.SDHASH:
				_, _ = fmt.Fprintf(cli.w, "%s  %s\n", sum, h)
			default:
				_, _ = fmt.Fprintf(cli.w, "%x  %s\n", sum, h)
			}
		}
	}

	return nil
}

func (cmd *Dump) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Dump.Paths)
	defer cli.ThrowAway()

	var n uint

	if cli.Tail {
		n = cli.Bytes
	}

	for _, h := range hs.Get() {
		if hs.Len() > 1 && !cli.NoFile {
			_, _ = fmt.Fprintf(cli.w, "%s\n", text.Header(h.String()))
		}

		for l := range buffer.Hex(h, n).Lines {
			_, _ = fmt.Fprintf(cli.w, "%s\n", l)
		}
	}

	return nil
}

func (cmd *Show) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Show.Paths)
	defer cli.ThrowAway()

	for _, h := range hs.Get() {
		if hs.Len() > 1 && !cli.NoFile {
			_, _ = fmt.Fprintf(cli.w, "%s\n", text.Header(h.String()))
		}

		for l := range buffer.Text(h, 2).Lines {
			if !cli.NoLine && l.Nr == buffer.Sep {
				_, _ = fmt.Fprintf(cli.w, "%s\n", buffer.Sep)
			} else if !cli.NoLine {
				_, _ = fmt.Fprintf(cli.w, "%s %s\n", l.Nr, l)
			} else {
				_, _ = fmt.Fprintf(cli.w, "%s\n", l)
			}
		}
	}

	return nil
}
