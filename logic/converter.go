// Package logic is the main logical package for the bot
package logic

import (
	"fmt"
	financie "github.com/pieterclaerhout/go-finance"
	"strings"
)

func ConvertMoney(amount float64, convertTo string) (float64, string) {

	var convertFrom string

	convertTo = strings.ToUpper(convertTo)
	if strings.Contains(convertTo, "CZK") == true {
		convertFrom = "EUR"
	} else {
		convertFrom = "CZK"
	}

	converted, err := financie.ConvertRate(amount, convertTo, convertFrom)
	if err != nil {
		fmt.Println("ERROR CONV: ", err)
	}

	return converted, convertFrom
}
