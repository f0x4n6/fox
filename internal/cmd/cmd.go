package cmd

import (
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/cuhsat/fox/v4/internal/pkg/files/stream"
	"github.com/cuhsat/fox/v4/internal/pkg/files/stream/ecs"
	"github.com/cuhsat/fox/v4/internal/pkg/files/stream/hec"
	"github.com/cuhsat/fox/v4/internal/pkg/files/stream/raw"
	"github.com/cuhsat/fox/v4/internal/pkg/hash"
	"github.com/cuhsat/fox/v4/internal/pkg/hunt"
	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/buffer"
	"github.com/cuhsat/fox/v4/internal/pkg/types/event"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heapset"
)

type Hunt struct {
	All   bool     `short:"a"`
	Ext   int      `short:"x" type:"counter"`
	Sort  bool     `short:"s"`
	Json  bool     `short:"j" xor:"json,jsonl"`
	Jsonl bool     `short:"J" xor:"json,jsonl"`
	Paths []string `arg:"" type:"path" optional:""`
}

type Info struct {
	Min   float64  `default:"0.0"`
	Max   float64  `default:"1.0"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

type Text struct {
	Min   uint     `default:"3"`
	Max   uint     `default:"256"`
	Paths []string `arg:"" type:"path" optional:""`
}

type Hash struct {
	Algo struct {
		Algo  string `arg:""`
		Paths struct {
			Paths []string `arg:"" type:"path" optional:""`
		} `arg:""`
	} `arg:""`

	// internal
	algos []string `kong:"-"`
}

type Hex struct {
	Mode  string   `short:"m" enum:"c,hd,od,xxd" default:"c"`
	Paths []string `arg:"" type:"path"`
}

type Cat struct {
	Paths []string `arg:"" type:"path" optional:""`
}

type Cli struct {
	// commands
	Hunt Hunt `cmd:"" aliases:"u"`
	Info Info `cmd:"" aliases:"i,wc"`
	Text Text `cmd:"" aliases:"t,strings"`
	Hash Hash `cmd:"" aliases:"h"`
	Hex  Hex  `cmd:"" aliases:"x"`
	Cat  Cat  `cmd:"" default:"withargs" aliases:"c,less"`

	// file limits
	Head  bool `short:"h" xor:"head,tail"`
	Tail  bool `short:"t" xor:"head,tail"`
	Lines uint `short:"n" xor:"lines,bytes"`
	Bytes uint `short:"c" xor:"lines,bytes"`

	// file loader
	Pass string `short:"p"`

	// line filter
	Regex   string `short:"e"`
	Context uint   `short:"C"`
	Before  uint   `short:"B"`
	After   uint   `short:"A"`

	// data stream
	File string `short:"f"`
	Url  string `short:"u"`
	Auth string `short:"T"`
	Ecs  bool   `short:"E" xor:"ecs,hec"`
	Hec  bool   `short:"H" xor:"ecs,hec" and:"hec,auth"`

	// disable
	Readonly  bool `short:"R"`
	Raw       bool `short:"r"`
	NoFile    bool `long:"no-file"`
	NoLine    bool `long:"no-line"`
	NoDeflate bool `long:"no-deflate"`
	NoConvert bool `long:"no-convert"`

	// aliases
	Logstash bool `short:"L" xor:"logstash,splunk"`
	Splunk   bool `short:"S" xor:"logstash,splunk"`

	// standard
	DryRun  bool `short:"d" long:"dry-run"`
	Verbose int  `short:"v" type:"counter"`

	// internal
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

	if len(cli.Hunt.Paths) == 0 {
		cli.Hunt.Paths = hunt.Paths
	}

	if len(cli.Hash.Algo.Algo) > 0 {
		cli.Hash.algos = strings.Split(cli.Hash.Algo.Algo, ",")
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

	if cli.DryRun {
		for _, h := range cli.hs.Get() {
			_, _ = fmt.Fprintf(cli.w, "%s\n", h.Name)
		}

		// exit early
		cli.hs.ThrowAway()
		os.Exit(0)
	}

	return cli.hs
}

func (cli *Cli) ThrowAway() {
	if len(cli.File) > 0 {
		_ = cli.w.Close()
	}

	cli.hs.ThrowAway()
}

func (cmd *Hunt) Run(cli *Cli) error {
	cli.NoConvert = true // force

	hs := cli.Bootstrap(cli.Hunt.Paths)
	defer cli.ThrowAway()

	n, sort := 0, func(in <-chan *event.Event) <-chan *event.Event {
		out := make(chan *event.Event)

		go func() {
			defer close(out)
			cache := make(map[int64]*event.Event)

			for e := range in {
				cache[e.Time.UnixNano()] = e
			}

			for _, k := range slices.Sorted(maps.Keys(cache)) {
				out <- cache[k]
			}
		}()

		return out
	}

	for _, h := range hs.Get() {
		ch := hunt.Hunt(h, &hunt.Options{
			Extensions: cli.Hunt.Ext,
			Verbose:    cli.Verbose,
		})

		if cli.Hunt.Sort {
			ch = sort(ch)
		}

		for e := range ch {
			if cli.Hunt.All || e.Severity >= hunt.Level {
				switch {
				case cli.Hunt.Jsonl:
					_, _ = fmt.Fprintln(cli.w, e.ToJSONL())
				case cli.Hunt.Json:
					_, _ = fmt.Fprintln(cli.w, e.ToJSON())
				default:
					_, _ = fmt.Fprintln(cli.w, e.ToCEF())
				}
				n++
			}
		}
	}

	if cli.Verbose > 1 {
		log.Printf("found %d events\n", n)
	}

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
	hs := cli.Bootstrap(cli.Hash.Algo.Paths.Paths)
	defer cli.ThrowAway()

	for _, a := range cli.Hash.algos {
		if len(cli.Hash.algos) > 1 && !cli.NoFile {
			_, _ = fmt.Fprintf(cli.w, "%s\n", text.Header(strings.ToUpper(a)))
		}

		for _, h := range hs.Get() {
			sum, err := hash.Sum(a, h.MMap())

			if err != nil {
				log.Println("could not compute hash: ", err.Error())
				continue
			}

			switch a {
			case types.SDHASH:
				_, _ = fmt.Fprintf(cli.w, "%s  %s\n", sum, h)
			default:
				_, _ = fmt.Fprintf(cli.w, "%x  %s\n", sum, h)
			}
		}
	}

	return nil
}

func (cmd *Hex) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Hex.Paths)
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

func (cmd *Cat) Run(cli *Cli) error {
	hs := cli.Bootstrap(cli.Cat.Paths)
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
