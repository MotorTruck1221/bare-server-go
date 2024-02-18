SHELL := bash
.DEFAULT_GOAL := default
.PHONY: all clean linux default windows mac run download clean compress-linux compress-windows compress-mac compress-default compress compress-all

all: clean linux windows mac compress-linux compress-windows compress-mac

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
	@go build -o bin/bare-server-go -ldflags "-s -w" main.go
	@GOOS=linux GOARCH=arm go build -o bin/bare-server-go-arm -ldflags "-s -w" main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/bare-server-go-arm64 -ldflags "-s -w" main.go
	@GOOS=linux GOARCH=386 go build -o bin/bare-server-go-386 -ldflags "-s -w" main.go

windows: download
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/bare-server-go.exe -ldflags "-s -w" main.go

mac: download
	@echo "Building for Mac..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/bare-server-go-mac -ldflags "-s -w" main.go
	@GOOS=darwin GOARCH=arm64 go build -o bin/bare-server-go-mac-arm64 -ldflags "-s -w" main.go

default: download 
	@echo "Building for default OS..."
	@go build -o bin/bare-server-go -ldflags "-s -w" main.go

compress-linux: 
	@echo "Compressing (Linux)..."
	@upx --brute bin/bare-server-go -o bin/bare-server-go-compressed
	@upx --brute bin/bare-server-go-arm -o bin/bare-server-go-arm-compressed
	@upx --brute bin/bare-server-go-arm64 -o bin/bare-server-go-arm64-compressed
	@upx --brute bin/bare-server-go-386 -o bin/bare-server-go-386-compressed

compress-windows:
	@echo "Compressing (Windows)..."
	@upx --brute bin/bare-server-go.exe -o bin/bare-server-go-compressed.exe

compress-mac:
	@echo "Compressing (Mac)..."
	@upx --brute bin/bare-server-go-mac -o bin/bare-server-go-mac-compressed
	@upx --brute bin/bare-server-go-mac-arm64 -o bin/bare-server-go-mac-arm64-compressed

compress-default:
	@echo "Compressing (Default OS)..."
	@upx --brute bin/bare-server-go -o bin/bare-server-go-compressed

compress:
	@echo "Compressing (no output file)..."
	@upx --brute bin/bare-server-go

compress-all: compress-linux compress-windows compress-mac
	@echo "Compressing all..."
