package stip_st_001_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/stip_mo_001/stip_mo_001_00_04_00/stip_st_001"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03/pain_001_001_03"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_stip_mo_001_00_04_00 = "example-document-stip_mo_001_00_04_00.xml"

func TestDocument_STIP_MO_001_00_04_00(t *testing.T) {
	d := stip_st_001.Document{
		GrpHdr: pain_001_001_03.GroupHeader32{
			MsgId:   common.MustToMax35Text(common.Max35TextSample),
			CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
			Authstn: []pain_001_001_03.Authorisation1Choice{{
				Cd:    common.MustToAuthorisation1Code(common.Authorisation1CodeSample),
				Prtry: common.MustToMax128Text(common.Max128TextSample)},
			},
			NbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
			CtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
			InitgPty: pain_001_001_03.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_001_001_03.PostalAddress6{
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
				Id: &pain_001_001_03.Party6Choice{
					OrgId: &pain_001_001_03.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_001_001_03.PersonIdentification5{
						DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_001_001_03.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_001_001_03.ContactDetails2{
					NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
					Nm:       common.MustToMax140Text(common.Max140TextSample),
					PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
					Othr:     common.MustToMax35Text(common.Max35TextSample),
				},
			},
			FwdgAgt: &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
				FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
					BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
					ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
					Othr: &pain_001_001_03.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_001_001_03.BranchData2{
					Id: common.MustToMax35Text(common.Max35TextSample),
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
		PmtInf: []pain_001_001_03.PaymentInstructionInformation3{{
			PmtInfId:  common.MustToMax35Text(common.Max35TextSample),
			PmtMtd:    common.MustToPaymentMethod3Code(common.PaymentMethod3CodeSample),
			BtchBookg: xsdt.MustToBoolean(xsdt.BooleanSample),
			NbOfTxs:   common.MustToMax15NumericText(common.Max15NumericTextSample),
			CtrlSum:   xsdt.MustToDecimal(xsdt.DecimalSample),
			PmtTpInf: &pain_001_001_03.PaymentTypeInformation19{
				InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
				SvcLvl: &pain_001_001_03.ServiceLevel8Choice{
					Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				LclInstrm: &pain_001_001_03.LocalInstrument2Choice{
					Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				CtgyPurp: &pain_001_001_03.CategoryPurpose1Choice{
					Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
			},
			ReqdExctnDt:     common.MustToISODate(common.ISODateSample),
			PoolgAdjstmntDt: common.MustToISODate(common.ISODateSample),
			Dbtr: pain_001_001_03.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_001_001_03.PostalAddress6{
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
				Id: &pain_001_001_03.Party6Choice{
					OrgId: &pain_001_001_03.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_001_001_03.PersonIdentification5{
						DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_001_001_03.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_001_001_03.ContactDetails2{
					NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
					Nm:       common.MustToMax140Text(common.Max140TextSample),
					PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
					MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
					EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
					Othr:     common.MustToMax35Text(common.Max35TextSample),
				},
			},
			DbtrAcct: pain_001_001_03.CashAccount16{
				Id: pain_001_001_03.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_001_001_03.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_001_001_03.CashAccountType2{
					Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
			},
			DbtrAgt: pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
				FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
					BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
					ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
					Othr: &pain_001_001_03.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_001_001_03.BranchData2{
					Id: common.MustToMax35Text(common.Max35TextSample),
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
			DbtrAgtAcct: &pain_001_001_03.CashAccount16{
				Id: pain_001_001_03.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_001_001_03.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_001_001_03.CashAccountType2{
					Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
			},
			UltmtDbtr: &pain_001_001_03.PartyIdentification32{
				Nm: common.MustToMax140Text(common.Max140TextSample),
				PstlAdr: &pain_001_001_03.PostalAddress6{
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
				Id: &pain_001_001_03.Party6Choice{
					OrgId: &pain_001_001_03.OrganisationIdentification4{
						BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
						Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
					PrvtId: &pain_001_001_03.PersonIdentification5{
						DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
							BirthDt:     common.MustToISODate(common.ISODateSample),
							PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
							CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
						},
						Othr: []pain_001_001_03.GenericPersonIdentification1{{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample)},
						},
					},
				},
				CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
				CtctDtls: &pain_001_001_03.ContactDetails2{
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
			ChrgsAcct: &pain_001_001_03.CashAccount16{
				Id: pain_001_001_03.AccountIdentification4Choice{
					IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
					Othr: &pain_001_001_03.GenericAccountIdentification1{
						Id: common.MustToMax34Text(common.Max34TextSample),
						SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
							Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Tp: &pain_001_001_03.CashAccountType2{
					Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
				Nm:  common.MustToMax70Text(common.Max70TextSample),
			},
			ChrgsAcctAgt: &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
				FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
					BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
					ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
						ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
							Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						MmbId: common.MustToMax35Text(common.Max35TextSample),
					},
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
					Othr: &pain_001_001_03.GenericFinancialIdentification1{
						Id: common.MustToMax35Text(common.Max35TextSample),
						SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
							Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Issr: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				BrnchId: &pain_001_001_03.BranchData2{
					Id: common.MustToMax35Text(common.Max35TextSample),
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
			CdtTrfTxInf: []pain_001_001_03.CreditTransferTransactionInformation10{{
				PmtId: pain_001_001_03.PaymentIdentification1{
					InstrId:    common.MustToMax35Text(common.Max35TextSample),
					EndToEndId: common.MustToMax35Text(common.Max35TextSample),
				},
				PmtTpInf: &pain_001_001_03.PaymentTypeInformation19{
					InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
					SvcLvl: &pain_001_001_03.ServiceLevel8Choice{
						Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					LclInstrm: &pain_001_001_03.LocalInstrument2Choice{
						Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					CtgyPurp: &pain_001_001_03.CategoryPurpose1Choice{
						Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				Amt: pain_001_001_03.AmountType3Choice{
					InstdAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					EqvtAmt: &pain_001_001_03.EquivalentAmount2{
						Amt: pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						CcyOfTrf: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					},
				},
				XchgRateInf: &pain_001_001_03.ExchangeRateInformation1{
					XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
					RateTp:   common.MustToExchangeRateType1Code(common.ExchangeRateType1CodeSample),
					CtrctId:  common.MustToMax35Text(common.Max35TextSample),
				},
				ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
				ChqInstr: &pain_001_001_03.Cheque6{
					ChqTp: common.MustToChequeType2Code(common.ChequeType2CodeSample),
					ChqNb: common.MustToMax35Text(common.Max35TextSample),
					ChqFr: &pain_001_001_03.NameAndAddress10{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						Adr: pain_001_001_03.PostalAddress6{
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
					DlvryMtd: &pain_001_001_03.ChequeDeliveryMethod1Choice{
						Cd:    common.MustToChequeDelivery1Code(common.ChequeDelivery1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					DlvrTo: &pain_001_001_03.NameAndAddress10{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						Adr: pain_001_001_03.PostalAddress6{
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
					InstrPrty:  common.MustToPriority2Code(common.Priority2CodeSample),
					ChqMtrtyDt: common.MustToISODate(common.ISODateSample),
					FrmsCd:     common.MustToMax35Text(common.Max35TextSample),
					MemoFld: []common.Max35Text{
						common.MustToMax35Text(common.Max35TextSample),
					},
					RgnlClrZone: common.MustToMax35Text(common.Max35TextSample),
					PrtLctn:     common.MustToMax35Text(common.Max35TextSample),
				},
				UltmtDbtr: &pain_001_001_03.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
					Id: &pain_001_001_03.Party6Choice{
						OrgId: &pain_001_001_03.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_001_001_03.PersonIdentification5{
							DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_001_001_03.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_001_001_03.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				IntrmyAgt1: &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
						Othr: &pain_001_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
				IntrmyAgt1Acct: &pain_001_001_03.CashAccount16{
					Id: pain_001_001_03.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_03.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_03.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				IntrmyAgt2: &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
						Othr: &pain_001_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
				IntrmyAgt2Acct: &pain_001_001_03.CashAccount16{
					Id: pain_001_001_03.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_03.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_03.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				IntrmyAgt3: &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
						Othr: &pain_001_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
				IntrmyAgt3Acct: &pain_001_001_03.CashAccount16{
					Id: pain_001_001_03.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_03.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_03.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				CdtrAgt: &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_001_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_001_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
						Othr: &pain_001_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_03.PostalAddress6{
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
				CdtrAgtAcct: &pain_001_001_03.CashAccount16{
					Id: pain_001_001_03.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_03.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_03.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				Cdtr: &pain_001_001_03.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
					Id: &pain_001_001_03.Party6Choice{
						OrgId: &pain_001_001_03.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_001_001_03.PersonIdentification5{
							DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_001_001_03.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_001_001_03.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				CdtrAcct: &pain_001_001_03.CashAccount16{
					Id: pain_001_001_03.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_03.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_03.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_03.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				UltmtCdtr: &pain_001_001_03.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_03.PostalAddress6{
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
					Id: &pain_001_001_03.Party6Choice{
						OrgId: &pain_001_001_03.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_001_001_03.PersonIdentification5{
							DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_001_001_03.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_001_001_03.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				InstrForCdtrAgt: []pain_001_001_03.InstructionForCreditorAgent1{{
					Cd:       common.MustToInstruction3Code(common.Instruction3CodeSample),
					InstrInf: common.MustToMax140Text(common.Max140TextSample)},
				},
				InstrForDbtrAgt: common.MustToMax140Text(common.Max140TextSample),
				Purp: &pain_001_001_03.Purpose2Choice{
					Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				RgltryRptg: []pain_001_001_03.RegulatoryReporting3{{
					DbtCdtRptgInd: common.MustToRegulatoryReportingType1Code(common.RegulatoryReportingType1CodeSample),
					Authrty: &pain_001_001_03.RegulatoryAuthority2{
						Nm:   common.MustToMax140Text(common.Max140TextSample),
						Ctry: common.MustToCountryCode(common.CountryCodeSample),
					},
					Dtls: []pain_001_001_03.StructuredRegulatoryReporting3{{
						Tp:   common.MustToMax35Text(common.Max35TextSample),
						Dt:   common.MustToISODate(common.ISODateSample),
						Ctry: common.MustToCountryCode(common.CountryCodeSample),
						Cd:   common.MustToMax10Text(common.Max10TextSample),
						Amt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Inf: []common.Max35Text{
							common.MustToMax35Text(common.Max35TextSample),
						}},
					}},
				},
				Tax: &pain_001_001_03.TaxInformation3{
					Cdtr: &pain_001_001_03.TaxParty1{
						TaxId:  common.MustToMax35Text(common.Max35TextSample),
						RegnId: common.MustToMax35Text(common.Max35TextSample),
						TaxTp:  common.MustToMax35Text(common.Max35TextSample),
					},
					Dbtr: &pain_001_001_03.TaxParty2{
						TaxId:  common.MustToMax35Text(common.Max35TextSample),
						RegnId: common.MustToMax35Text(common.Max35TextSample),
						TaxTp:  common.MustToMax35Text(common.Max35TextSample),
						Authstn: &pain_001_001_03.TaxAuthorisation1{
							Titl: common.MustToMax35Text(common.Max35TextSample),
							Nm:   common.MustToMax140Text(common.Max140TextSample),
						},
					},
					AdmstnZn: common.MustToMax35Text(common.Max35TextSample),
					RefNb:    common.MustToMax140Text(common.Max140TextSample),
					Mtd:      common.MustToMax35Text(common.Max35TextSample),
					TtlTaxblBaseAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					TtlTaxAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					Dt:    common.MustToISODate(common.ISODateSample),
					SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
					Rcrd: []pain_001_001_03.TaxRecord1{{
						Tp:       common.MustToMax35Text(common.Max35TextSample),
						Ctgy:     common.MustToMax35Text(common.Max35TextSample),
						CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
						DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
						CertId:   common.MustToMax35Text(common.Max35TextSample),
						FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
						Prd: &pain_001_001_03.TaxPeriod1{
							Yr: common.MustToISODate(common.ISODateSample),
							Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
							FrToDt: &pain_001_001_03.DatePeriodDetails{
								FrDt: common.MustToISODate(common.ISODateSample),
								ToDt: common.MustToISODate(common.ISODateSample),
							},
						},
						TaxAmt: &pain_001_001_03.TaxAmount1{
							Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
							TaxblBaseAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TtlAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Dtls: []pain_001_001_03.TaxRecordDetails1{{
								Prd: &pain_001_001_03.TaxPeriod1{
									Yr: common.MustToISODate(common.ISODateSample),
									Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
									FrToDt: &pain_001_001_03.DatePeriodDetails{
										FrDt: common.MustToISODate(common.ISODateSample),
										ToDt: common.MustToISODate(common.ISODateSample),
									},
								},
								Amt: pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								}},
							},
						},
						AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
					},
				},
				RltdRmtInf: []pain_001_001_03.RemittanceLocation2{{
					RmtId:             common.MustToMax35Text(common.Max35TextSample),
					RmtLctnMtd:        common.MustToRemittanceLocationMethod2Code(common.RemittanceLocationMethod2CodeSample),
					RmtLctnElctrncAdr: common.MustToMax2048Text(common.Max2048TextSample),
					RmtLctnPstlAdr: &pain_001_001_03.NameAndAddress10{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						Adr: pain_001_001_03.PostalAddress6{
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
				RmtInf: &pain_001_001_03.RemittanceInformation5{
					Ustrd: []common.Max140Text{
						common.MustToMax140Text(common.Max140TextSample),
					},
					Strd: []pain_001_001_03.StructuredRemittanceInformation7{{
						RfrdDocInf: []pain_001_001_03.ReferredDocumentInformation3{{
							Tp: &pain_001_001_03.ReferredDocumentType2{
								CdOrPrtry: pain_001_001_03.ReferredDocumentType1Choice{
									Cd:    common.MustToDocumentType5Code(common.DocumentType5CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Nb:     common.MustToMax35Text(common.Max35TextSample),
							RltdDt: common.MustToISODate(common.ISODateSample)},
						},
						RfrdDocAmt: &pain_001_001_03.RemittanceAmount1{
							DuePyblAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							DscntApldAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CdtNoteAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							TaxAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							AdjstmntAmtAndRsn: []pain_001_001_03.DocumentAdjustment1{{
								Amt: pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
								Rsn:       common.MustToMax4Text(common.Max4TextSample),
								AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
							},
							RmtdAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
						},
						CdtrRefInf: &pain_001_001_03.CreditorReferenceInformation2{
							Tp: &pain_001_001_03.CreditorReferenceType2{
								CdOrPrtry: pain_001_001_03.CreditorReferenceType1Choice{
									Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
							Ref: common.MustToMax35Text(common.Max35TextSample),
						},
						Invcr: &pain_001_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_03.PostalAddress6{
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
							Id: &pain_001_001_03.Party6Choice{
								OrgId: &pain_001_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_001_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_001_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_001_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Invcee: &pain_001_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_03.PostalAddress6{
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
							Id: &pain_001_001_03.Party6Choice{
								OrgId: &pain_001_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_001_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_001_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_001_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_001_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_001_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_001_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_001_001_03.ContactDetails2{
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

	err = ioutil.WriteFile(Example_stip_mo_001_00_04_00, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_stip_mo_001_00_04_00)

}
