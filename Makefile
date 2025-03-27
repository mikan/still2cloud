.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Resolve dependencies using Go Modules
	go mod download

.PHONY: fmt
fmt: ## Format code with goimports
	goimports -w .

.PHONY: test
test: ## Run go test with coverage
	go test -v -cover .

.PHONY: build
build: ## Create cross-compile binaries
ifdef v # usage: make build v=v?.?.?
	$(eval VER := ${v})
else
	$(eval VER := $(shell git describe --tags))
endif
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.ver=$(VER)" -o build/still2cloud.linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -X main.ver=$(VER)" -o build/still2cloud.linux-arm64 .
	GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-s -w -X main.ver=$(VER)" -o build/still2cloud.linux-arm6 .
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w -X main.ver=$(VER)" -o build/still2cloud.linux-arm7 .
	bunzip2 -zf build/still2cloud.linux-amd64
	bunzip2 -zf build/still2cloud.linux-arm64
	bunzip2 -zf build/still2cloud.linux-arm6
	bunzip2 -zf build/still2cloud.linux-arm7

.PHONY: count-go
count-go: ## Count number of lines of all go codes
	find . -name "*.go" -type f | xargs wc -l | tail -n 1

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
