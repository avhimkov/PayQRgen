package main

import (
	"encoding/json"
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
func (c CoreBankPay) QRgenPayCore(paycore CoreBankPay) string {

	cbpmap := &CoreBankPay{
		Name:        paycore.Name,
		PersonalAcc: paycore.PersonalAcc,
		BankName:    paycore.BankName,
		BIC:         paycore.BIC,
		CorrespAcc:  paycore.CorrespAcc,
	}

	b, err := json.Marshal(cbpmap)
	if err != nil {
		fmt.Println(err)
		// return
	}

	bytes := []byte(b)

	// Unmarshal string into structs.
	var coreBankPay []CoreBankPay
	json.Unmarshal(bytes, &coreBankPay)

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
		fmt.Fprintf(buff, "%s=%s|", b)
		return buff.String()
	*/

	return string(bytes)
}

// QRgenPay - generate pay info
func (c ExtendBankPay) QRgenPayExt(extpay ExtendBankPay) string {

	ebpmap := &ExtendBankPay{
		Sum:          extpay.Sum,
		Purpose:      extpay.Purpose,
		PayeeINN:     extpay.PayeeINN,
		PayerINN:     extpay.PayerINN,
		DrawerStatus: extpay.DrawerStatus,
		KPP:          extpay.KPP,
		CBC:          extpay.CBC,
		OKTMO:        extpay.OKTMO,
		PaytReason:   extpay.PaytReason,
		TaxPeriod:    extpay.TaxPeriod,
		DocNo:        extpay.DocNo,
		DocDate:      extpay.DocDate,
		TaxPaytKind:  extpay.TaxPaytKind,
	}

	b, err := json.Marshal(ebpmap)
	if err != nil {
		fmt.Println(err)
		// return
	}

	return string(b)

}

// QRgenPay - generate pay info
func (c AnotherExtendBankPay) QRgenPayAnotExt(advextpay AnotherExtendBankPay) string {

	ebpmap := &AnotherExtendBankPay{
		LastName:        advextpay.LastName,
		FirstName:       advextpay.FirstName,
		MiddleName:      advextpay.MiddleName,
		PayerAddress:    advextpay.PayerAddress,
		PersonalAccount: advextpay.PersonalAccount,
		DocIdx:          advextpay.DocIdx,
		PensAcc:         advextpay.PensAcc,
		Contract:        advextpay.Contract,
		PersAcc:         advextpay.PersAcc,
		Flat:            advextpay.Flat,
		Phone:           advextpay.Phone,
		PayerIDType:     advextpay.PayerIDType,
		PayerIDNum:      advextpay.PayerIDNum,
		ChildFio:        advextpay.ChildFio,
		BirthDate:       advextpay.BirthDate,
		PaymTerm:        advextpay.PaymTerm,
		PaymPeriod:      advextpay.PaymPeriod,
		Category:        advextpay.Category,
		ServiceName:     advextpay.ServiceName,
		CounterID:       advextpay.CounterID,
		CounterVal:      advextpay.CounterVal,
		QuittID:         advextpay.QuittID,
		QuittDate:       advextpay.QuittDate,
		InstNum:         advextpay.InstNum,
		ClassNum:        advextpay.ClassNum,
		SpecFio:         advextpay.SpecFio,
		AddAmount:       advextpay.AddAmount,
		RuleID:          advextpay.RuleID,
		ExecID:          advextpay.ExecID,
		RegType:         advextpay.RegType,
		UIN:             advextpay.UIN,
		TechCode:        advextpay.TechCode,
	}

	b, err := json.Marshal(ebpmap)
	if err != nil {
		fmt.Println(err)
		// return
	}

	return string(b)
}
