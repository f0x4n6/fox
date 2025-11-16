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

	"github.com/alecthomas/kong"
	"github.com/cuhsat/fox/v4/internal"
	"github.com/cuhsat/fox/v4/internal/pkg/run"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
)

var Usage = fox.Banner + `
The Swiss Army Knife for examining text files (%s)
Visit <https://%s>.

Usage:
  fox [COMMAND] [FLAG ...] [PATH ...]

Positional arguments:
  Path(s) to open or '-' for STDIN

Commands:
  hunt                     hunt down suspicious files
  hash                     show file content hashes
  stat                     show file content stats
  text                     show file ASCII strings
  dump                     show file in canonical hex
  show                     show file content (default)

Hash flags:
  -a, --type=ALGO[,ALGO]   use algorithm

Stat flags:
      --min=DECIMAL        minimum entropy value
      --max=DECIMAL        maximal entropy value

Text flags:
      --min=NUMBER         minimum string length
      --max=NUMBER         maximal string length

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

Evidence bag:
  -f, --file=FILE          evidence bag file name (default: YYYY-MM-DD)
  -m, --mode=MODE          evidence bag file mode (default: text)

Evidence sign:
  -s, --sign=PHRASE        key phrase to sign evidence bag via HMAC-SHA256

Evidence URL:
  -u, --url=SERVER         forward evidence to server address
  -a, --auth=TOKEN         forward evidence using auth token
      --ecs                use ECS schema for evidence
      --hec                use HEC schema for evidence

Turn off:
  -r, --raw                don't process files at all
  -R, --readonly           don't write anything at all
      --no-file            don't print filenames
      --no-line            don't print line numbers
      --no-deflate         don't deflate automatically
      --no-convert         don't convert automatically

Aliases:
  -L, --logstash           short for: --ecs --url=http://localhost:8080
  -S, --splunk             short for: --hec --url=http://localhost:8088/...
  -Q, --sqlite             short for: --mode=sqlite
  -J, --jsonl              short for: --mode=jsonl
  -j, --json               short for: --mode=json

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

Example: dump the image MBR in hex format
  $ fox dump -hc=512 image.dd > mbr

Example: find occurrences in all logs
  $ fox -e "login" ./**/*.log

Example: hunt down suspicious files
  $ fox hunt .

Type "man fox" for more help...
`

type Cli struct {
	run.Globals
	Help    bool
	Version bool
}

// Main start and catch.
func main() {
	defer sys.Recover()

	log.SetPrefix(sys.Prefix)

	cli := new(Cli)

	ctx := kong.Parse(cli,
		kong.NoDefaultHelp(),
		kong.UsageOnError(),
	)

	err := ctx.Run(&cli.Globals)

	if err != nil {
		fmt.Printf(Usage, fox.Version, fox.Website)
	}

	ctx.FatalIfErrorf(err)

	sys.Exit("x")

	switch {
	case cli.Version:
		fmt.Printf("%s %s\n", fox.Product, fox.Version)
	case ctx.Error != nil:
		fallthrough
	default:
		fmt.Printf(Usage, fox.Version, fox.Website)
	}
}
