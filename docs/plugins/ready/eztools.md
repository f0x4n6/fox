# Eric Zimmerman's tools

## [JLECmd](https://github.com/EricZimmerman/JLECmd)
```toml
[auto.jle]
name = 'JLECmd'
path = '.*\.(automatic|custom)Destination-ms'
exec = [
  'JLECmd.exe -f "FILE" --json "TEMP"'
]
```

## [LECmd](https://github.com/EricZimmerman/LECmd)
```toml
[auto.le]
name = 'LECmd'
path = '.*\.lnk'
exec = [
  'LECmd.exe -f "FILE" --json "TEMP"'
]
```

## [MFTECmd](https://github.com/EricZimmerman/MFTECmd)
```toml
[auto.mfte]
name = 'MFTECmd'
path = '\(Boot|LogFile|J|MFT|SDS)'
exec = [
  'MFTECmd.exe -f "FILE" --json "TEMP"'
]
```

## [PECmd](https://github.com/EricZimmerman/PECmd)
```toml
[auto.pe]
name = 'PECmd'
path = '.*\.pf'
exec = [
  'PECmd.exe -f "FILE" --json "TEMP"'
]
```

## [RBCmd](https://github.com/EricZimmerman/RBCmd)
```toml
[auto.rb]
name = 'RBCmd'
path = '(INFO2|\[0-9A-Z]{7}(\..+)?)'
exec = [
  'RBCmd.exe -f "FILE" --csv "TEMP"'
]
```

## [RECmd](https://github.com/EricZimmerman/RECmd)
```toml
[auto.re]
name = 'RECmd'
path = '.*\.dat'
exec = [
  'RECmd.exe -f "FILE" --json "TEMP"'
]
```

## [SQLECmd](https://github.com/EricZimmerman/SQLECmd)
```toml
[auto.sqle]
name = 'SQLECmd'
path = '.*\.db'
exec = [
  'SQLECmd.exe -f "FILE" --json "TEMP"'
]
```

## [SrumECmd](https://github.com/EricZimmerman/Srum)
```toml
[auto.srume]
name = 'SrumECmd'
path = 'SRUDB.dat'
exec = [
  'SrumECmd.exe -f "FILE" --csv "TEMP"'
]
```

## [WxTCmd](https://github.com/EricZimmerman/WxTCmd)
```toml
[auto.wxt]
name = 'WxTCmd'
path = '.*\ActivitiesCache.db'
exec = [
  'WxTCmd.exe -f "FILE" --csv "TEMP"'
]
```
