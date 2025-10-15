# File deflation
Common compressions and archive [formats](../files/loader.md) can be viewed or deflated using the [`fox deflate`](../start/usage/deflate.md) action. For archive files, a new directory will be created from the initial file name and used as the destination (*default*).

> Password-based decryption is only supported for **RAR** and **ZIP** archives.

## Example
```console
$ fox deflate testdata/test.zip
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  hello.txt.bz2
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  hello.txt.gz
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  hello.txt.lz4
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  hello.txt.xz
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  hello.txt.zst
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  hello.rar/hello.txt
6 file(s) written
```
