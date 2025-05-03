GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: install
install:
	go get github.com/mark3labs/mcp-go
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: build
build:
	go build -o bin/mcp-server-poc cmd/main.go

.PHONY: lint
lint:
	goimports -w .

.PHONY: gen-wire
gen-wire:
	wire gen di/wire.go

