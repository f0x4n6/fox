package run

import (
	"log"
	"time"

	"github.com/alecthomas/kong"
	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/v4/internal/pkg/types/loader"
	"github.com/cuhsat/fox/v4/internal/pkg/types/page"
)

type Globals struct {
	// Commands
	Hunt Hunt `cmd:"" aliases:"u"`
	Stat Stat `cmd:"" aliases:"c"`
	Text Text `cmd:"" aliases:"s"`
	Hash Hash `cmd:"" aliases:"h"`
	Dump Dump `cmd:"" aliases:"x"`
	Show Show `cmd:"" default:"withargs" aliases:"p"`

	// File limits
	Head  bool `short:"h" xor:"head,tail"`
	Tail  bool `short:"t" xor:"head,tail"`
	Lines int  `short:"n" default:"10"`
	Bytes int  `short:"c" default:"16"`

	// File loader
	Pass string `short:"p"`

	// Line filter
	Regex   []string `short:"e"`
	Context int      `short:"C"`
	Before  int      `short:"B"`
	After   int      `short:"A"`

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

	// Aliases
	Logstash bool `short:"L"`
	Splunk   bool `short:"S"`
	Sqlite   bool `short:"Q"`
	Jsonl    bool `short:"J"`
	Json     bool `short:"j"`
}

func (cli *Globals) Muster() {
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
}

type Hunt struct {
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

func (cmd *Hunt) Run() error {
	// TODO
	return nil
}

type Stat struct {
	Min   float32  `default:"0.0"`
	Max   float32  `default:"1.0"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

func (cmd *Stat) Run(ctx *kong.Context) error {
	//if v := h.Entropy(
	//	flg.Entropy.Min,
	//	flg.Entropy.Max,
	//); v != -1 {
	//	fmt.Printf("%8dL %8dB %.10f  %s\n", h.Length(), len(*h.MMap()), v, h.String())
	//}
	return nil
}

type Text struct {
	Min   int      `default:"3"`
	Max   int      `default:"-1"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

func (cmd *Text) Run(ctx *kong.Context) error {
	//for s := range h.Strings(
	//	flg.Strings.Min,
	//	flg.Strings.Max,
	//	flg.Strings.Class,
	//	flg.Strings.Re,
	//) {
	//	if !flg.NoLine {
	//		fmt.Printf("%08x  %s\n", s.Off, strings.TrimSpace(s.Str))
	//	} else {
	//		fmt.Println(strings.TrimSpace(s.Str))
	//	}
	//}
	return nil
}

type Hash struct {
	Type  string   `short:"a" sep:"," default:"SHA256"`
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

func (cmd *Hash) Run(ctx *kong.Context) error {
	//if !flg.NoFile {
	//	fmt.Println(text.Block(h.String(), page.TermW))
	//}
	//
	//for _, algo := range algos {
	//	if len(algos) > 1 {
	//		fmt.Println(text.Block(strings.ToUpper(algo), page.TermW))
	//	}
	//
	//	hs.Range(func(_ int, h *heap.Heap) bool {
	//		sum, err := h.HashSum(algo)
	//
	//		if err != nil {
	//			sys.Exit(fmt.Sprintf("could not compute hash: %s", err.Error()))
	//			return false
	//		}
	//
	//		switch algo {
	//		case types.SDHASH:
	//			fmt.Printf("%s  %s\n", sum, h.String())
	//		default:
	//			fmt.Printf("%x  %s\n", sum, h.String())
	//		}
	//		return true
	//	})
	//}
	return nil
}

type Dump struct {
	Paths []string `arg:"" name:"path" type:"path"`
}

func (cmd *Dump) Run(ctx *kong.Context) error {
	//				for l := range page.Hex(h).Lines {
	//					fmt.Println(l)
	//				}
	return nil
}

type Show struct {
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}

func (cmd *Show) Run(ctx *kong.Context, cli *Globals) error {
	hs := heapset.New(ctx.Args, &loader.Options{
		Password:  cli.Pass,
		NoDeflate: cli.NoDeflate,
		NoConvert: cli.NoConvert,
	})

	defer hs.ThrowAway()

	for _, h := range hs.Get() {
		if h.Type == types.Stdin {
			continue // TODO
		}

		if hs.Len() > 1 && !cli.NoFile {
			log.Println(text.Block(h.String(), page.TermW))
		}

		if h.Size() == 0 {
			continue // ignore empty files
		}

		for l := range page.Text(h, 2).Lines {
			switch l.Nr {
			case "--":
				if !cli.NoLine {
					log.Println("--")
				}
			default:
				if !cli.NoLine {
					log.Printf("%s %s\n", l.Nr, l)
				} else {
					log.Println(l)
				}
			}
		}
	}

	return nil
}
