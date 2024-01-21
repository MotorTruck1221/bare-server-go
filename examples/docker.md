# Docker (Docker and Docker Compose)

> [!IMPORTANT]  
> Currently, when using Docker **YOU MUST** clone the repository.

## Docker

### Build

```bash
make docker-build
```

### Run

```bash
make docker-run
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
make docker-compose-build
```

### Run

```bash
make docker-compose-up
```

### Stop

```bash
docker compose down
```

### Remove

```bash
docker compose down --rmi all
```
