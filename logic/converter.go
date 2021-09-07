// Package logic is the main logical package for the bot
package logic

import (
	"fmt"
	"github.com/almerlucke/go-iban/iban"
	"github.com/hrharder/go-gas"
	financie "github.com/pieterclaerhout/go-finance"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
	"strconv"
	"strings"
)

func ConvertMoney(amount float64, convertFrom string, convertTo string) (float64, string) {

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

// CheckBankNameByCode Checks the bankname by the code of the bank (a map)
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

// ConvertIBANtoNumber Converts the IBAN to a number
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

// ConvertNumberToIBAN TODO: finish this function
func ConvertNumberToIBAN(accountNumber string) string {

	accountNumberProcessed := strings.ReplaceAll(accountNumber, "/", "")

	return accountNumberProcessed

}

func GetCurrentGasPrice() []string {
	// get a gas price in base units with one of the exported priorities (fast, fastest, safeLow, average)
	fastestGasPrice, err := gas.SuggestGasPrice(gas.GasPriorityFastest)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	// convenience wrapper for getting the fast gas price
	fastGasPrice, err := gas.SuggestFastGasPrice()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	safeGasPrice, err := gas.SuggestGasPrice(gas.GasPrioritySafeLow)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	averageGasPrice, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	ethPrice := GetETHCoinPrice()

	fastestGasPriceFloatUSD := (float64(fastestGasPrice.Uint64()) * 0.00000000000001) * ethPrice
	fastGasPriceFloatUSD := (float64(fastGasPrice.Uint64()) * 0.00000000000001) * ethPrice
	safeGasPriceFloatUSD := (float64(safeGasPrice.Uint64()) * 0.00000000000001) * ethPrice
	averageGasPriceFloatUSD := (float64(averageGasPrice.Uint64()) * 0.00000000000001) * ethPrice

	var output []string
	output = append(output, "Fastest gas price: \t"+strconv.FormatFloat(fastestGasPriceFloatUSD, 'f', 2, 64)+" USD")
	output = append(output, "Fast gas price: \t"+strconv.FormatFloat(fastGasPriceFloatUSD, 'f', 2, 64)+" USD")
	output = append(output, "Safe low gas price: \t"+strconv.FormatFloat(safeGasPriceFloatUSD, 'f', 2, 64)+" USD")
	output = append(output, "Average gas price: \t"+strconv.FormatFloat(averageGasPriceFloatUSD, 'f', 2, 64)+" USD")
	return output

}

func GetETHCoinPrice() float64 {

	cg := gecko.NewClient(nil)

	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "eur"}
	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		log.Fatal(err)
	}
	eth := (*sp)["ethereum"]

	ethUSD, err := strconv.ParseFloat(fmt.Sprintf("%f", eth["usd"]), 64)
	if err != nil {
		fmt.Println("ERROR CONVERTING STRING TO FLOAT: ", err)
		return 0
	}

	return ethUSD
}
