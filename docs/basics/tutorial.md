# Tutorial

## 1. Open File
Type the line below into your shell to open the UI:
```console
$ fox testdata/test.evtx
```

![](../images/tutorial/step_1.png)

## 2. Search Artifacts
Start typing `winlogon` to switch to [GREP](../features/ui/mode/grep.md) mode. Then press <kbd>Enter</kbd> to filter the file only for lines, that contain this expression:

```
winlogon
```

![](../images/tutorial/step_2.png)

## 3. Analyse Artifacts
Use <kbd>F8</kbd> to switch to the [assistant](../features/ai/assistant.md) and specify your prompt. Then press <kbd>Enter</kbd> and wait for the assistant to answer:

```
analyse this
```

> If the model is not loaded into memory already, this step could take a few seconds.

![](../images/tutorial/step_3.png)

## 4. Save Evidence
Press <kbd>Tab</kbd> to go back. Then use <kbd>Ctrl</kbd>+<kbd>S</kbd> to save all filtered lines into the [Evidence Bag](../features/evidence.md#evidence-bag). Press <kbd>Esc</kbd> twice, to exit the Forensic Examiner.

![](../images/tutorial/step_4.png)
