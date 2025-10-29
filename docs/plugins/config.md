# Plugin configuration
The per-user configuration for plugins is located under `~/.config/fox/plugins`.

## Plugin types
A plugin is a specified list of commands, that can be triggered either by pressing a hotkey, or by opening a matching file path. The `STDOUT` output of the commands are than shown, instead of the original file contents.
Two different kinds of plugins can be defined.

### Automatic plugins
Automatic plugins will be executed, if the file path matches the specified regular expression under `path`.

> To prevent the automatic execution of plugins, use either the `-r` or `--no-plugins` flag.

### Hotkey plugins
Hotkey plugins will only be executed, if the specified hotkey is pressed. They also provide the ability to ask the user for a specific `INPUT`, which then can be used as a variable while executing commands. The function keys <kbd>F8</kbd> to <kbd>F24</kbd> are reserved for hotkey plugins.

## Defining a new plugin
A plugin consists at least of `name`, `exec` and either `path` (automatic) or `mode` (hotkey). The `exec` value can either be one command or a list of commands. The following variables are available for automatic plugins:

- `FILE` path
- `BASE` path of file
- `TEMP` folder (will be created then)

Additionally, the following variable is available only for hotkey plugins:

- `INPUT` by user

For example, a simple **stat** plugin can be defined as follows:
```toml
[auto.stat]
name = 'stat'
path = '.*\.log'
exec = 'stat "FILE"'
```

For example, a simple **echo** plugin can be defined as follows:
```toml
[hotkey.f8]
name = 'echo'
mode = 'echo'
exec = 'echo "INPUT"'
```
