// Package logic is the main logical package for the bot
package logic

import (
	"MedzernikTelegramPayBot/config"
	"fmt"
	"os/exec"
	"strconv"
)

type PaymentInfo struct {
	name   string
	amount string
	swift  string
	iban   string
}

func Ready(amount float64) {
	var amountString string = strconv.FormatFloat(amount, 'f', 2, 64)

	currentInfo := PaymentInfo{
		name:   config.Cfg.AccountInfo.Name,
		amount: amountString,
		swift:  config.Cfg.AccountInfo.Swift,
		iban:   config.Cfg.AccountInfo.Iban,
	}

	fmt.Println(currentInfo)
	GenerateCode(currentInfo)

}

func GenerateCode(currentInfo PaymentInfo) {
	args := []string{"logic/generateQR.py", currentInfo.amount, currentInfo.iban, currentInfo.swift, currentInfo.name}
	err := exec.Command("python", args...).Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}
