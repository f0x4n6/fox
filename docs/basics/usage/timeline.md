# fox timeline
Display super timeline.

## Usage
```console
fox timeline [FLAG ...] PATH ...
```

### Aliases
`super`, `tl`

### Arguments
Path(s) to open

### Global
- `-p`, `--print` — print directly to console

### Timeline
- `-c`, `--cef` — use Common Event Format

> Currently only [Windows Event Logs](../../features/files/evtx.md) and [Linux Systemd Journals](../../features/files/journal.md) are supported.

## Example
```console
$ fox timeline ./**/*.evtx
```
