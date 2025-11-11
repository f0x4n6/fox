# fox deflate
Deflate compressed files.

## Usage
```console
fox deflate [FLAG ...] PATH...
```

### Aliases
`extract`, `unzip`, `de`

### Arguments
Path(s) to open

### Additional flags
- `-l`, `--list` — don't deflate, only list files
- `-o`, `--out[=PATH]` — output to path (*default:* `.`)

## Example
```console
$ fox deflate --pass=infected ioc.zip
```
