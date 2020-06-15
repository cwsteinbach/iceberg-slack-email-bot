package main

import (
	"bytes"
	"fmt"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"regexp"
	"strings"
	"time"
)
import "github.com/nlopes/slack"
import "github.com/sendgrid/sendgrid-go"

var re = regexp.MustCompile(`(?m)<@(\w+)>`)

const initialNumBufSize = 24

func DigestMessage() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	t := time.Now()
	t = t.In(loc)
	title := fmt.Sprintf("Apache Pinot Daily Email Digest (%s)", t.Format("2006-01-02"))
	return fmt.Sprintf("Daily digest sent with the title: `%s`", title)
}

func RunDailyDigest(config *Config) {
	if config.From == "" || config.To == "" || config.SendgridToken == "" {
		log.Println("")
	}

	api := slack.New(config.SlackAppToken)

	userList, err := Users(api)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	pm := &slack.GetConversationsParameters{
		ExcludeArchived: "true",
		Limit: 1000,
	}

	channels, _, err := api.GetConversations(pm)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	buffer := bytes.NewBuffer(make([]byte, 0, initialNumBufSize))

	for _, channel := range channels {
		if channel.Name == "daily-digest" {
			continue
		}

		ch := &slack.GetConversationHistoryParameters{
			ChannelID: channel.ID,
			Oldest: fmt.Sprintf("%f", float64(time.Now().Add(-24 * time.Hour).Unix())),
			Latest: fmt.Sprintf("%f", float64(time.Now().Unix())),
			Limit: 10000,
		}

		history, err := api.GetConversationHistory(ch)
		if err != nil {
			log.Println("failed to get conversation history: ", channel.Name)
		}
		if len(history.Messages) > 0 {
			buffer.WriteString(fmt.Sprintf("<h3><u>#%s</u></h3>", channel.Name))
			buffer.WriteString("<br>")
			for i := len(history.Messages) -1; i >=0; i-- {
				m := history.Messages[i]
				buffer.WriteString(fmt.Sprintf("<strong>%s: </strong>", userList[fmt.Sprintf("<@%s>", m.User)]))
				buffer.WriteString(ReplaceMentionUser(userList, m.Text))
				buffer.WriteString("<br>")
			}
		}
	}
	log.Println(string(buffer.Bytes()))
	t := time.Now()
	title := fmt.Sprintf("Apache Pinot Daily Email Digest (%s)", t.Format("2006-01-02"))
	SendGridEmail(config, title, string(buffer.Bytes()))
}

func Users(api *slack.Client) (map[string]string, error) {
	users, err := api.GetUsers()
	if err != nil {
		return nil, err
	}
	var userList = make(map[string]string)
	for _, user := range users {
		userList[fmt.Sprintf("<@%s>", user.ID)] = fmt.Sprintf("@%s", user.Name)
	}
	return userList, nil
}

func ReplaceMentionUser(userList map[string]string, text string) string {
	var msg = text
	for _, match := range re.FindAllString(msg, -1) {
		msg = strings.ReplaceAll(msg, match, userList[match])
	}
	return msg
}

func SendGridEmail(c *Config, subject string, htmlContent string) {
	from := mail.NewEmail("Pinot Slack Email Digest", c.From)
	client := sendgrid.NewSendClient(c.SendgridToken)
	toMails := strings.Split(c.To, ",")
	for _, toMail := range toMails {
		to := mail.NewEmail("Apache Pinot Dev", toMail)
		message := mail.NewSingleEmail(from, subject, to, htmlContent, htmlContent)
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(response.StatusCode)
			log.Println(response.Body)
			log.Println(response.Headers)
		}
	}
}
