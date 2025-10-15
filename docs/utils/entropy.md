# File entropy
The **Shannon** entropy of a file can be calculated using the [`fox entropy`](../start/usage/entropy.md) action or using a [hotkey](../reference/keyboard.md#utils-plugins) while in the [Terminal UI](../features/ui/terminal.md).

> In the Terminal UI, the entropy of all open files will be calculated at once.

## Example
```console
$ fox entropy -p testdata/test.bin
0.9939355973  testdata/test.bin
```
