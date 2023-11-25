# Bare Server GO

- This is a complete bare server written in GO that is compliant with the [TompHTTP Server Specifications](https://github.com/tomphttp/specifications)

***NOTE: Unfinished*** *I skimmed the spec and wrote this only in a few hours. There are lots of unimplemented features and bugs. (and it's curtrently not very TOMP compatible)*
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
