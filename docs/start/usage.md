# fox
The Swiss Army Knife for examining text files.

## Usage
```console
fox [COMMAND] [FLAG ...] [PATH ...]
```


### Arguments
Path(s) to open or `-` for **STDIN**

### Print
- `--no-file` — don't print filenames
- `--no-line` — don't print line numbers

### File limits
- `-h`, `--head` — limit head of file by *...*
- `-t`, `--tail` — limit tail of file by *...*
- `-n`, `--lines[=NUMBER]` — number of lines (*default:* `10`)
- `-c`, `--bytes[=NUMBER]` — number of bytes (*default:* `16`)

### Line filter
- `-e`, `--regexp=PATTERN` — filter for lines that match pattern
- `-C`, `--context=NUMBER` — number of lines surrounding context of match
- `-B`, `--before=NUMBER` — number of lines leading context before match
- `-A`, `--after=NUMBER` — number of lines trailing context after match

### Evidence bag
- `-N`, `--case=NAME` — evidence bag case name (*default:* `YYYY-MM-DD`)
- `-F`, `--file=FILE` — evidence bag file name (*default:* `evidence`)
- `--mode=MODE` — evidence bag file mode (*default:* `text`)

Available evidence bag modes:
> none, text, json, jsonl, sqlite

### Evidence sign
- `-s, --sign=PHRASE` — key phrase to sign evidence bag via **HMAC-SHA256**

### Evidence URL
- `-u`, `--url=SERVER` — forward evidence to server address
- `-a`, `--auth=TOKEN` — forward evidence using auth token
- `--ecs` — use **[ECS](https://forensic-examiner.eu/features/evidence.html#ecs-schema)** schema for evidence
- `--hec` — use **[HEC](https://forensic-examiner.eu/features/evidence.html#hec-schema)** schema for evidence

### Deflate
- `--pass=PASSWORD` — password for decryption (only **RAR**, **ZIP**)

### Turn off
- `-R`, `--readonly` — don't write any new files
- `-r`, `--raw` — don't process files at all
- `--no-convert` — don't convert automatically
- `--no-deflate` — don't deflate automatically

### Aliases
- `-L`, `--logstash` — short for: `--ecs --url=http://localhost:8080`
- `-S`, `--splunk` — short for: `--hec --url=http://localhost:8088/...`
- `-j`, `--json` — short for: `--mode=json`
- `-J`, `--jsonl` — short for: `--mode=jsonl`
- `-Q`, `--sqlite` — short for: `--mode=sqlite`

### Standard
- `--help` — prints this message
- `--version` — prints the version

## Examples
Search for occurrences in all logs:
```console
$ fox -e "login" ./**/*.log
```

Export the disk MBR in hex format:
```console
$ fox -xhc=512 image.dd > mbr
```
