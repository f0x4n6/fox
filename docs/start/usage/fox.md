# fox
The Swiss Army Knife for examining text files.

## Usage
```console
fox [ACTION] [FLAG ...] [PATH ...]
```

### Actions
- `compare`
- `counts`
- `deflate`
- `entropy`
- `hash`
- `strings`
- `timeline`

### Arguments
Path(s) to open or `-` for **STDIN**

### Local
- `-b`, `--bag` — save into evidence bag
- `-x`, `--hex` — show file in canonical hex

### Print
- `-p`, `--print` — print only to console
- `--no-file` — don't print filenames
- `--no-line` — don't print line numbers

### Deflate
- `-P`, `--pass=PASSWORD` — password for decryption (only **RAR**, **ZIP**)

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

### AI assistant
- `-q`, `--query=QUERY` — query for the assistant to process
- `-m`, `--model=MODEL` — model for the assistant to use
- `--embed=MODEL` — embedding model for RAG

### AI options
- `--num-ctx=SIZE` — context window length (*default:* `4096`)
- `--temp=DECIMAL` — option for temperature (*default:* `0.2`)
- `--topp=DECIMAL` — option for model top_p (*default:* `0.5`)
- `--topk=NUMBER` — option for model top_k (*default:* `10`)
- `--seed=NUMBER` — option for random seed (*default:* `8211`)

### UI flags
- `--state={N|W|T|-}` — sets the used UI state flags
- `--theme=THEME` — sets the used UI theme
- `--space=NUMBER` — sets the used indentation space (*default:* `2`)
- `--legacy` — don't use any unicode decorations (**ISO 8859-1**)

### Evidence bag
- `-N`, `--case=NAME` — evidence bag case name (*default:* `YYYY-MM-DD`)
- `-f`, `--file=FILE` — evidence bag file name (*default:* `evidence`)
- `--mode=MODE` — evidence bag file mode (*default:* `plain`)

Available evidence bag modes:
> none, plain, text, json, jsonl, xml, sqlite

### Evidence sign
- `-s, --sign=PHRASE` — key phrase to sign evidence bag via **HMAC-SHA256**

### Evidence URL
- `-u`, `--url=SERVER` — forward evidence to server address
- `-a`, `--auth=TOKEN` — forward evidence using auth token
- `--ecs` — use **[ECS](https://forensic-examiner.eu/features/evidence.html#ecs-schema)** schema for evidence
- `--hec` — use **[HEC](https://forensic-examiner.eu/features/evidence.html#hec-schema)** schema for evidence

### Turn off
- `-R`, `--readonly` — don't write any new files
- `-r`, `--raw` — don't process files at all
- `--no-convert` — don't convert automatically
- `--no-deflate` — don't deflate automatically
- `--no-plugins` — don't run any plugins
- `--no-mouse` — don't use the mouse

### Aliases
- `-L`, `--logstash` — short for: `--ecs --url=http://localhost:8080`
- `-S`, `--splunk` — short for: `--hec --url=http://localhost:8088/...`
- `-T`, `--text` — short for: `--mode=text`
- `-j`, `--json` — short for: `--mode=json`
- `-J`, `--jsonl` — short for: `--mode=jsonl`
- `-Q`, `--sqlite` — short for: `--mode=sqlite`
- `-X`, `--xml` — short for: `--mode=xml`

### Standard
- `--help` — prints this message
- `--version` — prints the version

## Examples
Search for occurrences in all logs:
```console
$ fox -be "login" ./**/*.log
```

Export the disk MBR in hex format:
```console
$ fox -pxhc=512 image.dd > mbr
```

Analyse the given event log:
```console
$ fox -pq="analyse this" log.evtx
```
