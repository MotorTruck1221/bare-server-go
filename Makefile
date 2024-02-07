SHELL := bash
.PHONY: all build run download 
.DEFAULT_GOAL := all


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
