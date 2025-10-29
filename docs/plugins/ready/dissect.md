# Dissect framework
> Effortlessly extract and investigate forensic artefacts from any source with **Dissect**. Go from intake call to patient zero in a matter of hours, even in infrastructures with thousands of systems.

## [target-info](https://docs.dissect.tools/en/latest/tools/target-info.html)
`target-info` enables you to quickly view basic system information of your targets.
```toml
[auto.info]
name = 'target-info'
path = '.*\.(dd|img|raw|ad1|asdf|E0?|00?)'
exec = 'target-info "FILE"'
```

## [target-query](https://docs.dissect.tools/en/latest/tools/target-query.html)
`target-query` is a tool used to query specific data inside a one or more targets.
```toml
[hotkey.f8]
name = 'target-query'
mode = 'query'
exec = 'target-query -j -f "INPUT" "BASE"'
```

## [target-shell](https://docs.dissect.tools/en/latest/tools/target-shell.html)
`target-shell` gives you the ability to access a target using a virtual shell environment.
```toml
[hotkey.f9]
name = 'target-shell'
mode = 'shell'
exec = 'target-shell -c="INPUT" "BASE"'
```
