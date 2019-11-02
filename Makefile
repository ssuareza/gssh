EXECUTABLE=gssh
LINUX=$(EXECUTABLE)-linux-amd64
DARWIN=$(EXECUTABLE)-darwin-amd64
VERSION=$(shell git tag)

.PHONY: all clean

all: build ## Build

build: linux darwin ## Build binaries
	@echo version: $(VERSION)

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/gssh

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/gssh

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'