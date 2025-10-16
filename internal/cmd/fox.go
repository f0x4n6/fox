package cmd

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/cmd/actions"
	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ai"
	"github.com/cuhsat/fox/internal/opt/ai/chat"
	"github.com/cuhsat/fox/internal/opt/ui"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/types/mode"
	"github.com/cuhsat/fox/internal/pkg/types/page"
	"github.com/cuhsat/fox/internal/pkg/user/bag"
	"github.com/cuhsat/fox/internal/pkg/user/config"
)

var Usage = fmt.Sprintf(fox.Ascii+`
The Swiss Army Knife for examining text files (%s)
Visit https://%s for documentation.

Usage:
  fox [ACTION] [FLAG ...] [PATH ...]

Positional arguments:
  Path(s) to open or '-' for STDIN

Actions:
  compare                  compare two files
  counts                   display line and byte counts
  deflate                  deflate compressed files
  entropy                  display file entropy
  hash                     display file hash or checksums
  strings                  display ASCII and Unicode strings
  timeline                 display super timeline

Print:
  -p, --print              print only to console
      --no-file            don't print filenames
      --no-line            don't print line numbers

Seize:
  -b, --bag                save into evidence bag

Modes:
  -x, --hex                show file in canonical hex

Deflate:
  -P, --pass=PASSWORD      password for decryption (only RAR, ZIP)

File limits:
  -h, --head               limit head of file by ...
  -t, --tail               limit tail of file by ...
  -n, --lines[=NUMBER]     number of lines (default: 10)
  -c, --bytes[=NUMBER]     number of bytes (default: 16)

Line filter:
  -e, --regexp=PATTERN     filter for lines that match pattern
  -C, --context=NUMBER     number of lines surrounding context of match
  -B, --before=NUMBER      number of lines leading context before match
  -A, --after=NUMBER       number of lines trailing context after match

AI assistant:
  -q, --query=QUERY        query for the assistant to process
  -m, --model=MODEL        model for the assistant to use
      --embed=MODEL        embedding model used for RAG

AI options:
      --num-ctx=SIZE       context window length (default: 4096)
      --temp=DECIMAL       option for temperature (default: 0.2)
      --topp=DECIMAL       option for model top_p (default: 0.5)
      --topk=NUMBER        option for model top_k (default: 10)
      --seed=NUMBER        option for random seed (default: 8211)

UI flags:
      --state={N|W|T|-}    sets the used UI state flags
      --theme=THEME        sets the used UI theme
      --space=NUMBER       sets the used indentation space (default: 2)
      --legacy             don't use any unicode decorations (ISO 8859-1)

Evidence bag:
  -N, --case=NAME          evidence bag case name (default: YYYY-MM-DD)
  -f, --file=FILE          evidence bag file name (default: evidence)
      --mode=MODE          evidence bag file mode (default: plain):
                             none, plain, text, json, jsonl, xml, sqlite

  -k, --key=PHRASE         key phrase to sign evidence bag via HMAC-SHA256

Evidence URL:
  -u, --url=SERVER         forward evidence to server address
  -a, --auth=TOKEN         forward evidence using auth token
      --ecs                use ECS schema for evidence
      --hec                use HEC schema for evidence

Turn off:
  -R, --readonly           don't write any new files
  -r, --raw                don't process files at all
      --no-convert         don't convert automatically
      --no-deflate         don't deflate automatically
      --no-plugins         don't run any plugins
      --no-mouse           don't use the mouse

Aliases:
  -L, --logstash           short for: --ecs --url=http://localhost:8080
  -S, --splunk             short for: --hec --url=http://localhost:8088/...
  -T, --text               short for: --mode=text
  -j, --json               short for: --mode=json
  -J, --jsonl              short for: --mode=jsonl
  -s, --sqlite             short for: --mode=sqlite
  -X, --xml                short for: --mode=xml

Standard:
      --help               prints this message
      --version            prints the version

Example: search for occurrences in all logs
  $ fox -pe "John Doe" ./**/*.log

Example: export the disk MBR in hex format
  $ fox -pxhc=512 image.dd > mbr

Example: analyse the given event log
  $ fox -pq="analyse this" log.evtx

Type "fox help COMMAND" for more help...
`, fox.Version, fox.Website)

var Fox = &cobra.Command{
	Use:     "fox",
	Short:   "The Swiss Army Knife for examining text files",
	Long:    "The Swiss Army Knife for examining text files",
	Args:    cobra.ArbitraryArgs,
	Version: fox.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if flags.Get().Print {
			go log()
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		// print if output is piped
		if sys.Piped(os.Stdout) {
			flg.Print = true
		}

		if flg.Filters.Context > 0 {
			flg.Filters.Before = flg.Filters.Context
			flg.Filters.After = flg.Filters.Context
		}

		if flg.Optional.Raw {
			flg.NoFile = true
			flg.NoLine = true
			flg.Optional.NoConvert = true
			flg.Optional.NoDeflate = true
			flg.Optional.NoPlugins = true
		}

		if flg.Optional.Readonly {
			flg.Optional.NoPlugins = true
			flg.Evidence.Mode = flags.BagModeNone
		}

		if len(flg.Evidence.Case) == 0 {
			flg.Evidence.Case = time.Now().Format("2006-01-02")
		}

		if flg.Alias.Text {
			flg.Evidence.Mode = flags.BagModeText
		}

		if flg.Alias.Json {
			flg.Evidence.Mode = flags.BagModeJson
		}

		if flg.Alias.Jsonl {
			flg.Evidence.Mode = flags.BagModeJsonl
		}

		if flg.Alias.Xml {
			flg.Evidence.Mode = flags.BagModeXml
		}

		if flg.Alias.Sqlite {
			flg.Evidence.Mode = flags.BagModeSqlite
		}

		if flg.Alias.Logstash {
			flg.Evidence.Url = flags.BagUrlLogstash
			flg.Evidence.Ecs = true
		}

		if flg.Alias.Splunk {
			flg.Evidence.Url = flags.BagUrlSplunk
			flg.Evidence.Hec = true
		}

		// explicit set UI mode
		if flg.Hex {
			flg.UI.Mode = mode.Hex
		}

		// implicit set UI mode
		if len(flg.Filters.Patterns) > 0 {
			flg.UI.Mode = mode.Grep
		}

		if len(flg.AI.Query) > 0 && !flg.Print {
			sys.Exit("query requires print")
		}

		if len(flg.UI.State) > 0 {
			re := regexp.MustCompile("[^-nwtNWT]+")

			flg.UI.State = re.ReplaceAllString(flg.UI.State, "")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		if !flg.Print && !flg.Bag {
			ui.Start(args, types.None)
		} else if len(args) == 0 {
			fmt.Print(Usage)
			os.Exit(2)
		} else {
			run(args)
		}
	},
	SilenceUsage: true,
}

func init() {
	flg := flags.Get()

	Fox.Flags().BoolVarP(&flg.Print, "print", "p", false, "print only to console")
	Fox.Flags().BoolVarP(&flg.NoFile, "no-file", "", false, "don't print filenames")
	Fox.Flags().BoolVarP(&flg.NoLine, "no-line", "", false, "don't print line numbers")

	Fox.Flags().BoolVarP(&flg.Bag, "bag", "b", false, "save into evidence bag")

	Fox.Flags().BoolVarP(&flg.Hex, "hex", "x", false, "show file in canonical hex")

	Fox.PersistentFlags().StringVarP(&flg.Deflate.Pass, "pass", "P", "", "password for decryption")

	Fox.Flags().BoolVarP(&flg.Limits.IsHead, "head", "h", false, "limit head of file by ...")
	Fox.Flags().BoolVarP(&flg.Limits.IsTail, "tail", "t", false, "limit tail of file by ...")
	Fox.Flags().IntVarP(&flg.Limits.Lines, "lines", "n", 0, "number of lines (default: 10)")
	Fox.Flags().IntVarP(&flg.Limits.Bytes, "bytes", "c", 0, "number of bytes (default: 16)")

	Fox.Flags().Lookup("lines").NoOptDefVal = "10"
	Fox.Flags().Lookup("bytes").NoOptDefVal = "16"

	Fox.Flags().VarP(&flg.Filters, "regexp", "e", "filter for lines that match pattern")
	Fox.Flags().IntVarP(&flg.Filters.Context, "context", "C", 0, "number of lines surrounding context of match")
	Fox.Flags().IntVarP(&flg.Filters.Before, "before", "B", 0, "number of lines leading context before match")
	Fox.Flags().IntVarP(&flg.Filters.After, "after", "A", 0, "number of lines trailing context after match")

	Fox.Flags().StringVarP(&flg.AI.Query, "query", "q", "", "query for the assistant to process")
	Fox.Flags().StringVarP(&flg.AI.Model, "model", "m", "", "model for the assistant to use")
	Fox.Flags().StringVarP(&flg.AI.Embed, "embed", "", "", "embedding model used for RAG")
	Fox.Flags().IntVarP(&flg.AI.NumCtx, "num-ctx", "", 4096, "context window length")
	Fox.Flags().Float64VarP(&flg.AI.Temp, "temp", "", 0.2, "option for temperature")
	Fox.Flags().Float64VarP(&flg.AI.TopP, "topp", "", 0.5, "option for model top_p")
	Fox.Flags().IntVarP(&flg.AI.TopK, "topk", "", 10, "option for model top_k")
	Fox.Flags().IntVarP(&flg.AI.Seed, "seed", "", 8211, "option for random seed")

	Fox.Flags().StringVarP(&flg.UI.State, "state", "", "", "sets the used UI state flags")
	Fox.Flags().StringVarP(&flg.UI.Theme, "theme", "", themes.Default, "sets the used UI theme")
	Fox.Flags().IntVarP(&flg.UI.Space, "space", "", 2, "sets the used indentation space")
	Fox.Flags().BoolVarP(&flg.UI.Legacy, "legacy", "", false, "don't use any unicode decorations")

	Fox.Flags().StringVarP(&flg.Evidence.Case, "case", "N", "", "evidence bag case name")
	Fox.Flags().StringVarP(&flg.Evidence.File, "file", "f", flags.BagFile, "evidence bag file name")
	Fox.Flags().VarP(&flg.Evidence.Mode, "mode", "", "evidence bag file mode")
	Fox.Flags().StringVarP(&flg.Evidence.Key, "key", "k", "", "key phrase to sign evidence bag via HMAC-SHA256")
	Fox.Flags().StringVarP(&flg.Evidence.Url, "url", "u", "", "forward evidence to server address")
	Fox.Flags().StringVarP(&flg.Evidence.Auth, "auth", "a", "", "forward evidence using auth token")
	Fox.Flags().BoolVarP(&flg.Evidence.Ecs, "ecs", "", false, "use ECS schema for evidence")
	Fox.Flags().BoolVarP(&flg.Evidence.Hec, "hec", "", false, "use HEC schema for evidence")

	Fox.Flags().Lookup("mode").NoOptDefVal = string(flags.BagModeText)

	Fox.Flags().BoolVarP(&flg.Optional.Raw, "raw", "r", false, "don't process files at all")
	Fox.Flags().BoolVarP(&flg.Optional.Readonly, "readonly", "R", false, "don't write any new files")
	Fox.Flags().BoolVarP(&flg.Optional.NoConvert, "no-convert", "", false, "don't convert automatically")
	Fox.Flags().BoolVarP(&flg.Optional.NoDeflate, "no-deflate", "", false, "don't deflate automatically")
	Fox.Flags().BoolVarP(&flg.Optional.NoPlugins, "no-plugins", "", false, "don't run any plugins")
	Fox.Flags().BoolVarP(&flg.Optional.NoMouse, "no-mouse", "", false, "don't use the mouse")

	Fox.Flags().BoolVarP(&flg.Alias.Logstash, "logstash", "L", false, "short for: --ecs --url=http://localhost:8080")
	Fox.Flags().BoolVarP(&flg.Alias.Splunk, "splunk", "S", false, "short for: --hec --url=http://localhost:8088/...")
	Fox.Flags().BoolVarP(&flg.Alias.Text, "text", "T", false, "short for: --mode=text")
	Fox.Flags().BoolVarP(&flg.Alias.Json, "json", "j", false, "short for: --mode=json")
	Fox.Flags().BoolVarP(&flg.Alias.Jsonl, "jsonl", "J", false, "short for: --mode=jsonl")
	Fox.Flags().BoolVarP(&flg.Alias.Sqlite, "sqlite", "s", false, "short for: --mode=sqlite")
	Fox.Flags().BoolVarP(&flg.Alias.Xml, "xml", "X", false, "short for: --mode=xml")

	Fox.PersistentFlags().BoolP("version", "", false, "prints the version")
	Fox.PersistentFlags().BoolP("help", "", false, "prints this message")

	Fox.MarkFlagsRequiredTogether("hec", "auth")

	Fox.MarkFlagsMutuallyExclusive("head", "tail")
	Fox.MarkFlagsMutuallyExclusive("ecs", "hec")

	Fox.SetErrPrefix(sys.Prefix)
	Fox.SetHelpTemplate(Usage)
	Fox.SetVersionTemplate(fmt.Sprintf("%s %s\n", fox.Product, fox.Version))

	Fox.AddCommand(actions.Compare)
	Fox.AddCommand(actions.Counts)
	Fox.AddCommand(actions.Deflate)
	Fox.AddCommand(actions.Entropy)
	Fox.AddCommand(actions.Hash)
	Fox.AddCommand(actions.Strings)
	Fox.AddCommand(actions.Timeline)

	Fox.Flags()

	config.Load(Fox.Flags())

	cobra.MousetrapHelpText = "" // disable
}

func run(args []string) {
	var ctx = opt.NewState(nil)
	var flg = flags.Get()
	var b *bag.Bag

	if len(flg.AI.Query) > 0 && !ai.Check() {
		sys.Exit("Assistant is not available")
	}

	if flg.Bag {
		b = bag.New()
	}

	hs := heapset.New(args)
	defer hs.ThrowAway()

	hs.Range(func(_ int, h *heap.Heap) bool {
		if h.Type != types.Stdin {
			buf := page.NewContext(h)

			if hs.Len() > 1 && !flg.NoFile {
				fmt.Println(text.Block(h.String(), page.TermW))
			}

			if flg.Bag {
				b.Put(h)
			} else if len(flg.AI.Query) > 0 {
				c := chat.New(ctx, h)
				defer c.Close()

				c.Query(flg.AI.Query, false)
			} else if flg.Hex {
				buf.W = page.TermW

				for l := range page.Hex(buf).Lines {
					fmt.Println(l)
				}
			} else {
				if buf.Heap.Len() == 0 {
					return true // ignore empty files
				}

				for l := range page.Text(buf).Lines {
					if l.Nr == "--" {
						if !flg.NoLine {
							fmt.Println("--")
						}
					} else {
						if !flg.NoLine {
							fmt.Printf("%s %s\n", l.Nr, l)
						} else {
							fmt.Println(l)
						}
					}
				}
			}
		}
		return true
	})
}

func log() {
	for {
		_, _ = fmt.Fprintf(os.Stderr, sys.Prefix+" %s\n", <-sys.Logs)
	}
}
