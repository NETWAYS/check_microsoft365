.PHONY: build

VERSION := $(shell git rev-parse HEAD)
DATE := $(shell date --iso-8601)

build:
	go build -ldflags "-X main.commit=$(VERSION) -X main.date=$(DATE)"
fmt:
	go fmt *.go
