// Package logic is the main logical package for the bot
package logic

import (
	"fmt"
	financie "github.com/pieterclaerhout/go-finance"
	"golang.org/x/text/currency"
)

func initialization() {
	fmt.Println()

	financie.ConvertRate(10, currency.TRY.String(), currency.AUD.String())
}
