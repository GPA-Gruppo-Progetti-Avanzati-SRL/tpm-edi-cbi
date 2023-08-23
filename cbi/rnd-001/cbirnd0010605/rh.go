package cbirnd0010605

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile"
)

const (
	StartFiller              = "start-filler"
	RecordType               = "record-type"
	Sender                   = "sender"
	Recipient                = "recipient"
	CreationDate             = "creation-date"
	SupportName              = "support-name"
	Filler2                  = "filler-2"
	FieldNa                  = "field-na"
	NoStatements             = "no-statements"
	Filler3                  = "filler-3"
	NoRecords                = "no-records"
	Filler4                  = "filler-4"
	ProgrNumber              = "progr-number"
	OrigBankAbi              = "orig-bank-abi"
	Reason                   = "reason"
	Description              = "description"
	AccountType              = "account-type"
	Cin                      = "cin"
	BankAbi                  = "bank-abi"
	BankCab                  = "bank-cab"
	CurrentAccountCode       = "current-account-code"
	CurrencyCode             = "currency-code"
	AccountingDate           = "accounting-date"
	Sign                     = "sign"
	OpeningBalance           = "opening-balance"
	CountryCode              = "country-code"
	CheckDigit               = "check-digit"
	MovmntProgrNumber        = "movmnt-progr-number"
	ValueDate                = "value-date"
	MovmntSign               = "movmnt-sign"
	Movmntamount             = "movmntamount"
	CbiReason                = "cbi-reason"
	InternalReason           = "internal-reason"
	ChequeNumber             = "cheque-number"
	BankRef                  = "bank-ref"
	CustRefType              = "cust-ref-type"
	CustRefMovmntDescr       = "cust-ref-movmnt-descr"
	StructureFlag            = "structure-flag"
	TypeIdentifier           = "type-identifier"
	OrderDate                = "order-date"
	OrderingPrtyTaxpayerCode = "ordering-prty--taxpayer-code"
	OrderingPrtyDescr        = "ordering-prty-descr"
	Country                  = "country"
	OrderingPrtyAddr         = "ordering-prty-addr"
	OrderingPrtyIban         = "ordering-prty-iban"
	OrigAmnt                 = "orig-amnt"
	OrigAmntCurrencyCode     = "orig-amnt-currency-code"
	PaidAmnt                 = "paid-amnt"
	PaidAmntCurrencyCode     = "paid-amnt-currency-code"
	TrxAmnt                  = "trx-amnt"
	TrxAmntCurrencyCode      = "trx-amnt-currency-code"
	ExchgRate                = "exchg-rate"
	CommissionAmnt           = "commission-amnt"
	CommissionFeesAmnt       = "commission-fees-amnt"
	OrderingPrty             = "ordering-prty"
	Payee                    = "payee"
	PaymentReason            = "payment-reason"
	MsgId                    = "msg-id"
	End2EndId                = "end-2-end-id"
	Filler                   = "filler"
	Descr                    = "descr"
	AccountsBalanceSign      = "accounts-balance-sign"
	AccountsBalance          = "accounts-balance"
	CashBalanceSign          = "cash-balance-sign"
	CashBalance              = "cash-balance"
	FirstCashOnHandDate      = "first-cash-on-hand-date"
	FirstCashSign            = "first-cash-sign"
	FirstCashBalance         = "first-cash-balance"
	SecondCashOnHandDate     = "second-cash-on-hand-date"
	SecondCashSign           = "second-cash-sign"
	SecondCashBalance        = "second-cash-balance"
	ThirdCashOnHandDate      = "third-cash-on-hand-date"
	ThirdCashSign            = "third-cash-sign"
	ThirdCashBalance         = "third-cash-balance"
	FourthCashOnHandDate     = "fourth-cash-on-hand-date"
	FourthCashSign           = "fourth-cash-sign"
	FourthCashBalance        = "fourth-cash-balance"
	FifthCashOnHandDate      = "fifth-cash-on-hand-date"
	FifthCashSign            = "fifth-cash-sign"
	fifthCashBalance         = "fifth-cash-balance"
)

var RHDefinition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH",
	PrefixDiscriminator: " RH",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: Sender, Name: Sender, Length: 5},
		{Trim: true, Drop: false, Id: Recipient, Name: Recipient, Length: 5},
		{Trim: true, Drop: false, Id: CreationDate, Name: CreationDate, Length: 6},
		{Trim: true, Drop: false, Id: SupportName, Name: SupportName, Length: 20},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 76},
		{Trim: true, Drop: true, Id: FieldNa, Name: FieldNa, Length: 5},
	},
}

var RHEFDefinition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-EF",
	PrefixDiscriminator: " EF",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: Sender, Name: Sender, Length: 5},
		{Trim: true, Drop: false, Id: Recipient, Name: Recipient, Length: 5},
		{Trim: true, Drop: false, Id: CreationDate, Name: CreationDate, Length: 6},
		{Trim: true, Drop: false, Id: SupportName, Name: SupportName, Length: 20},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 6},
		{Trim: true, Drop: false, Id: NoStatements, Name: NoStatements, Length: 7},
		{Trim: true, Drop: true, Id: Filler3, Name: Filler3, Length: 30},
		{Trim: true, Drop: false, Id: NoRecords, Name: NoRecords, Length: 7},
		{Trim: true, Drop: true, Id: Filler4, Name: Filler4, Length: 25},
		{Trim: true, Drop: true, Id: FieldNa, Name: FieldNa, Length: 6},
	},
}

var RH61Definition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-61",
	PrefixDiscriminator: " 61",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 13},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 5},
		{Trim: true, Drop: false, Id: Reason, Name: Reason, Length: 5},
		{Trim: true, Drop: false, Id: Description, Name: Description, Length: 16},
		{Trim: true, Drop: false, Id: AccountType, Name: AccountType, Length: 2},
		// {Trim: true, Drop: false,  Id: "bank-details", Name: "bank-details", Length: 23},
		{Trim: true, Drop: false, Id: Cin, Name: Cin, Length: 1},
		{Trim: true, Drop: false, Id: BankAbi, Name: BankAbi, Length: 5},
		{Trim: true, Drop: false, Id: BankCab, Name: BankCab, Length: 5},
		{Trim: true, Drop: false, Id: CurrentAccountCode, Name: CurrentAccountCode, Length: 12},
		{Trim: true, Drop: false, Id: CurrencyCode, Name: CurrencyCode, Length: 3},
		{Trim: true, Drop: false, Id: AccountingDate, Name: AccountingDate, Length: 6},
		{Trim: true, Drop: false, Id: Sign, Name: Sign, Length: 1},
		{Trim: true, Drop: false, Id: OpeningBalance, Name: OpeningBalance, Length: 15},
		// {Trim: true, Drop: false,  Id: "more-iban-details", Name: "more-iban-details", Length: 4},
		{Trim: true, Drop: false, Id: CountryCode, Name: CountryCode, Length: 2},
		{Trim: true, Drop: false, Id: CheckDigit, Name: CheckDigit, Length: 2},
		{Trim: true, Drop: true, Id: Filler3, Name: Filler3, Length: 17},
	},
}

var RH62Definition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-62",
	PrefixDiscriminator: " 62",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: ValueDate, Name: ValueDate, Length: 6},
		{Trim: true, Drop: false, Id: AccountingDate, Name: AccountingDate, Length: 6},
		{Trim: true, Drop: false, Id: MovmntSign, Name: MovmntSign, Length: 1},
		{Trim: true, Drop: false, Id: Movmntamount, Name: Movmntamount, Length: 15},
		{Trim: true, Drop: false, Id: CbiReason, Name: CbiReason, Length: 2},
		{Trim: true, Drop: false, Id: InternalReason, Name: InternalReason, Length: 2},
		{Trim: true, Drop: false, Id: ChequeNumber, Name: ChequeNumber, Length: 16},
		{Trim: true, Drop: false, Id: BankRef, Name: BankRef, Length: 16},
		{Trim: true, Drop: false, Id: CustRefType, Name: CustRefType, Length: 9},
		{Trim: true, Drop: false, Id: CustRefMovmntDescr, Name: CustRefMovmntDescr, Length: 34},
	},
}

var RH63Definition_KKK = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-KKK",
	PrefixDiscriminator: " 63**********KKK",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "type-identifier", Name: "type-identifier", Length: 23},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 81},
	},
}

var RH63Definition_YYY = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-YYY",
	PrefixDiscriminator: " 63**********YYY",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "order-date", Name: "order-date", Length: 8},
		{Trim: true, Drop: false, Id: "ordering-prty--taxpayer-code", Name: "ordering-prty--taxpayer-code", Length: 16},
		{Trim: true, Drop: false, Id: "ordering-prty-descr", Name: "ordering-prty-descr", Length: 40},
		{Trim: true, Drop: false, Id: "country", Name: "country", Length: 40},
	},
}

var RH63Definition_YY2 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-YY2",
	PrefixDiscriminator: " 63**********YY2",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "ordering-prty-addr", Name: "ordering-prty-addr", Length: 50},
		{Trim: true, Drop: false, Id: "ordering-prty-iban", Name: "ordering-prty-iban", Length: 34},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 20},
	},
}

var RH63Definition_ZZ1 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-ZZ1",
	PrefixDiscriminator: " 63**********ZZ1",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "orig-amnt", Name: "orig-amnt", Length: 18},
		{Trim: true, Drop: false, Id: "orig-amnt-currency-code", Name: "orig-amnt-currency-code", Length: 3},
		{Trim: true, Drop: false, Id: "paid-amnt", Name: "paid-amnt", Length: 18},
		{Trim: true, Drop: false, Id: "paid-amnt-currency-code", Name: "paid-amnt-currency-code", Length: 3},
		{Trim: true, Drop: false, Id: TrxAmnt, Name: TrxAmnt, Length: 18},
		{Trim: true, Drop: false, Id: "trx-amnt-currency-code", Name: "trx-amnt-currency-code", Length: 3},
		{Trim: true, Drop: false, Id: "exchg-rate", Name: "exchg-rate", Length: 12},
		{Trim: true, Drop: false, Id: "commission-amnt", Name: "commission-amnt", Length: 13},
		{Trim: true, Drop: false, Id: "commission-fees-amnt", Name: "commission-fees-amnt", Length: 13},
		{Trim: true, Drop: false, Id: CountryCode, Name: CountryCode, Length: 3},
	},
}

var RH63Definition_ZZ2 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-ZZ2",
	PrefixDiscriminator: " 63**********ZZ2",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "ordering-prty", Name: "ordering-prty", Length: 104},
	},
}

var RH63Definition_ZZ3 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-ZZ3",
	PrefixDiscriminator: " 63**********ZZ3",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "payee", Name: "payee", Length: 50},
		{Trim: true, Drop: false, Id: "payment-reason", Name: "payment-reason", Length: 54},
	},
}

var RH63Definition_ID1 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-ID1",
	PrefixDiscriminator: " 63**********ID1",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "msg-id", Name: "msg-id", Length: 35},
		{Trim: true, Drop: false, Id: "end-2-end-id", Name: "end-2-end-id", Length: 35},
		{Trim: true, Drop: true, Id: "filler", Name: "filler", Length: 34},
	},
}

var RH63Definition_RI1 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-RI1",
	PrefixDiscriminator: " 63**********RI1",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "reconc-data", Name: "reconc-data", Length: 104},
	},
}

var RH63Definition_RI2 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-RI2",
	PrefixDiscriminator: " 63**********RI2",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: "reconc-data", Name: "reconc-data", Length: 104},
	},
}

var RH63Definition_Else = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-63-Else",
	PrefixDiscriminator: " 63",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: "descr", Name: "descr", Length: 107},
	},
}

var RH64Definition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-64",
	PrefixDiscriminator: " 64",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: CurrencyCode, Name: CurrencyCode, Length: 3},
		{Trim: true, Drop: false, Id: AccountingDate, Name: AccountingDate, Length: 6},
		{Trim: true, Drop: false, Id: "accounts-balance-sign", Name: "accounts-balance-sign", Length: 1},
		{Trim: true, Drop: false, Id: "accounts-balance", Name: "accounts-balance", Length: 15},
		{Trim: true, Drop: false, Id: "cash-balance-sign", Name: "cash-balance-sign", Length: 1},
		{Trim: true, Drop: false, Id: "cash-balance", Name: "cash-balance", Length: 15},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 54},
		{Trim: true, Drop: true, Id: Filler3, Name: Filler3, Length: 15},
	},
}

var RH65Definition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  "RH-65",
	PrefixDiscriminator: " 65",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		// {Trim: true, Drop: false,  Id: "first-cash-balance", Name: "first-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: "first-cash-on-hand-date", Name: "first-cash-on-hand-date", Length: 6},
		{Trim: true, Drop: false, Id: "first-cash-sign", Name: "first-cash-sign", Length: 1},
		{Trim: true, Drop: false, Id: "first-cash-balance", Name: "first-cash-balance", Length: 15},
		// {Trim: true, Drop: false,  Id: "second-cash-balance", Name: "second-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: "second-cash-on-hand-date", Name: "second-cash-on-hand-date", Length: 6},
		{Trim: true, Drop: false, Id: "second-cash-sign", Name: "second-cash-sign", Length: 1},
		{Trim: true, Drop: false, Id: "second-cash-balance", Name: "second-cash-balance", Length: 15},
		// {Trim: true, Drop: false,  Id: "third-cash-balance", Name: "third-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: "third-cash-on-hand-date", Name: "third-cash-on-hand-date", Length: 6},
		{Trim: true, Drop: false, Id: "third-cash-sign", Name: "third-cash-sign", Length: 1},
		{Trim: true, Drop: false, Id: "third-cash-balance", Name: "third-cash-balance", Length: 15},
		// {Trim: true, Drop: false,  Id: "fourth-cash-balance", Name: "fourth-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: "fourth-cash-on-hand-date", Name: "fourth-cash-on-hand-date", Length: 6},
		{Trim: true, Drop: false, Id: "fourth-cash-sign", Name: "fourth-cash-sign", Length: 1},
		{Trim: true, Drop: false, Id: "fourth-cash-balance", Name: "fourth-cash-balance", Length: 15},
		//{Trim: true, Drop: false,  Id: "fifth-cash-balance", Name: "fifth-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: "fifth-cash-on-hand-date", Name: "fifth-cash-on-hand-date", Length: 6},
		{Trim: true, Drop: false, Id: "fifth-cash-sign", Name: "fifth-cash-sign", Length: 1},
		{Trim: true, Drop: false, Id: "fifth-cash-balance", Name: "fifth-cash-balance", Length: 15},
	},
}
