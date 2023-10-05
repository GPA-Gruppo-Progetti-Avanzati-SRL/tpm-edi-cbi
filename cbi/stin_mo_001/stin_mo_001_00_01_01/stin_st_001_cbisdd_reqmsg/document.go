package stin_st_001_cbisdd_reqmsg

import (
	"bytes"
	"encoding/xml"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.08/pain_008_001_08"
)

type DocumentAdapter func(d *Document) (*Document, error)

type Document struct {
	XMLName xml.Name                               `xml:"urn:CBI:xsd:CBISDDReqLogMsg.00.01.01 CBISDDReqLogMsg"`
	GrpHdr  *pain_008_001_08.GroupHeader83         `xml:"GrpHdr"`
	PmtInf  []pain_008_001_08.PaymentInstruction29 `xml:"PmtInf"`
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
