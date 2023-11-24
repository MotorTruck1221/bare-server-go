SHELL := bash

all: download build

build:
	@echo "Building..."
	@go build -o bare-server main.go

run:
	@echo "Running..."
	@go run main.go

download:
	@echo "Downloading..."
	@go get

docker-compose-build:
	@echo "Building docker-compose..."
	@docker compose build 

docker-compose-up:
	@echo "Running docker-compose..."
	@docker compose up -d

docker-compose-down:
	@echo "Stopping docker-compose..."
	@docker compose down

docker-compose: docker-compose-build docker-compose-up
