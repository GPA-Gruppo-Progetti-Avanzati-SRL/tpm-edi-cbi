package stin_st_002_cbisdd_techvalstsmsg

import (
	"bytes"
	"encoding/xml"
)

type DocumentAdapter func(d *Document) (*Document, error)

type Document struct {
	XMLName                     xml.Name                          `xml:"urn:CBI:xsd:CBIBdySDDTechValSts.00.01.01 CBIBdySDDTechValSts"`
	PhyMsgInf                   *PhyMsgInf                        `xml:"PhyMsgInf,omitempty"`
	CBIEnvelSDDTechValStsLogMsg []CBIEnvelSDDTechValStsLogMsgType `xml:"CBIEnvelSDDTechValStsLogMsg,omitempty"`
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
