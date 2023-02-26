#!/usr/bin/env bash

BUILD_VERSION := $(shell git describe --always --tags)
BUILD_TIME=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
DOCKER_IMAGE_NAME="SamiraAfrin/To-Do-App"
BINARY_NAME=to-do-app
BIN_OUT_DIR=bin
GO_VERSION=1.19

export PATH=$(shell go env GOPATH)/bin:$(shell echo $$PATH)

.PHONY: all

all: dl-deps build test-unit ## Build binary (with unit tests)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: build ## Run lint checks
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(shell go env GOPATH)/bin/golangci-lint run

fmt: ## Refactor go files with gofmt and goimports
	go install golang.org/x/tools/cmd/goimports@latest
	find . -name '*.go' | while read -r file; do goimports -w "$$file"; done

test:  ## Run unit tests
	go test -v -coverprofile=coverage.txt -covermode=atomic -cover ./...


clean: ## Cleans output directory
	$(shell rm -rf $(BIN_OUT_DIR)/*)

dl-deps: ## Get dependencies
	go mod vendor

build: clean ## Build binary
	go build -v -o $(BIN_OUT_DIR)/$(BINARY_NAME)

serve: build ## Run http server
	./$(BIN_OUT_DIR)/$(BINARY_NAME)

docker-build: ## Build docker image
	docker build --build-arg GO_VERSION=${GO_VERSION} --build-arg BUILD_VERSION=${BUILD_VERSION} --build-arg BUILD_TIME=${BUILD_TIME} --tag ${DOCKER_IMAGE_NAME} .

docker-push: ## Push docker image
	docker push

docker-run: ## Run docker image
	docker run --name todoapp --rm -it -p 8000:8000 \
		-v $$(pwd)/config.json:/config.json \
		$(DOCKER_IMAGE_NAME):latest

