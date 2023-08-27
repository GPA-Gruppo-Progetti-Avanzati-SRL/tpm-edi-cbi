package cbirnd0010605_test

import (
	_ "embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-cbi/cbi/rnd-001/cbirnd0010605"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

//go:embed example-cbi-rnd-001-06-05.txt
var example []byte

func TestReader(t *testing.T) {

	r, err := cbirnd0010605.NewReader(example)
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
