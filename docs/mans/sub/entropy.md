% FOX(1) Version 3.0 | Forensic Examiner Documentation

NAME
====

**fox** — The Swiss Army Knife for examining text files

SYNOPSIS
========

| **fox** **entropy** \[_options_] \[_path_ ...]

DESCRIPTION
===========

Display file entropy.

Alias
-----

en

Options
-------

**--min**=[_score_]

:   Minimum _score_ (default 0.8).

**--max**=[_score_]

:   Maximum _score_ (default 0.8).

EXAMPLES
========

fox entropy --min ./**/*

BUGS
====

See GitHub Issues: <_https://github.com/cuhsat/fox/issues_>

AUTHOR
======

Christian Uhsat <christian at uhsat dot de>

SEE ALSO
========

**fox(1)**
