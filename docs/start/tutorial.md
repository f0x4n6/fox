# Tutorial

## 1. Open a file
Type the line below into your shell to open the [Terminal UI](../features/ui/terminal.md) on the test file:

```console
fox testdata/test.evtx
```

As you can see, the [Windows Event Log](../files/logs/windows.md) is opened and will automatically be converted to [JSON Lines](../files/text/json.md) for easier processing. In the next step, we will start searching for specific artifacts.

## 2. Search for artifacts
Start typing `winlogon` to automatically switch to [Grep](../features/ui/grep.md) mode. Then press <kbd>Enter</kbd> to filter the file only for lines, that contain this exact expression:

```
winlogon
```

As you see, there is only one entry, containing the text `winlogon` in a binary name. In the next step, we will analyse this artifact using a Large Language Model. 

## 3. Analyse the artifacts
Press <kbd>F2</kbd> to switch to the [AI Assistant](../features/ai/assistant.md). If this is your first time using the assistant, we need to specify the models we want to use. Order the assistant to use the `nomic-embed-text` model for text embedding:

```
set embed nomic-embed-text
```

Now, we need to specify also a model, that analyses the embedded text for us. Order the assistant to use the `mistral` model for that:

```
set model mistral
```

Depending on your network connection, it could take a while to download the two models. After the models have downloaded, type the text below. Then press <kbd>Enter</kbd> and wait for the assistant to answer:

```
analyse this
```

If the models are not loaded into memory already, this could take a few seconds extra. As you can see, the assistant has formed an opinion about the artifact we found earlier. In the next step we will save the artifact into the evidence bag.

## 4. Save the evidence
Press <kbd>Tab</kbd> to switch back from the assistant to the original file. To tag all filtered lines as evidence, simply press <kbd>Ctrl</kbd>+<kbd>A</kbd> followed by <kbd>Ctrl</kbd>+<kbd>S</kbd> to save the formerly tagged lines into the [Evidence Bag](../features/evidence.md#evidence-bag). Press <kbd>Esc</kbd> twice, to exit the Forensic Examiner. This concludes this tutorial.
