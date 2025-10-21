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

### Additional flags
- `--min[=DECIMAL]` — minimum score (*default:* `0.8`)
- `--max[=DECIMAL]` — maximum score (*default:* `0.8`)

## Example
```console
$ fox entropy --min ./**/*
```
