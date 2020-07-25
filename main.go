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
	}

	fmt.Println(paycore.QRgenPay(paycore))

}
