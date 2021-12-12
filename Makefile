SRC=$(shell find . -name "*.go")

ifeq (, $(shell which golangci-lint))
$(warning "Could not find golangci-lint in PATH, run: asdf plugin add golangci-lint && asdf install golangci-lint latest && asdf global golangci-lint latest")
endif

ifeq (, $(shell which richgo))
$(warning "Could not find richgo in PATH, run: asdf plugin add richgo && asdf install richgo latest && asdf global richgo latest")
endif

ifeq (, $(shell which goreleaser))
$(warning "Could not find goreleaser in PATH, run: asdf plugin add goreleaser && asdf install goreleaser latest && asdf global goreleaser latest")
endif

.PHONY: install lint test build

default: all

all: lint test release

install:
	@go mod tidy

lint:
	@golangci-lint run

test: install lint
	$(info ******************** running tests ********************)
	@richgo test -v ./...

build:
	@go build cmd/main.go -o ./dist/gwatch
	@cp ./dist/gwatch /usr/local/bin/gwatch
