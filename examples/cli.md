# Usage as a CLI tool

## Prerequisites 
- [GO](https://golang.org/doc/install)

> [!IMPORTANT]  
> Currently, the CLI tool is only installable via GO or the [AUR](https://aur.archlinux.org/packages/bare-server-go/)

## Installation

- With GO:
```bash
go install github.com/tomphttp/bare-server-go@latest
```

- AUR (yay):
```bash
yay -S bare-server-go
```

## Usage

For use with defaults run:
```bash
bare-server-go start
```

### Start Command Flags

| Flag | Description | Default |
| ---- | ----------- | ------- |
| -p, --port   | Port to listen on | 8080 |
| -H, --host   | Host to listen on | 0.0.0.0 |
| -d, --directory | Directory to serve on | / |
| -h, --help | Help for bare-go start command | N/A |
