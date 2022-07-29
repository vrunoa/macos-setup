# A helper to setup a new laptop

## Requirements

* [chrome](https://www.google.ca/chrome/)
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
go run cmd/macos-setup/macos-setup.go setup -c config.yaml
```
