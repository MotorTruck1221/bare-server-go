SHELL := bash
.DEFAULT_GOAL := default
.PHONY: all clean linux default windows mac run download clean

all: clean linux windows mac

clean:
	@echo "Cleaning..."
	@rm -rf bin

download:
	@echo "Downloading..."
	@go get

run:
	@echo "Running..."
	@go run main.go

linux: download
	@echo "Building for Linux..."
	@go build -o bin/bare-server-go main.go
	@GOOS=linux GOARCH=arm go build -o bin/bare-server-go-arm main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/bare-server-go-arm64 main.go
	@GOOS=linux GOARCH=386 go build -o bin/bare-server-go-386 main.go

windows: download
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/bare-server-go.exe main.go

mac: download
	@echo "Building for Mac..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/bare-server-go-mac main.go
	@GOOS=darwin GOARCH=arm64 go build -o bin/bare-server-go-mac-arm64 main.go

default: download 
	@echo "Building for default OS..."
	@go build -o bin/bare-server-go main.go
