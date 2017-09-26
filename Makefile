GOPATH?=$(shell go env GOPATH)
GOX=$(GOPATH)/bin/gox
GOLINT=$(GOPATH)/bin/golint
PREFIX?=/usr/local/bin
GOFILES=misto.go
BINARY_NAME?=misto

.PHONY: all
all: build

.PHONY: clean
clean:
	$(RM) -r build

build: build/$(BINARY_NAME)

build/$(BINARY_NAME): $(GOFILES)
	mkdir -p build
	go build -o $@ $?

$(GOX):
	go get github.com/mitchellh/gox

.PHONY: cross
cross: $(GOX)
	mkdir -p build
	gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"


.PHONY: clean-vendor
clean-vendor:
	$(RM) -r vendor

vendor:
	glide update

$(GOLINT):
	go get github.com/golang/lint/golint

.PHONY: lint
lint: $(GOFILES) $(GOLINT)
	golint $<

.PHONY: test
test:
	go test -v

.PHONY: install
install: $(PREFIX)/$(BINARY_NAME)

.PHONY: uninstall
uninstall:
	$(RM) $(PREFIX)/$(BINARY_NAME)

$(PREFIX)/$(BINARY_NAME): build/$(BINARY_NAME)
	mkdir -p $(PREFIX)
	mv -v $<  $@
