# Dissect framework

## [target-info](https://docs.dissect.tools/en/latest/tools/target-info.html)
```toml
[auto.info]
name = 'target-info'
path = '.*\.(dd|img|raw|ad1|asdf|E0?|00?)'
exec = [
  'target-info "FILE"'
]
```

## [target-query](https://docs.dissect.tools/en/latest/tools/target-query.html)
```toml
[hotkey.f8]
name = 'target-query'
mode = 'query'
exec = [
  'target-query -j -f "INPUT" "BASE"'
]
```

## [target-shell](https://docs.dissect.tools/en/latest/tools/target-shell.html)
```toml
[hotkey.f9]
name = 'target-shell'
mode = 'shell'
exec = [
  'target-shell -c="INPUT" "BASE"'
]
```
