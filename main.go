// The Swiss Army Knife for examining text files.
//
// Copyright 2025 Christian Uhsat. All rights reserved.
// Use of this source code is governed by the GPL-3.0
// license that can be found in the LICENSE.md file.
//
// For more information, please consult:
//
//	https://foxhunt.dev
package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/alecthomas/kong"

	"github.com/cuhsat/fox/v4/internal"
	"github.com/cuhsat/fox/v4/internal/cmd"
)

var usage = ` ____ _____  __  _  _ _   _ _  _ _____
|  __/ _ \ \/ / | || | | | | \| |_   _|
|  _| (_) >  <  | __ | |_| | .' | | |
|__| \___/_/\_\ |_||_|\___/|_|\_| |_|

The Swiss Army Knife for examining text files (%s)
Visit <https://%s>.

Usage:
  fox [COMMAND] [ARG] [FLAG ...] <PATH> ...

Positional arguments:
  Path(s) to open or '-' for STDIN

Commands:
  HUNT     hunt suspicious activities
  HASH     prints file content hashes
  INFO     prints file content infos
  TEXT     prints file ASCII strings
  HEX      prints file in hex format
  CAT      prints file (default)

Hunt flags:
  -a, --all                show logs with lower severity
  -s, --sort               sort logs by timestamp (slower)

Hash args:
  ALGO[,ALGO]              use hash algorithm(s)

Info flags:
      --min=DECIMAL        minimum entropy value (default: 0.0)
      --max=DECIMAL        maximal entropy value (default: 1.0)

Text flags:
      --min=NUMBER         minimum string length (default: 3)
      --max=NUMBER         maximal string length (default: 256)

File limits:
  -h, --head               limit head of file by ...
  -t, --tail               limit tail of file by ...
  -n, --lines=NUMBER       number of lines
  -c, --bytes=NUMBER       number of bytes

File loader:
  -p, --pass=PASSWORD      password for decryption (only RAR, ZIP)

Line filter:
  -e, --regexp=PATTERN     filter for lines that match pattern
  -C, --context=NUMBER     number of lines surrounding context of match
  -B, --before=NUMBER      number of lines leading context before match
  -A, --after=NUMBER       number of lines trailing context after match

Data stream:
  -f, --file=FILE          stream data to file name
  -u, --url=SERVER         stream data to server address
  -T, --auth=TOKEN         stream data using auth token
  -E, --ecs                use ECS schema for streaming
  -H, --hec                use HEC schema for streaming

Turn off:
  -r, --raw                don't process files at all
  -R, --readonly           don't write anything at all
      --no-file            don't print filenames
      --no-line            don't print line numbers
      --no-deflate         don't deflate automatically
      --no-convert         don't convert automatically

Localhost:
  -L, --logstash           short for: --ecs --url=http://localhost:8080
  -S, --splunk             short for: --hec --url=http://localhost:8088/...

Standard:
  -v, --verbose[=LEVEL]    prints more details (alt: -v, -vv, -vvv)
      --version            prints the version number
      --help               prints this help message

Hashes (cryptographic):
  SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512,
  MD5, BLAKE3-256, BLAKE3-512

Hashes (performance):
  FNV-1, FNV-1A, XXH64, XXH3

Hashes (similarity):
  SIMHASH, SDHASH, SSDEEP, TLSH

Checksums:
  ADLER32, CRC32-IEEE, CRC64-ECMA, CRC64-ISO

Example: dump the image MBR in hex format
  $ fox hex -hc512 image.dd > mbr

Example: find occurrences in all logs
  $ fox cat -elogin ./**/*.log

Example: hunt down suspicious files
  $ fox hunt .

Type "man fox" for more help...
`

// Main start and catch.
func main() {
	log.SetFlags(0)
	log.SetPrefix("fox: ")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v\n\n%s\n", err, debug.Stack())
		}
	}()

	fox := new(struct {
		Help, Version bool
		cmd.Cli
	})

	ctx := kong.Parse(fox,
		kong.NoDefaultHelp(),
		kong.DefaultEnvars("FOX"),
		kong.ConfigureHelp(kong.HelpOptions{}),
	)

	switch {
	case fox.Version:
		fmt.Printf("%s %s\n", app.Product, app.Version)
	case fox.Help || ctx.Error != nil || len(ctx.Args) == 0:
		fmt.Printf(usage, app.Version, app.Website)
	default:
		if fox.Cli.Verbose > 1 {
			defer func(start time.Time) {
				log.Printf("took %v\n", time.Since(start))
			}(time.Now())
		}

		if err := ctx.Run(&fox.Cli); err != nil {
			log.Fatal(err)
		}
	}
}
