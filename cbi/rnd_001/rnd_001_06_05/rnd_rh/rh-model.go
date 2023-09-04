package rnd_rh

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
	"strings"
)

/*
type BatchOfStatementsBORecord reader.Record

	func (r BatchOfStatementsBORecord) String() string {
		var r1 = reader.Record(r)
		return fmt.Sprintf("%-20s - %s", "Batch header", (&r1).String())
	}

	func (s BatchOfStatementsBORecord) IsEmpty() bool {
		return s.RecordId == ""
	}

type BatchOfStatementsEORecord reader.Record

	func (r BatchOfStatementsEORecord) String() string {
		var r1 = reader.Record(r)
		return fmt.Sprintf("%-20s - %s", "Batch footer", (&r1).String())
	}
*/

type Statement struct {
	OpeningBalance reader.Record
	ClosingBalance reader.Record
	Movements      []MovementRecord
	ExpCashOnHand  reader.Record
}

func (s Statement) IsEmpty() bool {
	return s.OpeningBalance.RecordId == ""
}

func (s Statement) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-20s - ", "Opening balance"))
	sb.WriteString(s.OpeningBalance.String())
	for _, m := range s.Movements {
		// sb.WriteString(fmt.Sprintf("%-20s - ", "Movement"))
		sb.WriteString("\n")
		sb.WriteString(m.String())
	}
	sb.WriteString(fmt.Sprintf("\n%-20s - ", "Closing balance"))
	sb.WriteString(s.ClosingBalance.String())
	sb.WriteString(fmt.Sprintf("\n%-20s - ", "Expected cash on hnd"))
	sb.WriteString(s.ExpCashOnHand.String())
	return sb.String()
}

/*
type OpeningBalanceRecord reader.Record

	func (r OpeningBalanceRecord) String() string {
		var r1 = reader.Record(r)
		return (&r1).String()
	}
*/
type MovementRecord struct {
	reader.Record
	Details []reader.Record
}

func (r MovementRecord) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-20s - ", "Movement"))
	sb.WriteString(r.Record.String())
	for _, d := range r.Details {
		sb.WriteString(fmt.Sprintf("\n%-20s - ", "Detail"))
		sb.WriteString(d.String())
	}

	return sb.String()
}

func (r MovementRecord) FirstByStructureFlag(structureFlag string) reader.Record {
	for _, d := range r.Details {
		if d.Get(StructureFlag) == structureFlag {
			return d
		}
	}

	return reader.Record{}
}

func (r MovementRecord) AdditionalInfo() string {
	var sb strings.Builder
	for _, d := range r.Details {
		if d.RecordId == Rh63MvmntDetElseRecId {
			sb.WriteString(d.Get(Descr))
		}
	}

	return sb.String()
}

/*
type MovementDetailRecord reader.Record

func (r MovementDetailRecord) String() string {
	var r1 = reader.Record(r)
	return (&r1).String()
}


type ClosingBalanceRecord reader.Record

func (r ClosingBalanceRecord) String() string {
	var r1 = reader.Record(r)
	return (&r1).String()
}
*/

/*
type CashInHandRecord reader.Record

func (r CashInHandRecord) String() string {
	var r1 = reader.Record(r)
	return (&r1).String()
}
*/
