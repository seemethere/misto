BINARY_NAME?=misto

.PHONY: all
all: build

.PHONY: clean
clean:
	$(RM) -r build

build: build/$(BINARY_NAME)

build/$(BINARY_NAME): *.go
	mkdir -p build
	go build -o $@ $?

.PHONY: clean-vendor
clean-vendor:
	$(RM) -r vendor

vendor:
	glide update

lint:
	gometalinter --vendor ./...
