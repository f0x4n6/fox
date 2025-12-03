![Fox](assets/logo.png "Forensic Examiner")

The Forensic Examiners Swiss Army Knife.

![Release](https://img.shields.io/github/release/cuhsat/fox.svg?style=flat-square&label=Release)
![Status](https://img.shields.io/github/actions/workflow/status/cuhsat/fox/ci.yaml?style=flat-square&label=Status)

## Features
TODO

## Examples
TODO

## Install
Stable releases are available for both AMD64 and ARM64 architectures, including signed SBOMs.

### Install using Make
Execute the following command:

```console
sudo make install
```

### Install using Go
Execute the following command as user:

```console
go install github.com/cuhsat/fox/v4@latest
```

### Install binary
> No installation is required, as the binaries are standalone native executables.

This software has no dependencies to any other operating system packages. Standalone binaries are available for the following operating systems:

* Linux 
* macOS
* Windows

### Install packages
Linux packages are available for:

Alpine Linux
```console
sudo apk add --allow-untrusted fox*.apk
```

Arch Linux
```console
sudo pacman -U fox*.pkg.tar.zst
```

Debian Linux
```console
sudo dpkg -i fox*.deb
```

Red Hat Linux
```console
sudo rpm -i fox*.rpm
```

## License
🦊 is released under the [GPL-3.0](LICENSE.md)