# File hashes
Different types of file hashes and checksums can be generated using the [`fox hash`](../start/usage/hash.md) action or using a [hotkey](../reference/keyboard.md#utils-plugins) while in the [Terminal UI](../features/ui/terminal.md).

> In the Terminal UI, the hashes for all open files will be calculated at once.

## Algorithms
The following algorithms are grouped by their *operational purposes*.

### Cryptographic Hashes
Built-in cryptographic hashes *(for chain-of-custody)*:

- `MD5`
- `SHA1`
- `SHA256`
- `SHA3`
- `SHA3-224`
- `SHA3-256`
- `SHA3-384`
- `SHA3-512`

### Performance Hashes
Built-in performance hashes *(for large files)*:

- `FNV-1`
- `FNV-1A`
- `XXH64`
- `XXH3`

### Similarity Hashes
Built-in similarity hashes *(for malware detection)*:
 
- `SDHASH`
- `SSDEEP`
- `TLSH`

### Checksums
Built-in checksums *(for file verification)*:

- `CRC32-IEEE`
- `CRC64-ECMA`
- `CRC64-ISO`

## Example
```console
$ fox hashes -pt=md5 testdata/test.bin
534a6a08a693b374803b6eda2bf8baab  testdata/test.bin
```
