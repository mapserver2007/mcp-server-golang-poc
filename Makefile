GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: install
install:
	go get github.com/mark3labs/mcp-go
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: lint
lint:
	goimports -w .