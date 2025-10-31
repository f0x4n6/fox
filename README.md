![](docs/assets/img/examiner.png "Forensic Examiner")

The Swiss Army Knife for examining text files. Combining the power of many traditional tools like **grep**, **diff**, **hexdump** and **strings** with the abilities of modern Large Language Models, to leverage your forensic examination process. Standalone binaries are available for Windows, Linux and macOS.

![](docs/assets/img/terminal.png)

## Key Features
* Read-only [forensic filesystem](https://forensic-examiner.eu/features/filesystem) abstraction
* Multibyte support with [bidirectional character](https://nvd.nist.gov/vuln/detail/CVE-2021-42574) detection
* Built-in `grep`, `diff`, `hexdump` and `strings` like [abilities](https://forensic-examiner.eu/start/usage/fox)
* Built-in parsing of [Linux Journals](https://forensic-examiner.eu/files/logs/linux) and [Windows Event Logs](https://forensic-examiner.eu/files/logs/windows)
* Built-in popular [cryptographic](https://forensic-examiner.eu/utils/hashes#cryptographic-hashes) and [similarity](https://forensic-examiner.eu/utils/hashes#similarity-hashes) hashes
* Extraction and deflation of many [archive](https://forensic-examiner.eu/files/loader) formats
* Evidence bag with [Chain of Custody](https://forensic-examiner.eu/features/evidence) signing
* Evidence streaming using [Splunk HEC](https://docs.splunk.com/Documentation/Splunk/latest/RESTREF/RESTinput) or [ECS](https://www.elastic.co/docs/reference/ecs)
* Integrated super timeline in [Common Event Format](https://www.microfocus.com/documentation/arcsight/arcsight-smartconnectors-8.3/cef-implementation-standard/Content/CEF/Chapter%201%20What%20is%20CEF.htm)
* Integrated plugin support for e.g. [Sleuth Kit](https://www.sleuthkit.org/) or [Eric Zimmerman's tools](https://ericzimmerman.github.io/)
* Integrated assistant using local [Ollama LLMs](https://ollama.com/search) like *DeepSeek R1*

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
