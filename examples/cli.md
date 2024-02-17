# Usage as a CLI tool

## Prerequisites 
- [GO](https://golang.org/doc/install) (when building from source/installing via GO)

> [!IMPORTANT]  
> Currently, the CLI tool is only installable via GO or the [AUR](https://aur.archlinux.org/packages/bare-server-go/)

## Installation

- From Source (linux only):
1.  ```bash
    git clone https://github.com/tomphttp/bare-server-go
    ```

2.  ```bash
    cd bare-server-go
    ```

3.  ```bash
    make
    ```

4.  ```bash
    ./bin/bare-server-go start
    ```

##### Optional 

5. Compress the binary
```bash
make compress
```

---
- With GO:
```bash
go install github.com/tomphttp/bare-server-go@latest
```

---
- AUR (binary):
```bash
yay -S bare-server-go
```

- AUR (build from source):
```bash
yay -S bare-server-go-git
```
---
## CLI Usage

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
