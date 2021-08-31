// Package logic is the main logical package for the bot
package logic

import (
	"fmt"
	"github.com/almerlucke/go-iban/iban"
	financie "github.com/pieterclaerhout/go-finance"
	"strings"
)

func ConvertMoney(amount float64, convertFrom string, convertTo string) (float64, string) {
	//TODO: HESLO DIVE MAKY

	convertTo = strings.ToUpper(convertTo)

	converted, err := financie.ConvertRate(amount, convertFrom, convertTo)
	if err != nil {
		fmt.Println("ERROR CONV: ", err)
	}

	return converted, convertTo
}

// CheckIBAN This is a function to check whether an IBAN is valid at all
func CheckIBAN(ibanString string) (iban.IBAN, error) {

	ibanNumber, err := iban.NewIBAN(ibanString)

	if err != nil {
		fmt.Printf("%v\n", err)

	}
	return *ibanNumber, err
}

func CheckBankNameByCode(bankID string) string {

	//Updated in August 2021
	banks := make(map[string]string)
	banks["0100"] = "Komerční banka"
	banks["0300"] = "ČSOB"
	banks["0600"] = "MONETA"
	banks["0710"] = "Česká národní banka"
	banks["0800"] = "Česká spořitelna"
	banks["2010"] = "Fio banka"
	banks["2020"] = "MUFG Bank (Europe) N.V. Prague Branch"
	banks["2030"] = "AKCENTA, spořitelní a úvěrní družstvo"
	banks["2060"] = "Citfin, spořitelní družstvo"
	banks["2070"] = "test"
	banks["2100"] = "test"
	banks["2200"] = "test"
	banks["2220"] = "test"
	banks["2240"] = "test"
	banks["2250"] = "test"
	banks["2260"] = "test"
	banks["2310"] = "test"
	banks["2600"] = "test"
	banks["2700"] = "UniCredit Bank Czech Republic and Slovakia, a.s."
	banks["3030"] = "Air Bank a.s."
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"
	banks["2310"] = "test"

	//HESLO JE DIVE MAKY
	bankName := banks[bankID]

	return bankName
}

type NonIBANBankInfo struct {
	OriginalIBAN            string
	BankCode                string
	BankName                string
	AccountNumberPredcislie string
	AccountNumberMain       string
}

func ConvertIBANtoNumber(ibanString string) (NonIBANBankInfo, error) {
	var BankInfo NonIBANBankInfo

	VerifiedIBAN, err := CheckIBAN(ibanString)
	if err != nil {
		fmt.Println("ERROR VALIDATING IBAN: ", err)
		return BankInfo, err
	}

	BankInfo.OriginalIBAN = VerifiedIBAN.Code
	BankInfo.AccountNumberMain = VerifiedIBAN.BBAN[10:]
	BankInfo.AccountNumberPredcislie = VerifiedIBAN.BBAN[4:10]
	BankInfo.BankCode = VerifiedIBAN.BBAN[0:4]
	BankInfo.BankName = CheckBankNameByCode(BankInfo.BankCode)

	fmt.Printf("%+v\n", BankInfo)

	return BankInfo, err

}

func ConvertNumbertoIBAN(accountNumber string) {

}
