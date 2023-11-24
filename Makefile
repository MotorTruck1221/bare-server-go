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
