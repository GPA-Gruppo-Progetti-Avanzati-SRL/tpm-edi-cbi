package rnd_rh_test

import (
	_ "embed"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/rnd_001/rnd_001_06_05/rnd_rh"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

//go:embed example-cbi-rnd-001-06-05.txt
var example []byte

func TestReader(t *testing.T) {

	r, err := rnd_rh.NewReader(example)
	require.NoError(t, err)

	var rec reader.Record
	for err == nil {
		rec, err = r.Read()
		if err == nil {
			t.Log(rec.String())
		}
	}

	if err != io.EOF {
		require.NoError(t, err)
	}
}

func TestStatementReader(t *testing.T) {
	r, err := rnd_rh.NewReader(example)
	require.NoError(t, err)

	var rhBofRec reader.Record
	var rhEofRec reader.Record
	for err == nil {
		rhBofRec, err = r.ReadBatchOfStatementsBORecord()
		handleError(t, err)
		if rhBofRec.IsEmpty() {
			t.Log("###### EOF stream")
			break
		}

		fmt.Printf("%-20s - %s\n", "Batch header", rhBofRec.String())
		var rhStmtRec rnd_rh.Statement
		for err == nil {
			rhStmtRec, err = r.ReadStatement()
			handleError(t, err)
			if rhStmtRec.IsEmpty() {
				break
			}
			fmt.Println(rhStmtRec.String())
		}

		rhEofRec, err = r.ReadBatchOfStatementsEORecord()
		handleError(t, err)
		fmt.Printf("%-20s - %s\n", "Batch footer", rhEofRec.String())
	}
}

func handleError(t *testing.T, err error) {
	if err != nil {
		if err != io.EOF {
			// it just panics...
			t.Fatal(err)
		}
	}

	return
}
