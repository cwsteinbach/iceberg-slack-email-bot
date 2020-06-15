Pinot Bot
====
![Go](https://github.com/snleee/pinot-bot/workflows/Go/badge.svg)

Simple Pinot Bot for Apache Pinot Slack Workspace.

# Build and run locally
```
# Install dependencies
$ go get -d ./...

# Running pinot-bot locally
$ go run pinot-bot.go digest.go
```

# Build docker image
```
# Install gox (cross compilation tool)
$ go get github.com/mitchellh/gox

# Add GOROOT/GOPATH to the environmental path.
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Run makefile script to build docker image
$ make build

# Publish docker image to docker hub
$ make push
```
