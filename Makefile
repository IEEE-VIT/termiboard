# Build variables
PROJECT_NAME = $(shell basename "$(PWD)")
BASE=$(shell pwd)
BUILD_DIR=$(BASE)/bin
VERSION ?= $(shell git tag --points-at HEAD | tail -n 1)
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT_SHA = $(shell git rev-parse --short HEAD)
LDFLAGS = -ldflags "-w -X main.version=${VERSION} -X main.buildDate=${BUILD_DATE} -X main.commit=${COMMIT_SHA}"
PACKAGE = $(shell go list -m)

.PHONY: dep
dep: ## Install dependencies
	@echo "  >  Install dependencies"
	@go mod tidy
	@go mod download
	@go mod vendor

.PHONY: test
test: dep ## Run unit tests
	@echo "  >  Run tests"
	@go test -v ./...

.PHONY: build
build: test ## Build a binary executable file
	@echo "  >  Build binary"
	go build ${LDFLAGS} -o ${BUILD_DIR}/${PROJECT_NAME}

clean: ## Clean build cache
	@echo "  >  Cleaning build cache"
	go clean

.PHONY: run
run: ## Run current version
	./bin/${PROJECT_NAME}

.PHONY: rerun
rerun: build run ## Build and Run

.PHONY: printvars
printvars: ## Print variables (including Makefile)
	$(foreach v, $(filter-out .VARIABLES,$(.VARIABLES)), $(info $(v) = $($(v))))
	@echo "  >  Print vars done"

.PHONY: help
.DEFAULT_GOAL := help
help: ## Print this message
	@echo "Makefile usage help:"
	@echo "------"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

