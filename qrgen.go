package main

import (
	"bytes"
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

	buff := new(bytes.Buffer)

	s := reflect.ValueOf(core).Elem()

	for i := 0; i < s.NumField(); i++ {
		nameField := s.Type().Field(i).Name
		valueField := s.Field(i).Interface()
		// typeField := s.Type().Field(i).Type
		// Field := s.Type().Field(i)
		// varTag := Field.Tag

		fmt.Fprintf(buff, "%v=%v|", nameField, valueField)
	}

	// return string(data)
	return buff.String()
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

func fieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))
	for _, s := range fields {
		set[s] = true
	}
	return set
}

func (c *CoreBankPay) SelectFields(fields ...string) map[string]interface{} {
	fs := fieldSet(fields...)
	rt, rv := reflect.TypeOf(*c), reflect.ValueOf(*c)
	out := make(map[string]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")
		if fs[jsonKey] {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}
	// return buff.String()
	return out
}

func CreateKeyValuePairs(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=%s|", key, value)
	}
	return b.String()
}
