package main

import (
	"context"
	"github.com/go-joe/cron"
	"github.com/go-joe/http-server"
	"github.com/go-joe/joe"
	"github.com/go-joe/slack-adapter"
	"log"
	"os"
)

type PinotBot struct {
	*joe.Bot
	Config *Config
}

type DailyDigestEvent struct {
}

type Config struct {
	SlackAppToken     string
	SlackBotUserToken string
	From              string
	To                string
	SendgridToken     string
	Port              string
}

func NewConfig() (*Config, error) {
	config := &Config{
		SlackAppToken:     os.Getenv("SLACK_APP_TOKEN"),
		SlackBotUserToken: os.Getenv("SLACK_BOT_USER_TOKEN"),
		From:              os.Getenv("FROM"),
		To:                os.Getenv("TO"),
		SendgridToken:     os.Getenv("SENDGRID_TOKEN"),
		Port:              os.Getenv("PORT"),
	}
	log.Println(config)
	if config.Port == "" {
		config.Port = "80"
	}
	return config, nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	modules := []joe.Module {
		joehttp.Server(":" + config.Port),
		// Schedule the daily digest cron job at 2:00:00 AM (UTC)
		cron.ScheduleEvent("0 0 2 * * *", DailyDigestEvent{}),
	}
	if config.SlackAppToken != "" && config.SlackBotUserToken != ""  {
		modules = append(modules, slack.Adapter(config.SlackBotUserToken))
	}

	b := &PinotBot{
		Bot: joe.New("pinot-bot", modules...),
		Config: config,
	}

	// Register event handlers
	b.Brain.RegisterHandler(b.HandleDailyDigestEvent)
	b.Brain.RegisterHandler(b.HandleHTTP)
	b.Respond("daily-digest", b.DailyDigest)
	b.Respond("ping", Pong)

	b.Say("daily-digest", "Pinot bot is starting..")
	err = b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func (b *PinotBot) HandleDailyDigestEvent(evt DailyDigestEvent) {
	RunDailyDigest(b.Config)
	b.Say("daily-digest", DigestMessage())
}

func (b *PinotBot) DailyDigest(msg joe.Message) error {
	RunDailyDigest(b.Config)
	msg.Respond(DigestMessage())
	return nil
}

func Pong(msg joe.Message) error {
	msg.Respond("PONG")
	return nil
}

func (b *PinotBot) HandleHTTP(c context.Context, r joehttp.RequestEvent) {
	if r.URL.Path == "/" {
		b.Say("daily-digest", "Pinot bot is running..")
	}
}
