package main

// StdUtf8 - Standart coding UTF8
const StdUtf8 = "ST00012"

// StdWin1251 - Standart coding Win1251
const StdWin1251 = "ST00011"

// CoreBankPay - Обязательные реквизиты (блок «Payee» УФЭБС[5])
type CoreBankPay struct {
	// ST          string // Идентификатор формата (формат «свой»)
	// FormatVer   string // Версия формата
	// CodeText    string // кодировка
	ST          string // Служебный блок
	Name        string // Наименование получателя платежа - Макс. 160 знаков (имя тега по [5] : Payee/ Name)
	PersonalAcc string // Номер счета получателя платежа - Макс. 20 знаков (имя тега по [5] : Payee/ PersonalAcc)
	BankName    string // Наименование банка получателя платежа - Макс. 45 знаков (не определен [5])
	BIC         string // БИК - Макс. 9 знаков (имя тега по [5] : Payee/ Bank/ BIC)
	CorrespAcc  string // Номер кор./сч. банка получателя платежа - Макс. 20 знаков (имя тега по УФЭБС: Payee/ Bank/ CorrespAcc)
}

// ExtendBankPay - Дополнительные реквизиты, формат значений которых определяется Альбомом [5].
type ExtendBankPay struct {
	Sum          string // Сумма платежа, в копейках -	Макс. 18 знаков (имя тега по [5] : Sum)
	Purpose      string // Наименование платежа (назначение) - Макс. 210 знаков (имя тега по [5]: Purpose)
	PayeeINN     string // ИНН получателя платежа - Макс. 12 знаков 	(имя тега по [5]: Payee/ INN)
	PayerINN     string // ИНН плательщика - Макс. 12 знаков имя тега по [5] : Payer/ INN)
	DrawerStatus string // Статус составителя платежного документа - Макс. 2 знака (имя тега по [5]:	DepartmentalInfo/ DrawerStatus)
	KPP          string // КПП получателя платежа - Макс. 9 знаков имя тега по [5]: Payee/ KPP)
	CBC          string // КБК - Макс. 20 знаков имя тега по [5]: DepartmentalInfo/ CBC)
	OKTMO        string // Общероссийский классификатор территорий муниципальных образований (ОКТМО) - Макс. 11 знаков (имя тега по [5]: DepartmentalInfo/ OKATO, поле 105)
	PaytReason   string // Основание налогового платежа - Макс. 2 знака (имя тега по [5]: DepartmentalInfo/ PaytReason)
	TaxPeriod    string // Налоговый период - Макс. 10 знаков (имя тега по [5]: DepartmentalInfo/ TaxPeriod)
	DocNo        string // Номер документа - Макс. 15 знаков (имя тега по [5]: DepartmentalInfo/ DocNo)
	DocDate      string // Дата документа - Макс. 10 знаков (имя тега по [5]: DepartmentalInfo/ DocDate)
	TaxPaytKind  string // Тип платежа - Макс. 2 знака (имя тега по [5]: sDepartmentalInfo/ TaxPaytKind)
}

// AnotherExtendBankPay - Прочие дополнительные реквизиты
type AnotherExtendBankPay struct {
	LastName        string // Фамилия плательщика
	FirstName       string // Имя плательщика
	MiddleName      string // Отчество плательщика
	PayerAddress    string // Адрес плательщика
	PersonalAccount string // Лицевой счет бюджетного получателя
	DocIdx          string // Индекс платежного документа
	PensAcc         string // No лицевого счета в системе персонифицированного учета в ПФР - СНИЛС
	Contract        string // Номер договора
	PersAcc         string // Номер лицевого счета плательщика в организации (в системе учета ПУ)
	Flat            string // Номер квартиры
	Phone           string // Номер телефона
	PayerIDType     string // PayerIdType - Вид ДУЛ плательщика
	PayerIDNum      string // PayerIdNum - Номер ДУЛ плательщика
	ChildFio        string // Ф.И.О. ребенка/учащегося
	BirthDate       string // Дата рождения
	PaymTerm        string // Срок платежа/дата выставления счета
	PaymPeriod      string // Период оплаты
	Category        string // Вид платежа
	ServiceName     string // Код услуги/название прибора учета
	CounterID       string // CounterId - Номер прибора учета
	CounterVal      string // Показание прибора учета
	QuittID         string // QuittId - Номер извещения, начисления, счета
	QuittDate       string // Дата извещения/начисления/счета/постановления (для ГИБДД)
	InstNum         string // Номер учреждения (образовательного, медицинского)
	ClassNum        string // Номер группы детсада/класса школы
	SpecFio         string // ФИО преподавателя, специалиста, оказывающего услугу
	AddAmount       string // Сумма страховки/дополнительной услуги/Сумма пени (в копейках)
	RuleID          string // RuleId - Номер постановления (для ГИБДД)
	ExecID          string // ExecId - Номер исполнительного производства
	RegType         string // Код вида платежа (например, для платежей в адрес Росреестра)
	UIN             string //Уникальный идентификатор начисления
	TechCode        string /* Технический код, рекомендуемый для заполнения поставщиком услуг. Может
	использоваться принимающей организацией для вызова соответствующей
	обрабатывающей ИТ-системы.*/
}
