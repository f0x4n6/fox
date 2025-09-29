![](docs/fox.png "Forensic Examiner")

The Swiss Army Knife for examining text files. Combining the power of many traditional tools like **grep**, **hexdump** and **strings** with the abilities of modern **LLMs**, to leverage your forensic examination process. Standalone native binaries are available for Windows, Linux and macOS.

![](docs/images/terminal.png)

## Key Features
* Read-only in-memory [filesystem abstraction](https://forensic-examiner.eu/features/memory)
* Multibyte support with [bidirectional character](https://nvd.nist.gov/vuln/detail/CVE-2021-42574) detection
* Built-in `grep`, `hexdump`, `diff` and `strings` like [abilities](https://forensic-examiner.eu/basics/usage/fox)
* Built-in parsing of [Linux Journals](https://forensic-examiner.eu/features/files/journal) and [Windows Event Logs](https://forensic-examiner.eu/features/files/evtx)
* Built-in popular [cryptography](https://forensic-examiner.eu/features/utils/hashes#cryptographic-hashes) and [similarity](https://forensic-examiner.eu/features/utils/hashes#similarity-hashes) hashes
* Deflation and extraction of many [archive formats](https://forensic-examiner.eu/features/loader)
* Evidence streaming using [Splunk HEC](https://docs.splunk.com/Documentation/Splunk/latest/RESTREF/RESTinput) or [ECS](https://www.elastic.co/docs/reference/ecs)
* Evidence bag with [Chain of Custody](https://forensic-examiner.eu/features/evidence) signing
* Integrated plugin support for [Dissect](https://docs.dissect.tools) or [Eric Zimmerman's tools](https://ericzimmerman.github.io/)
* Integrated assistant using [Ollama LLMs](https://ollama.com/search) like *DeepSeek R1*

## Install
Install directly using Go:
```console
go install github.com/cuhsat/fox@latest
```

## Build
Build a full-featured version:
```console
go build -o fox main.go
```

Build a minimal version with stripped AI and UI:
```console
go build -o fox -tags minimal main.go
```

## License
🦊 [Forensic Examiner](https://forensic-examiner.eu) is released under the [GPL-3.0](LICENSE.md).
