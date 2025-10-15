# File counts
Unix `wc` like file counts, can be calculated using the [`fox counts`](../start/usage/counts.md) action or using a [hotkey](../reference/keyboard.md#utils-plugins) while in the [Terminal UI](../features/ui/terminal.md). 

> In the Terminal UI, the counts of all open files will be calculated at once.

## Example
```console
$ fox counts -p testdata/test.txt
   31108L  4633983B  testdata/test.txt
```
