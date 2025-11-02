# The Sleuth Kit
> The Sleuth Kit® (*TSK*) is a library and collection of command line tools that allow you to investigate disk images. The core functionality of *TSK* allows you to analyze volume and file system data.

## [blkcalc](https://www.sleuthkit.org/sleuthkit/man/blkcalc.html)
Converts between unallocated disk unit numbers and regular disk unit numbers.
```toml
[auto.blkcalc]
name = 'blkcalc'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'blkcalc "FILE"'
```

## [blkcat](https://www.sleuthkit.org/sleuthkit/man/blkcat.html)
Display the contents of file system data unit in a disk image.
```toml
[hotkey.f9]
name = 'blkcat'
mode = 'unit'
exec = 'blkcat "BASE" "INPUT"'
```

## [blkls](https://www.sleuthkit.org/sleuthkit/man/blkls.html)
List or output file system data units.
```toml
[auto.blkls]
name = 'blkls'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'blkls "FILE"'
```

## [blkstat](https://www.sleuthkit.org/sleuthkit/man/blkstat.html)
Display details of a file system data unit (i.e. block or sector).
```toml
[hotkey.f10]
name = 'blkstat'
mode = 'addr'
exec = 'blkstat "BASE" "INPUT"'
```

## [fcat](https://www.sleuthkit.org/sleuthkit/man/fcat.html)
Output the contents of a file based on its name.
```toml
[hotkey.f11]
name = 'fcat'
mode = 'file'
exec = 'fcat "INPUT" "BASE"'
```

## [ffind](https://www.sleuthkit.org/sleuthkit/man/ffind.html)
Finds the name of the file or directory using a given inode.
```toml
[hotkey.f10]
name = 'ffind'
mode = 'inode'
exec = 'ffind "BASE" "INPUT"'
```

## [fls](https://www.sleuthkit.org/sleuthkit/man/fls.html)
List file and directory names in a disk image.
```toml
[hotkey.f9]
name = 'fls'
exec = 'fls "BASE"'
```

## [fsstat](https://www.sleuthkit.org/sleuthkit/man/fsstat.html)
Display general details of a file system.
```toml
[auto.fsstat]
name = 'fsstat'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'fsstat "FILE"'
```

## [hfind](https://www.sleuthkit.org/sleuthkit/man/hfind.html)
Lookup a hash value in a hash database.
```toml
[hotkey.f9]
name = 'hfind'
mode = 'hash'
exec = 'hfind "BASE" "INPUT"'
```

## [icat](https://www.sleuthkit.org/sleuthkit/man/icat.html)
Output the contents of a file based on its inode number.
```toml
[hotkey.f9]
name = 'icat'
mode = 'inode'
exec = 'icat "BASE" "INPUT"'
```

## [ifind](https://www.sleuthkit.org/sleuthkit/man/ifind.html)
Find the meta-data structure that has allocated a given disk unit or file name.
```toml
[hotkey.f10]
name = 'ifind'
mode = 'file'
exec = 'ifind -n "INPUT" "BASE"'
```

## [ils](https://www.sleuthkit.org/sleuthkit/man/ils.html)
List inode information.
```toml
[auto.img_cat]
name = 'ils'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'ils "FILE"'
```

## [img_cat](https://www.sleuthkit.org/sleuthkit/man/img_cat.html)
Output contents of an image file.
```toml
[auto.img_cat]
name = 'img_cat'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'img_cat "FILE"'
```

## [img_stat](https://www.sleuthkit.org/sleuthkit/man/img_stat.html)
Display details of an image file.
```toml
[auto.img_stat]
name = 'img_stat'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'img_stat "FILE"'
```

## [istat](https://www.sleuthkit.org/sleuthkit/man/istat.html)
Display details of a meta-data structure (i.e. inode).
```toml
[hotkey.f11]
name = 'istat'
mode = 'inode'
exec = 'istat "BASE" "INPUT"'
```

## [jcat](https://www.sleuthkit.org/sleuthkit/man/jcat.html)
Show the contents of a block in the file system journal.
```toml
[hotkey.f9]
name = 'jcat'
mode = 'block'
exec = 'jcat "BASE" "INPUT"'
```

## [jls](https://www.sleuthkit.org/sleuthkit/man/jls.html)
List the contents of a file system journal.
```toml
[auto.jls]
name = 'jls'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'jls "FILE"'
```

## [mactime](https://www.sleuthkit.org/sleuthkit/man/mactime.html)
Create an ASCII time line of file activity.
```toml
[hotkey.f9]
name = 'mactime'
mode = 'zone'
exec = 'mactime -b "BASE" -dhm -z "INPUT"'
```

## [mmcat](https://www.sleuthkit.org/sleuthkit/man/mmcat.html)
Output the contents of a partition to stdout.
```toml
[auto.mmcat]
name = 'mmcat'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'mmcat "FILE"'
```

## [mmls](https://www.sleuthkit.org/sleuthkit/man/mmls.html)
Display the partition layout of a volume system (partition tables).
```toml
[auto.mmls]
name = 'mmls'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'mmls "FILE"'
```

## [mmstat](https://www.sleuthkit.org/sleuthkit/man/mmstat.html)
Display details about the volume system (partition tables).
```toml
[auto.mmstat]
name = 'mmstat'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'mmstat "FILE"'
```

## [sigfind](https://www.sleuthkit.org/sleuthkit/man/sigfind.html)
Find a binary signature in a file.
```toml
[hotkey.f9]
name = 'sigfind'
mode = 'hex'
exec = 'sigfind "INPUT" "BASE"'
```

## [sorter](https://www.sleuthkit.org/sleuthkit/man/sorter.html)
Sort files in an image into categories based on file type.
```toml
[hotkey.f10]
name = 'sorter'
mode = 'path'
exec = 'sorter -d "INPUT" "BASE"'
```

## [tsk_comparedir](https://www.sleuthkit.org/sleuthkit/man/tsk_comparedir.html)
Compare the contents of a directory with the contents of an image or local device.
```toml
[hotkey.f11]
name = 'tsk_comparedir'
mode = 'path'
exec = 'tsk_comparedir "BASE" "INPUT"'
```

## [tsk_gettimes](https://www.sleuthkit.org/sleuthkit/man/tsk_gettimes.html)
Collect MAC times from a disk image into a body file.
```toml
[auto.tsk_gettimes]
name = 'tsk_gettimes'
path = '.*\.(dd|img|raw|afd|aff|afflib|afm|ewf|vhd|vmdk|00?)'
exec = 'tsk_gettimes "FILE"'
```

## [tsk_loaddb](https://www.sleuthkit.org/sleuthkit/man/tsk_loaddb.html)
Populate a SQLite database with metadata from a disk image.
```toml
[hotkey.f9]
name = 'tsk_loaddb'
exec = 'tsk_loaddb "BASE"'
```

## [tsk_recover](https://www.sleuthkit.org/sleuthkit/man/tsk_recover.html)
Export files from an image into a local directory.
```toml
[hotkey.f10]
name = 'tsk_recover'
mode = 'path'
exec = 'tsk_recover "BASE" "INPUT"'
```
