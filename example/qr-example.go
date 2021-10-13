package main

import (
	"fmt"
	qrGost "github.com/ofstudio/qr-gost-56042"
)

var p = qrGost.Payment{
	Name:        `АНО "Танцующий Петербург"`,
	PersonalAcc: `40703810990200000032`,
	BankName:    `ПАО "БАНК "САНКТ-ПЕТЕРБУРГ"`,
	BIC:         `044030790`,
	CorrespAcc:  `30101810900000000790`,
	KPP:         `780201001`,
	PayeeINN:    `7802532605`,
	Purpose:     `Оплата за участие в Swing & Snow 2022 REG 6821320/1106`,
	Sum:         2250,
}

func main() {
	var str string
	var err error

	if str, err = p.String(qrGost.UTF8); err != nil {
		panic(err)
	}
	fmt.Println(str)

	if err = p.PngFile("qr-example.png", qrGost.Windows1251, 512); err != nil {
		panic(err)
	}
}
