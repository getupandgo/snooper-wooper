SHELL := /bin/bash
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

all: deps test

deps: go.mod go.sum
	@go mod download

test:
	@go test -v $(GOPACKAGES)