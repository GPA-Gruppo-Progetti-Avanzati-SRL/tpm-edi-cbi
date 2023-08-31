package strn_st_001_cbibktocstmrstmt_reqmsg_test

import (
	_ "embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/strn_mo_001/strn_mo_001_00_01_02/strn_st_001_cbibktocstmrstmt_reqmsg"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed example-in-strn-st-001-cbibktocstmrstmt-reqmsg.xml
var exampleInCBIDlyStmtReqLogMsg []byte

const Example_out_stin_st_002_cbisdd_techvalstsmsg = "example-out-strn-st-001-cbibktocstmrstmt-reqmsg.xml"

func TestDocument_Read_STRN_ST_001_CBIDlyStmtReqLogMsg(t *testing.T) {
	d, err := strn_st_001_cbibktocstmrstmt_reqmsg.NewDocumentFromXML(exampleInCBIDlyStmtReqLogMsg)
	require.NoError(t, err)

	xml, err := d.ToXML()
	require.NoError(t, err)

	t.Log(string(xml))
}
