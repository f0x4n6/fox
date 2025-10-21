# fox counts
Display line and byte counts.

## Usage
```console
fox counts [FLAG ...] PATH ...
```

### Aliases
`co`, `wc`

### Arguments
Path(s) to open

### Global
- `-p`, `--print` — print directly to console
- `-h`, `--head` — limit head of file by *...*
- `-t`, `--tail` — limit tail of file by *...*
- `-n`, `--lines[=NUMBER]` — number of lines (*default:* `10`)
- `-c`, `--bytes[=NUMBER]` — number of bytes (*default:* `16`)

## Example
```console
$ fox counts ./**/*.txt
```
