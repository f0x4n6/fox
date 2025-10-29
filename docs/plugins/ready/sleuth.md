# Sleuth Kit

## [blkcalc](https://www.sleuthkit.org/sleuthkit/man/blkcalc.html)
```toml
[auto.blkcalc]
name = 'blkcalc'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'blkcalc "FILE"'
]
```

## [blkcat](https://www.sleuthkit.org/sleuthkit/man/blkcat.html)
```toml
[hotkey.f8]
name = 'blkcat'
mode = 'unit'
exec = [
  'blkcat "BASE" "INPUT"'
]
```

## [blkls](https://www.sleuthkit.org/sleuthkit/man/blkls.html)
```toml
[auto.blkls]
name = 'blkls'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'blkls "FILE"'
]
```

## [blkstat](https://www.sleuthkit.org/sleuthkit/man/blkstat.html)
```toml
[hotkey.f9]
name = 'blkstat'
mode = 'addr'
exec = [
  'blkstat "BASE" "INPUT"'
]
```

## [fls](https://www.sleuthkit.org/sleuthkit/man/fls.html)
```toml
[hotkey.f8]
name = 'fls'
exec = [
  'fls "BASE"'
]
```

## [ffind](https://www.sleuthkit.org/sleuthkit/man/ffind.html)
```toml
[hotkey.f9]
name = 'ffind'
mode = 'inode'
exec = [
  'ffind "BASE" "INPUT"'
]
```

## [fsstat](https://www.sleuthkit.org/sleuthkit/man/fsstat.html)
```toml
[auto.fsstat]
name = 'fsstat'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'fsstat "FILE"'
]
```

## [hfind](https://www.sleuthkit.org/sleuthkit/man/hfind.html)
```toml
[hotkey.f8]
name = 'hfind'
mode = 'hash'
exec = [
  'hfind "BASE" "INPUT"'
]
```

## [icat](https://www.sleuthkit.org/sleuthkit/man/icat.html)
```toml
[hotkey.f8]
name = 'icat'
mode = 'inode'
exec = [
  'icat "BASE" "INPUT"'
]
```

## [ifind](https://www.sleuthkit.org/sleuthkit/man/ifind.html)
```toml
[hotkey.f9]
name = 'ifind'
mode = 'file'
exec = [
  'ifind -n "INPUT" "BASE"'
]
```

## [istat](https://www.sleuthkit.org/sleuthkit/man/istat.html)
```toml
[hotkey.f10]
name = 'istat'
mode = 'inode'
exec = [
  'istat "BASE" "INPUT"'
]
```

## [ils](https://www.sleuthkit.org/sleuthkit/man/ils.html)
```toml
[auto.img_cat]
name = 'ils'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'ils "FILE"'
]
```

## [img_cat](https://www.sleuthkit.org/sleuthkit/man/img_cat.html)
```toml
[auto.img_cat]
name = 'img_cat'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'img_cat "FILE"'
]
```

## [img_stat](https://www.sleuthkit.org/sleuthkit/man/img_stat.html)
```toml
[auto.img_stat]
name = 'img_stat'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'img_stat "FILE"'
]
```

## [jls](https://www.sleuthkit.org/sleuthkit/man/jls.html)
```toml
[auto.jls]
name = 'jls'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'jls "FILE"'
]
```

## [jcat](https://www.sleuthkit.org/sleuthkit/man/jcat.html)
```toml
[hotkey.f8]
name = 'jcat'
mode = 'block'
exec = [
  'jcat "BASE" "INPUT"'
]
```

## [mactime](https://www.sleuthkit.org/sleuthkit/man/mactime.html)
```toml
[hotkey.f8]
name = 'mactime'
mode = 'zone'
exec = [
  'mactime -b "BASE" -dhm -z "INPUT"'
]
```

## [mmcat](https://www.sleuthkit.org/sleuthkit/man/mmcat.html)
```toml
[auto.mmcat]
name = 'mmcat'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'mmcat "FILE"'
]
```

## [mmls](https://www.sleuthkit.org/sleuthkit/man/mmls.html)
```toml
[auto.mmls]
name = 'mmls'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'mmls "FILE"'
]
```

## [mmstat](https://www.sleuthkit.org/sleuthkit/man/mmstat.html)
```toml
[auto.mmstat]
name = 'mmstat'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'mmstat "FILE"'
]
```

## [sigfind](https://www.sleuthkit.org/sleuthkit/man/sigfind.html)
```toml
[hotkey.f8]
name = 'sigfind'
mode = 'hex'
exec = [
  'sigfind "INPUT" "BASE"'
]
```

## [sorter](https://www.sleuthkit.org/sleuthkit/man/sorter.html)
```toml
[hotkey.f9]
name = 'sorter'
mode = 'path'
exec = [
  'sorter -d "INPUT" "BASE"'
]
```

## [tsk_comparedir](https://www.sleuthkit.org/sleuthkit/man/tsk_comparedir.html)
```toml
[hotkey.f10]
name = 'tsk_comparedir'
mode = 'path'
exec = [
  'tsk_comparedir "BASE" "INPUT"'
]
```

## [tsk_gettimes](https://www.sleuthkit.org/sleuthkit/man/tsk_gettimes.html)
```toml
[auto.tsk_gettimes]
name = 'tsk_gettimes'
path = '.*\.(dd|img|raw|afd|aff||afflib|afm|ewf|vhd|vmdk|00?)'
exec = [
  'tsk_gettimes "FILE"'
]
```

## [tsk_loaddb](https://www.sleuthkit.org/sleuthkit/man/tsk_loaddb.html)
```toml
[hotkey.f8]
name = 'tsk_loaddb'
exec = [
  'tsk_loaddb "BASE"'
]
```

## [tsk_recover](https://www.sleuthkit.org/sleuthkit/man/tsk_recover.html)
```toml
[hotkey.f9]
name = 'tsk_recover'
mode = 'path'
exec = [
  'tsk_recover "BASE" "INPUT"'
]
```
