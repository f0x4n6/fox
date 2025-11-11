% FOX(1) Version 3.0 | Forensic Examiner Documentation

NAME
====

**fox** — The Swiss Army Knife for examining text files

SYNOPSIS
========

| **fox** **strings** \[_options_] \[_path_ ...]

DESCRIPTION
===========

Display ASCII and Unicode strings.

Alias
-----

carve, st

Options
-------

**--min**=[_length_]

:   Minimum _length_ (default 3).

**--max**=[_length_]

:   Maximum _length_ (default unlimited).

**--ascii**

:   Carve only _ASCII_ strings.

**--class**

:   Run built-in classification

Classes
-------

IPv4, IPv6, MAC, Mail, URL, UUID

EXAMPLES
========

fox strings --class malware.exe

BUGS
====

See GitHub Issues: <_https://github.com/cuhsat/fox/issues_>

AUTHOR
======

Christian Uhsat <christian at uhsat dot de>

SEE ALSO
========

**fox(1)**
