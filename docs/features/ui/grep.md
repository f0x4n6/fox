# Grep mode
All line-based files can be filtered with **Regular Expressions** by switching to Grep mode. The applied filters will be displayed in the lower area, while the matching parts will be highlighted. 

!!! tip "Tip"

    Use <kbd>Ctrl</kbd> + <kbd>F</kbd> to switch to Grep mode while in the Terminal UI, or type any letter or number while in Less mode.

## Lines and ranges
Specific lines and line range can be picked using the syntax defined below.

!!! tip "Tip"

    Use <kbd>Ctrl</kbd> + <kbd>P</kbd> to switch to Pick mode while in the Terminal UI.

| Syntax | Meaning               |
|--------|-----------------------|
| `1`    | Pick line `1`         |
| `1,2`  | Pick lines `1` & `2`  |
| `1-3`  | Pick lines `1` to `3` |
| `%4`   | Pick every `4`th line |

## Context window
This mode offers also a dynamic **Context Window** for filtered lines. This context window can be changed on-the-fly for all open files alike:

| Key                            | Action                  |
|--------------------------------|-------------------------|
| <kbd>Ctrl</kbd> + <kbd>K</kbd> | Increase context window |
| <kbd>Ctrl</kbd> + <kbd>J</kbd> | Decrease context window |

## Keymap
Available mode specific keys:

| Key                                | Action                         |
|------------------------------------|--------------------------------|
| <kbd>Esc</kbd>                     | Switch to [Less](less.md) mode |
| <kbd>Enter</kbd>                   | Append filter                  |
| <kbd>Backspace</kbd>               | Remove filter                  |
| <kbd>Up</kbd>                      | Prev value in history          |
| <kbd>Down</kbd>                    | Next value in history          |
| <kbd>Left</kbd>                    | Move cursor left               |
| <kbd>Right</kbd>                   | Move cursor right              |
| <kbd>Right</kbd> at the end        | Complete suggestion            |
| <kbd>Ctrl</kbd> + <kbd>Left</kbd>  | Move cursor to start           |
| <kbd>Ctrl</kbd> + <kbd>Right</kbd> | Move cursor to end             |
| <kbd>Ctrl</kbd> + <kbd>V</kbd>     | Paste from clipboard           |
| *Any other key*                    | Define a regular expression    |

## Example
![](../../assets/img/modes/grep.png "Grep mode")
