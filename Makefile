.PHONY: default check test lint lint-gofmt build install

default: build

check: test lint

test:
	go test ./...

lint: lint-gofmt

lint-gofmt:
	gofmt -l .

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-domainname ~/.tflint.d/plugins
