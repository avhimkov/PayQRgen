package main

// ST          string // Идентификатор формата (формат «свой»)
// FormatVer   string // Версия формата
// CodeText    string // кодировка
// ST          string // Служебный блок

// 1 (соответствует Windows-1251) – будет выводиться ST00011;
// 2 (соответствует UTF-8) – будет выводиться ST00012.

// StdUtf8 - Standart coding UTF8
const StdUtf8 = "ST00012"

// StdWin1251 - Standart coding Win1251
const StdWin1251 = "ST00011"

// CoreBankPay - Обязательные реквизиты (блок «Payee» УФЭБС[5])
type CoreBankPay struct {
	Name        string `json:"Name,omitempty"`        // Наименование получателя платежа - Макс. 160 знаков (имя тега по [5] : Payee/ Name)
	PersonalAcc string `json:"PersonalAcc,omitempty"` // Номер счета получателя платежа - Макс. 20 знаков (имя тега по [5] : Payee/ PersonalAcc)
	BankName    string `json:"BankName,omitempty"`    // Наименование банка получателя платежа - Макс. 45 знаков (не определен [5])
	BIC         string `json:"BIC,omitempty"`         // БИК - Макс. 9 знаков (имя тега по [5] : Payee/ Bank/ BIC)
	CorrespAcc  string `json:"CorrespAcc,omitempty"`  // Номер кор./сч. банка получателя платежа - Макс. 20 знаков (имя тега по УФЭБС: Payee/ Bank/ CorrespAcc)
	// ExBankPay   ExtendBankPay
}

// ExtendBankPay - Дополнительные реквизиты, формат значений которых определяется Альбомом [5].
type ExtendBankPay struct {
	Sum          string `json:"Sum,omitempty"`          // Сумма платежа, в копейках -	Макс. 18 знаков (имя тега по [5] : Sum)
	Purpose      string `json:"Purpose,omitempty"`      // Наименование платежа (назначение) - Макс. 210 знаков (имя тега по [5]: Purpose)
	PayeeINN     string `json:"PayeeINN,omitempty"`     // ИНН получателя платежа - Макс. 12 знаков 	(имя тега по [5]: Payee/ INN)
	PayerINN     string `json:"PayerINN,omitempty"`     // ИНН плательщика - Макс. 12 знаков имя тега по [5] : Payer/ INN)
	DrawerStatus string `json:"DrawerStatus,omitempty"` // Статус составителя платежного документа - Макс. 2 знака (имя тега по [5]:	DepartmentalInfo/ DrawerStatus)
	KPP          string `json:"KPP,omitempty"`          // КПП получателя платежа - Макс. 9 знаков имя тега по [5]: Payee/ KPP)
	CBC          string `json:"CBC,omitempty"`          // КБК - Макс. 20 знаков имя тега по [5]: DepartmentalInfo/ CBC)
	OKTMO        string `json:"OKTMO,omitempty"`        // Общероссийский классификатор территорий муниципальных образований (ОКТМО) - Макс. 11 знаков (имя тега по [5]: DepartmentalInfo/ OKATO, поле 105)
	PaytReason   string `json:"PaytReason,omitempty"`   // Основание налогового платежа - Макс. 2 знака (имя тега по [5]: DepartmentalInfo/ PaytReason)
	TaxPeriod    string `json:"TaxPeriod,omitempty"`    // Налоговый период - Макс. 10 знаков (имя тега по [5]: DepartmentalInfo/ TaxPeriod)
	DocNo        string `json:"DocNo,omitempty"`        // Номер документа - Макс. 15 знаков (имя тега по [5]: DepartmentalInfo/ DocNo)
	DocDate      string `json:"DocDate,omitempty"`      // Дата документа - Макс. 10 знаков (имя тега по [5]: DepartmentalInfo/ DocDate)
	TaxPaytKind  string `json:"TaxPaytKind,omitempty"`  // Тип платежа - Макс. 2 знака (имя тега по [5]: sDepartmentalInfo/ TaxPaytKind)
	// AExBankPay   AnotherExtendBankPay
}

// AnotherExtendBankPay - Прочие дополнительные реквизиты
type AnotherExtendBankPay struct {
	LastName        string `json:"Sum,omitempty"`             // Фамилия плательщика
	FirstName       string `json:"FirstName,omitempty"`       // Имя плательщика
	MiddleName      string `json:"MiddleName,omitempty"`      // Отчество плательщика
	PayerAddress    string `json:"PayerAddress,omitempty"`    // Адрес плательщика
	PersonalAccount string `json:"PersonalAccount,omitempty"` // Лицевой счет бюджетного получателя
	DocIdx          string `json:"DocIdx,omitempty"`          // Индекс платежного документа
	PensAcc         string `json:"PensAcc,omitempty"`         // No лицевого счета в системе персонифицированного учета в ПФР - СНИЛС
	Contract        string `json:"Contract,omitempty"`        // Номер договора
	PersAcc         string `json:"PersAcc,omitempty"`         // Номер лицевого счета плательщика в организации (в системе учета ПУ)
	Flat            string `json:"Flat,omitempty"`            // Номер квартиры
	Phone           string `json:"Phone,omitempty"`           // Номер телефона
	PayerIDType     string `json:"PayerIdType,omitempty"`     // PayerIdType - Вид ДУЛ плательщика
	PayerIDNum      string `json:"PayerIdNum,omitempty"`      // PayerIdNum - Номер ДУЛ плательщика
	ChildFio        string `json:"ChildFio,omitempty"`        // Ф.И.О. ребенка/учащегося
	BirthDate       string `json:"BirthDate,omitempty"`       // Дата рождения
	PaymTerm        string `json:"PaymTerm,omitempty"`        // Срок платежа/дата выставления счета
	PaymPeriod      string `json:"PaymPeriod,omitempty"`      // Период оплаты
	Category        string `json:"Category,omitempty"`        // Вид платежа
	ServiceName     string `json:"ServiceName,omitempty"`     // Код услуги/название прибора учета
	CounterID       string `json:"CounterID,omitempty"`       // CounterId - Номер прибора учета
	CounterVal      string `json:"CounterVal,omitempty"`      // Показание прибора учета
	QuittID         string `json:"QuittId,omitempty"`         // QuittId - Номер извещения, начисления, счета
	QuittDate       string `json:"QuittDate,omitempty"`       // Дата извещения/начисления/счета/постановления (для ГИБДД)
	InstNum         string `json:"InstNum,omitempty"`         // Номер учреждения (образовательного, медицинского)
	ClassNum        string `json:"ClassNum,omitempty"`        // Номер группы детсада/класса школы
	SpecFio         string `json:"SpecFio,omitempty"`         // ФИО преподавателя, специалиста, оказывающего услугу
	AddAmount       string `json:"AddAmount,omitempty"`       // Сумма страховки/дополнительной услуги/Сумма пени (в копейках)
	RuleID          string `json:"RuleId,omitempty"`          // RuleId - Номер постановления (для ГИБДД)
	ExecID          string `json:"ExecId,omitempty"`          // ExecId - Номер исполнительного производства
	RegType         string `json:"RegType,omitempty"`         // Код вида платежа (например, для платежей в адрес Росреестра)
	UIN             string `json:"UIN,omitempty"`             //Уникальный идентификатор начисления
	TechCode        string `json:"TechCode,omitempty"`        /* Технический код, рекомендуемый для заполнения поставщиком услуг. Может
	использоваться принимающей организацией для вызова соответствующей
	обрабатывающей ИТ-системы.*/
}
