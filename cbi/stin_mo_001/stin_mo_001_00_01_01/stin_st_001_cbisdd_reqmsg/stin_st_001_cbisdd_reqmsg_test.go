package stin_st_001_cbisdd_reqmsg_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/stin_mo_001/stin_mo_001_00_01_01/stin_st_001_cbisdd_reqmsg"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.08/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.08/pain_008_001_08"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"os"
	"testing"
)

const Example_stin_st_001_cbisdd_reqmsg = "example-stin_st_001_cbisdd_reqmsg.xml"

func TestDocument_STIN_ST_001_CbiSddReqMsg(t *testing.T) {
	d := stin_st_001_cbisdd_reqmsg.Document{
		GrpHdr: &pain_008_001_08.GroupHeader83{
			MsgId:   common.MustToMax35Text(common.Max35TextSample),
			CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
			Authstn: []pain_008_001_08.Authorisation1Choice{{
				Cd:    common.MustToAuthorisation1Code(common.Authorisation1CodeSample),
				Prtry: common.MustToMax128Text(common.Max128TextSample)},
			},
			NbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
			CtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
			InitgPty: &pain_008_001_08.PartyIdentification135{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_08.PostalAddress24{
					AdrTp: &pain_008_001_08.AddressType3Choice{
						Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
						Prtry: &pain_008_001_08.GenericIdentification30{
							Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
							Issr:    common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					BldgNm:      common.MustToMax35Text(common.Max35TextSample),
					Flr:         common.MustToMax70Text(common.Max70TextSample),
					PstBx:       common.MustToMax16Text(common.Max16TextSample),
					Room:        common.MustToMax70Text(common.Max70TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
					DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_08.Party38Choice{
					OrgId: &pain_008_001_08.OrganisationIdentification29{
						AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
						LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_08.PersonIdentification13{
						DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_08.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_08.Contact4{
					NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
					Nm:        common.MustToMax140Text(common.Max140TextSample),
					PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
					EmailPurp: common.MustToMax35Text(common.Max35TextSample),
					JobTitl:   common.MustToMax35Text(common.Max35TextSample),
					Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
					Dept:      common.MustToMax70Text(common.Max70TextSample),
					Othr: []pain_008_001_08.OtherContact1{{
						ChanlTp: common.MustToMax4Text(common.Max4TextSample),
						Id:      common.MustToMax128Text(common.Max128TextSample)},
					},
					PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
				},
			},
			FwdgAgt: &pain_008_001_08.BranchAndFinancialInstitutionIdentification6{
				FinInstnId: &pain_008_001_08.FinancialInstitutionIdentification18{
					BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
					ClrSysMmbId: &pain_008_001_08.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_008_001_08.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
					Nm:  common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Othr: &pain_008_001_08.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_008_001_08.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_008_001_08.BranchData3{
					Id:  common.MustToMax35Text(common.Max35TextSample),
					LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
					Nm:  common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
				},
			},
		},
		PmtInf: []pain_008_001_08.PaymentInstruction29{{
			PmtInfId:  common.MustToMax35Text(common.Max35TextSample),
			PmtMtd:    common.MustToPaymentMethod2Code(common.PaymentMethod2CodeSample),
			BtchBookg: xsdt.MustToBoolean(xsdt.BooleanSample),
			NbOfTxs:   common.MustToMax15NumericText(common.Max15NumericTextSample),
			CtrlSum:   xsdt.MustToDecimal(xsdt.DecimalSample),
			PmtTpInf: &pain_008_001_08.PaymentTypeInformation29{
				InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
				SvcLvl: []pain_008_001_08.ServiceLevel8Choice{{
					Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample)},
				},
				LclInstrm: &pain_008_001_08.LocalInstrument2Choice{
					Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				SeqTp: common.MustToSequenceType3Code(common.SequenceType3CodeSample),
				CtgyPurp: &pain_008_001_08.CategoryPurpose1Choice{
					Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
			},
			ReqdColltnDt: common.MustToISODate(common.ISODateSample),
			Cdtr: &pain_008_001_08.PartyIdentification135{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_08.PostalAddress24{
					AdrTp: &pain_008_001_08.AddressType3Choice{
						Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
						Prtry: &pain_008_001_08.GenericIdentification30{
							Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
							Issr:    common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					BldgNm:      common.MustToMax35Text(common.Max35TextSample),
					Flr:         common.MustToMax70Text(common.Max70TextSample),
					PstBx:       common.MustToMax16Text(common.Max16TextSample),
					Room:        common.MustToMax70Text(common.Max70TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
					DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_08.Party38Choice{
					OrgId: &pain_008_001_08.OrganisationIdentification29{
						AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
						LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_08.PersonIdentification13{
						DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_08.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_08.Contact4{
					NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
					Nm:        common.MustToMax140Text(common.Max140TextSample),
					PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
					EmailPurp: common.MustToMax35Text(common.Max35TextSample),
					JobTitl:   common.MustToMax35Text(common.Max35TextSample),
					Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
					Dept:      common.MustToMax70Text(common.Max70TextSample),
					Othr: []pain_008_001_08.OtherContact1{{
						ChanlTp: common.MustToMax4Text(common.Max4TextSample),
						Id:      common.MustToMax128Text(common.Max128TextSample)},
					},
					PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
				},
			},
			CdtrAcct: &pain_008_001_08.CashAccount38{
				Id: &pain_008_001_08.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_008_001_08.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_008_001_08.CashAccountType2Choice{
					Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
				Prxy: &pain_008_001_08.ProxyAccountIdentification1{
					Tp: &pain_008_001_08.ProxyAccountType1Choice{
						Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Id: common.MustToMax2048Text(common.Max2048TextSample),
				},
			},
			CdtrAgt: &pain_008_001_08.BranchAndFinancialInstitutionIdentification6{
				FinInstnId: &pain_008_001_08.FinancialInstitutionIdentification18{
					BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
					ClrSysMmbId: &pain_008_001_08.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_008_001_08.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
					Nm:  common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Othr: &pain_008_001_08.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_008_001_08.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_008_001_08.BranchData3{
					Id:  common.MustToMax35Text(common.Max35TextSample),
					LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
					Nm:  common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
				},
			},
			CdtrAgtAcct: &pain_008_001_08.CashAccount38{
				Id: &pain_008_001_08.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_008_001_08.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_008_001_08.CashAccountType2Choice{
					Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
				Prxy: &pain_008_001_08.ProxyAccountIdentification1{
					Tp: &pain_008_001_08.ProxyAccountType1Choice{
						Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Id: common.MustToMax2048Text(common.Max2048TextSample),
				},
			},
			UltmtCdtr: &pain_008_001_08.PartyIdentification135{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_08.PostalAddress24{
					AdrTp: &pain_008_001_08.AddressType3Choice{
						Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
						Prtry: &pain_008_001_08.GenericIdentification30{
							Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
							Issr:    common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					BldgNm:      common.MustToMax35Text(common.Max35TextSample),
					Flr:         common.MustToMax70Text(common.Max70TextSample),
					PstBx:       common.MustToMax16Text(common.Max16TextSample),
					Room:        common.MustToMax70Text(common.Max70TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
					DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_08.Party38Choice{
					OrgId: &pain_008_001_08.OrganisationIdentification29{
						AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
						LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_08.PersonIdentification13{
						DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_08.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_08.Contact4{
					NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
					Nm:        common.MustToMax140Text(common.Max140TextSample),
					PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
					EmailPurp: common.MustToMax35Text(common.Max35TextSample),
					JobTitl:   common.MustToMax35Text(common.Max35TextSample),
					Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
					Dept:      common.MustToMax70Text(common.Max70TextSample),
					Othr: []pain_008_001_08.OtherContact1{{
						ChanlTp: common.MustToMax4Text(common.Max4TextSample),
						Id:      common.MustToMax128Text(common.Max128TextSample)},
					},
					PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
				},
			},
			ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
			ChrgsAcct: &pain_008_001_08.CashAccount38{
				Id: &pain_008_001_08.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_008_001_08.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_008_001_08.CashAccountType2Choice{
					Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
				Prxy: &pain_008_001_08.ProxyAccountIdentification1{
					Tp: &pain_008_001_08.ProxyAccountType1Choice{
						Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Id: common.MustToMax2048Text(common.Max2048TextSample),
				},
			},
			ChrgsAcctAgt: &pain_008_001_08.BranchAndFinancialInstitutionIdentification6{
				FinInstnId: &pain_008_001_08.FinancialInstitutionIdentification18{
					BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
					ClrSysMmbId: &pain_008_001_08.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_008_001_08.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
					Nm:  common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Othr: &pain_008_001_08.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_008_001_08.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_008_001_08.BranchData3{
					Id:  common.MustToMax35Text(common.Max35TextSample),
					LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
					Nm:  common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
				},
			},
			CdtrSchmeId: &pain_008_001_08.PartyIdentification135{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_008_001_08.PostalAddress24{
					AdrTp: &pain_008_001_08.AddressType3Choice{
						Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
						Prtry: &pain_008_001_08.GenericIdentification30{
							Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
							Issr:    common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Dept:        common.MustToMax70Text(common.Max70TextSample),
					SubDept:     common.MustToMax70Text(common.Max70TextSample),
					StrtNm:      common.MustToMax70Text(common.Max70TextSample),
					BldgNb:      common.MustToMax16Text(common.Max16TextSample),
					BldgNm:      common.MustToMax35Text(common.Max35TextSample),
					Flr:         common.MustToMax70Text(common.Max70TextSample),
					PstBx:       common.MustToMax16Text(common.Max16TextSample),
					Room:        common.MustToMax70Text(common.Max70TextSample),
					PstCd:       common.MustToMax16Text(common.Max16TextSample),
					TwnNm:       common.MustToMax35Text(common.Max35TextSample),
					TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
					DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
					CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
					Ctry:        common.MustToCountryCode(common.CountryCodeSample),
					AdrLine: []common.Max70Text{
						common.MustToMax70Text(common.Max70TextSample),
					},
				},
				Id: &pain_008_001_08.Party38Choice{
					OrgId: &pain_008_001_08.OrganisationIdentification29{
						AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
						LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_008_001_08.PersonIdentification13{
						DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_008_001_08.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_008_001_08.Contact4{
					NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
					Nm:        common.MustToMax140Text(common.Max140TextSample),
					PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
					EmailPurp: common.MustToMax35Text(common.Max35TextSample),
					JobTitl:   common.MustToMax35Text(common.Max35TextSample),
					Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
					Dept:      common.MustToMax70Text(common.Max70TextSample),
					Othr: []pain_008_001_08.OtherContact1{{
						ChanlTp: common.MustToMax4Text(common.Max4TextSample),
						Id:      common.MustToMax128Text(common.Max128TextSample)},
					},
					PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
				},
			},
			DrctDbtTxInf: []pain_008_001_08.DirectDebitTransactionInformation23{{
				PmtId: &pain_008_001_08.PaymentIdentification6{
					InstrId:    common.MustToMax35Text(common.Max35TextSample),
					EndToEndId: common.MustToMax35Text(common.Max35TextSample),
					UETR:       common.MustToUUIDv4Identifier(common.UUIDv4IdentifierSample),
				},
				PmtTpInf: &pain_008_001_08.PaymentTypeInformation29{
					InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
					SvcLvl: []pain_008_001_08.ServiceLevel8Choice{{
						Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample)},
					},
					LclInstrm: &pain_008_001_08.LocalInstrument2Choice{
						Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					SeqTp: common.MustToSequenceType3Code(common.SequenceType3CodeSample),
					CtgyPurp: &pain_008_001_08.CategoryPurpose1Choice{
						Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				InstdAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
					Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Value: xsdt.MustToDecimal(xsdt.DecimalSample),
				},
				ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
				DrctDbtTx: &pain_008_001_08.DirectDebitTransaction10{
					MndtRltdInf: &pain_008_001_08.MandateRelatedInformation14{
						MndtId:    common.MustToMax35Text(common.Max35TextSample),
						DtOfSgntr: common.MustToISODate(common.ISODateSample),
						AmdmntInd: xsdt.MustToBoolean(xsdt.BooleanSample),
						AmdmntInfDtls: &pain_008_001_08.AmendmentInformationDetails13{
							OrgnlMndtId: common.MustToMax35Text(common.Max35TextSample),
							OrgnlCdtrSchmeId: &pain_008_001_08.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_008_001_08.PostalAddress24{
									AdrTp: &pain_008_001_08.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_008_001_08.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_008_001_08.Party38Choice{
									OrgId: &pain_008_001_08.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_008_001_08.PersonIdentification13{
										DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_008_001_08.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_008_001_08.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_008_001_08.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							OrgnlCdtrAgt: &pain_008_001_08.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_008_001_08.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_008_001_08.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_008_001_08.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_08.PostalAddress24{
										AdrTp: &pain_008_001_08.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_008_001_08.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_008_001_08.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_08.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_008_001_08.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_08.PostalAddress24{
										AdrTp: &pain_008_001_08.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_008_001_08.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							OrgnlCdtrAgtAcct: &pain_008_001_08.CashAccount38{
								Id: &pain_008_001_08.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_008_001_08.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_008_001_08.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_008_001_08.ProxyAccountIdentification1{
									Tp: &pain_008_001_08.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							OrgnlDbtr: &pain_008_001_08.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_008_001_08.PostalAddress24{
									AdrTp: &pain_008_001_08.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_008_001_08.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_008_001_08.Party38Choice{
									OrgId: &pain_008_001_08.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_008_001_08.PersonIdentification13{
										DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_008_001_08.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_008_001_08.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_008_001_08.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							OrgnlDbtrAcct: &pain_008_001_08.CashAccount38{
								Id: &pain_008_001_08.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_008_001_08.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_008_001_08.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_008_001_08.ProxyAccountIdentification1{
									Tp: &pain_008_001_08.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							OrgnlDbtrAgt: &pain_008_001_08.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_008_001_08.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_008_001_08.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_008_001_08.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_08.PostalAddress24{
										AdrTp: &pain_008_001_08.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_008_001_08.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_008_001_08.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_08.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_008_001_08.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_008_001_08.PostalAddress24{
										AdrTp: &pain_008_001_08.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_008_001_08.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							OrgnlDbtrAgtAcct: &pain_008_001_08.CashAccount38{
								Id: &pain_008_001_08.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_008_001_08.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_008_001_08.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_008_001_08.ProxyAccountIdentification1{
									Tp: &pain_008_001_08.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							OrgnlFnlColltnDt: common.MustToISODate(common.ISODateSample),
							OrgnlFrqcy: &pain_008_001_08.Frequency36Choice{
								Tp: common.MustToFrequency6Code(common.Frequency6CodeSample),
								Prd: &pain_008_001_08.FrequencyPeriod1{
									Tp:        common.MustToFrequency6Code(common.Frequency6CodeSample),
									CntPerPrd: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								PtInTm: &pain_008_001_08.FrequencyAndMoment1{
									Tp:     common.MustToFrequency6Code(common.Frequency6CodeSample),
									PtInTm: common.MustToExact2NumericText(common.Exact2NumericTextSample),
								},
							},
							OrgnlRsn: &pain_008_001_08.MandateSetupReason1Choice{
								Cd:    common.MustToExternalMandateSetupReason1Code(common.ExternalMandateSetupReason1CodeSample),
								Prtry: common.MustToMax70Text(common.Max70TextSample),
							},
							OrgnlTrckgDays: common.MustToExact2NumericText(common.Exact2NumericTextSample),
						},
						ElctrncSgntr: common.MustToMax1025Text(common.Max1025TextSample),
						FrstColltnDt: common.MustToISODate(common.ISODateSample),
						FnlColltnDt:  common.MustToISODate(common.ISODateSample),
						Frqcy: &pain_008_001_08.Frequency36Choice{
							Tp: common.MustToFrequency6Code(common.Frequency6CodeSample),
							Prd: &pain_008_001_08.FrequencyPeriod1{
								Tp:        common.MustToFrequency6Code(common.Frequency6CodeSample),
								CntPerPrd: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							PtInTm: &pain_008_001_08.FrequencyAndMoment1{
								Tp:     common.MustToFrequency6Code(common.Frequency6CodeSample),
								PtInTm: common.MustToExact2NumericText(common.Exact2NumericTextSample),
							},
						},
						Rsn: &pain_008_001_08.MandateSetupReason1Choice{
							Cd:    common.MustToExternalMandateSetupReason1Code(common.ExternalMandateSetupReason1CodeSample),
							Prtry: common.MustToMax70Text(common.Max70TextSample),
						},
						TrckgDays: common.MustToExact2NumericText(common.Exact2NumericTextSample),
					},
					CdtrSchmeId: &pain_008_001_08.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_008_001_08.PostalAddress24{
							AdrTp: &pain_008_001_08.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_008_001_08.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Id: &pain_008_001_08.Party38Choice{
							OrgId: &pain_008_001_08.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_008_001_08.PersonIdentification13{
								DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_008_001_08.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_008_001_08.Contact4{
							NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
							Nm:        common.MustToMax140Text(common.Max140TextSample),
							PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
							EmailPurp: common.MustToMax35Text(common.Max35TextSample),
							JobTitl:   common.MustToMax35Text(common.Max35TextSample),
							Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
							Dept:      common.MustToMax70Text(common.Max70TextSample),
							Othr: []pain_008_001_08.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					PreNtfctnId: common.MustToMax35Text(common.Max35TextSample),
					PreNtfctnDt: common.MustToISODate(common.ISODateSample),
				},
				UltmtCdtr: &pain_008_001_08.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_008_001_08.Party38Choice{
						OrgId: &pain_008_001_08.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_008_001_08.PersonIdentification13{
							DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_008_001_08.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_008_001_08.Contact4{
						NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
						Nm:        common.MustToMax140Text(common.Max140TextSample),
						PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
						EmailPurp: common.MustToMax35Text(common.Max35TextSample),
						JobTitl:   common.MustToMax35Text(common.Max35TextSample),
						Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
						Dept:      common.MustToMax70Text(common.Max70TextSample),
						Othr: []pain_008_001_08.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				DbtrAgt: &pain_008_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_008_001_08.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_008_001_08.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_008_001_08.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_008_001_08.PostalAddress24{
							AdrTp: &pain_008_001_08.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_008_001_08.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &pain_008_001_08.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_008_001_08.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_008_001_08.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_008_001_08.PostalAddress24{
							AdrTp: &pain_008_001_08.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_008_001_08.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				DbtrAgtAcct: &pain_008_001_08.CashAccount38{
					Id: &pain_008_001_08.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_008_001_08.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_008_001_08.CashAccountType2Choice{
						Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Prxy: &pain_008_001_08.ProxyAccountIdentification1{
						Tp: &pain_008_001_08.ProxyAccountType1Choice{
							Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Id: common.MustToMax2048Text(common.Max2048TextSample),
					},
				},
				Dbtr: &pain_008_001_08.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_008_001_08.Party38Choice{
						OrgId: &pain_008_001_08.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_008_001_08.PersonIdentification13{
							DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_008_001_08.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_008_001_08.Contact4{
						NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
						Nm:        common.MustToMax140Text(common.Max140TextSample),
						PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
						EmailPurp: common.MustToMax35Text(common.Max35TextSample),
						JobTitl:   common.MustToMax35Text(common.Max35TextSample),
						Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
						Dept:      common.MustToMax70Text(common.Max70TextSample),
						Othr: []pain_008_001_08.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				DbtrAcct: &pain_008_001_08.CashAccount38{
					Id: &pain_008_001_08.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_008_001_08.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_008_001_08.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_008_001_08.CashAccountType2Choice{
						Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Prxy: &pain_008_001_08.ProxyAccountIdentification1{
						Tp: &pain_008_001_08.ProxyAccountType1Choice{
							Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Id: common.MustToMax2048Text(common.Max2048TextSample),
					},
				},
				UltmtDbtr: &pain_008_001_08.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_008_001_08.PostalAddress24{
						AdrTp: &pain_008_001_08.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_008_001_08.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_008_001_08.Party38Choice{
						OrgId: &pain_008_001_08.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_008_001_08.PersonIdentification13{
							DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_008_001_08.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_008_001_08.Contact4{
						NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
						Nm:        common.MustToMax140Text(common.Max140TextSample),
						PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
						EmailPurp: common.MustToMax35Text(common.Max35TextSample),
						JobTitl:   common.MustToMax35Text(common.Max35TextSample),
						Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
						Dept:      common.MustToMax70Text(common.Max70TextSample),
						Othr: []pain_008_001_08.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				InstrForCdtrAgt: common.MustToMax140Text(common.Max140TextSample),
				Purp: &pain_008_001_08.Purpose2Choice{
					Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				RgltryRptg: []pain_008_001_08.RegulatoryReporting3{{
					DbtCdtRptgInd: common.MustToRegulatoryReportingType1Code(common.RegulatoryReportingType1CodeSample),
					Authrty: &pain_008_001_08.RegulatoryAuthority2{
						Nm:   common.MustToMax140Text(common.Max140TextSample),
						Ctry: common.MustToCountryCode(common.CountryCodeSample),
					},
					Dtls: []pain_008_001_08.StructuredRegulatoryReporting3{{
						Tp:   common.MustToMax35Text(common.Max35TextSample),
						Dt:   common.MustToISODate(common.ISODateSample),
						Ctry: common.MustToCountryCode(common.CountryCodeSample),
						Cd:   common.MustToMax10Text(common.Max10TextSample),
						Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Inf: []common.Max35Text{
							common.MustToMax35Text(common.Max35TextSample),
						}},
					}},
				},
				Tax: &pain_008_001_08.TaxInformation8{
					Cdtr: &pain_008_001_08.TaxParty1{
						TaxId:  common.MustToMax35Text(common.Max35TextSample),
						RegnId: common.MustToMax35Text(common.Max35TextSample),
						TaxTp:  common.MustToMax35Text(common.Max35TextSample),
					},
					Dbtr: &pain_008_001_08.TaxParty2{
						TaxId:  common.MustToMax35Text(common.Max35TextSample),
						RegnId: common.MustToMax35Text(common.Max35TextSample),
						TaxTp:  common.MustToMax35Text(common.Max35TextSample),
						Authstn: &pain_008_001_08.TaxAuthorisation1{
							Titl: common.MustToMax35Text(common.Max35TextSample),
							Nm:   common.MustToMax140Text(common.Max140TextSample),
						},
					},
					AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
					RefNb:      common.MustToMax140Text(common.Max140TextSample),
					Mtd:        common.MustToMax35Text(common.Max35TextSample),
					TtlTaxblBaseAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					TtlTaxAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					Dt:    common.MustToISODate(common.ISODateSample),
					SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
					Rcrd: []pain_008_001_08.TaxRecord2{{
						Tp:       common.MustToMax35Text(common.Max35TextSample),
						Ctgy:     common.MustToMax35Text(common.Max35TextSample),
						CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
						DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
						CertId:   common.MustToMax35Text(common.Max35TextSample),
						FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
						Prd: &pain_008_001_08.TaxPeriod2{
							Yr: common.MustToISODate(common.ISODateSample),
							Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
							FrToDt: &pain_008_001_08.DatePeriod2{
								FrDt: common.MustToISODate(common.ISODateSample),
								ToDt: common.MustToISODate(common.ISODateSample),
							},
						},
						TaxAmt: &pain_008_001_08.TaxAmount2{
							Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
							TaxblBaseAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TtlAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Dtls: []pain_008_001_08.TaxRecordDetails2{{
								Prd: &pain_008_001_08.TaxPeriod2{
									Yr: common.MustToISODate(common.ISODateSample),
									Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
									FrToDt: &pain_008_001_08.DatePeriod2{
										FrDt: common.MustToISODate(common.ISODateSample),
										ToDt: common.MustToISODate(common.ISODateSample),
									},
								},
								Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								}},
							},
						},
						AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
					},
				},
				RltdRmtInf: []pain_008_001_08.RemittanceLocation7{{
					RmtId: common.MustToMax35Text(common.Max35TextSample),
					RmtLctnDtls: []pain_008_001_08.RemittanceLocationData1{{
						Mtd:        common.MustToRemittanceLocationMethod2Code(common.RemittanceLocationMethod2CodeSample),
						ElctrncAdr: common.MustToMax2048Text(common.Max2048TextSample),
						PstlAdr: &pain_008_001_08.NameAndAddress16{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							Adr: &pain_008_001_08.PostalAddress24{
								AdrTp: &pain_008_001_08.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_008_001_08.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
						}},
					}},
				},
				RmtInf: &pain_008_001_08.RemittanceInformation16{
					Ustrd: []common.Max140Text{
						common.MustToMax140Text(common.Max140TextSample),
					},
					Strd: []pain_008_001_08.StructuredRemittanceInformation16{{
						RfrdDocInf: []pain_008_001_08.ReferredDocumentInformation7{{
							Tp: &pain_008_001_08.ReferredDocumentType4{
								CdOrPrtry: &pain_008_001_08.ReferredDocumentType3Choice{
									Cd:    common.MustToDocumentType6Code(common.DocumentType6CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Nb:     common.MustToMax35Text(common.Max35TextSample),
							RltdDt: common.MustToISODate(common.ISODateSample),
							LineDtls: []pain_008_001_08.DocumentLineInformation1{{
								Id: []pain_008_001_08.DocumentLineIdentification1{{
									Tp: &pain_008_001_08.DocumentLineType1{
										CdOrPrtry: &pain_008_001_08.DocumentLineType1Choice{
											Cd:    common.MustToExternalDocumentLineType1Code(common.ExternalDocumentLineType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
									Nb:     common.MustToMax35Text(common.Max35TextSample),
									RltdDt: common.MustToISODate(common.ISODateSample)},
								},
								Desc: common.MustToMax2048Text(common.Max2048TextSample),
								Amt: &pain_008_001_08.RemittanceAmount3{
									DuePyblAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									DscntApldAmt: []pain_008_001_08.DiscountAmountAndType1{{
										Tp: &pain_008_001_08.DiscountAmountType1Choice{
											Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										}},
									},
									CdtNoteAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									TaxAmt: []pain_008_001_08.TaxAmountAndType1{{
										Tp: &pain_008_001_08.TaxAmountType1Choice{
											Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										}},
									},
									AdjstmntAmtAndRsn: []pain_008_001_08.DocumentAdjustment1{{
										Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
										Rsn:       common.MustToMax4Text(common.Max4TextSample),
										AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
									},
									RmtdAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
								}},
							}},
						},
						RfrdDocAmt: &pain_008_001_08.RemittanceAmount2{
							DuePyblAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							DscntApldAmt: []pain_008_001_08.DiscountAmountAndType1{{
								Tp: &pain_008_001_08.DiscountAmountType1Choice{
									Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								}},
							},
							CdtNoteAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TaxAmt: []pain_008_001_08.TaxAmountAndType1{{
								Tp: &pain_008_001_08.TaxAmountType1Choice{
									Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								}},
							},
							AdjstmntAmtAndRsn: []pain_008_001_08.DocumentAdjustment1{{
								Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
								Rsn:       common.MustToMax4Text(common.Max4TextSample),
								AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
							},
							RmtdAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
						},
						CdtrRefInf: &pain_008_001_08.CreditorReferenceInformation2{
							Tp: &pain_008_001_08.CreditorReferenceType2{
								CdOrPrtry: &pain_008_001_08.CreditorReferenceType1Choice{
									Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Ref: common.MustToMax35Text(common.Max35TextSample),
						},
						Invcr: &pain_008_001_08.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_008_001_08.PostalAddress24{
								AdrTp: &pain_008_001_08.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_008_001_08.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &pain_008_001_08.Party38Choice{
								OrgId: &pain_008_001_08.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_008_001_08.PersonIdentification13{
									DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_008_001_08.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_008_001_08.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []pain_008_001_08.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						Invcee: &pain_008_001_08.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_008_001_08.PostalAddress24{
								AdrTp: &pain_008_001_08.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_008_001_08.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &pain_008_001_08.Party38Choice{
								OrgId: &pain_008_001_08.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_008_001_08.PersonIdentification13{
									DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_008_001_08.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_008_001_08.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []pain_008_001_08.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						TaxRmt: &pain_008_001_08.TaxInformation7{
							Cdtr: &pain_008_001_08.TaxParty1{
								TaxId:  common.MustToMax35Text(common.Max35TextSample),
								RegnId: common.MustToMax35Text(common.Max35TextSample),
								TaxTp:  common.MustToMax35Text(common.Max35TextSample),
							},
							Dbtr: &pain_008_001_08.TaxParty2{
								TaxId:  common.MustToMax35Text(common.Max35TextSample),
								RegnId: common.MustToMax35Text(common.Max35TextSample),
								TaxTp:  common.MustToMax35Text(common.Max35TextSample),
								Authstn: &pain_008_001_08.TaxAuthorisation1{
									Titl: common.MustToMax35Text(common.Max35TextSample),
									Nm:   common.MustToMax140Text(common.Max140TextSample),
								},
							},
							UltmtDbtr: &pain_008_001_08.TaxParty2{
								TaxId:  common.MustToMax35Text(common.Max35TextSample),
								RegnId: common.MustToMax35Text(common.Max35TextSample),
								TaxTp:  common.MustToMax35Text(common.Max35TextSample),
								Authstn: &pain_008_001_08.TaxAuthorisation1{
									Titl: common.MustToMax35Text(common.Max35TextSample),
									Nm:   common.MustToMax140Text(common.Max140TextSample),
								},
							},
							AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
							RefNb:      common.MustToMax140Text(common.Max140TextSample),
							Mtd:        common.MustToMax35Text(common.Max35TextSample),
							TtlTaxblBaseAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TtlTaxAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Dt:    common.MustToISODate(common.ISODateSample),
							SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
							Rcrd: []pain_008_001_08.TaxRecord2{{
								Tp:       common.MustToMax35Text(common.Max35TextSample),
								Ctgy:     common.MustToMax35Text(common.Max35TextSample),
								CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
								DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
								CertId:   common.MustToMax35Text(common.Max35TextSample),
								FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
								Prd: &pain_008_001_08.TaxPeriod2{
									Yr: common.MustToISODate(common.ISODateSample),
									Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
									FrToDt: &pain_008_001_08.DatePeriod2{
										FrDt: common.MustToISODate(common.ISODateSample),
										ToDt: common.MustToISODate(common.ISODateSample),
									},
								},
								TaxAmt: &pain_008_001_08.TaxAmount2{
									Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
									TaxblBaseAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									TtlAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									Dtls: []pain_008_001_08.TaxRecordDetails2{{
										Prd: &pain_008_001_08.TaxPeriod2{
											Yr: common.MustToISODate(common.ISODateSample),
											Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
											FrToDt: &pain_008_001_08.DatePeriod2{
												FrDt: common.MustToISODate(common.ISODateSample),
												ToDt: common.MustToISODate(common.ISODateSample),
											},
										},
										Amt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										}},
									},
								},
								AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
							},
						},
						GrnshmtRmt: &pain_008_001_08.Garnishment3{
							Tp: &pain_008_001_08.GarnishmentType1{
								CdOrPrtry: &pain_008_001_08.GarnishmentType1Choice{
									Cd:    common.MustToExternalGarnishmentType1Code(common.ExternalGarnishmentType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Grnshee: &pain_008_001_08.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_008_001_08.PostalAddress24{
									AdrTp: &pain_008_001_08.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_008_001_08.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_008_001_08.Party38Choice{
									OrgId: &pain_008_001_08.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_008_001_08.PersonIdentification13{
										DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_008_001_08.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_008_001_08.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_008_001_08.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							GrnshmtAdmstr: &pain_008_001_08.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_008_001_08.PostalAddress24{
									AdrTp: &pain_008_001_08.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_008_001_08.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_008_001_08.Party38Choice{
									OrgId: &pain_008_001_08.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_008_001_08.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_008_001_08.PersonIdentification13{
										DtAndPlcOfBirth: &pain_008_001_08.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_008_001_08.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_008_001_08.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_008_001_08.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_008_001_08.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							RefNb: common.MustToMax140Text(common.Max140TextSample),
							Dt:    common.MustToISODate(common.ISODateSample),
							RmtdAmt: &pain_008_001_08.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							FmlyMdclInsrncInd: xsdt.MustToBoolean(xsdt.BooleanSample),
							MplyeeTermntnInd:  xsdt.MustToBoolean(xsdt.BooleanSample),
						},
						AddtlRmtInf: []common.Max140Text{
							common.MustToMax140Text(common.Max140TextSample),
						}},
					},
				},
				SplmtryData: []pain_008_001_08.SupplementaryData1{{
					PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
					Envlp: &pain_008_001_08.SupplementaryDataEnvelope1{
						Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
					}},
				}},
			},
		},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = os.WriteFile(Example_stin_st_001_cbisdd_reqmsg, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_stin_st_001_cbisdd_reqmsg)

}
