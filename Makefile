.PHONY: test

default: test

test:
	go test ./...

fmt:
	go fmt ./...

build:
	builddir=$(shell mktemp -d)
	$(builddir)
	rmdir $(builddir)
