SHELL=/bin/bash

IMAGE=shlee0605/pinot-bot

.PHONY: build
build:
	mkdir -p build
	gox -osarch="linux/amd64" --output="build/pinot-bot"
	docker build -t $(IMAGE) .
	rm -rf build

.PHONY: push
push:
	docker push $(IMAGE)
