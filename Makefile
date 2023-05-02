# makefile to format, test, and build the project
.PHONY: all

all: format test build

format:
	@echo "Formatting..."
	@go fmt ./...
	
test:
	@echo "Testing..."
	@go test ./...

build:
	@echo "Building..."
	@go build -o bin/ ./...