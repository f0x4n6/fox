# String carving
The **ASCII** and **Unicode** strings contained in binary files can be carved using the [`fox strings`](../start/usage/strings.md) action or using a [hotkey](../reference/keyboard.md#utils-plugins) while in the [Terminal UI](../features/ui/terminal.md).

> In the Terminal UI, the strings of all open files will be carved at once.

## Classifications
The string carving util has also a built-in classification of found strings. These classifications can be used to filter the source for possible **Indicators of Compromise**.

| Class  | Description                   |
|--------|-------------------------------|
| `ipv4` | IPv4 address                  |
| `ipv6` | IPv6 address                  |
| `mac`  | MAC address                   |
| `mail` | E-Mail address                |
| `url`  | Uniform Resource Locator      |
| `uuid` | Universally Unique Identifier |

## Example
```console
$ fox strings -pi testdata/test.ioc
00000000  data  Hello World
0000000c  ipv4  127.0.0.1
00000016  ipv6  2001:0db8:85a3:08d3:1319:8a2e:0370:7344
0000003e  mac   00:80:41:ae:fd:7e
00000050  mail  test@example.org
00000061  url   https://example.org
00000075  uuid  550e8400-e29b-11d4-a716-446655440000
```
