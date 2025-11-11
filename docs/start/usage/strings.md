# fox strings
Display ASCII and Unicode strings.

## Usage
```console
fox strings [FLAG ...] PATH ...
```

### Aliases
`carve`, `st`

### Arguments
Path(s) to open

### Additional flags
- `--min=LENGTH` — minimum length (*default:* `3`)
- `--max=LENGTH` — maximum length (*default:* unlimited)
- `--ascii` — carve only ASCII strings
- `--class` — run built-in classification

Built-in classifications:
> IPv4, IPv6, MAC, Mail, URL, UUID

## Example
```console
$ fox strings --class malware.exe
```
