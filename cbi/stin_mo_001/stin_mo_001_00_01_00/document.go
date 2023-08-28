package stin_mo_001_00_01_00

import (
	"bytes"
	"encoding/xml"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02/pain_008_001_02"
)

type Document struct {
	XMLName xml.Name                                         `xml:"urn:CBI:xsd:CBISDDReqLogMsg.00.01.00 CBISDDReqLogMsg"`
	GrpHdr  pain_008_001_02.GroupHeader39                    `xml:"GrpHdr"`
	PmtInf  []pain_008_001_02.PaymentInstructionInformation4 `xml:"PmtInf"`
}

func (d *Document) ToXML() ([]byte, error) {
	w := &bytes.Buffer{}
	w.Write([]byte(xml.Header))

	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	err := enc.Encode(d)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
