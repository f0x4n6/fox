[![](docs/assets/img/title.png "Forensic Examiner")](https://forensic-examiner.eu)

The Swiss Army Knife for examining text files. Combining the power of many traditional tools like **grep**, **hexdump** and **strings** with the abilities of modern Large Language Models, to leverage your forensic examination process. Standalone binaries are available for Windows, Linux and macOS.

## Key Features
* Read-only [filesystem](https://forensic-examiner.eu/features/filesystem) access only
* Multibyte support with [bidirectional character](https://nvd.nist.gov/vuln/detail/CVE-2021-42574) detection
* Built-in `grep`, `hexdump` and `strings` like [abilities](https://forensic-examiner.eu/start/usage/fox)
* Built-in parsing of [Linux Journals](https://forensic-examiner.eu/files/logs/linux) and [Windows Event Logs](https://forensic-examiner.eu/files/logs/windows)
* Built-in popular [cryptographic](https://forensic-examiner.eu/utils/hashes#cryptographic-hashes) and [similarity](https://forensic-examiner.eu/utils/hashes#similarity-hashes) hashes
* Extraction and deflation of many [archive](https://forensic-examiner.eu/files/loader) formats
* Evidence bag with [Chain of Custody](https://forensic-examiner.eu/features/evidence) signing
* Evidence streaming using [Splunk HEC](https://docs.splunk.com/Documentation/Splunk/latest/RESTREF/RESTinput) or [ECS](https://www.elastic.co/docs/reference/ecs)

## Install
Install directly using Go:
```console
go install github.com/cuhsat/fox/v4@latest
```

## Build
Build a full-featured version:
```console
go build -o fox main.go
```

## License
This software is released under the [GPL-3.0](LICENSE.md)

![Status](https://img.shields.io/github/actions/workflow/status/cuhsat/fox/ci.yaml?style=flat-square&label=Status)
![Date](https://img.shields.io/github/release-date/cuhsat/fox.svg?style=flat-square&label=Date)
![Release](https://img.shields.io/github/release/cuhsat/fox.svg?style=flat-square&label=Release)
