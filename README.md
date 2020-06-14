Pinot Bot
====

Simple Pinot Bot for Apache Pinot Slack Workspace.

```
# Install dependencies
$ go get -d ./...

# Running pinot-bot locally
$ go run pinot-bot.go digest.go
```

```
# Install gox (cross compilation tool)
$ go get github.com/mitchellh/gox
$ export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
$ make build
```