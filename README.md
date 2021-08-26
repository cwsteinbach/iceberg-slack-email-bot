Pinot Bot
====
![Go](https://github.com/snleee/pinot-bot/workflows/Go/badge.svg)

Simple Pinot Bot for Apache Pinot Slack Workspace.

# Build and run locally
```
# Initialize the module
$ go mod init iceberg-bot

# Install dependencies
$ go get -d ./...

# Running iceberg-bot locally
$ go run iceberg-bot.go digest.go

# Running with all configs
$ SLACK_BOT_USER_TOKEN=<>  FROM=iceberg.slack.bot@gmail.com TO=dev@iceberg.apache.org MAIL_CLIENT_TYPE=gmail GMAIL_ACCOUNT=iceberg.slack.bot GMAIL_APP_PASSWORD=<> PORT=8989 go run iceberg-bot.go digest.go
```

# Build docker image
```
# Install gox (cross compilation tool)
$ go get github.com/mitchellh/gox

# Add GOROOT/GOPATH to the environmental path.
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Run makefile script to build docker image
$ make build

# Run locally
$ docker run -a stdin -a stdout -i -t cwsteinbach/iceberg-bot

# Run with all configs
$ docker run -p 5005:80 -e PORT=80 -e SLACK_BOT_USER_TOKEN=<bot_token> -e TO=xxx@gmail.com cwsteinbach/iceberg-bot

# Publish docker image to docker hub
$ make push
```

# Set up with Gmail Account
1. Setup the app account by following: https://support.google.com/accounts/answer/185833?hl=en
2. Configure `MAIL_CLIENT_TYPE=GMAIL,GMAIL_ACCOUNT=<gmail_account>,GMAIL_APP_PASSWORD=<gmail_apppassword>`
