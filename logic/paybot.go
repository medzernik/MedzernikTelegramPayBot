// Package logic Logic is the main logical package for the bot
package logic

import (
	"fmt"
	pay "github.com/jovandeginste/payme/payment"
)

func Ready() {
	fmt.Println()

	var myTempAccount = pay.Payment{
		ServiceTag:             "BCD",
		Version:                2,
		CharacterSet:           1,
		IdentificationCode:     "SCT",
		BICBeneficiary:         "",
		NameBeneficiary:        "TestName",
		IBANBeneficiary:        "SK3327182322399811163967",
		EuroAmount:             10,
		Purpose:                "",
		Remittance:             "test",
		B2OInformation:         "",
		RemittanceIsStructured: false,
	}

	test, err := myTempAccount.ToQRPNG(64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(test)

}
