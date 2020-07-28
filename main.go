package main

import "fmt"

// github.com/divan/qrlogo

// 1 (соответствует Windows-1251) – будет выводиться ST00011;
// 2 (соответствует UTF-8) – будет выводиться ST00012.

func main() {
	// QRgen("http://yandex.ru", "gen/qr.png")

	cbpmap := CoreBankPay{
		Name:        "ООО «Три кита»",
		PersonalAcc: "40702810138250123017",
		BankName:    "ОАО 'БАНК'",
		BIC:         "044525225",
		CorrespAcc:  "30101810400000000225",
	}

	/* 	extpay := ExtendBankPay{
	   		Sum:          "",
	   		Purpose:      "",
	   		PayeeINN:     "",
	   		PayerINN:     "",
	   		DrawerStatus: "",
	   		KPP:          "",
	   		CBC:          "",
	   		OKTMO:        "",
	   		PaytReason:   "",
	   		TaxPeriod:    "",
	   		DocNo:        "",
	   		DocDate:      "",
	   		TaxPaytKind:  "",
	   	}

	   	anotextpay := AnotherExtendBankPay{
	   		LastName:        "",
	   		FirstName:       "",
	   		MiddleName:      "",
	   		PayerAddress:    "",
	   		PersonalAccount: "",
	   		DocIdx:          "",
	   		PensAcc:         "",
	   		Contract:        "",
	   		PersAcc:         "",
	   		Flat:            "",
	   		Phone:           "",
	   		PayerIDType:     "",
	   		PayerIDNum:      "",
	   		ChildFio:        "",
	   		BirthDate:       "",
	   		PaymTerm:        "",
	   		PaymPeriod:      "",
	   		Category:        "",
	   		ServiceName:     "",
	   		CounterID:       "",
	   		CounterVal:      "",
	   		QuittID:         "",
	   		QuittDate:       "",
	   		InstNum:         "",
	   		ClassNum:        "",
	   		SpecFio:         "",
	   		AddAmount:       "",
	   		RuleID:          "",
	   		ExecID:          "",
	   		RegType:         "",
	   		UIN:             "",
	   		TechCode:        "",
	   	} */

	// fmt.Println(StdUtf8 + "|" + paycore.QRgenPayCore(paycore) + "|" + extpay.QRgenPayExt(extpay) + "|" + anotextpay.QRgenPayAnotExt(anotextpay))
	// QRgen(StdUtf8+"|"+paycore.QRgenPayCore(paycore)+extpay.QRgenPayExt(extpay)+anotextpay.QRgenPayAnotExt(anotextpay), "gen/qr.png")
	// fmt.Println(paycore.QRgenPayCore(paycore))
	// fmt.Println(extpay.QRgenPayExt(extpay))
	// fmt.Println(anotextpay.QRgenPayAnotExt(anotextpay))

	// QRreader("gen/qr.png")

	// fmt.Println(StdUtf8 + "|" + cbpmap.QRgenPayCore(cbpmap) + extpay.QRgenPayExt(extpay) + anotextpay.QRgenPayAnotExt(anotextpay))
	fmt.Println(cbpmap.QRgenPayCore(cbpmap))
}
