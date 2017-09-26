IMAGE?=golang:1.8.3
PROJECT_PATH=/go/src/github.com/seemethere/misto
MOUNT="$(CURDIR):$(PROJECT_PATH)"
DOCKER_RUN=docker run --rm -it -v $(MOUNT) -w $(PROJECT_PATH) $(IMAGE)

%:
	$(DOCKER_RUN) make $@

vendor:
	$(DOCKER_RUN) sh -c 'curl https://glide.sh/get | sh; make $@'

.PHONY: shell
shell:
	$(DOCKER_RUN) bash
