package qrGost

import "fmt"

var p = Payment{
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

func ExamplePayment_String() {
	s, _ := p.String(UTF8)
	fmt.Println(s)
	// Output:
	// ST00012|Name=АНО "Танцующий Петербург"|PersonalAcc=40703810990200000032|BankName=ПАО "БАНК "САНКТ-ПЕТЕРБУРГ"|BIC=044030790|CorrespAcc=30101810900000000790|KPP=780201001|PayeeINN=7802532605|Purpose=Оплата за участие в Swing & Snow 2022 REG 6821320/1106|Sum=225000
}

func ExamplePayment_String_win1251() {
	s, _ := p.String(Windows1251)
	fmt.Printf("%x", s)
	// Output:
	// 535430303031317c4e616d653dc0cdce2022d2e0edf6f3fef9e8e920cfe5f2e5f0e1f3f0e3227c506572736f6e616c4163633d34303730333831303939303230303030303033327c42616e6b4e616d653dcfc0ce2022c1c0cdca2022d1c0cdcad22dcfc5d2c5d0c1d3d0c3227c4249433d3034343033303739307c436f72726573704163633d33303130313831303930303030303030303739307c4b50503d3738303230313030317c5061796565494e4e3d373830323533323630357c507572706f73653dceefebe0f2e020e7e020f3f7e0f1f2e8e520e2205377696e67202620536e6f7720323032322052454720363832313332302f313130367c53756d3d323235303030
}

func ExamplePayment_String_koi8r() {
	s, _ := p.String(KOI8R)
	fmt.Printf("%x", s)
	// Output:
	// 535430303031337c4e616d653de1eeef2022f4c1cec3d5c0ddc9ca20f0c5d4c5d2c2d5d2c7227c506572736f6e616c4163633d34303730333831303939303230303030303033327c42616e6b4e616d653df0e1ef2022e2e1eeeb2022f3e1eeebf42df0e5f4e5f2e2f5f2e7227c4249433d3034343033303739307c436f72726573704163633d33303130313831303930303030303030303739307c4b50503d3738303230313030317c5061796565494e4e3d373830323533323630357c507572706f73653defd0ccc1d4c120dac120d5dec1d3d4c9c520d7205377696e67202620536e6f7720323032322052454720363832313332302f313130367c53756d3d323235303030
}
