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

.PHONY: clean-vendor
clean-vendor:
	$(RM) -r vendor

vendor:
	glide update

lint: $(GOFILES)
	golint $?

.PHONY: test
test:
	go test -v
