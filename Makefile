.PHONY: default build test install

default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-domainname ~/.tflint.d/plugins
