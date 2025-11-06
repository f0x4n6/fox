# Evidence handling
File contents, including [AI Assistant](ai/assistant.md) or [Plugin](../plugins/config.md) output, can be saved as digital evidence along with metadata into an **Evidence Bag**. Please refer to the specific schema for details about the applied data and file format. Timestamps will be normalized to **Coordinated Universal Time (UTC)**.

Using the [Terminal UI](ui/terminal.md), all filtered lines of the current file will be saved into an evidence bag by pressing <kbd>Ctrl</kbd> + <kbd>S</kbd>. The evidence bag then can be viewed by pressing <kbd>Ctrl</kbd> + <kbd>B</kbd>.

## Evidence bag
The following evidence bag formats are available:

| Type      | Details                                                                         |
|-----------|---------------------------------------------------------------------------------|
| `plain`   | Plain text                                                                      |
| `text`    | Plain text with metadata (*default*)                                            |
| `json`    | [JSON Schema](https://github.com/cuhsat/fox/blob/main/api/evidence.schema.json) |
| `jsonl`   | [JSON Schema](https://github.com/cuhsat/fox/blob/main/api/evidence.schema.json) |
| `xml`     | [XML Schema](https://github.com/cuhsat/fox/blob/main/api/evidence.schema.xsd)   |
| `sqlite3` | [SQL Schema](https://github.com/cuhsat/fox/blob/main/api/evidence.schema.sql)   |

Example of an evidence bag in `text` format (including metadata):
```
 ___ ___  ___ ___ _  _ ___ ___ ___   _____  __  _   __  __ ___ _  _ ___ ___
| __/ _ \| _ \ __| \| / __|_ _/ __| | __\ \/ / /_\ |  \/  |_ _| \| | __| _ \
| _| (_) |   / _||  ' \__ \| | (__  | _| >  < / _ \| |\/| || ||  ' | _||   /
|_| \___/|_|_\___|_|\_|___/___\___| |___/_/\_\_/ \_\_|  |_|___|_|\_|___|_|_\

Forensic Examiner Evidence Bag (v1.2.3)
==============================================================================
File: /hello.txt > Hello (12 bytes)
User: jd (John Doe)
Time: 2025-12-06T12:34:56Z modified, 2025-12-24T12:34:56Z seized
Hash: SHA256 d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26
------------------------------------------------------------------------------
1:0 Hello World
==============================================================================
```

## Evidence signing
While saving, the evidence bag will be cryptographically signed using `SHA256` to a separate file with the `.sig` file extension. This is done to guarantee the juristic **Chain of Custody**.

> To use `HMAC-SHA256` for signing, specific a key phrase using the `--sign` flag.

## Evidence streaming
All evidence can additionally be streamed to an HTTP server in various formats using the `--url` flag. This is done immediately after the time of saving the evidence to the specified bag. 

### Raw text
If no schema was specified, the raw data will be streamed as `text/plain` to the given server alongside with the file metadata as custom `x-evidence-*` HTTP headers.

```
Hello World
```

### ECS schema
The evidence can also be streamed as `application/json` to a given URL using the [Elastic Common Schema](https://www.elastic.co/docs/reference/ecs) version `9.1.0`.

!!! tip "Tip"

    Use the `-L` flag as a shortcut to specify a local running **Logstash** server.

```json
{
  "@timestamp": "2025-12-24T12:34:56.789000Z",
  "message": "Hello World\n",
  "labels": {
    "case": "demo-case",
    "filters": ""
  },
  "agent": {
    "type": "Forensic Examiner",
    "version": "1.2.3"
  },
  "ecs": {
    "version": "9.1.0"
  },
  "file": {
    "mtime": "2025-12-06T12:34:56.789000Z",
    "path": "/hello.txt",
    "size": 12,
    "hash": {
      "sha256": "d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26"
    }
  },
  "user": {
    "name": "jd",
    "full_name": "John Doe"
  }
}
```

### HEC schema
The evidence can also be streamed as `application/json` to a [Splunk HTTP Event Collector](https://docs.splunk.com/Documentation/Splunk/latest/RESTREF/RESTinput), an **Authorization Token** is required for this.

!!! tip "Tip"

    Use the `-S` flag as a shortcut to specify a local running **Splunk** server.

```json
{
  "time": 1766576096000,
  "source": "Forensic Examiner",
  "sourcetype": "_json",
  "index": "demo-case",
  "event": {
    "user": "jd (John Doe)",
    "path": "/hello.txt",
    "hash": "d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26",
    "time": 1765020896000,
    "size": 12,
    "lines": [
      "Hello World\n"
    ]
  }
}
```
