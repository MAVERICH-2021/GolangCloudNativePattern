# forces the make utility (regardless of whether you run it in PowerShell, CMD, or a Unix terminal) 
#to use a specific shell environment that supports the standard GOOS=value syntax.
# SHELL := /bin/bash
# The name of the output binary file
BIN_NAME := main

# --- Standard Targets ---
.PHONY: 
	all fmt vet test build clean

# The default target runs linting/checking, testing, and then builds the project.

all: 
	fmt vet test build

fmt: 
	@echo "-> Running go fmt..." 
	go fmt ./...

# Runs the Go vet tool to check for suspicious constructs.
vet: 
	@echo "-> Running go vet..." 
	go vet ./...

test: 
	@echo "-> Running go test..." 
	go test -v ./...

build:
	go build -o $(BIN_NAME).exe .

clean:
	@echo "-> Cleaning up binaries..."
	rm -f $(BIN_NAME)
	rm -rf dist/