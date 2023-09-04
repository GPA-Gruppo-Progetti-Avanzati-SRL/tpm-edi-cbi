package strn_st_001_cbibktocstmrstmt_reqmsg

import (
	"bytes"
	"encoding/xml"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02/camt_053_001_02"
)

type Document struct {
	XMLName                        xml.Name                         `xml:"urn:CBI:xsd:CBIBdyBkToCstmrStmtReq.00.01.02 CBIBdyBkToCstmrStmtReq"`
	PhyMsgInf                      PhyMsgInf                        `xml:"PhyMsgInf,omitempty"`
	CBIEnvelBkToCstmrStmtReqLogMsg []CBIEnvelBkToCstmrStmtReqLogMsg `xml:"CBIEnvelBkToCstmrStmtReqLogMsg,omitempty"`
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

type PhyMsgInf struct {
	PhyMsgTpCd string `xml:"PhyMsgTpCd,omitempty"`
	NbOfLogMsg int    `xml:"NbOfLogMsg,omitempty"`
}

type CBIEnvelBkToCstmrStmtReqLogMsg struct {
	CBIBkToCstmrStmtReqLogMsg CBIBkToCstmrStmtReqLogMsg `xml:"CBIBkToCstmrStmtReqLogMsg,omitempty"`
}

type CBIBkToCstmrStmtReqLogMsg struct {
	CBIDlyStmtReqLogMsg CBIDlyStmtReqLogMsg `xml:"CBIDlyStmtReqLogMsg,omitempty"`
}

// Note. In this current implementation I start from CBIDlyStmtReqLogMsg foor the RH type of files.
// Don't know if the file starts from the parent CBIBkToCstmrStmtReqLogMsg that is achoic type of message.
type CBIDlyStmtReqLogMsg struct {
	// XMLName xml.Name                            `xml:"urn:CBI:xsd:CBIDlyStmtReqLogMsg.00.01.02 CBIDlyStmtReqLogMsg"`
	GrpHdr camt_053_001_02.GroupHeader42       `xml:"GrpHdr"`
	Stmt   []camt_053_001_02.AccountStatement2 `xml:"Stmt"`
}

/*
 * To use these methods need to uncomment the xmlName of CBIDlyStmtReqLogMsg in order to set the proper nmaespace....
 */
func NewCBIDlyStmtReqLogMsgFromXML(b []byte) (*CBIDlyStmtReqLogMsg, error) {
	d := NewCBIDlyStmtReqLogMsg()
	err := xml.Unmarshal(b, &d)
	return &d, err
}

func NewCBIDlyStmtReqLogMsg() CBIDlyStmtReqLogMsg {
	d := CBIDlyStmtReqLogMsg{}
	return d
}

func (d *CBIDlyStmtReqLogMsg) ToXML() ([]byte, error) {
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
