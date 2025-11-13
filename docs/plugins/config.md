# Plugin configuration
The per-user configuration for plugins is located under `~/.config/fox/plugins`.

## Plugin types
A plugin is a specified list of commands, that can be triggered either by pressing a hotkey, or by opening a matching file path. The `STDOUT` output of the commands are than shown, instead of the original file contents.
Two different kinds of plugins can be defined.

### Automatic plugins
Automatic plugins will be executed, if the file path matches the specified regular expression under `path`.

> To prevent the automatic execution of plugins, use either the `-r` or `--no-plugins` flag.

## Defining a new plugin
A plugin consists at least of `name`, `exec` and `path`. The `exec` value can either be one command or a list of commands. The following variables are available for automatic plugins:

- `FILE` path
- `TEMP` folder (will be created then)

For example, a simple **stat** plugin can be defined as follows:
```toml
[plugin.stat]
name = 'stat'
path = '.*\.log'
exec = 'stat "FILE"'
```
