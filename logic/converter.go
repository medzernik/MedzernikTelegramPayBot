// Package logic is the main logical package for the bot
package logic

import (
	"fmt"
	"github.com/almerlucke/go-iban/iban"
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

// CheckIBAN This is a function to check whether an IBAN is valid at all
func CheckIBAN(ibanString string) (iban.IBAN, error) {

	ibanNumber, err := iban.NewIBAN(ibanString)

	if err != nil {
		fmt.Printf("%v\n", err)

	} else {
		fmt.Printf("%v\n", ibanNumber.PrintCode)
		fmt.Printf("%v\n", ibanNumber.BBAN)

	}
	return *ibanNumber, err
}

func ConvertIBANtoNumber(ibanString string) error {
	test, err := CheckIBAN(ibanString)
	if err != nil {
		fmt.Println("ERROR VALIDATING IBAN: ", err)
		return err
	}

	//na bankcode vytvorit mapu bank s cisly
	bankCode := test.BBAN[0:4]
	accountNumberPredcisli := test.BBAN[4:10]
	accountNumberMain := test.BBAN[10:]

	fmt.Println(bankCode, accountNumberPredcisli, accountNumberMain)

	return err

}

func ConvertNumbertoIBAN(accountNumber string) {

}
