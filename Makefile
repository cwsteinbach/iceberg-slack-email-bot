SHELL=/bin/bash

IMAGE=pinot-bot
VERSION=0.1

.PHONY: build
build:
	mkdir -p build
	gox -osarch="linux/amd64" --output="build/pinot-bot"
	docker build -t $(IMAGE):$(VERSION) .
	rm -rf build

.PHONY: push
push:
	docker push $(IMAGE):$(VERSION)
