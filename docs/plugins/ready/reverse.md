# Reverse Engineering tools

## [Mandiant capa](https://github.com/mandiant/capa)
```toml
[auto.capa]
name = 'capa'
path = '.*\.(bin|dll|exe|scr|sys)'
exec = [
  'capa "FILE"'
]
```

## [objdump](https://linux.die.net/man/1/objdump)
```toml
[auto.objdump]
name = 'objdump'
path = '.*\.(bin|dll|exe|scr|sys)'
exec = [
  'objdump --disassemble "FILE"'
]
```
