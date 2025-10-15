# Firewall log files

## [FortiAnalyzer](https://github.com/GDATAAdvancedAnalytics/FortilogDecoder)
```toml
[auto.fortilog]
name = 'fortilog'
path = '.*\.log.(gz|zst)'
exec = [
  'python fortilog_decoder.py "FILE"'
]
```
