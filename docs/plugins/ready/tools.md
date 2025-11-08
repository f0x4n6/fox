# Various tools

## [Mandiant capa](https://github.com/mandiant/capa)
The FLARE team's open-source tool to identify capabilities in executable files.
```toml
[auto.capa]
name = 'capa'
path = '.*\.(bin|dll|exe|scr|sys)'
exec = 'capa "FILE"'
```

## [objdump](https://linux.die.net/man/1/objdump)
Display information from object files.
```toml
[auto.objdump]
name = 'objdump'
path = '.*\.(bin|dll|exe|scr|sys)'
exec = 'objdump --disassemble "FILE"'
```

## [RepCheck](https://github.com/riotsecurity/RepCheck)
A reputation checker for IPs, URLs, hosts or domains.
```toml
[auto.repcheck]
name = 'RepCheck'
path = '.*check\.txt'
exec = 'python3 repcheck.py -b -I "FILE"'
```
