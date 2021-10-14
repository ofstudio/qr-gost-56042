# qr-gost-56042
QR-код для оплаты по банковским реквизитам по 
[ГОСТ Р 56042-2014](https://files.stroyinf.ru/Data2/1/4293771/4293771168.htm).

Документация: [pkg.go.dev/github.com/ofstudio/qr-gost-56042](https://pkg.go.dev/github.com/ofstudio/qr-gost-56042)

```go
package main

import qrGost "github.com/ofstudio/qr-gost-56042"

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
	err := p.PngFile("qr-example.png", qrGost.Windows1251, 512) 
	if err != nil {
		panic(err)
	}
}
```

© Oleg Fomin 2021

Licence: MIT
