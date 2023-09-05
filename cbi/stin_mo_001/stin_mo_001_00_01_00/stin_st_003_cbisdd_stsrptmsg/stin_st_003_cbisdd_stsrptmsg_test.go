package stin_st_003_cbisdd_stsrptmsg_test

import (
	_ "embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/stin_mo_001/stin_mo_001_00_01_00/stin_st_003_cbisdd_stsrptmsg"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed example_stin_st_003_cbisdd_stsrptmsg.xml
var exampleCBIBdySDDStsRpt []byte

func TestDocument_Read_STRN_ST_001_CBIBdyBkToCstmrStmtReq(t *testing.T) {
	d, err := stin_st_003_cbisdd_stsrptmsg.NewDocumentFromXML(exampleCBIBdySDDStsRpt)
	require.NoError(t, err)

	xml, err := d.ToXML()
	require.NoError(t, err)

	t.Log(string(xml))
}
