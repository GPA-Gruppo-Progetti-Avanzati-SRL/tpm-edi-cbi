package rnd_rh

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
)

var fixedLengthReaderConfig = reader.Config{
	FileName:       "",
	Discriminator:  "prefix",
	EmptyLinesMode: reader.EmptyLinesModeKeep,
	Records: []fixedlengthfile.FixedLengthRecordDefinition{
		RHDefinition,
		RHEFDefinition,
		RH61Definition,
		RH62Definition,
		RH63Definition_KKK,
		RH63Definition_YYY,
		RH63Definition_YY2,
		RH63Definition_ZZ1,
		RH63Definition_ZZ2,
		RH63Definition_ZZ3,
		RH63Definition_ID1,
		RH63Definition_RI1,
		RH63Definition_RI2,
		RH63Definition_Else,
		RH64Definition,
		RH65Definition,
	},
}
