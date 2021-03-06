package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// github.com/divan/qrlogo

func main() {

	cbpmap := &CoreBankPay{
		Name:        "ООО «Три кита»",
		PersonalAcc: "40702810138250123017",
		BankName:    "ОАО \"БАНК\"",
		BIC:         "044525225",
		CorrespAcc:  "",
	}

	extpay := ExtendBankPay{
		Sum:          "100000",
		Purpose:      "Оплата членского взноса",
		PayeeINN:     "6200098765",
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

	anothextpay := AnotherExtendBankPay{
		LastName:        "Иванов",
		FirstName:       "Иван",
		MiddleName:      "Иванович",
		PayerAddress:    "г.Рязань ул.Ленина д.10 кв.15",
		PersonalAccount: "",
		DocIdx:          "",
		PensAcc:         "",
		Contract:        "",
		PersAcc:         "",
		Flat:            "",
		Phone:           "79101234567",
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
	}

	core := cbpmap.SelectFields("Name", "BankName")

	ext := extpay.SelectFields("Sum", "PayeeINN")

	anoth := anothextpay.SelectFields("LastName", "Phone")

	cod := StdUtf8 + "|" + CreateKeyValuePairs(core) + CreateKeyValuePairs(ext) + CreateKeyValuePairs(anoth)

	fmt.Print(cod + "\n")

	QRgen(cod, "gen/qr.png")

	QRreader("gen/qr.png")

	/*
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run()
	*/
	g := SetupRouter()
	g.GET("/", indexPageGet)
	g.POST("/", indexPagePost)

	// 404 page
	g.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{})
	})

	// Start serving the application
	g.Run(":3000")

}
