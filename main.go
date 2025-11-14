// The Swiss Army Knife for examining text files.
//
// Copyright 2025 Christian Uhsat. All rights reserved.
// Use of this source code is governed by the GPL-3.0
// license that can be found in the LICENSE.md file.
//
// For more information, please consult:
//
//	https://forensic-examiner.eu
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alecthomas/kong"

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
)

var Usage = fox.Banner + `
The Swiss Army Knife for examining text files (%s)
Visit <https://%s>.

Usage:
  fox [ACTION] [FLAG ...] PATH ...

Positional arguments:
  Path(s) to open or '-' for STDIN

File infos:
  -x, --hex                show file in canonical hex
  -w, --count              show file counts
  -a, --hash=ALGO[,ALGO]   show file hashes
  -y, --entropy[=MIN:MAX]  show file entropy
  -s, --strings[=MIN:MAX]  show carved strings (only ASCII)

File limits:
  -h, --head               limit head of file by ...
  -t, --tail               limit tail of file by ...
  -n, --lines[=NUMBER]     number of lines (default: 10)
  -c, --bytes[=NUMBER]     number of bytes (default: 16)

File loader:
  -p, --pass=PASSWORD      password for decryption (only RAR, ZIP)

Line filter:
  -e, --regexp=PATTERN     filter for lines that match pattern
  -C, --context=NUMBER     number of lines surrounding context of match
  -B, --before=NUMBER      number of lines leading context before match
  -A, --after=NUMBER       number of lines trailing context after match

LLM parser:
  -q, --query=QUERY        assistant query to process
  -m, --model=MODEL        assistant model (https://ollama.com/library)
      --embed=MODEL        embedding model (https://ollama.com/library)

LLM options:
      --num-ctx=SIZE       context window length (default: 4096)
      --temp=DECIMAL       option for temperature (default: 0.2)
      --topp=DECIMAL       option for model top_p (default: 0.5)
      --topk=NUMBER        option for model top_k (default: 10)
      --seed=NUMBER        option for random seed (default: 8211)

Evidence bag:
  -N, --case=NAME          evidence bag case name (default: YYYY-MM-DD)
  -F, --file=FILE          evidence bag file name (default: evidence)
      --mode=MODE          evidence bag file mode (default: text)

Evidence sign:
      --sign=PHRASE        key phrase to sign evidence bag via HMAC-SHA256

Evidence URL:
  -u, --url=SERVER         forward evidence to server address
      --auth=TOKEN         forward evidence using auth token
      --ecs                use ECS schema for evidence
      --hec                use HEC schema for evidence

Turn off:
  -r, --raw                don't process files at all
  -R, --readonly           don't write anything at all
      --no-file            don't print filenames
      --no-line            don't print line numbers
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

Evidence bag modes:
  Text, JSON, JSONL, SQLite

Hashes (cryptographic):
  MD5, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512,
  BLAKE3-256, BLAKE3-512

Hashes (performance):
  FNV-1, FNV-1A, XXH64, XXH3

Hashes (similarity):
  SDHASH, SSDEEP, TLSH

Checksums:
  CRC32-IEEE, CRC64-ECMA, CRC64-ISO

Example: search for occurrences in all logs
  $ fox -e "login" ./**/*.log

Example: export the disk MBR in hex format
  $ fox -xhc=512 image.dd > mbr

Example: analyse the given event log
  $ fox -q="analyse this" log.evtx

Type "man fox" for more help...
`

// Main start and catch.
func main() {
	defer sys.Recover()

	log.SetPrefix(sys.Prefix)

	cli := flags.CLI

	ctx := kong.Parse(&cli,
		kong.NoDefaultHelp(),
		kong.UsageOnError(),
	)

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

	if cli.Readonly {
		cli.Mode = types.NONE
	}

	if len(cli.Case) == 0 {
		cli.Case = time.Now().Format("2006-01-02")
	}

	if cli.Logstash {
		cli.Url = types.LOGSTASH
		cli.Ecs = true
	}

	if cli.Splunk {
		cli.Url = types.SPLUNK
		cli.Hec = true
	}

	if cli.Text {
		cli.Mode = types.TEXT
	}

	if cli.Json {
		cli.Mode = types.JSON
	}

	if cli.Jsonl {
		cli.Mode = types.JSONL
	}

	if cli.Sqlite {
		cli.Mode = types.SQLITE
	}

	switch {
	case cli.Version:
		fmt.Printf("%s %s\n", fox.Product, fox.Version)
	case ctx.Error != nil:
		fallthrough
	default:
		fmt.Printf(Usage, fox.Version, fox.Website)
	}

	if len(cli.Query) > 0 && !ai.Check() {
		sys.Exit("assistant is not available")
	}

	hs := heapset.New(ctx.Args)
	defer hs.ThrowAway()

	hs.Range(func(_ int, h *heap.Heap) bool {
		if h.Type != types.Stdin {
			if hs.Len() > 1 && !cli.NoFile {
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

			if len(cli.Query) > 0 {
				c := chat.New(h)
				defer c.Close()

				c.Query(cli.Query)
			} else if cli.Hex {
				for l := range page.Hex(h).Lines {
					fmt.Println(l)
				}
			} else {
				if h.Size() == 0 {
					return true // ignore empty files
				}

				for l := range page.Text(h, 2).Lines {
					if l.Nr == "--" {
						if !cli.NoLine {
							fmt.Println("--")
						}
					} else {
						if !cli.NoLine {
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
