IMAGE?=golang:1.8.3
PROJECT_PATH=/go/src/github.com/seemethere/misto
MOUNT="$(CURDIR):$(PROJECT_PATH)"
DOCKER_RUN=docker run --rm -i -v $(MOUNT) -w $(PROJECT_PATH) $(IMAGE)

.PHONY: clean
clean:
	$(DOCKER_RUN) make $@

build:
	$(DOCKER_RUN) make $@

.PHONY: clean-vendor
clean-vendor:
	$(DOCKER_RUN) make $@

vendor:
	$(DOCKER_RUN) sh -c 'curl https://glide.sh/get | sh; make $@'
