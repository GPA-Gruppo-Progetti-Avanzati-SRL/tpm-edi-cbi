package cbirnd0010605

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
)

var fixedLengthReaderConfig = reader.Config{
	FileName:       "",
	Discriminator:  "prefix",
	EmptyLinesMode: reader.EmptyLinesModeKeep,
	Records: []fixedlengthfile.FixedLengthRecordDefinition{
		reader.RHDefinition,
		reader.RHEFDefinition,
		reader.RH61Definition,
		reader.RH62Definition,
		reader.RH63Definition_KKK,
		reader.RH63Definition_YYY,
		reader.RH63Definition_YY2,
		reader.RH63Definition_ZZ1,
		reader.RH63Definition_ZZ2,
		reader.RH63Definition_ZZ3,
		reader.RH63Definition_ID1,
		reader.RH63Definition_RI1,
		reader.RH63Definition_RI2,
		reader.RH63Definition_Else,
		reader.RH64Definition,
		reader.RH65Definition,
	},
}
