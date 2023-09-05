package stin_st_003_cbisdd_stsrptmsg

import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/pain_002_001_03"

type PhyMsgInf struct {
	PhyMsgTpCd string `xml:"PhyMsgTpCd,omitempty"`
	NbOfLogMsg int    `xml:"NbOfLogMsg,omitempty"`
}

type CBIEnvelSDDStsRptLogMsg struct {
	CBISDDStsRptLogMsg CBISDDStsRptLogMsg `xml:"CBISDDStsRptLogMsg,omitempty"`
}

type CBISDDStsRptLogMsg struct {
	GrpHdr            pain_002_001_03.GroupHeader36                 `xml:"GrpHdr"`
	OrgnlGrpInfAndSts pain_002_001_03.OriginalGroupInformation20    `xml:"OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []pain_002_001_03.OriginalPaymentInformation1 `xml:"OrgnlPmtInfAndSts,omitempty"`
}
