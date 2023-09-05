package stin_st_003_cbisdd_stsrptmsg

import (
	"bytes"
	"encoding/xml"
)

type Document struct {
	XMLName                 xml.Name                  `xml:"urn:CBI:xsd:CBIBdySDDStsRpt.00.01.00 CBIBdySDDStsRpt"`
	PhyMsgInf               PhyMsgInf                 `xml:"PhyMsgInf,omitempty"`
	CBIEnvelSDDStsRptLogMsg []CBIEnvelSDDStsRptLogMsg `xml:"CBIEnvelSDDStsRptLogMsg,omitempty"`
}

func NewDocumentFromXML(b []byte) (*Document, error) {
	d := NewDocument()
	err := xml.Unmarshal(b, &d)
	return &d, err
}

func NewDocument() Document {
	d := Document{}
	return d
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
