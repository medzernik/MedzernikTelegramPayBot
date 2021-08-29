// A custom telegram bot made to help with managing of finances while living in a country with a different currency than Euro (and still using Euro primarily).
package main

import (
	"MedzernikTelegramPayBot/config"
	"MedzernikTelegramPayBot/logic"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
func main() {
	config.Initialization()
	logic.ConvertIBANtoNumber("CZ54 0800 0000 0045 1328 7033")
	//logic.ConvertIBANtoNumber("SK7609000000005176795612")
	bot, err := tgbotapi.NewBotAPI(config.Cfg.Server.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Debug = true

	//log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	var amount float64 = 0.0

	go logic.Ready(amount)

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
				bot.Send(msg)
			case "sayhi":
				msg.Text = "Hi :)"
				bot.Send(msg)
			case "status":
				msg.Text = "I'm ok."
				bot.Send(msg)
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
				arguments := strings.Split(update.Message.CommandArguments(), " ")
				fmt.Println("ARGUMENT\n\n", arguments[0])
				bot.Send(msg)
			case "conv":
				//TODO: HESLO: ZASTAVENY CAS
				arguments := strings.Split(update.Message.CommandArguments(), " ")

				var currencyTo string
				if len(arguments) < 2 {
					fmt.Println("Insufficient argument length")
					return
				} else if len(arguments) == 2 && strings.ToUpper(arguments[1]) == "CZK" {
					currencyTo = "EUR"
				} else if len(arguments) == 2 && strings.ToUpper(arguments[1]) == "EUR" {
					currencyTo = "CZK"
				} else if len(arguments) > 2 {
					currencyTo = strings.ToUpper(arguments[2])
				}

				floatAmount, err1 := strconv.ParseFloat(arguments[0], 64)
				if err1 != nil {
					msg.Text = "Invalid amount format"
					bot.Send(msg)
					fmt.Println("Invalid amount format")
					break
				}

				currencyFrom := strings.ToUpper(arguments[1])

				converted, currency_conv := logic.ConvertMoney(floatAmount, currencyFrom, currencyTo)
				msg.Text = "Conversion is: " + currency_conv + " " + strconv.FormatFloat(converted, 'f', 2, 64) + " "
				bot.Send(msg)

			case "pay":
				arguments := strings.Split(update.Message.CommandArguments(), " ")
				floatAmount, err := strconv.ParseFloat(arguments[0], 64)
				if err != nil {
					msg.Text = "Invalid amount format"
					bot.Send(msg)
					fmt.Println("Invalid amount format")
					break
				}

				logic.Ready(floatAmount)
				msg.Text = "Paying MEDZERNIK " + arguments[0] + " EUROS INTO ACCOUNT. SCAN QR CODE IN BANK APP TO PAY."
				bot.Send(msg)

				msg1 := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "payCodeQR.png")

				_, err = bot.Send(msg1)
				if err != nil {
					fmt.Println("ERROR: ", err)
					return
				}

			default:
				msg.Text = "Error in syntax"
				bot.Send(msg)
			}

		}

	}

}
