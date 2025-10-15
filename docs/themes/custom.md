# Custom themes
The per-user configuration for custom themes is located under `~/.config/fox/themes`.

## Defining a new theme
All colors are hexadecimal encoded RGB values, where the first value defines the foreground color and the second one the background color.

For example, the built-in [Monochrome](builtin.md#monochrome) theme is specified as:
```toml
[theme.monochrome]
name = 'Monochrome'
terminal = [ 0xFFFFFF, 0x000000 ] # terminal
surface0 = [ 0xFFFFFF, 0x000000 ] # info panes
surface1 = [ 0xFFFFFF, 0x000000 ] # info areas
surface2 = [ 0xFFFFFF, 0x000000 ] # current file
surface3 = [ 0x000000, 0xFFFFFF ] # current mode
overlay0 = [ 0x000000, 0xFFFFFF ] # failure messages
overlay1 = [ 0x000000, 0xFFFFFF ] # success messages
subtext0 = [ 0xFFFFFF, 0x000000 ] # navigation items
subtext1 = [ 0x000000, 0xFFFFFF ] # separator lines
subtext2 = [ 0x000000, 0xFFFFFF ] # highlights
subtext3 = [ 0x000000, 0xFFFFFF ] # completion
```
