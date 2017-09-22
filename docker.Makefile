IMAGE?=golang:1.8.3
PROJECT_PATH=/go/src/github.com/seemethere/misto
MOUNT="$(CURDIR):$(PROJECT_PATH)"
DOCKER_RUN=docker run --rm -i -v $(MOUNT) -w $(PROJECT_PATH) $(IMAGE)

.PHONY: clean
clean:
	$(DOCKER_RUN) make $@

build:
	$(DOCKER_RUN) make $@
