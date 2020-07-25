package main

import (
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

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
func (c CoreBankPay) QRgenPay(paycore CoreBankPay) string {

	var core []CoreBankPay
	core = append(core, paycore)

	var s []string
	for _, v := range core {
		s = append(s, v.ST, v.Name, v.PersonalAcc, v.BankName, v.BIC, v.CorrespAcc)
	}

	return strings.Join(s, "|")
}
