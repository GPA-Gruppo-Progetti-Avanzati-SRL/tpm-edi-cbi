package strn_st_001_cbibktocstmrstmt_reqmsg_test

import (
	_ "embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/strn_mo_001/strn_mo_001_00_01_02/strn_st_001_cbibktocstmrstmt_reqmsg"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed example-in-strn-st-001-cbi-dly-stmt-req-log-msg.xml
var exampleInCBIDlyStmtReqLogMsg []byte

func TestDocument_Read_STRN_ST_001_CBIDlyStmtReqLogMsg(t *testing.T) {
	d, err := strn_st_001_cbibktocstmrstmt_reqmsg.NewCBIDlyStmtReqLogMsgFromXML(exampleInCBIDlyStmtReqLogMsg)
	require.NoError(t, err)

	xml, err := d.ToXML()
	require.NoError(t, err)

	t.Log(string(xml))
}

//go:embed example-in-strn-st-001-cbi-bdy_bk-to-cstmr-stmt-req.xml
var exampleInCBIBdyBkToCstmrStmtReq []byte

func TestDocument_Read_STRN_ST_001_CBIBdyBkToCstmrStmtReq(t *testing.T) {
	d, err := strn_st_001_cbibktocstmrstmt_reqmsg.NewDocumentFromXML(exampleInCBIBdyBkToCstmrStmtReq)
	require.NoError(t, err)

	xml, err := d.ToXML()
	require.NoError(t, err)

	t.Log(string(xml))
}
