# Roadmap

## 1. Bug Fixes
* ~~Done~~

## 2. Refactorings
* Remap some key bindings

## 3. Features
* Add new image README as GIF
* Add more global flags to actions
* Add tab tab cycling
* Add custom render layout, like text, hex, csv, log
* Add macOS Brew cask
  * https://goreleaser.com/customization/homebrew_casks/
  * https://docs.brew.sh/Cask-Cookbook
* Add *Yara* rules scan

## 4. Optimizations
* Check replacement with Kong / Koanf / Coral
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
* Add search to Hex mode?
* Add `readline` config support?
  * https://github.com/chzyer/readline
* Add `find` command ability?
* Use multiple parallel filters?
* Use reflow algos?
  * https://github.com/muesli/reflow
* Generic syntax highlighting?
  * `Start Color [ … ] End Color`
  * `{}[]<>()““‘‘:;`
