# Eric Zimmerman's tools
> These open source digital forensics tools can be used in a wide variety of investigations including cross validation of tools, providing insight into technical details not exposed by other tools, and more.

## [JLECmd](https://github.com/EricZimmerman/JLECmd)
Automatic and Custom Destinations jump list parser with Windows 10 support.
```toml
[plugin.jle]
name = 'JLECmd'
path = '.*\.(automatic|custom)Destination-ms'
exec = 'JLECmd.exe -f "FILE" --json "TEMP"'
```

## [LECmd](https://github.com/EricZimmerman/LECmd)
Lnk Explorer Command line edition!!
```toml
[plugin.le]
name = 'LECmd'
path = '.*\.lnk'
exec = 'LECmd.exe -f "FILE" --json "TEMP"'
```

## [MFTECmd](https://github.com/EricZimmerman/MFTECmd)
Parses `$MFT` from NTFS file systems.
```toml
[plugin.mfte]
name = 'MFTECmd'
path = '\(Boot|LogFile|J|MFT|SDS)'
exec = 'MFTECmd.exe -f "FILE" --json "TEMP"'
```

## [PECmd](https://github.com/EricZimmerman/PECmd)
Prefetch Explorer Command Line.
```toml
[plugin.pe]
name = 'PECmd'
path = '.*\.pf'
exec = 'PECmd.exe -f "FILE" --json "TEMP"'
```

## [RBCmd](https://github.com/EricZimmerman/RBCmd)
Recycle bin artifact parser.
```toml
[plugin.rb]
name = 'RBCmd'
path = '(INFO2|\[0-9A-Z]{7}(\..+)?)'
exec = 'RBCmd.exe -f "FILE" --csv "TEMP"'
```

## [RECmd](https://github.com/EricZimmerman/RECmd)
Command line access to the Registry.
```toml
[plugin.re]
name = 'RECmd'
path = '.*\.dat'
exec = 'RECmd.exe -f "FILE" --json "TEMP"'
```

## [SQLECmd](https://github.com/EricZimmerman/SQLECmd)
SQLECmd parses any SQLite database from any OS.
```toml
[plugin.sqle]
name = 'SQLECmd'
path = '.*\.db'
exec = 'SQLECmd.exe -f "FILE" --json "TEMP"'
```

## [SrumECmd](https://github.com/EricZimmerman/Srum)
A SRUM parser!
```toml
[plugin.srume]
name = 'SrumECmd'
path = 'SRUDB.dat'
exec = 'SrumECmd.exe -f "FILE" --csv "TEMP"'
```

## [WxTCmd](https://github.com/EricZimmerman/WxTCmd)
WxTCmd is a parser for the Windows 10 Timeline feature database.
```toml
[plugin.wxt]
name = 'WxTCmd'
path = '.*\ActivitiesCache.db'
exec = 'WxTCmd.exe -f "FILE" --csv "TEMP"'
```
