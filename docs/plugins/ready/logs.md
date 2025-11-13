# Various log files

## [FortiAnalyzer](https://github.com/GDATAAdvancedAnalytics/FortilogDecoder)
Script to decode Fortinet binary firewall logs.
```toml
[plugin.fortilog]
name = 'fortilog'
path = '.*\.log.(gz|zst)'
exec = 'python fortilog_decoder.py "FILE"'
```
