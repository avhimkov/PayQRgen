package main

import (
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

// GetCoreBankPay - generate pay info
func (c *CoreBankPay) GetCoreBankPay() string {
	return StdUtf8 + "|" + c.Name + "|" + c.PersonalAcc + "|" + c.BankName + "|" + c.BIC + "|" + c.CorrespAcc
}

// QRgenPayCore - generate pay info
func (c CoreBankPay) QRgenPayCore() string {

	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		// return
	}

	// example 1
	core := CoreBankPay{}
	e := reflect.ValueOf(&core).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
	// example 2
	s := reflect.ValueOf(&core).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	// The output of this program is
	//
	// 0: A int = 23
	// 1: B string = skidoo

	/*
		bytes := []byte(b)

		// Unmarshal string into structs.
		var coreBankPay []CoreBankPay
		json.Unmarshal(bytes, &coreBankPay)
	*/

	// Loop over structs and display them.

	/*
		for c := range coreBankPay {
		fmt.Printf("%v=%v|", coreBankPay[c], coreBankPay[c])
		fmt.Println()
		}
	*/

	// New Buffer.

	/*
	   buff := new(bytes.Buffer)
	   	fmt.Fprintf(buff, "%s=%s|", cbpmap)
	   	return buff.String()
	*/

	/*
		b := []byte(bytes)
		err1 := json.Unmarshal(b, &cbpmap)
		if err1 != nil {
			panic(err1)
		}
	*/
	// return core
	return string(bytes)
}

// QRgenPayExt - generate pay info
func (e ExtendBankPay) QRgenPayExt() string {

	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		// return
	}

	return string(b)

}

// QRgenPayAnotExt - generate pay info
func (a AnotherExtendBankPay) QRgenPayAnotExt() string {

	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		// return
	}

	return string(b)
}
