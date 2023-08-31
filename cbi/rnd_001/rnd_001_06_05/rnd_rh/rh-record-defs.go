package rnd_rh

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
	ReconcData               = "reconc-data"

	RhRecId               = "RH"
	RhEfRecId             = "RH-EF"
	Rh61OpenBalRecId      = "RH-61"
	Rh62MvmntRecId        = "RH-62"
	Rh63MvmntDetKKKRecId  = "RH-63-KKK"
	Rh63MvmntDetYYYRecId  = "RH-63-YYY"
	Rh63MvmntDetYY2RecId  = "RH-63-YY2"
	Rh63MvmntDetZZ1RecId  = "RH-63-ZZ1"
	Rh63MvmntDetZZ2RecId  = "RH-63-ZZ2"
	Rh63MvmntDetZZ3RecId  = "RH-63-ZZ3"
	Rh63MvmntDetID1RecId  = "RH-63-ID1"
	Rh63MvmntDetRI1RecId  = "RH-63-RI1"
	Rh63RI2MvmntRecId     = "RH-63-RI2"
	Rh63MvmntDetElseRecId = "RH-63-Else"
	Rh64ClosBalRecId      = "RH-64"
	Rh65CashOnHandRecId   = "RH-65"
)

var RHDefinition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  RhRecId,
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
	Id:                  RhEfRecId,
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
	Id:                  Rh61OpenBalRecId,
	PrefixDiscriminator: " 61",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 13},
		{Trim: true, Drop: false, Id: OrigBankAbi, Name: OrigBankAbi, Length: 5},
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
	Id:                  Rh62MvmntRecId,
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
	Id:                  Rh63MvmntDetKKKRecId,
	PrefixDiscriminator: " 63**********KKK",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: TypeIdentifier, Name: TypeIdentifier, Length: 23},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 81},
	},
}

var RH63Definition_YYY = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetYYYRecId,
	PrefixDiscriminator: " 63**********YYY",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: OrderDate, Name: OrderDate, Length: 8},
		{Trim: true, Drop: false, Id: OrderingPrtyTaxpayerCode, Name: OrderingPrtyTaxpayerCode, Length: 16},
		{Trim: true, Drop: false, Id: OrderingPrtyDescr, Name: OrderingPrtyDescr, Length: 40},
		{Trim: true, Drop: false, Id: Country, Name: Country, Length: 40},
	},
}

var RH63Definition_YY2 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetYY2RecId,
	PrefixDiscriminator: " 63**********YY2",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: OrderingPrtyAddr, Name: OrderingPrtyAddr, Length: 50},
		{Trim: true, Drop: false, Id: OrderingPrtyIban, Name: OrderingPrtyIban, Length: 34},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 20},
	},
}

var RH63Definition_ZZ1 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetZZ1RecId,
	PrefixDiscriminator: " 63**********ZZ1",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: OrigAmnt, Name: OrigAmnt, Length: 18},
		{Trim: true, Drop: false, Id: OrigAmntCurrencyCode, Name: OrigAmntCurrencyCode, Length: 3},
		{Trim: true, Drop: false, Id: PaidAmnt, Name: PaidAmnt, Length: 18},
		{Trim: true, Drop: false, Id: PaidAmntCurrencyCode, Name: PaidAmntCurrencyCode, Length: 3},
		{Trim: true, Drop: false, Id: TrxAmnt, Name: TrxAmnt, Length: 18},
		{Trim: true, Drop: false, Id: TrxAmntCurrencyCode, Name: TrxAmntCurrencyCode, Length: 3},
		{Trim: true, Drop: false, Id: ExchgRate, Name: ExchgRate, Length: 12},
		{Trim: true, Drop: false, Id: CommissionAmnt, Name: CommissionAmnt, Length: 13},
		{Trim: true, Drop: false, Id: CommissionFeesAmnt, Name: CommissionFeesAmnt, Length: 13},
		{Trim: true, Drop: false, Id: CountryCode, Name: CountryCode, Length: 3},
	},
}

var RH63Definition_ZZ2 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetZZ2RecId,
	PrefixDiscriminator: " 63**********ZZ2",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: OrderingPrty, Name: OrderingPrty, Length: 104},
	},
}

var RH63Definition_ZZ3 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetZZ3RecId,
	PrefixDiscriminator: " 63**********ZZ3",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: Payee, Name: Payee, Length: 50},
		{Trim: true, Drop: false, Id: PaymentReason, Name: PaymentReason, Length: 54},
	},
}

var RH63Definition_ID1 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetID1RecId,
	PrefixDiscriminator: " 63**********ID1",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: MsgId, Name: MsgId, Length: 35},
		{Trim: true, Drop: false, Id: End2EndId, Name: End2EndId, Length: 35},
		{Trim: true, Drop: true, Id: Filler, Name: Filler, Length: 34},
	},
}

var RH63Definition_RI1 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetRI1RecId,
	PrefixDiscriminator: " 63**********RI1",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: ReconcData, Name: ReconcData, Length: 104},
	},
}

var RH63Definition_RI2 = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63RI2MvmntRecId,
	PrefixDiscriminator: " 63**********RI2",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: StructureFlag, Name: StructureFlag, Length: 3},
		{Trim: true, Drop: false, Id: ReconcData, Name: ReconcData, Length: 104},
	},
}

var RH63Definition_Else = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh63MvmntDetElseRecId,
	PrefixDiscriminator: " 63",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: MovmntProgrNumber, Name: MovmntProgrNumber, Length: 3},
		{Trim: true, Drop: false, Id: Descr, Name: Descr, Length: 107},
	},
}

var RH64Definition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh64ClosBalRecId,
	PrefixDiscriminator: " 64",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		{Trim: true, Drop: false, Id: CurrencyCode, Name: CurrencyCode, Length: 3},
		{Trim: true, Drop: false, Id: AccountingDate, Name: AccountingDate, Length: 6},
		{Trim: true, Drop: false, Id: AccountsBalanceSign, Name: AccountsBalanceSign, Length: 1},
		{Trim: true, Drop: false, Id: AccountsBalance, Name: AccountsBalance, Length: 15},
		{Trim: true, Drop: false, Id: CashBalanceSign, Name: CashBalanceSign, Length: 1},
		{Trim: true, Drop: false, Id: CashBalance, Name: CashBalance, Length: 15},
		{Trim: true, Drop: true, Id: Filler2, Name: Filler2, Length: 54},
		{Trim: true, Drop: true, Id: Filler3, Name: Filler3, Length: 15},
	},
}

var RH65Definition = fixedlengthfile.FixedLengthRecordDefinition{
	Id:                  Rh65CashOnHandRecId,
	PrefixDiscriminator: " 65",
	Fields: []fixedlengthfile.FixedLengthFieldDefinition{
		{Trim: true, Drop: true, Id: StartFiller, Name: StartFiller, Length: 1},
		{Trim: true, Drop: false, Id: RecordType, Name: RecordType, Length: 2},
		{Trim: true, Drop: false, Id: ProgrNumber, Name: ProgrNumber, Length: 7},
		// {Trim: true, Drop: false,  Id: "first-cash-balance", Name: "first-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: FirstCashOnHandDate, Name: FirstCashOnHandDate, Length: 6},
		{Trim: true, Drop: false, Id: FirstCashSign, Name: FirstCashSign, Length: 1},
		{Trim: true, Drop: false, Id: FirstCashBalance, Name: FirstCashBalance, Length: 15},
		// {Trim: true, Drop: false,  Id: "second-cash-balance", Name: "second-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: SecondCashOnHandDate, Name: SecondCashOnHandDate, Length: 6},
		{Trim: true, Drop: false, Id: SecondCashSign, Name: SecondCashSign, Length: 1},
		{Trim: true, Drop: false, Id: SecondCashBalance, Name: SecondCashBalance, Length: 15},
		// {Trim: true, Drop: false,  Id: "third-cash-balance", Name: "third-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: ThirdCashOnHandDate, Name: ThirdCashOnHandDate, Length: 6},
		{Trim: true, Drop: false, Id: ThirdCashSign, Name: ThirdCashSign, Length: 1},
		{Trim: true, Drop: false, Id: ThirdCashBalance, Name: ThirdCashBalance, Length: 15},
		// {Trim: true, Drop: false,  Id: "fourth-cash-balance", Name: "fourth-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: FourthCashOnHandDate, Name: FourthCashOnHandDate, Length: 6},
		{Trim: true, Drop: false, Id: FourthCashSign, Name: FourthCashSign, Length: 1},
		{Trim: true, Drop: false, Id: FourthCashBalance, Name: FourthCashBalance, Length: 15},
		//{Trim: true, Drop: false,  Id: "fifth-cash-balance", Name: "fifth-cash-balance", Length: 22},
		{Trim: true, Drop: false, Id: FifthCashOnHandDate, Name: FifthCashOnHandDate, Length: 6},
		{Trim: true, Drop: false, Id: FifthCashSign, Name: FifthCashSign, Length: 1},
		{Trim: true, Drop: false, Id: fifthCashBalance, Name: fifthCashBalance, Length: 15},
	},
}
