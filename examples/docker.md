# Docker (Docker and Docker Compose)

## Docker (Prebuilt, and from Source)

### Build from Source

- Clone the repository
```bash
git clone https://github.com/tomphttp/bare-server-go.git 
```

- Change to the directory
```bash
cd bare-server-go
```

- Build the image
```bash
docker build -t bare-server-go .
```

- Run the image
```bash
docker run -p 8080:8080 bare-server-go
```

### Run (Prebuilt)

1. Pull the image
```bash
docker pull motortruck1221/bare-server-go:main
```

2. Run the image
```bash
docker run -p 8080:8080 motortruck1221/bare-server-go:main
```

### Stop (BOTH)

To stop your going to have to get the container id and then stop it.

```bash
docker ps
docker stop <container id>
```

### Remove (BOTH)

To remove your going to have to get the container id and then remove it.

```bash
docker ps -a
docker rm <container id>
```


## Docker Compose (Prebuilt and Build from Source)

### Build from Source

- Clone the repository
```bash
git clone https://github.com/tomphttp/bare-server-go.git 
```

- Change to the directory
```bash
cd bare-server-go
```

- Build the image
```bash
docker compose -f docker-compose.build.yml build
```

- Run the image
```bash
docker compose -f docker-compose.build.yml up
```

### Run (Prebuilt)

1. Aquire the docker-compose.yml file
  - `wget` (Linux)
  ```bash
    wget https://raw.githubusercontent.com/tomphttp/bare-server-go/latest/docker-compose.yml
```
  - `curl` (Linux)
  ```bash
    curl -O https://raw.githubusercontent.com/tomphttp/bare-server-go/latest/docker-compose.yml
```
  - Download the file from the repository [here](https://github.com/tomphttp/bare-server-go/tree/latest/docker-compose.yml)

2. Run the image
```bash
docker compose up
```

### Stop (BOTH)

```bash
docker compose down
```

### Remove (BOTH)

```bash
docker compose down --rmi all
```
