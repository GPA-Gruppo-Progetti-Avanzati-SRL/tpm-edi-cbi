package stin_st_002_cbisdd_techvalstsmsg

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/pain_002_001_03"
)

type PhyMsgInf struct {
	PhyMsgTpCd           string                    `xml:"PhyMsgTpCd,omitempty"`
	NbOfLogMsg           int                       `xml:"NbOfLogMsg,omitempty"`
	OrgnlPhyMsgInfAndSts *OrgnlPhyMsgInfAndStsType `xml:"OrgnlPhyMsgInfAndSts,omitempty"`
}

type OrgnlPhyMsgInfAndStsType struct {
	IdE2EMsg   string         `xml:"IdE2EMsg,omitempty"`
	XMLCreDt   string         `xml:"XMLCreDt,omitempty"`
	PhyMsgTpCd string         `xml:"PhyMsgTpCd,omitempty"`
	Sts        string         `xml:"Sts,omitempty"`
	StsRsnInf  *StsRsnInfType `xml:"StsRsnInf,omitempty"`
}

type StsRsnInfType struct {
	StsRsn         *StsRsnType `xml:"StsRsn,omitempty"`
	AddtlStsRsnInf []string    `xml:"AddtlStsRsnInf,omitempty"`
}

type StsRsnType struct {
	Cd     string `xml:"Cd,omitempty"`
	Prtry  string `xml:"Prtry,omitempty"`
	ElmRfc string `xml:"ElmRfc,omitempty"`
}

type CBIEnvelSDDTechValStsLogMsgType struct {
	CBISDDTechValStsLogMsg *CBISDDTechValStsLogMsgType `xml:"CBISDDTechValStsLogMsg,omitempty"`
}

type CBISDDTechValStsLogMsgType struct {
	GrpHdr            *pain_002_001_03.GroupHeader36                `xml:"GrpHdr,omitempty"`
	OrgnlGrpInfAndSts *pain_002_001_03.OriginalGroupInformation20   `xml:"OrgnlGrpInfAndSts,omitempty"`
	OrgnlPmtInfAndSts []pain_002_001_03.OriginalPaymentInformation1 `xml:"OrgnlPmtInfAndSts,omitempty"`
}
