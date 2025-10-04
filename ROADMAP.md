# Roadmap

## 1. Bug Fixes
* Fix partial highlighting

## 2. Refactorings
* Add keychain support
  * Remap some key bindings

## 3. Features
* Add Super Timeline ability
  * Specify *Common Log Format* (CLF)
    * Encoding utf-8
    * Line ending `\n`
    * Add to docs under *Resource* section
    * Timestamps as ISO-8601 or RFC 3339
    * CEF format?
  * Extract timestamps per RegEx
  * Sort lines by timestamp
* Add additional documentation
  * Add manpage(s)
    * https://appliedgo.net/spotlight/install-with-manpage/
    * https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
  * Add `bash` and `zsh` autocompletion files
    * https://applejag.eu/blog/go-spf13-cobra-custom-flag-types/
    * https://cobra.dev/docs/how-to-guides/shell-completion/
    * https://github.com/spf13/cobra/blob/v1.8.0/site/content/completions/_index.md#completions-for-flags
* Add *Yara* rules scan

## 4. Optimizations
* Add debug output (`-v` flag)
* SMap speed
  * https://dev.to/moseeh_52/efficient-file-reading-in-go-mastering-bufionewscanner-vs-osreadfile-4h05
  * https://dave.cheney.net/high-performance-json.html

## 5. Ideas
* Colorize CLI output?
  * https://github.com/logrusorgru/aurora
  * https://github.com/cyucelen/marker
* Watch configs for changes?
  * `viper.WatchConfig()`
  * `viper.OnConfigChange(func(e fsnotify.Event)`
* Add auto-completion from history?
* Add search to HEX mode?
* Add `readline` config support?
  * https://github.com/chzyer/readline
* Add `find` command ability?
* Use multiple parallel filters?
* Use reflow algos?
  * https://github.com/muesli/reflow
* Generic syntax highlighting?
  * `Start Color [ … ] End Color`
  * `{}[]<>()““‘‘:;`
