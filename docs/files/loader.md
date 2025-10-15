# File loader
The built-in file loader applies matching and globbing [rules](https://github.com/bmatcuk/doublestar?tab=readme-ov-file#patterns) to any given path, before opening and processing the resulting files further. The processing is done in three steps, described below.

## 1. Deflate
The file loader is capable of reading nested archives/compressed files in many different formats. This is done by parsing the file beginning for various **magic numbers**, like the Unix `file` utility does.

> To prevent automatic file deflation, use either the `-r` or `--no-deflate` flag.

## 2. Execute
If the file to be opened matches a `path` specified by an [Automatic Plugin](../plugins/config.md#automatic-plugins), the plugin is executed and its output is shown instead of the files contents. This step may yield additional files to be opened automatically.

> To prevent automatic plugin execution, use either the `-r` or `--no-plugins` flag.

## 3. Convert
The file loader then detects and converts binary files like [Windows Event Logs](logs/windows.md) or [Linux Systemd Journals](logs/linux.md) to [JSON Lines](https://jsonlines.org/) upon loading for further processing. 

> To prevent automatic file conversion, use either the `-r` or `--no-convert` flag.

## Supported

### Compression formats
Built-in deflation of files is supported for the following formats:

| Format   | Specification                                                       |
|----------|---------------------------------------------------------------------|
| `brotli` | [Brotli Compression Format](https://github.com/google/brotli)       |
| `bzip2`  | [bzip2 Data Compressor](https://sourceware.org/bzip2/)              |
| `gzip`   | [Gzip Data Compression](https://www.gnu.org/software/gzip/)         |
| `lz4`    | [LZ4 Compression](https://github.com/lz4/lz4)                       |
| `xz`     | [xz File Format](https://tukaani.org/xz/format.html)                |
| `zlib`   | [zlib Compression](https://zlib.net/)                               |
| `zstd`   | [Zstandard Compression Algorithm](https://github.com/facebook/zstd) |

### Archive formats
Built-in extraction of files is supported for the following formats:

| Format | Specification                                                                           |
|--------|-----------------------------------------------------------------------------------------|
| `cab`  | [Microsoft Cabinet Format](https://msdn.microsoft.com/en-us/library/bb267310.aspx)      |
| `rar`  | [RAR Archive Format](https://www.rarlab.com/technote.htm)                               |
| `tar`  | [Basic Tar Format](https://www.gnu.org/software/tar/manual/html_node/Standard.html)     |
| `zip`  | [ZIP File Format](https://www.loc.gov/preservation/digital/formats/fdd/fdd000354.shtml) |

## Examples
Load all files with `log.gz` extension:
```console
$ fox *.log.gz
```

Load all files with `evtx` extension, from all subfolders:
```console
$ fox ./**/*.evtx
```

Load all files from all subfolders:
```console
$ fox ./**/*
```
