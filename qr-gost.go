// Copyright 2021 Oleg Fomin, ofstudio@gmail.com

// Package qrGost генерирует QR-код платежа по банковским реквизитам по ГОСТ Р 56042-2014.
// См: https://files.stroyinf.ru/Data2/1/4293771/4293771168.htm
//
// QR-код может создаваться в одной из 3 кодировок: Windows1251, UTF8, KOI8R
//
package qrGost

import (
	"errors"
	"github.com/skip2/go-qrcode"
	"golang.org/x/text/encoding/charmap"
	"strconv"
)

type codePage int

const (
	Windows1251 codePage = 1
	UTF8        codePage = 2
	KOI8R       codePage = 3
)

var (
	encodeWindows1251 = charmap.Windows1251.NewEncoder()
	encodeKOI8R       = charmap.KOI8R.NewEncoder()
)

// Payment - платеж по банковским реквизитам
type Payment struct {
	Name        string // Наименование получателя платежа
	PersonalAcc string // Счет получателя платежа
	BankName    string // Наименование банка получателя платежа
	BIC         string // БИК банка получателя платежа
	CorrespAcc  string // Корреспондентский счет банка получателя платежа
	KPP         string // КПП получателя платежа либо 0, если нет или неизвестен. У ИП нет КПП
	PayeeINN    string // ИНН получателя платежа
	Purpose     string // Назначение платежа
	Sum         int    // Сумма платежа, ₽
}

// String - сериализует платежные данные в строку в указанной кодировке
func (p *Payment) String(c codePage) (string, error) {
	// ВАЖНО!
	// Для мобильного приложения Сбербанка необходимо,
	// чтобы поля в QR-коде шли строго в таком порядке
	var result = "ST0001" + strconv.Itoa(int(c)) +
		"|" + "Name=" + p.Name +
		"|" + "PersonalAcc=" + p.PersonalAcc +
		"|" + "BankName=" + p.BankName +
		"|" + "BIC=" + p.BIC +
		"|" + "CorrespAcc=" + p.CorrespAcc +
		"|" + "KPP=" + p.KPP +
		"|" + "PayeeINN=" + p.PayeeINN +
		"|" + "Purpose=" + p.Purpose +
		"|" + "Sum=" + strconv.Itoa(p.Sum*100)

	switch c {
	case Windows1251:
		return encodeWindows1251.String(result)
	case UTF8:
		return result, nil
	case KOI8R:
		return encodeKOI8R.String(result)
	default:
		return "", errors.New("unknown codepage")
	}
}

// Png - создает QR-код в формате png размера size пикселей
func (p *Payment) Png(c codePage, size int) ([]byte, error) {
	s, err := p.String(c)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(s, qrcode.Medium, size)
}

// PngFile - сохраняет файл с QR-кодом
func (p *Payment) PngFile(filename string, c codePage, size int) error {
	s, err := p.String(c)
	if err != nil {
		return err
	}
	return qrcode.WriteFile(s, qrcode.Medium, size, filename)
}
