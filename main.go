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
  fox [COMMAND] [FLAGS...] <PATHS...>

Commands:
  hunt [FLAGS] <PATHS>     hunt suspicious activities
    -a, --all              show logs with all severities
    -s, --sort             show logs sorted by timestamp (slow)

  hash ALGO[,..] <PATHS>   prints file hash using algorithm(s)

  info [FLAGS] <PATHS>     prints file info and entropy
        --min=DECIMAL      minimum entropy value (default: 0.0)
        --max=DECIMAL      maximal entropy value (default: 1.0)

  text [FLAGS] <PATHS>     prints file ASCII strings
        --min=NUMBER       minimum string length (default: 3)
        --max=NUMBER       maximal string length (default: 256)

  hex [FLAGS] <PATHS>      prints file in hex format
  cat [FLAGS] <PATHS>      prints file (default)

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

Disable:
  -r, --raw                don't process files at all
  -R, --readonly           don't write anything at all
      --no-file            don't print filenames
      --no-line            don't print line numbers
      --no-deflate         don't deflate automatically
      --no-convert         don't convert automatically

Aliases:
  -L, --logstash           alias for: -E -uhttp://localhost:8080
  -S, --splunk             alias for: -H -uhttp://localhost:8088/...

Standard:
  -v, --verbose[=LEVEL]    prints more details (v/vv/vvv)
      --version            prints the version number
      --help               prints this help message

Positional arguments:
  Path(s) to open or '-' for STDIN

Hashes (cryptographic):
  SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512,
  MD5, BLAKE3-256, BLAKE3-512

Hashes (performance):
  FNV-1, FNV-1A, XXH64, XXH3

Hashes (similarity):
  SIMHASH, SDHASH, SSDEEP, TLSH

Checksums:
  ADLER32, CRC32-IEEE, CRC64-ECMA, CRC64-ISO

Example: Dump the image MBR in hex format
  $ fox hex -hc512 image.dd > mbr

Example: Find occurrences in all logs
  $ fox cat -elogin ./**/*.log

Example: Hunt down suspicious files
  $ fox hunt -as .

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
