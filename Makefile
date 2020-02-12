# -*- mode:makefile-gmake -*-
GO111MODULE := on
REVISION := $(shell git rev-parse HEAD)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS := -s -w -X main.buildVersion=${REVISION} -X main.buildDate=${DATE}

OS := $(shell uname -s)
ifeq ($(OS),Darwin)
	TARGET := mac
else
	TARGET := linux
endif

.PHONY: build
build: $(TARGET) ## Build for local environment.

.PHONY: linux
linux: ## Build for Linux.
	GOOS=linux GOARCH=amd64 go build -o ./bin/kadm -ldflags "$(LDFLAGS)"

.PHONY: mac
mac: ## Build for macosx.
	go build -o ./bin/kadm -ldflags "$(LDFLAGS)"

.PHONY: install
install: ## install
	go install -ldflags "$(LDFLAGS)"

.PHONY: clean
clean: ## clean
	rm -f ./bin/kadm
	go mod tidy
	go clean -testcache

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this message.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
