# Installation
Stable [releases](https://github.com/cuhsat/fox/releases) are available for both `AMD64` and `ARM64` architectures, including signed SBOMs.

> To use the [AI Assistant](../features/ai/assistant.md), a connection to a local or remote running [Ollama](https://ollama.com) server is required.

## Install from source
Simply execute as user the following command:

```console
go install github.com/cuhsat/fox@latest
```

## Install from binaries
> No installation is required, as the binaries are standalone native executables.

Standalone binaries are available for the following operating systems:
 
- Linux
- macOS
- Windows

## Install from packages
Linux packages are available for:

### Alpine Linux
```console
# sudo apk add --allow-untrusted fox*.apk
```

### Arch Linux
```console
# sudo pacman -U fox*.pkg.tar.zst
```

### Debian Linux
```console
# sudo dpkg -i fox*.deb
```

### Red Hat Linux
```console
# sudo rpm -i fox*.rpm
```

## Dependencies
This software has no dependencies to any other operating system packages.
