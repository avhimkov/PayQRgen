package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"

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

	cbpmap := map[string]string{
		"Name":        paycore.Name,
		"PersonalAcc": paycore.PersonalAcc,
		"BankName":    paycore.BankName,
		"BIC":         paycore.BIC,
		"CorrespAcc":  paycore.CorrespAcc,
	}

	b := new(bytes.Buffer)

	for k, v := range cbpmap {
		fmt.Fprintf(b, "%s=%s|", k, v)
	}

	return b.String()
}
