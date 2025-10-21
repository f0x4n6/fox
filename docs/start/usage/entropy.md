# fox entropy
Display file entropy.

## Usage
```console
fox entropy [FLAG ...] PATH ...
```

### Aliases
`en`

### Arguments
Path(s) to open

### Global
- `-p`, `--print` — print directly to console
- `-h`, `--head` — limit head of file by *...*
- `-t`, `--tail` — limit tail of file by *...*
- `-n`, `--lines[=NUMBER]` — number of lines (*default:* `10`)
- `-c`, `--bytes[=NUMBER]` — number of bytes (*default:* `16`)

### Entropy
- `--min[=DECIMAL]` — minimum score (*default:* `0.8`)
- `--max[=DECIMAL]` — maximum score (*default:* `0.8`)

## Example
```console
$ fox entropy --min ./**/*
```
