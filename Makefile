.PHONY: all

GIT_VERSION = $(shell git describe --tags --dirty)
VERSION ?= $(GIT_VERSION)

OS := $(shell uname)
SHELL := /bin/bash
BASE = $(GOPATH)/src/github.com/mahakamcloud/netd

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

fmt: ## Run go fmt against code
	go fmt ./...

vet: ## Run go vet against code
	go vet ./...

test: ## run tests
	@echo running tests...
	go test -v $(shell go list -v ./... | grep -v /vendor/ | grep -v integration | grep -v /playground )
