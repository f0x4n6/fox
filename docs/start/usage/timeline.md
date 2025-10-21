# fox timeline
Display super timeline.

> Currently only [Windows Event Logs](../../files/logs/windows.md) and [Linux Systemd Journals](../../files/logs/linux.md) are supported.

## Usage
```console
fox timeline [FLAG ...] PATH ...
```

### Aliases
`time`, `tl`

### Arguments
Path(s) to open

### Additional flags
- `--cef` — use the Common Event Format

## Example
```console
$ fox timeline --cef ./**/*.evtx
```
