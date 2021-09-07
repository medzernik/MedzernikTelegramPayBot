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
	//logic.ConvertNumberToIBAN("5176795612/0900")
	logic.GetCurrentGasPrice()
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
			case "ethPrice":
				msg.Text = "Hi :)"
				bot.Send(msg)
			case "gas":
				price := logic.GetCurrentGasPrice()

				for _, i := range price {
					msg.Text += i + "\n"
				}
				bot.Send(msg)
			case "gas2min":
				//TODO: make a greenthread response parallel to the main catcher
				/*
					price := logic.GetCurrentGasPrice()

					for _, i := range price {
						msg.Text += i + "\n"
					}
					bot.Send(msg)

				*/
			case "ibantoacc":
				arguments := update.Message.CommandArguments()

				var ibanString string

				if len(arguments) < 2 {
					fmt.Println("Insufficient argument length")
					break

				} else {
					ibanString = strings.ToUpper(arguments)
				}
				fmt.Println(ibanString)

				ibanConverted, err1 := logic.ConvertIBANtoNumber(ibanString)
				if err1 != nil {
					msg.Text = err1.Error()
					bot.Send(msg)
					fmt.Println(err1)
					break
				}

				msg.Text = "Forenumber: \t" + ibanConverted.AccountNumberPredcislie + "\n" +
					"Number: \t" + ibanConverted.AccountNumberMain + "\n" +
					"Bank code: \t" + ibanConverted.BankCode + "\n" +
					"Bank name: \t" + ibanConverted.BankName + "\n" +
					"-----------\n" +
					"Copyfriendly: \t" + ibanConverted.AccountNumberPredcislie + ibanConverted.AccountNumberMain + "/" + ibanConverted.BankCode
				bot.Send(msg)
			case "conv":

				arguments := strings.Split(update.Message.CommandArguments(), " ")

				var currencyTo string
				if len(arguments) < 2 {
					fmt.Println("Insufficient argument length")
					break
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
					break
				}

			default:
				msg.Text = "Error in syntax"
				bot.Send(msg)
			}

		}

	}

}
