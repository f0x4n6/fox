% FOX(1) Version 2.0 | Forensic Examiner Documentation

NAME
====

**fox** — The Swiss Army Knife for examining text files

SYNOPSIS
========

| **fox** \[_action_] \[_options_] \[_path_ ...]

DESCRIPTION
===========

The Swiss Army Knife for examining text files. Combining the power of many traditional tools like **grep(1)**, **diff(1)**, **hexdump(1)** and **strings(1)** with the abilities of modern Large Language Models, to leverage your forensic examination process.

All files are processed in the order specified. Use **-** to read from **STDIN(4)**.

Actions
-------

compare

:   Compare two files.

counts

:   Display line and byte counts.

deflate

:   Deflate compressed files.

entropy

:   Display file entropy.

hash

:   Display file hash or checksums.

strings

:   Display ASCII and Unicode strings.

timeline

:   Display super timeline.

Options
-------

**-b, --bag**

:   Seize into evidence bag.

**-p, --print**

:   Print directly to console.

**--help**

:   Prints brief usage information.

**--version**

:   Prints the current version number.

Options Modes
-------------

**-x, --hex**

:   Show file in canonical hex.

Options Deflate
---------------

**-P, --pass**=_password_

:   _Password_ for decryption (only _RAR_, _ZIP_).

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

Options AI
----------

**-q, --query**=_query_

:   _Query_ for the assistant to process.

**-m, --model**=_model_

:   _Model_ for the assistant to use.

**--embed**=_model_

:   Embedding _model_ used for RAG.

**--num-ctx**=_length_

:   Context window _length_ (default 4096).

**--temp**=decimal

:   Option for temperature (default 0.2).

**--topp**=decimal

:   Option for model top_p (default 0.5).

**--topk**=number

:   Option for model top_k (default 10).

**--seed**=number

:   Option for random seed (default 8211).

Options UI
----------

**--state**=_state_

:   Sets the used UI _state_ flags.

    The state must be either be **N**, **W**, **T** or a combination. Use **-** to reset the state.

**--theme**=_theme_

:   Sets the used UI _theme_.

**--space**=number

:   Sets the used indentation space (default 2).

**--legacy**

:   Don't use any unicode decorations (_ISO-8859-1_)

Options Evidence
----------------

**-N, --case**=_name_

:   Evidence bag case _name_ (default YYYY-MM-DD).

**-f, --file**=_file_

:   Evidence bag _file_ name (default evidence).

**--mode**=_mode_

:   Evidence bag file _mode_ (default **plain**):

    Modes are **none**, **plain**, **text**, **json**, **jsonl**, **xml**, **sqlite**.

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

**--no-plugins**

:   Don't run any plugins.

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

**-T, --text**

:   Short for **--mode=text**.

**-j, --json**

:   Short for **--mode=json**.

**-J, --jsonl**

:   Short for **--mode=jsonl**.

**-Q, --sqlite**

:   Short for **--mode=sqlite**.

**-X, --xml**

:   Short for **--mode=xml**.

FILES
=====

All configuration files will also be searched for in the _/etc/fox_ and _/usr/local/etc/fox_ directories.

*~/.config/fox/foxrc*

:   Per-user configuration file.

*~/.config/fox/history*

:   Per-user input history file.

*~/.config/fox/plugins*

:   Per-user plugin configuration file.

*~/.config/fox/themes*

:   Per-user theme configuration file.

ENVIRONMENT
===========

**FOX_VAR_NAME**

:   Every configuration file setting can also be set through an environment variable. The variable name must be prefixed with _FOX_ followed by an underscore, dots must be replaced also with underscores.

EXAMPLES
========

% fox -be "login" ./**/*.log

:   Search for occurrences in all logs.

% fox -pxhc=512 image.dd > mbr

:   Export the disk MBR in hex format.

% fox -pq="analyse this" log.evtx

:   Analyse the given event log.

BUGS
====

See GitHub Issues: <https://github.com/cuhsat/fox/issues>

AUTHOR
======

Christian Uhsat <christian@uhsat.de>

SEE ALSO
========

**grep(1)**, **diff(1)**, **hexdump(1)**, **strings(1)**
