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
Type `fox --help` for more help:
```console
$ fox [COMMAND] [FLAGS] <PATHS>
```

## Examples
Find occurrences in event logs:
```console
$ fox cat -eWinlogon ./**/*.evtx
```

Show the MBR in canonical hex:
```console
$ fox hex -mc -hc512 disk.bin
```

Find ASCII strings in binaries:
```console
$ fox text -ra8 download.exe
```

Hash the archive contents:
```console
$ fox hash -amd5,sha1 files.zip
```

Hunt down suspicious events:
```console
$ fox hunt -sxv ./**/*.dd
```

## Supports

### File Formats
BROTLI, BZIP2, CAB, GZIP, EVTX, JSONL, JOURNAL, LZ4, LZW, MINLZ, RAR, S2, SNAPPY, TAR, XZ, ZIP, ZLIB, ZSTD

### Algorithms
ADLER32, BLAKE3-256, BLAKE3-512, CRC32-IEEE, CRC64-ECMA, CRC64-ISO, FNV-1, FNV-1A, MD5, SDHASH, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512, SSDEEP, TLSH, XXH3, XXH64

## License
🦊 is released under the [GPL-3.0](LICENSE.md)