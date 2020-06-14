package main

import (
	"github.com/go-joe/cron"
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
	SlackAppToken    string `yaml:"slack_app_token"`
	SlackBotUserToken string `yaml:"slack_bot_user_token""`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
	Sendgrid string `yaml:"sendgrid"`
}

func NewConfig() (*Config, error) {
	config := &Config{
		SlackAppToken: os.Getenv("SLACK_APP_TOKEN"),
		SlackBotUserToken: os.Getenv("SLACK_BOT_USER_TOKEN"),
		From: os.Getenv("FROM"),
		To: os.Getenv("TO"),
		Sendgrid: os.Getenv("SENDGRID_TOKEN"),
	}
	return config, nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	b := &PinotBot{
		Bot: joe.New("pinot-bot",
			slack.Adapter(config.SlackBotUserToken),

			// Schedule the daily digest cron job at 9:00:00 am
			cron.ScheduleEvent("0 0 9 * * *", DailyDigestEvent{}),
		),
		Config: config,
	}

	// Register event handlers
	b.Brain.RegisterHandler(b.HandleDailyDigestEvent)
	b.Respond("daily-digest", b.DailyDigest)
	b.Respond("ping", Pong)

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
