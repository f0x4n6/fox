package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/v4/internal"
	"github.com/cuhsat/fox/v4/internal/opt/ai"
	"github.com/cuhsat/fox/v4/internal/opt/ai/chat"
	"github.com/cuhsat/fox/v4/internal/pkg/flags"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/v4/internal/pkg/types/page"
	"github.com/cuhsat/fox/v4/internal/pkg/user/bag"
	"github.com/cuhsat/fox/v4/internal/pkg/user/config"
)

var Usage = fmt.Sprintf(
	` ___ ___  ___ ___ _  _ ___ ___ ___   _____  __  _   __  __ ___ _  _ ___ ___
| __/ _ \| _ \ __| \| / __|_ _/ __| | __\ \/ / /_\ |  \/  |_ _| \| | __| _ \
| _| (_) |   / _||  ' \__ \| | (__  | _| >  < / _ \| |\/| || ||  ' | _||   /
|_| \___/|_|_\___|_|\_|___/___\___| |___/_/\_\_/ \_\_|  |_|___|_|\_|___|_|_\

The Swiss Army Knife for examining text files (%s)
Visit <https://%s>.

Usage:
  fox [FLAG ...] [PATH ...]

Positional arguments:
  Path(s) to open or '-' for STDIN

Flags:
  -b, --bag                save into evidence bag
  -x, --hex                show file in canonical hex
      --no-file            don't print filenames
      --no-line            don't print line numbers

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

Evidence bag:
  -N, --case=NAME          evidence bag case name (default: YYYY-MM-DD)
  -F, --file=FILE          evidence bag file name (default: evidence)
      --mode=MODE          evidence bag file mode (default: text):
                             none, text, json, jsonl, sqlite

Evidence sign:
  -s, --sign=PHRASE        key phrase to sign evidence bag via HMAC-SHA256

Evidence URL:
  -u, --url=SERVER         forward evidence to server address
  -a, --auth=TOKEN         forward evidence using auth token
      --ecs                use ECS schema for evidence
      --hec                use HEC schema for evidence

Deflate:
      --pass=PASSWORD      password for decryption (only RAR, ZIP)

Strings:

Entropy:

Hashsum:

Counts:

Turn off:
  -R, --readonly           don't write any new files
  -r, --raw                don't process files at all
      --no-convert         don't convert automatically
      --no-deflate         don't deflate automatically
      --no-plugins         don't run any plugins

Aliases:
  -L, --logstash           short for: --ecs --url=http://localhost:8080
  -S, --splunk             short for: --hec --url=http://localhost:8088/...
  -T, --text               short for: --mode=text
  -j, --json               short for: --mode=json
  -J, --jsonl              short for: --mode=jsonl
  -Q, --sqlite             short for: --mode=sqlite

Standard:
      --help               prints this message
      --version            prints the version

Example: search for occurrences in all logs
  $ fox -pe "login" ./**/*.log

Example: export the disk MBR in hex format
  $ fox -pxhc=512 image.dd > mbr

Example: analyse the given event log
  $ fox -pq="analyse this" log.evtx

Type "man fox" for more help...
`, fox.Version, fox.Website)

var Fox = &cobra.Command{
	Use:     "fox",
	Short:   "The Swiss Army Knife for examining text files",
	Long:    "The Swiss Army Knife for examining text files",
	Args:    cobra.ArbitraryArgs,
	Version: fox.Version,
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

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

		if len(flg.AI.Query) > 0 {
			sys.Exit("query requires print")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
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

	Fox.Flags().BoolVarP(&flg.Bag, "bag", "b", false, "save into evidence bag")
	Fox.Flags().BoolVarP(&flg.Hex, "hex", "x", false, "show file in canonical hex")

	Fox.Flags().BoolVar(&flg.NoFile, "no-file", false, "don't print filenames")
	Fox.Flags().BoolVar(&flg.NoLine, "no-line", false, "don't print line numbers")

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
	Fox.Flags().StringVar(&flg.AI.Embed, "embed", "", "embedding model used for RAG")
	Fox.Flags().IntVar(&flg.AI.NumCtx, "num-ctx", 4096, "context window length")
	Fox.Flags().Float64Var(&flg.AI.Temp, "temp", 0.2, "option for temperature")
	Fox.Flags().Float64Var(&flg.AI.TopP, "topp", 0.5, "option for model top_p")
	Fox.Flags().IntVar(&flg.AI.TopK, "topk", 10, "option for model top_k")
	Fox.Flags().IntVar(&flg.AI.Seed, "seed", 8211, "option for random seed")

	Fox.Flags().StringVarP(&flg.Evidence.Case, "case", "N", "", "evidence bag case name")
	Fox.Flags().StringVarP(&flg.Evidence.File, "file", "F", flags.BagFile, "evidence bag file name")
	Fox.Flags().Var(&flg.Evidence.Mode, "mode", "evidence bag file mode")
	Fox.Flags().StringVarP(&flg.Evidence.Sign, "sign", "s", "", "key phrase to sign evidence bag via HMAC-SHA256")
	Fox.Flags().StringVarP(&flg.Evidence.Url, "url", "u", "", "forward evidence to server address")
	Fox.Flags().StringVarP(&flg.Evidence.Auth, "auth", "a", "", "forward evidence using auth token")
	Fox.Flags().BoolVar(&flg.Evidence.Ecs, "ecs", false, "use ECS schema for evidence")
	Fox.Flags().BoolVar(&flg.Evidence.Hec, "hec", false, "use HEC schema for evidence")

	Fox.Flags().Lookup("mode").NoOptDefVal = string(flags.BagModeText)

	Fox.Flags().StringVar(&flg.Deflate.Pass, "pass", "", "password for decryption")

	///
	//
	//Fox.Flags().Float64Var(&flg.Entropy.Min, "min", 0.0, "minimum score")
	//Fox.Flags().Float64Var(&flg.Entropy.Max, "max", 1.0, "maximum score")
	//Fox.Flags().Lookup("min").NoOptDefVal = "0.8"
	//Fox.Flags().Lookup("max").NoOptDefVal = "0.8"
	//
	//Fox.Flags().IntVar(&flg.Strings.Min, "min", 3, "minimum length")
	//Fox.Flags().IntVar(&flg.Strings.Max, "max", math.MaxInt, "maximum length")
	//Fox.Flags().BoolVar(&flg.Strings.Class, "class", false, "run built-in classification")
	//
	//Fox.Flags().Var(&flg.Hash.Algos, "type", "hash algorithm")
	//
	//Cryptographic hash algorithms:
	//MD5, SHA1, SHA256, SHA3
	//
	//Sha3 hash algorithms:
	//SHA3-224, SHA3-256, SHA3-384, SHA3-512
	//
	//Blake3 hash algorithms:
	//BLAKE3-256, BLAKE3-512
	//
	//Performance hash algorithms:
	//FNV-1, FNV-1A, XXH64, XXH3
	//
	//Similarity hash algorithms:
	//SDHASH, SSDEEP, TLSH
	//
	//Checksum algorithms:
	//CRC32-IEEE, CRC64-ECMA, CRC64-ISO

	///

	Fox.Flags().BoolVarP(&flg.Optional.Raw, "raw", "r", false, "don't process files at all")
	Fox.Flags().BoolVarP(&flg.Optional.Readonly, "readonly", "R", false, "don't write any new files")
	Fox.Flags().BoolVar(&flg.Optional.NoConvert, "no-convert", false, "don't convert automatically")
	Fox.Flags().BoolVar(&flg.Optional.NoDeflate, "no-deflate", false, "don't deflate automatically")
	Fox.Flags().BoolVar(&flg.Optional.NoPlugins, "no-plugins", false, "don't run any plugins")

	Fox.Flags().BoolVarP(&flg.Alias.Logstash, "logstash", "L", false, "short for: --ecs --url=http://localhost:8080")
	Fox.Flags().BoolVarP(&flg.Alias.Splunk, "splunk", "S", false, "short for: --hec --url=http://localhost:8088/...")
	Fox.Flags().BoolVarP(&flg.Alias.Text, "text", "T", false, "short for: --mode=text")
	Fox.Flags().BoolVarP(&flg.Alias.Json, "json", "j", false, "short for: --mode=json")
	Fox.Flags().BoolVarP(&flg.Alias.Jsonl, "jsonl", "J", false, "short for: --mode=jsonl")
	Fox.Flags().BoolVarP(&flg.Alias.Sqlite, "sqlite", "Q", false, "short for: --mode=sqlite")

	Fox.Flags().Bool("version", false, "prints the version")
	Fox.Flags().Bool("help", false, "prints this message")

	Fox.MarkFlagsRequiredTogether("hec", "auth")

	Fox.MarkFlagsMutuallyExclusive("head", "tail")
	Fox.MarkFlagsMutuallyExclusive("ecs", "hec")

	Fox.SetErrPrefix(sys.Prefix)
	Fox.SetHelpTemplate(Usage)
	Fox.SetVersionTemplate(fmt.Sprintf("%s %s\n", fox.Product, fox.Version))

	Fox.CompletionOptions.HiddenDefaultCmd = true

	config.Load(Fox.Flags())
	cobra.MousetrapHelpText = "Usage: fox [FLAG ...] [PATH ...]"
}

func run(args []string) {
	var flg = flags.Get()
	var b *bag.Bag

	log.SetPrefix(sys.Prefix)

	if len(flg.AI.Query) > 0 && !ai.Check() {
		sys.Exit("assistant is not available")
	}

	if flg.Bag {
		b = bag.New()
	}

	hs := heapset.New(args)
	defer hs.ThrowAway()

	hs.Range(func(_ int, h *heap.Heap) bool {
		if h.Type != types.Stdin {
			if hs.Len() > 1 && !flg.NoFile {
				fmt.Println(text.Block(h.String(), page.TermW))
			}

			///

			//fmt.Printf("%8dL %8dB  %s\n", h.Length(), len(*h.MMap()), h.String())
			//
			//if v := h.Entropy(
			//	flg.Entropy.Min,
			//	flg.Entropy.Max,
			//); v != -1 {
			//	fmt.Printf("%.10f  %s\n", v, h.String())
			//}
			//
			//fmt.Print(text.Diff(
			//	a[0].String(),
			//	a[1].String(),
			//	a[0].SMap().Lines(),
			//	a[1].SMap().Lines(),
			//	false,
			//))
			//
			//hs.Unique().CloseOther()
			//
			//for l := range page.Text(hs.LoadHeap(), 2).Lines {
			//	fmt.Println(l)
			//}
			//
			//if !flg.NoFile {
			//	fmt.Println(text.Block(h.String(), page.TermW))
			//}
			//
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

			///

			if flg.Bag {
				b.Put(h)
			} else if len(flg.AI.Query) > 0 {
				c := chat.New(h)
				defer c.Close()

				c.Query(flg.AI.Query)
			} else if flg.Hex {
				for l := range page.Hex(h).Lines {
					fmt.Println(l)
				}
			} else {
				if h.Size() == 0 {
					return true // ignore empty files
				}

				for l := range page.Text(h, 2).Lines {
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
