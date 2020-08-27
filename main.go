package main

import (
	"fmt"
)

// github.com/divan/qrlogo

func main() {
	// QRgen("http://yandex.ru", "gen/qr.png")

	cbpmap := &CoreBankPay{
		Name:        "ООО «Три кита»",
		PersonalAcc: "40702810138250123017",
		BankName:    "ОАО \"БАНК\"",
		BIC:         "044525225",
		CorrespAcc:  "",
	}

	/* 	extpay := ExtendBankPay{
	   		Sum:      "100000",
	   		Purpose:  "Оплата членского взноса",
	   		PayeeINN: "6200098765",
	   		// PayerINN:     "",
	   		// DrawerStatus: "",
	   		// KPP:          "",
	   		// CBC:          "",
	   		// OKTMO:        "",
	   		// PaytReason:   "",
	   		// TaxPeriod:    "",
	   		// DocNo:        "",
	   		// DocDate:      "",
	   		// TaxPaytKind:  "",
	   	}

	   	anotextpay := AnotherExtendBankPay{
	   		LastName:     "Иванов",
	   		FirstName:    "Иван",
	   		MiddleName:   "Иванович",
	   		PayerAddress: "г.Рязань ул.Ленина д.10 кв.15",
	   		// PersonalAccount: "",
	   		// DocIdx:          "",
	   		// PensAcc:         "",
	   		// Contract:        "",
	   		// PersAcc:         "",
	   		// Flat:            "",
	   		Phone: "79101234567",
	   		// PayerIDType:     "",
	   		// PayerIDNum:      "",
	   		// ChildFio:        "",
	   		// BirthDate:       "",
	   		// PaymTerm:        "",
	   		// PaymPeriod:      "",
	   		// Category:        "",
	   		// ServiceName:     "",
	   		// CounterID:       "",
	   		// CounterVal:      "",
	   		// QuittID:         "",
	   		// QuittDate:       "",
	   		// InstNum:         "",
	   		// ClassNum:        "",
	   		// SpecFio:         "",
	   		// AddAmount:       "",
	   		// RuleID:          "",
	   		// ExecID:          "",
	   		// RegType:         "",
	   		// UIN:             "",
	   		// TechCode:        "",
	   	}
	*/

	// QRgen(StdUtf8+"|"+paycore.QRgenPayCore(paycore)+extpay.QRgenPayExt(extpay)+anotextpay.QRgenPayAnotExt(anotextpay), "gen/qr.png")

	// QRreader("gen/qr.png")

	fmt.Println(StdUtf8 + "|" + cbpmap.QRgenPayCore(cbpmap))

	d := cbpmap.SelectFields("Name", "BankName")

	fmt.Print(StdUtf8 + "|" + CreateKeyValuePairs(d))

}
