package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"os"
	"reflect"

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

// QRgenPayCore - generate pay info
func (c *CoreBankPay) QRgenPayCore(core *CoreBankPay) string {

	coreBankPay := core

	data, err := json.Marshal(coreBankPay)
	if err != nil {
		fmt.Println("An error occured: %v", err)
		os.Exit(1)
	}

	// fmt.Println(string(data))

	// var bytes bytes.Buffer
	/* buff := new(bytes.Buffer)

	s := reflect.ValueOf(c).Elem()

	for i := 0; i < s.NumField(); i++ {
		nameField := s.Type().Field(i).Name
		valueField := s.Field(i).Interface()
		// typeField := s.Type().Field(i).Type
		// Field := s.Type().Field(i)
		// varTag := Field.Tag

		fmt.Fprintf(buff, "%v=%v|", nameField, valueField)
	}
	*/
	/*
		b, err := json.Marshal(c)
		if err != nil {
			fmt.Println(err)
			// return
		}
	*/

	/*
		bytes := []byte(b)

		// Unmarshal string into structs.
		var coreBankPay []CoreBankPay
		json.Unmarshal(bytes, &coreBankPay)
	*/

	// Loop over structs and display them.

	/*
		b := []byte(bytes)
		err1 := json.Unmarshal(b, &cbpmap)
		if err1 != nil {
			panic(err1)
		}
	*/

	return string(data)
	// return buff.String()
}

// QRgenPayExt - generate pay info
func (e ExtendBankPay) QRgenPayExt() string {

	buff := new(bytes.Buffer)

	s := reflect.ValueOf(e).Elem()

	for i := 0; i < s.NumField(); i++ {
		nameField := s.Type().Field(i).Name
		valueField := s.Field(i).Interface()

		fmt.Fprintf(buff, "%v=%v|", nameField, valueField)
	}

	return buff.String()
}

// QRgenPayAnotExt - generate pay info
func (a AnotherExtendBankPay) QRgenPayAnotExt() string {

	buff := new(bytes.Buffer)

	s := reflect.ValueOf(a).Elem()

	for i := 0; i < s.NumField(); i++ {
		nameField := s.Type().Field(i).Name
		valueField := s.Field(i).Interface()

		fmt.Fprintf(buff, "%v=%v|", nameField, valueField)
	}

	return buff.String()
}
