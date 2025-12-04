![Fox](assets/logo.png "Forensic Examiner")

The Cyber Forensic Swiss Army Knife. Standalone binaries available for Windows, Linux and macOS.

![Release](https://img.shields.io/github/release/cuhsat/fox.svg?style=flat-square&label=Release)
![Status](https://img.shields.io/github/actions/workflow/status/cuhsat/fox/ci.yaml?style=flat-square&label=Status)

```console
go install github.com/cuhsat/fox/v4@latest
```

## Features
* Hunt suspicious system activities with:
  * an integrated super timeline in [Common Event Format](https://www.microfocus.com/documentation/arcsight/arcsight-smartconnectors-8.3/cef-implementation-standard/Content/CEF/Chapter%201%20What%20is%20CEF.htm) 
  * build from carved [Linux Journals](https://systemd.io/JOURNAL_FILE_FORMAT/) and [Windows Event Logs](https://learn.microsoft.com/en-us/windows/win32/eventlog/event-log-file-format)
  * with an extensive translation list of Windows Event IDs
  * and preconfigured critical events to examine 
  * support for `JSON`, `JSON Lines` and `SQLite3` output
* Enforced read-only filesystem access
* [Bidirectional character](https://nvd.nist.gov/vuln/detail/CVE-2021-42574) detection
* Fast file entropy calculation
* Built-in `grep`, `head`, `tail`, `hexdump`, `strings` and `wc` like abilities
* Many built-in archive compression formats*
* Many built-in cryptographic, fuzzy and fast hashes**
* (TODO) Evidence saving with Chain of Custody signing
* (TODO) Evidence streaming using [Splunk HEC](https://help.splunk.com/en/splunk-enterprise/leverage-rest-apis/rest-api-reference/10.0/input-endpoints/input-endpoint-descriptions) or [ECS](https://www.elastic.co/docs/reference/ecs)

## Usage
```console
Usage:
  fox [COMMAND] [FLAGS] <PATHS>

Commands:
  hunt [FLAGS] <PATHS>     hunt suspicious activities
    -a, --all              show logs with all severities
    -x, --ext              show logs with all extensions (slow)
    -s, --sort             show logs sorted by timestamp (slow)
    -j, --json             show logs as JSON objects
    -J, --jsonl            show logs as JSON lines
    -D, --sqlite           save logs to SQLite3 DB

  hash [FLAGS] <PATHS>     prints file hashes and checksums
    -a, --algo=ALGO[,]     use algorithm(s) (default: SHA256)
    -F, --find=HASH[,]     show only files that match

  info [FLAGS] <PATHS>     prints file infos and entropy
        --min=DECIMAL      minimum entropy value (default: 0.0)
        --max=DECIMAL      maximal entropy value (default: 1.0)

  text [FLAGS] <PATHS>     prints file ASCII strings
        --min=NUMBER       minimum string length (default: 3)
        --max=NUMBER       maximal string length (default: 256)

  hex [FLAGS] <PATHS>      prints file in hex format
    -m, --mode=[c|hd|xxd]  use compatible mode for output 

  cat [FLAGS] <PATHS>      prints file (default)

File limits:
  -h, --head               limit head of file by ...
  -t, --tail               limit tail of file by ...
  -n, --lines=NUMBER       number of lines
  -c, --bytes=NUMBER       number of bytes

File loader:
  -p, --pass=PASSWORD      password for decryption (only RAR and ZIP)

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
  -q, --quiet              don't print anything
      --no-file            don't print filenames
      --no-line            don't print line numbers
      --no-color           don't colorize the output
      --no-deflate         don't deflate automatically
      --no-convert         don't convert automatically

Aliases:
  -L, --logstash           alias for: -E -uhttp://localhost:8080
  -S, --splunk             alias for: -H -uhttp://localhost:8088/...

Standard:
  -d, --dry-run            prints only the found filenames
  -v, --verbose[=LEVEL]    prints more details (v/vv/vvv)
      --version            prints the version number
      --help               prints this help message

Positional arguments:
  Path(s) to open or '-' for STDIN
```

## Commands
* `hunt`   hunt suspicious activities
* `hash`   prints file hash using algorithms
* `info`   prints file info and entropy
* `text`   prints file ASCII strings
* `hex`    prints file in hex format
* `cat`    prints file (default)

## Formats
Archive formats:
```
CAB, RAR, TAR, ZIP
```

Compression formats:
```
Brotli, BZip2, Gzip, lz4, LZW, MinLZ, S2, Snappy, xz, zlib, Zstandard
```

## Hashes
Cryptographic hashes:
```
BLAKE3-256, BLAKE3-512, MD5, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512
```

Performance hashes:
```
FNV-1, FNV-1A, XXH64, XXH3
```

Similarity hashes:
```
SDHASH, SSDEEP, TLSH
```

Checksums:
```
ADLER32, CRC32-IEEE, CRC64-ECMA, CRC64-ISO
```

## Examples
Find occurrences in event logs:
```console
$ fox cat -elogin ./**/*.evtx
```

Show the MBR in canonical hex:
```console
$ fox hex -mc -hc512 disk.bin
```

Hunt down suspicious events:
```console
$ fox hunt -sxv ./**/*.dd
```

## License
🦊 is released under the [GPL-3.0](LICENSE.md)