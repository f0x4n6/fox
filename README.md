![Fox](assets/logo.png "Forensic Examiner")

Fox is the [Fo]rensic E[x]aminers Swiss Army Knife. Available for Windows, Linux and macOS as AMD64 and ARM64 standalone native binaries.

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
* Built-in archive formats:
    * `CAB`, `RAR`, `TAR`, `ZIP`
* Built-in compression formats:
    * `Brotli`, `BZip2`, `Gzip`, `lz4`, `LZW`, `MinLZ`, `S2`, `Snappy`, `xz`, `zlib`, `Zstandard`
* Built-in cryptographic hashes:
  * `BLAKE3-256`, `BLAKE3-512`, `MD5`, `SHA1`, `SHA256`, `SHA3`, `SHA3-224`, `SHA3-256`, `SHA3-384`, `SHA3-512`
* Built-in similarity / fuzzy hashes:
  * `sdhash`, `SSDeep`, `TLSH`
* Built-in performance hashes:
  * `FNV-1`, `FNV-1a`, `XXH64`, `XXH3`
* Built-in checksums:
  * `Adler-32`, `CRC32-IEEE`, `CRC64-ECMA`, `CRC64-ISO`
* (TODO) Evidence saving with Chain of Custody signing
* (TODO) Evidence streaming using [Splunk HEC](https://help.splunk.com/en/splunk-enterprise/leverage-rest-apis/rest-api-reference/10.0/input-endpoints/input-endpoint-descriptions) or [ECS](https://www.elastic.co/docs/reference/ecs)

## Examples
TODO

Example: Find occurrences in event logs
$ fox cat -elogin ./**/*.evtx

Example: Show the MBR as canonical hex
$ fox hex -mc -hc512 disk.bin

Example: Hunt down suspicious events
$ fox hunt -sxv ./**/*.dd

## License
🦊 is released under the [GPL-3.0](LICENSE.md)