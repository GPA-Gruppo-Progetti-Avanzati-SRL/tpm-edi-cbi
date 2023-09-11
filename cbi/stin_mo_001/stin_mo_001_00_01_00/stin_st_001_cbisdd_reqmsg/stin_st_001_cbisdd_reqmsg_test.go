package stin_st_001_cbisdd_reqmsg_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/stin_mo_001/stin_mo_001_00_01_00/stin_st_001_cbisdd_reqmsg"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02/pain_008_001_02"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"os"
	"testing"
)

const Example_stin_st_001_cbisdd_reqmsg = "example-stin_st_001_cbisdd_reqmsg.xml"

func TestDocument_STIN_ST_001_CbiSddReqMsg(t *testing.T) {
	d := stin_st_001_cbisdd_reqmsg.Document{
		GrpHdr: &pain_008_001_02.GroupHeader39{
			MsgId:   common.MustToMax35Text(common.Max35TextSample),
			CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
			Authstn: []pain_008_001_02.Authorisation1Choice{{
				Cd:    common.MustToAuthorisation1Code(common.Authorisation1CodeSample),
				Prtry: common.MustToMax128Text(common.Max128TextSample)},
			},
			NbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
			CtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
			InitgPty: &pain_008_001_02.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_02.PostalAddress6{
					AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_02.Party6Choice{
					OrgId: &pain_008_001_02.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_02.PersonIdentification5{
						DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_02.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_02.ContactDetails2{
					NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
					Nm:       common.MustToMax140Text(common.Max140TextSample),
					PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
					Othr:     common.MustToMax35Text(common.Max35TextSample),
				},
			},
			FwdgAgt: &pain_008_001_02.BranchAndFinancialInstitutionIdentification4{
				FinInstnId: &pain_008_001_02.FinancialInstitutionIdentification7{
					BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
					ClrSysMmbId: &pain_008_001_02.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_008_001_02.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Othr: &pain_008_001_02.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_008_001_02.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_008_001_02.BranchData2{
					Id: common.MustToMax35Text(common.Max35TextSample),
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
				},
			},
		},
		PmtInf: []pain_008_001_02.PaymentInstructionInformation4{{
			PmtInfId:  common.MustToMax35Text(common.Max35TextSample),
			PmtMtd:    common.MustToPaymentMethod2Code(common.PaymentMethod2CodeSample),
			BtchBookg: xsdt.MustToBoolean(xsdt.BooleanSample),
			NbOfTxs:   common.MustToMax15NumericText(common.Max15NumericTextSample),
			CtrlSum:   xsdt.MustToDecimal(xsdt.DecimalSample),
			PmtTpInf: &pain_008_001_02.PaymentTypeInformation20{
				InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
				SvcLvl: &pain_008_001_02.ServiceLevel8Choice{
					Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				LclInstrm: &pain_008_001_02.LocalInstrument2Choice{
					Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				SeqTp: common.MustToSequenceType1Code(common.SequenceType1CodeSample),
				CtgyPurp: &pain_008_001_02.CategoryPurpose1Choice{
					Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
			},
			ReqdColltnDt: common.MustToISODate(common.ISODateSample),
			Cdtr: &pain_008_001_02.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_02.PostalAddress6{
					AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_02.Party6Choice{
					OrgId: &pain_008_001_02.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_02.PersonIdentification5{
						DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_02.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_02.ContactDetails2{
					NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
					Nm:       common.MustToMax140Text(common.Max140TextSample),
					PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
					Othr:     common.MustToMax35Text(common.Max35TextSample),
				},
			},
			CdtrAcct: &pain_008_001_02.CashAccount16{
				Id: &pain_008_001_02.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_008_001_02.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_008_001_02.CashAccountType2{
					Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
			},
			CdtrAgt: &pain_008_001_02.BranchAndFinancialInstitutionIdentification4{
				FinInstnId: &pain_008_001_02.FinancialInstitutionIdentification7{
					BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
					ClrSysMmbId: &pain_008_001_02.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_008_001_02.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Othr: &pain_008_001_02.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_008_001_02.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_008_001_02.BranchData2{
					Id: common.MustToMax35Text(common.Max35TextSample),
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
				},
			},
			CdtrAgtAcct: &pain_008_001_02.CashAccount16{
				Id: &pain_008_001_02.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_008_001_02.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_008_001_02.CashAccountType2{
					Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
			},
			UltmtCdtr: &pain_008_001_02.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_02.PostalAddress6{
					AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_02.Party6Choice{
					OrgId: &pain_008_001_02.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_02.PersonIdentification5{
						DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_02.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_02.ContactDetails2{
					NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
					Nm:       common.MustToMax140Text(common.Max140TextSample),
					PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
					Othr:     common.MustToMax35Text(common.Max35TextSample),
				},
			},
			ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
			ChrgsAcct: &pain_008_001_02.CashAccount16{
				Id: &pain_008_001_02.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_008_001_02.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_008_001_02.CashAccountType2{
					Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
			},
			ChrgsAcctAgt: &pain_008_001_02.BranchAndFinancialInstitutionIdentification4{
				FinInstnId: &pain_008_001_02.FinancialInstitutionIdentification7{
					BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
					ClrSysMmbId: &pain_008_001_02.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_008_001_02.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Othr: &pain_008_001_02.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_008_001_02.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_008_001_02.BranchData2{
					Id: common.MustToMax35Text(common.Max35TextSample),
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
				},
			},
			CdtrSchmeId: &pain_008_001_02.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_02.PostalAddress6{
					AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_02.Party6Choice{
					OrgId: &pain_008_001_02.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_02.PersonIdentification5{
						DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_02.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_02.ContactDetails2{
					NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
					Nm:       common.MustToMax140Text(common.Max140TextSample),
					PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
					Othr:     common.MustToMax35Text(common.Max35TextSample),
				},
			},
			DrctDbtTxInf: []pain_008_001_02.DirectDebitTransactionInformation9{{
				PmtId: &pain_008_001_02.PaymentIdentification1{
					InstrId:    common.MustToMax35Text(common.Max35TextSample),
					EndToEndId: common.MustToMax35Text(common.Max35TextSample),
				},
				PmtTpInf: &pain_008_001_02.PaymentTypeInformation20{
					InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
					SvcLvl: &pain_008_001_02.ServiceLevel8Choice{
						Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					LclInstrm: &pain_008_001_02.LocalInstrument2Choice{
						Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					SeqTp: common.MustToSequenceType1Code(common.SequenceType1CodeSample),
					CtgyPurp: &pain_008_001_02.CategoryPurpose1Choice{
						Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				InstdAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
					Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Value: xsdt.MustToDecimal(xsdt.DecimalSample),
				},
				ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
				DrctDbtTx: &pain_008_001_02.DirectDebitTransaction6{
					MndtRltdInf: &pain_008_001_02.MandateRelatedInformation6{
						MndtId:    common.MustToMax35Text(common.Max35TextSample),
						DtOfSgntr: common.MustToISODate(common.ISODateSample),
						AmdmntInd: xsdt.MustToBoolean(xsdt.BooleanSample),
						AmdmntInfDtls: &pain_008_001_02.AmendmentInformationDetails6{
							OrgnlMndtId: common.MustToMax35Text(common.Max35TextSample),
							OrgnlCdtrSchmeId: &pain_008_001_02.PartyIdentification32{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_008_001_02.PostalAddress6{
									AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_008_001_02.Party6Choice{
									OrgId: &pain_008_001_02.OrganisationIdentification4{
										BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
										Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_008_001_02.PersonIdentification5{
										DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_008_001_02.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_008_001_02.ContactDetails2{
									NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
									Nm:       common.MustToMax140Text(common.Max140TextSample),
									PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
									Othr:     common.MustToMax35Text(common.Max35TextSample),
								},
							},
							OrgnlCdtrAgt: &pain_008_001_02.BranchAndFinancialInstitutionIdentification4{
								FinInstnId: &pain_008_001_02.FinancialInstitutionIdentification7{
									BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
									ClrSysMmbId: &pain_008_001_02.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_008_001_02.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_008_001_02.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_02.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_008_001_02.BranchData2{
									Id: common.MustToMax35Text(common.Max35TextSample),
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							OrgnlCdtrAgtAcct: &pain_008_001_02.CashAccount16{
								Id: &pain_008_001_02.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_008_001_02.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_008_001_02.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							OrgnlDbtr: &pain_008_001_02.PartyIdentification32{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_008_001_02.PostalAddress6{
									AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_008_001_02.Party6Choice{
									OrgId: &pain_008_001_02.OrganisationIdentification4{
										BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
										Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_008_001_02.PersonIdentification5{
										DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_008_001_02.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_008_001_02.ContactDetails2{
									NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
									Nm:       common.MustToMax140Text(common.Max140TextSample),
									PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
									Othr:     common.MustToMax35Text(common.Max35TextSample),
								},
							},
							OrgnlDbtrAcct: &pain_008_001_02.CashAccount16{
								Id: &pain_008_001_02.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_008_001_02.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_008_001_02.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							OrgnlDbtrAgt: &pain_008_001_02.BranchAndFinancialInstitutionIdentification4{
								FinInstnId: &pain_008_001_02.FinancialInstitutionIdentification7{
									BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
									ClrSysMmbId: &pain_008_001_02.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_008_001_02.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_008_001_02.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_02.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_008_001_02.BranchData2{
									Id: common.MustToMax35Text(common.Max35TextSample),
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							OrgnlDbtrAgtAcct: &pain_008_001_02.CashAccount16{
								Id: &pain_008_001_02.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_008_001_02.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_008_001_02.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							OrgnlFnlColltnDt: common.MustToISODate(common.ISODateSample),
							OrgnlFrqcy:       common.MustToFrequency1Code(common.Frequency1CodeSample),
						},
						ElctrncSgntr: common.MustToMax1025Text(common.Max1025TextSample),
						FrstColltnDt: common.MustToISODate(common.ISODateSample),
						FnlColltnDt:  common.MustToISODate(common.ISODateSample),
						Frqcy:        common.MustToFrequency1Code(common.Frequency1CodeSample),
					},
					CdtrSchmeId: &pain_008_001_02.PartyIdentification32{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_008_001_02.PostalAddress6{
							AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Id: &pain_008_001_02.Party6Choice{
							OrgId: &pain_008_001_02.OrganisationIdentification4{
								BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
								Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_008_001_02.PersonIdentification5{
								DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_008_001_02.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_008_001_02.ContactDetails2{
							NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
							Nm:       common.MustToMax140Text(common.Max140TextSample),
							PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
							Othr:     common.MustToMax35Text(common.Max35TextSample),
						},
					},
					PreNtfctnId: common.MustToMax35Text(common.Max35TextSample),
					PreNtfctnDt: common.MustToISODate(common.ISODateSample),
				},
				UltmtCdtr: &pain_008_001_02.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_008_001_02.Party6Choice{
						OrgId: &pain_008_001_02.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_008_001_02.PersonIdentification5{
							DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_008_001_02.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_008_001_02.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				DbtrAgt: &pain_008_001_02.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: &pain_008_001_02.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_008_001_02.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_008_001_02.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_008_001_02.PostalAddress6{
							AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &pain_008_001_02.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_02.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_008_001_02.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_008_001_02.PostalAddress6{
							AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				DbtrAgtAcct: &pain_008_001_02.CashAccount16{
					Id: &pain_008_001_02.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_008_001_02.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_008_001_02.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				Dbtr: &pain_008_001_02.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_008_001_02.Party6Choice{
						OrgId: &pain_008_001_02.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_008_001_02.PersonIdentification5{
							DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_008_001_02.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_008_001_02.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				DbtrAcct: &pain_008_001_02.CashAccount16{
					Id: &pain_008_001_02.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_008_001_02.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_008_001_02.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_008_001_02.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				UltmtDbtr: &pain_008_001_02.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_008_001_02.Party6Choice{
						OrgId: &pain_008_001_02.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_008_001_02.PersonIdentification5{
							DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_008_001_02.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_008_001_02.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				InstrForCdtrAgt: common.MustToMax140Text(common.Max140TextSample),
				Purp: &pain_008_001_02.Purpose2Choice{
					Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				RgltryRptg: []pain_008_001_02.RegulatoryReporting3{{
					DbtCdtRptgInd: common.MustToRegulatoryReportingType1Code(common.RegulatoryReportingType1CodeSample),
					Authrty: &pain_008_001_02.RegulatoryAuthority2{
						Nm:   common.MustToMax140Text(common.Max140TextSample),
						Ctry: common.MustToCountryCode(common.CountryCodeSample),
					},
					Dtls: []pain_008_001_02.StructuredRegulatoryReporting3{{
						Tp:   common.MustToMax35Text(common.Max35TextSample),
						Dt:   common.MustToISODate(common.ISODateSample),
						Ctry: common.MustToCountryCode(common.CountryCodeSample),
						Cd:   common.MustToMax10Text(common.Max10TextSample),
						Amt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Inf: []common.Max35Text{
							common.MustToMax35Text(common.Max35TextSample),
						}},
					}},
				},
				Tax: &pain_008_001_02.TaxInformation3{
					Cdtr: &pain_008_001_02.TaxParty1{
						TaxId:  common.MustToMax35Text(common.Max35TextSample),
						RegnId: common.MustToMax35Text(common.Max35TextSample),
						TaxTp:  common.MustToMax35Text(common.Max35TextSample),
					},
					Dbtr: &pain_008_001_02.TaxParty2{
						TaxId:  common.MustToMax35Text(common.Max35TextSample),
						RegnId: common.MustToMax35Text(common.Max35TextSample),
						TaxTp:  common.MustToMax35Text(common.Max35TextSample),
						Authstn: &pain_008_001_02.TaxAuthorisation1{
							Titl: common.MustToMax35Text(common.Max35TextSample),
							Nm:   common.MustToMax140Text(common.Max140TextSample),
						},
					},
					AdmstnZn: common.MustToMax35Text(common.Max35TextSample),
					RefNb:    common.MustToMax140Text(common.Max140TextSample),
					Mtd:      common.MustToMax35Text(common.Max35TextSample),
					TtlTaxblBaseAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					TtlTaxAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					Dt:    common.MustToISODate(common.ISODateSample),
					SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
					Rcrd: []pain_008_001_02.TaxRecord1{{
						Tp:       common.MustToMax35Text(common.Max35TextSample),
						Ctgy:     common.MustToMax35Text(common.Max35TextSample),
						CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
						DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
						CertId:   common.MustToMax35Text(common.Max35TextSample),
						FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
						Prd: &pain_008_001_02.TaxPeriod1{
							Yr: common.MustToISODate(common.ISODateSample),
							Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
							FrToDt: &pain_008_001_02.DatePeriodDetails{
								FrDt: common.MustToISODate(common.ISODateSample),
								ToDt: common.MustToISODate(common.ISODateSample),
							},
						},
						TaxAmt: &pain_008_001_02.TaxAmount1{
							Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
							TaxblBaseAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TtlAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Dtls: []pain_008_001_02.TaxRecordDetails1{{
								Prd: &pain_008_001_02.TaxPeriod1{
									Yr: common.MustToISODate(common.ISODateSample),
									Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
									FrToDt: &pain_008_001_02.DatePeriodDetails{
										FrDt: common.MustToISODate(common.ISODateSample),
										ToDt: common.MustToISODate(common.ISODateSample),
									},
								},
								Amt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								}},
							},
						},
						AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
					},
				},
				RltdRmtInf: []pain_008_001_02.RemittanceLocation2{{
					RmtId:             common.MustToMax35Text(common.Max35TextSample),
					RmtLctnMtd:        common.MustToRemittanceLocationMethod2Code(common.RemittanceLocationMethod2CodeSample),
					RmtLctnElctrncAdr: common.MustToMax2048Text(common.Max2048TextSample),
					RmtLctnPstlAdr: &pain_008_001_02.NameAndAddress10{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						Adr: &pain_008_001_02.PostalAddress6{
							AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					}},
				},
				RmtInf: &pain_008_001_02.RemittanceInformation5{
					Ustrd: []common.Max140Text{
						common.MustToMax140Text(common.Max140TextSample),
					},
					Strd: []pain_008_001_02.StructuredRemittanceInformation7{{
						RfrdDocInf: []pain_008_001_02.ReferredDocumentInformation3{{
							Tp: &pain_008_001_02.ReferredDocumentType2{
								CdOrPrtry: &pain_008_001_02.ReferredDocumentType1Choice{
									Cd:    common.MustToDocumentType5Code(common.DocumentType5CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Nb:     common.MustToMax35Text(common.Max35TextSample),
							RltdDt: common.MustToISODate(common.ISODateSample)},
						},
						RfrdDocAmt: &pain_008_001_02.RemittanceAmount1{
							DuePyblAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							DscntApldAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CdtNoteAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TaxAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							AdjstmntAmtAndRsn: []pain_008_001_02.DocumentAdjustment1{{
								Amt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
								Rsn:       common.MustToMax4Text(common.Max4TextSample),
								AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
							},
							RmtdAmt: &pain_008_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
						},
						CdtrRefInf: &pain_008_001_02.CreditorReferenceInformation2{
							Tp: &pain_008_001_02.CreditorReferenceType2{
								CdOrPrtry: &pain_008_001_02.CreditorReferenceType1Choice{
									Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Ref: common.MustToMax35Text(common.Max35TextSample),
						},
						Invcr: &pain_008_001_02.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_008_001_02.PostalAddress6{
								AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &pain_008_001_02.Party6Choice{
								OrgId: &pain_008_001_02.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_008_001_02.PersonIdentification5{
									DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_008_001_02.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_008_001_02.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Invcee: &pain_008_001_02.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_008_001_02.PostalAddress6{
								AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &pain_008_001_02.Party6Choice{
								OrgId: &pain_008_001_02.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_008_001_02.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_02.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_008_001_02.PersonIdentification5{
									DtAndPlcOfBirth: &pain_008_001_02.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_008_001_02.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_02.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_008_001_02.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						AddtlRmtInf: []common.Max140Text{
							common.MustToMax140Text(common.Max140TextSample),
						}},
					},
				}},
			}},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = os.WriteFile(Example_stin_st_001_cbisdd_reqmsg, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_stin_st_001_cbisdd_reqmsg)

}
