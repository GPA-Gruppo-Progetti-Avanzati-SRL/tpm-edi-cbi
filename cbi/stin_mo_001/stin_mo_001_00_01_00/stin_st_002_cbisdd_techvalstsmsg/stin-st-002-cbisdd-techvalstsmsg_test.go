package stin_st_002_cbisdd_techvalstsmsg_test

import (
	_ "embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/stin_mo_001/stin_mo_001_00_01_00/stin_st_002_cbisdd_techvalstsmsg"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/pain_002_001_03"
	"github.com/stretchr/testify/require"
	"io/fs"
	"os"
	"testing"
)

//go:embed example-in-stin_st_002_cbisdd_techvalstsmsg.xml
var exampleInStinSt002CbiSddTechValStsMsg []byte

const Example_out_stin_st_002_cbisdd_techvalstsmsg = "example-out-stin_st_002_cbisdd_techvalstsmsg.xml"

func TestDocument_Read_STIN_ST_002_CbiSddTechValStsMsg(t *testing.T) {
	d, err := stin_st_002_cbisdd_techvalstsmsg.NewDocumentFromXML(exampleInStinSt002CbiSddTechValStsMsg)
	require.NoError(t, err)

	xml, err := d.ToXML()
	require.NoError(t, err)

	t.Log(string(xml))
}

func TestDocument_STIN_ST_002_CbiSddTechValStsMsg(t *testing.T) {
	d := stin_st_002_cbisdd_techvalstsmsg.Document{
		PhyMsgInf: stin_st_002_cbisdd_techvalstsmsg.PhyMsgInf{},
		CBIEnvelSDDTechValStsLogMsg: stin_st_002_cbisdd_techvalstsmsg.CBIEnvelSDDTechValStsLogMsgType{
			CBISDDTechValStsLogMsg: []stin_st_002_cbisdd_techvalstsmsg.CBISDDTechValStsLogMsgType{
				{
					GrpHdr: pain_002_001_03.GroupHeader36{
						MsgId: "message-id",
					},
					OrgnlGrpInfAndSts: pain_002_001_03.OriginalGroupInformation20{},
					OrgnlPmtInfAndSts: nil,
				},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = os.WriteFile(Example_out_stin_st_002_cbisdd_techvalstsmsg, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_out_stin_st_002_cbisdd_techvalstsmsg)
}
