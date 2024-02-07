# Docker (Docker and Docker Compose)

> [!IMPORTANT]  
> Currently, when using Docker **YOU MUST** clone the repository.

## Docker

### Build

```bash
docker build -t bare-server-go .
```

### Run

```bash
docker run -p 8080:8080 bare-server-go
```

### Stop

To stop your going to have to get the container id and then stop it.

```bash
docker ps
docker stop <container id>
```

### Remove

To remove your going to have to get the container id and then remove it.

```bash
docker ps -a
docker rm <container id>
```


## Docker Compose

### Build

```bash
docker compose build
```

### Run

```bash
docker compose up 
```

Or in the background

```bash
docker compose up -d
```

### Stop

```bash
docker compose down
```

### Remove

```bash
docker compose down --rmi all
```
