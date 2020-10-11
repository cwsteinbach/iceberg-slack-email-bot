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

# Running with all configs
$ SLACK_APP_TOKEN=<app_token> SLACK_BOT_USER_TOKEN=<bot_token> FROM=xxx@gmail.com TO=shlee0605@gmail.com SENDGRID_TOKEN=<sendgrid_token> PORT=5005 go run pinot-bot.go digest.go
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
$ docker run -a stdin -a stdout -i -t shlee0605/pinot-bot

# Run with all configs
$ docker run -p 5005:80 -e PORT=80 -e SLACK_APP_TOKEN=<slack_token> -e SLACK_BOT_USER_TOKEN=<bot_token> -e TO=xxx@gmail.com -e SENDGRID_TOKEN=<sendgrid_token> shlee0605/pinot-bot

# Publish docker image to docker hub
$ make push
```

# Set up with Gmail Account
1. Setup the app account by following: https://support.google.com/accounts/answer/185833?hl=en
2. Configure `MAIL_CLIENT_TYPE=GMAIL,GMAIL_ACCOUNT=<gmail_account>,GMAIL_APP_PASSWORD=<gmail_apppassword>`