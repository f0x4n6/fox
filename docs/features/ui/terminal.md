# Terminal UI
The terminal user interface is [compatible](https://github.com/gdamore/tcell) with many different terminal applications across all supported operating systems, including the popular [iTerm2](https://iterm2.com/) and the new [Windows Terminal](https://github.com/microsoft/terminal).

!!! tip "Tip"

    Use <kbd>Ctrl</kbd> + <kbd>Z</kbd> to suspend to shell while in the Terminal UI. Then type `exit` to resume the software.

## Overview
![](../../assets/img/terminal.png "Terminal UI")

### Title
The upper area of the terminal shows the currently displayed file, any executed plugin after the dash, the current file index and the count of open files.

### Status
The lower area of the terminal shows the current terminal mode, any applied filters, the line count and context window as well as the current UI state.

### States
| State | Meaning         |
|-------|-----------------|
| `N`   | Show navigation |
| `W`   | Wrap text       |
| `T`   | Tail file       |

## Themes
The user interface is fully themeable and has many popular color [themes](../../themes/builtin.md) are already integrated.

## History
All user input will be saved in timestamped history file located under `~/.config/fox/history`.

> To prevent the history file of being written, use the `--readonly` flag.

```
1766579696;analyse this
```

## Config
All relevant settings will be saved in configuration file located under `~/.config/fox/foxrc`.

> To prevent the config file of being written, use the `--readonly` flag.

```toml
[ai]
model = 'mistral'
embed = 'nomic-embed-text'
num_ctx = 4096
seed = 8211
temp = '0.2'
topk = 10
topp = '0.5'

[ui]
theme = 'examiner-dark'
space = 2

[ui.state]
n = true
w = true
y = true
r = true
```
