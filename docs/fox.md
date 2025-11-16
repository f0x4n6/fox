% FOX(1) Version 4.0 | Forensic Examiner Documentation

NAME
====

**fox** — The Swiss Army Knife for examining text files

SYNOPSIS
========

| **fox** \[_command_] \[_options_] \[_path_ ...]

DESCRIPTION
===========

The Swiss Army Knife for examining text files. Combining the power of many traditional tools like **grep(1)**, **hexdump(1)** and **strings(1)** with the abilities of modern Large Language Models, to leverage your forensic examination process.

All files are processed in the order specified. Use **-** to read from **STDIN(4)**.

Commands
--------

counts

:   Display line and byte counts.

entropy

:   Display file entropy.

hash

:   Display file hash or checksums.

strings

:   Display ASCII and Unicode strings.

Options
-------

**--help**

:   Prints brief usage information.

**--version**

:   Prints the current version number.

Options Limits
--------------

**-h, --head**

:   Limit head of file by **lines** or **bytes**.

**-t, --tail**

:   Limit tail of file by **lines** or **bytes**.

**-n, --lines**=[_number_]

:   _Number_ of lines (default 10).

**-c, --bytes**=[_number_]

:   _Number_ of bytes (default 16).

Options Filters
---------------

**-e, --regexp**=_pattern_

:   Filter for lines that match _pattern_.

**-C, --context**=_number_

:   _Number_ of lines surrounding context of match.

**-B, --before**=_number_

:   _Number_ of lines leading context before match.

**-A, --after**=_number_

:   _Number_ of lines trailing context after match.

Options Evidence
----------------

**-b, --bag**

:   Save into evidence bag.

**-N, --case**=_name_

:   Evidence bag case _name_ (default YYYY-MM-DD).

**-F, --file**=_file_

:   Evidence bag _file_ name (default evidence).

**--mode**=_mode_

:   Evidence bag file _mode_ (default **text**):

    Modes are **none**, **plain**, **text**, **json**, **jsonl**, **sqlite**.

**-s, --sign**=_phrase_

:   Key _phrase_ to sign evidence bag via _HMAC-SHA256_.

**-u, --url**=_server_

:   Forward evidence to _server_ address.

**-a, --auth**=_token_

:   Forward evidence using auth _token_.

**--ecs**

:   Use _ECS_ schema for evidence.

**--hec**

:   Use _HEC_ schema for evidence.

Options Deflate
---------------

**--pass**=_password_

:   _Password_ for decryption (only _RAR_, _ZIP_).

Options Misc
------------

**-R, --readonly**

:   Don't write any new files.

**-r, --raw**

:   Don't process files at all.

**--no-convert**

:   Don't convert automatically.

**--no-deflate**

:   Don't deflate automatically.

**--no-mouse**

:   Don't use the mouse.

**--no-file**

:   Don't print filenames.

**--no-line**

:   Don't print line numbers.

Aliases
-------

**-L, --logstash**

:   Short for **--ecs --url=http://localhost:8080**.

**-S, --splunk**

:   Short for **--hec --url=http://localhost:8088/services/collector/event**.

**-P, --plain**

:   Short for **--mode=plain**.

**-j, --json**

:   Short for **--mode=json**.

**-J, --jsonl**

:   Short for **--mode=jsonl**.

**-Q, --sqlite**

:   Short for **--mode=sqlite**.

FILES
=====

All configuration files will also be searched for in the _/etc/fox_ and _/usr/local/etc/fox_ directories.

*~/.config/fox/foxrc*

:   Per-user configuration file.

ENVIRONMENT
===========

**FOX_VAR_NAME**

:   Every configuration file setting can also be set through an environment variable. The variable name must be prefixed with _FOX_ followed by an underscore, dots must be replaced also with underscores.

EXAMPLES
========

fox -e "login" ./**/*.log

:   Search for occurrences in all logs.

fox dump -hc=512 image.dd > mbr

:   Export the disk MBR in hex format.

BUGS
====

See GitHub Issues: <_https://github.com/cuhsat/fox/issues_>

AUTHOR
======

Christian Uhsat <christian at uhsat dot de>

SEE ALSO
========

**grep(1)**, **hexdump(1)**, **strings(1)**
