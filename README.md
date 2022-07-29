# A guide to setup a new laptop

## Requirements

* [chrome](https://www.google.ca/chrome/)
* [brew](https://brew.sh/)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
* [golang](https://go.dev/doc/install)

## Setup

Once installed the minimal requirements

```bash
git clone git@github.com:vrunoa/macos-setup.git
cd dotfiles
go mod tidy
```

## Installing brew packages

```bash
go run cmd/macos-setup/macos-setup.go install -c brew.yaml
```
