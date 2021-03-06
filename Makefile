ifeq ($(shell uname -s), Darwin)
    shasum=shasum -a256
else
    shasum=sha256sum
endif

repo=github.com/ssuareza/gssh
version=$(shell git describe --all --dirty --long | awk -F"-|/" '/^heads/ {print $$2 "-" substr($$4, 2) "-" $$5}; /^tags/ { print $$2 }')
build_args=-ldflags "-X main.versionString=$(version)" ./cmd/gssh
files=$(shell find cmd -type f)

.PHONY: test

all: test build checksums

build: build-linux build-darwin

build-linux: build/gssh-$(version)-linux-amd64

build/gssh-$(version)-linux-amd64: ${files}
	GOARCH=amd64 GOOS=linux go build -o $@ $(build_args)

build-darwin: build/gssh-$(version)-darwin-amd64
build/gssh-$(version)-darwin-amd64: ${files}
	GOARCH=amd64 GOOS=darwin go build -o $@ $(build_args)

checksums: build
	cd build/ && ${shasum} * > $(version)-SHA256SUMS

test:
	@go test -v pkg/gssh/*