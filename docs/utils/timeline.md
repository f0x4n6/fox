# Super timeline
A combined **Super Timeline** of all open log files can be created using the [`fox timeline`](../start/usage/timeline.md) action or using a [hotkey](../reference/keyboard.md#utils-plugins) while in the [Terminal UI](../features/ui/terminal.md). A comprehensive list of `EventID` descriptions is used for Windows Event Logs.

> In the Terminal UI, the timeline of all open files will be created.

## Common Event Format
A super timeline using the [Common Event Format](https://www.microfocus.com/documentation/arcsight/arcsight-smartconnectors-8.3/cef-implementation-standard/Content/CEF/Chapter%201%20What%20is%20CEF.htm), can be created for [Linux Systemd Journals](../files/logs/linux.md) or [Windows Event Logs](../files/logs/windows.md).

## Example
```console
$ fox timeline -pc testdata/test.journal
2022-10-24T00:00:01.009823Z fedora CEF:1|Forensic Examiner|fox|v1.2.3|100|Booting Linux on physical CPU 0x0000000000 [0x410fd083]|Unknown
2022-10-24T00:00:01.009952Z fedora CEF:1|Forensic Examiner|fox|v1.2.3|100|Linux version 6.0.7-301.fc37.aarch64 (mockbuild@buildvm-a64-05.iad2.fedoraproject.org) (gcc (GCC) 12.2.1 20220819 (Red Hat 12.2.1-2), GNU ld version 2.38-24.fc37) #1 SMP PREEMPT_DYNAMIC Fri Nov 4 18:13:35 UTC 2022|Unknown
2022-10-24T00:00:01.010016Z fedora CEF:1|Forensic Examiner|fox|v1.2.3|100|Machine model: Raspberry Pi 4 Model B Rev 1.1|Unknown
...
```
