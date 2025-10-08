# Timeline
A **Super Timeline** of all open log files can be created with the [`fox timeline`](../../basics/usage/timeline.md) command or using a [hotkey](../ui/keymap.md). For Windows Event Logs, a comprehensive list of translations from `EventID` to their description is used.  

## Common Event Format
For [Windows Event Logs](../../features/files/evtx.md) or [Linux Systemd Journals](../../features/files/journal.md) the timeline can be created in the [Common Event Format](https://www.microfocus.com/documentation/arcsight/arcsight-smartconnectors-8.3/cef-implementation-standard/Content/CEF/Chapter%201%20What%20is%20CEF.htm).

## Example
```console
$ fox timeline -pc testdata/test.journal
2022-10-24T00:00:01.009823Z fedora CEF:1|Forensic Examiner|fox|v1.2.3|100|Booting Linux on physical CPU 0x0000000000 [0x410fd083]|Unknown
2022-10-24T00:00:01.009952Z fedora CEF:1|Forensic Examiner|fox|v1.2.3|100|Linux version 6.0.7-301.fc37.aarch64 (mockbuild@buildvm-a64-05.iad2.fedoraproject.org) (gcc (GCC) 12.2.1 20220819 (Red Hat 12.2.1-2), GNU ld version 2.38-24.fc37) #1 SMP PREEMPT_DYNAMIC Fri Nov 4 18:13:35 UTC 2022|Unknown
2022-10-24T00:00:01.010016Z fedora CEF:1|Forensic Examiner|fox|v1.2.3|100|Machine model: Raspberry Pi 4 Model B Rev 1.1|Unknown
...
```
