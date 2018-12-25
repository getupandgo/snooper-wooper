SHELL := /bin/bash
TEAM = "gosenior"
SERVICE = "snooper-wooper"
VERSION = $(shell git describe --tags)
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

all: deps test

deps: go.mod go.sum
	@go mod download

test:
	@go test -v $(GOPACKAGES)

build-container:
	docker build -t ${TEAM}/${SERVICE}:${VERSION} .

push-container:
	docker push ${TEAM}/${SERVICE}:${VERSION}