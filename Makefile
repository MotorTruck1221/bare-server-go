SHELL := bash

all: download build

build:
	@echo "Building..."
	@go build -o bin/bare-server-go main.go

run:
	@echo "Running..."
	@go run main.go

download:
	@echo "Downloading..."
	@go get

docker-build:
	@echo "Building docker image..."
	@docker build -t bare-server-go .

docker-run:
	@echo "Running docker image..."
	@docker run -d -p 8080:8080 bare-server-go


docker: docker-build docker-run

docker-compose-build:
	@echo "Building docker-compose..."
	@docker compose build 

docker-compose-up:
	@echo "Running docker-compose..."
	@docker compose up -d

docker-compose: docker-compose-build docker-compose-up
