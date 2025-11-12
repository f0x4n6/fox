# Building
To build a full-featured version:
```console
go build -o fox main.go
```

## Minimal build
To build a `minimal` version, stripped off [AI Assistant](../features/ai/assistant.md) support:
```console
go build -o fox -tags minimal main.go
```
