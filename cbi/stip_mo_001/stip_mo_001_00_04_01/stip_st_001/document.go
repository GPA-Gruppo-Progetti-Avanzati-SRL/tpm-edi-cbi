package stip_st_001

import (
	"bytes"
	"encoding/xml"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.09/pain_001_001_09"
)

type DocumentAdapter func(stipDocument *Document) (*Document, error)
type DocumentsAdapter func(stipDocuments []*Document) ([]*Document, error)

type Document struct {
	XMLName xml.Name                             `xml:"urn:CBI:xsd:CBIPaymentRequest.00.04.01 CBIPaymentRequest"`
	GrpHdr  *pain_001_001_09.GroupHeader85       `xml:"GrpHdr"`
	PmtInf  pain_001_001_09.PaymentInstruction30 `xml:"PmtInf"`
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
