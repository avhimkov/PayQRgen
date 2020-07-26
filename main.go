package main

import "fmt"

// github.com/divan/qrlogo

// 1 (соответствует Windows-1251) – будет выводиться ST00011;
// 2 (соответствует UTF-8) – будет выводиться ST00012.

func main() {
	QRgen("http://yandex.ru", "gen/qr.png")
	QRreader("gen/qr.png")

	paycore := CoreBankPay{
		ST:          StdUtf8,
		Name:        "ООО «Три кита»",
		PersonalAcc: "40702810138250123017",
		BankName:    "ОАО 'БАНК'",
		BIC:         "044525225",
		CorrespAcc:  "30101810400000000225",

		/*
			Name:        Field{"Name=", "ООО «Три кита»"},
			PersonalAcc: Field{"PersonalAcc=", "40702810138250123017"},
			BankName:    Field{"BankName=", "ОАО 'БАНК'"},
			BIC:         Field{"BIC=", "044525225"},
			CorrespAcc:  Field{"CorrespAcc=", "30101810400000000225"},
		*/
	}

	fmt.Println(StdUtf8 + "|" + paycore.QRgenPay(paycore))

}
