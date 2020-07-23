package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

//CoreBankPay - Обязательные реквизиты (блок «Payee» УФЭБС[5])
type CoreBankPay struct {
	Name        string // Наименование получателя платежа (Макс. 160 знаков (имя тега по [5] : Payee/ Name))
	PersonalAcc string // Номер счета получателя платежа (Макс. 20 знаков (имя тега по [5] : Payee/ PersonalAcc))
	BankName    string // Наименование банка получателя платежа (Макс. 45 знаков (не определен [5]))
	BIC         string // БИК (Макс. 9 знаков (имя тега по [5] : Payee/ Bank/ BIC))
	CorrespAcc  string // Номер кор./сч. банка получателя платежа (Макс. 20 знаков (имя тега по УФЭБС: Payee/ Bank/ CorrespAcc))
}

// StdUtf8 - Standart coding UTF8
const StdUtf8 = "ST00012"

// StdWin1251 - Standart coding Win1251
const StdWin1251 = "ST00011"

//QRgen - Create the qr, barcode
func QRgen(text, path string) error {
	// Create the barcode
	qrCode, _ := qr.Encode(text, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

	return err
}

// QRgenPay - generate pay info
func (c *CoreBankPay) QRgenPay() {
	fmt.Println(StdUtf8, "|", c.Name, "|", c.PersonalAcc, "|", c.BankName, "|", c.BIC, "|", c.CorrespAcc)
}
