% FOX(1) Version 3.0 | Forensic Examiner Documentation

NAME
====

**fox** — The Swiss Army Knife for examining text files

SYNOPSIS
========

| **fox** **deflate** \[_options_] \[_path_ ...]

DESCRIPTION
===========

Deflate compressed files.

Alias
-----

extract, unzip, de

Options
-------

**-l, --list**

:   Don't deflate, only list files.

**-o, --out**=[_path_]

:   Output to _path_ (default .).

EXAMPLES
========

fox deflate --pass=infected ioc.zip

BUGS
====

See GitHub Issues: <_https://github.com/cuhsat/fox/issues_>

AUTHOR
======

Christian Uhsat <christian at uhsat dot de>

SEE ALSO
========

**fox(1)**
