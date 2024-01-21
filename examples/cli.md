# Usage as a CLI tool

## Prerequisites 
- [GO](https://golang.org/doc/install)

> [!IMPORTANT]  
> Currently, the CLI tool is only installable via GO itself.

## Installation

- With GO:
```bash
go install github.com/ruby-network/bare-server-go@latest
```

## Usage

For use with defaults run:
```bash
bare-go start
```

### Start Command Flags

| Flag | Description | Default |
| ---- | ----------- | ------- |
| -p, --port   | Port to listen on | 8080 |
| -H, --host   | Host to listen on | 0.0.0.0 |
| -d, --directory | Directory to serve on | / |
| -h, --help | Help for bare-go start command | N/A |
