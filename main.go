package main

import (
	"MedzernikTelegramPayBot/config"
	"MedzernikTelegramPayBot/logic"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
func main() {
	config.Initialization()
	bot, err := tgbotapi.NewBotAPI(config.Cfg.Server.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	go logic.Ready()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /sayhi or /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			case "html":
				msg.ParseMode = "html"
				msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)

		}

	}

}
