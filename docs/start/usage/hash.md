# fox hash
Display file hash or checksums.

## Usage
```console
fox hash [FLAG ...] PATH ...
```

### Aliases
`ha`

### Arguments
Path(s) to open

### Additional flags
- `--type=ALGORITHM` — hash algorithm (*default:* `SHA256`)

Cryptographic hash algorithms:
> MD5, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512

Performance hash algorithms:
> FNV-1, FNV-1A, XXH64, XXH3

Similarity hash algorithms:
> SDHASH, SSDEEP, TLSH

Checksum algorithms:
> CRC32-IEEE, CRC64-ECMA, CRC64-ISO

## Example
```console
$ fox hash --type=md5 --type=sha1 artifacts.zip
```
