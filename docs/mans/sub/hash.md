% FOX(1) Version 3.0 | Forensic Examiner Documentation

NAME
====

**fox** — The Swiss Army Knife for examining text files

SYNOPSIS
========

| **fox** **hash** \[_options_] \[_path_ ...]

DESCRIPTION
===========

Display file hash or checksums.

Alias
-----

ha

Options
-------

**--type**=[_algorithm_]

:   hash _algorithm_ (default _SHA256_).

ALGORITHMS
==========

Cryptographic hash algorithms
-----------------------------

MD5, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512, BLAKE3-256, BLAKE3-512

Performance hash algorithms
---------------------------

FNV-1, FNV-1A, XXH64, XXH3

Similarity hash algorithms
--------------------------

SDHASH, SSDEEP, TLSH

Checksum algorithms
-------------------

CRC32-IEEE, CRC64-ECMA, CRC64-ISO

EXAMPLES
========

fox hash --type=md5 --type=sha1 artifacts.zip

BUGS
====

See GitHub Issues: <_https://github.com/cuhsat/fox/issues_>

AUTHOR
======

Christian Uhsat <christian at uhsat dot de>

SEE ALSO
========

**fox(1)**
