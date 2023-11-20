# Bare Server GO

- This is a complete bare server written in GO that is compliant with the [TompHTTP Server Specifications](https://github.com/tomphttp/specifications)
- This is used in conjunction with [Prism](https://github.com/ruby-network/prism) a frontend that focuses on being lightweight and extremely fast,
## Usage

### Prerequisites

- [GO](https://golang.org/doc/install)

### Installation

- Clone this repository
- To build run:
```bash
make all
```
- To run:
```bash
./bare-server
```

#### Docker 

- To build AND run:
```bash
make docker
```

- To JUST build:
```bash
make docker-build
```

- To JUST run:
```bash
make docker-run
```

#### Docker Compose

- To build AND run:
```bash
make docker-compose
```

- To JUST build:
```bash
make docker-compose-build
```

- To JUST run:
```bash
make docker-compose-up
```

## Todo 

- [ ] V3 implementation
- [ ] Websocket Support
- [ ] Installable in your own project
- [ ] Older versions of the spec
