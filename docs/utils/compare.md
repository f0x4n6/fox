# File comparison
Two files can easily be compared line-by-line using the [`fox compare`](../start/usage/compare.md) action. The [Unified Format](https://git-scm.com/docs/diff-format), which is commonly used by the Unix `git diff` command, can also be emulated to provide a compatible output.

## Example
```console
$ fox compare -p testdata/test.cls testdata/test.diff
testdata/test.class
testdata/test.diff
- 1 Hello World
+ 1 Hello Earth
  2 127.0.0.1
- 3 2001:0db8:85a3:08d3:1319:8a2e:0370:7344
  4 00:80:41:ae:fd:7e
  5 test@example.org
  6 https://example.org
- 7 550e8400-e29b-11d4-a716-446655440000
+ 6 EOF
```
