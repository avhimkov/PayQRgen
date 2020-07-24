package main

import (
	"fmt"
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
func (c *CoreBankPay) QRgenPay() {
	fmt.Println(StdUtf8, "|", c.Name, "|", c.PersonalAcc, "|", c.BankName, "|", c.BIC, "|", c.CorrespAcc)

	// paycore := CoreBankPay{}

	paycore := CoreBankPay{
		ST:          StdUtf8,
		Name:        "ООО «Три кита»",
		PersonalAcc: "40702810138250123017",
		BankName:    "ОАО 'БАНК'",
		BIC:         "044525225",
		CorrespAcc:  "30101810400000000225",
	}

	var s []string
	for _, v := range paycore {
		s = append(s, v.Tag1, v.Tag2, v.Tag3)
	}

	// weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	// there is a space after comma
	fmt.Println(strings.Join(s, "|"))
}
