# Unique lines
The unique lines combined of all files can be displayed using the [`fox unique`](../start/usage/entropy.md) action or using a [hotkey](../reference/keyboard.md#files) while in the [Terminal UI](../features/ui/terminal.md).

## Example
```console
$ fox unique -p testdata/test.evtx
{"Event":{"EventData":{"Binary":null,"Data":["WIN-TIP3N90KK74","IEUSER-S9CSARBA"]},"System":{"Channel":"System","Computer":"WIN-TIP3N90KK74","EventID":{"Qualifiers":"32768","Value":"6011"},"EventLog":"","EventRecordID":"1","Keywords":"0x80000000000000","Level":"4","Security":{},"Task":"0","TimeCreated":{"SystemTime":"2018-01-03T06:28:19Z"}}}}
{"Event":{"EventData":{"Binary":null,"Data":["6.03.","9600","Multiprocessor Free","16384"]},"System":{"Channel":"System","Computer":"WIN-TIP3N90KK74","EventID":{"Qualifiers":"32768","Value":"6009"},"EventLog":"","EventRecordID":"2","Keywords":"0x80000000000000","Level":"4","Security":{},"Task":"0","TimeCreated":{"SystemTime":"2018-01-03T06:28:19Z"}}}}
{"Event":{"EventData":{"Binary":"E20701000300030006001C001300A3030000000000000000","Data":null},"System":{"Channel":"System","Computer":"WIN-TIP3N90KK74","EventID":{"Qualifiers":"32768","Value":"6005"},"EventLog":"","EventRecordID":"3","Keywords":"0x80000000000000","Level":"4","Security":{},"Task":"0","TimeCreated":{"SystemTime":"2018-01-03T06:28:19Z"}}}}
...
```
