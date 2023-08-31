package rnd_rh

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
	"github.com/rs/zerolog/log"
)

type RNDRHReader0010605 struct {
	isError         bool
	fileReader      reader.Reader
	backTrackRecord reader.Record
}

func NewReader(dataContent []byte) (*RNDRHReader0010605, error) {
	const semLogContext = "cbi-rnd-001.06.05::new-rh-reader"

	strReader := bytes.NewReader(dataContent)
	flr, err := reader.NewReader(fixedLengthReaderConfig, reader.WithIoReader(strReader))
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	r := RNDRHReader0010605{fileReader: flr}
	return &r, nil
}

func (r *RNDRHReader0010605) Read() (reader.Record, error) {
	const semLogContext = "cbi-rnd-001.06.05::read-rh"

	if r.isError {
		return reader.Record{}, errors.New("cbi-rnd-001.06.05 already in error state")
	}

	var rec reader.Record
	var err error

	if r.backTrackRecord.RecordId != "" {
		rec = r.backTrackRecord
		r.backTrackRecord = reader.Record{}
		return rec, nil
	}

	rec, err = r.fileReader.Read()
	if err != nil {
		r.isError = true
		return reader.Record{}, err
	}

	return rec, nil
}

const (
	StatusStart        = "start"
	StatusEnd          = "end"
	StatusOpenBalance  = "opening-balance"
	StatusCloseBalance = "closing-balance"
	StatusMvmnt        = "movement"
	StatusMvmntDetail  = "movement-detail"
	StatusCashOnHand   = "cash-on-hand"

	ActionConsume  = "consume"
	ActionPushBack = "push-back"
	ActionReturn   = "return"
)

func (r *RNDRHReader0010605) ReadBatchOfStatementsBORecord() (BatchOfStatementsBORecord, error) {
	rec, err := r.readRecordOfType(RhRecId)
	return BatchOfStatementsBORecord(rec), err
}

func (r *RNDRHReader0010605) ReadBatchOfStatementsEORecord() (BatchOfStatementsEORecord, error) {
	rec, err := r.readRecordOfType(RhEfRecId)
	return BatchOfStatementsEORecord(rec), err
}

func (r *RNDRHReader0010605) readRecordOfType(recId string) (reader.Record, error) {

	rec, err := r.Read()
	if err != nil {
		r.isError = true
		return reader.Record{}, err
	}

	if rec.RecordId != recId {
		err = fmt.Errorf("unexpected record: %s, wanted: %s", rec.RecordId, recId)
		return reader.Record{}, err
	}

	return rec, nil
}

func (r *RNDRHReader0010605) ReadStatement() (Statement, error) {
	const semLogContext = "cbi-rnd-001.06.05::read-statement"

	stmt := Statement{}
	if r.isError {
		return stmt, errors.New("cbi-rnd-001.06.05 already in error state")
	}

	status := StatusStart
	for {
		rec, err := r.Read()
		if err != nil {
			r.isError = true
			return stmt, err
		}

		var act string
		switch status {
		case StatusStart:
			status, act, err = r.OnStatusStart(rec.RecordId)
		case StatusOpenBalance:
			status, act, err = r.OnStatusOpenBalance(rec.RecordId)
		case StatusCloseBalance:
			status, act, err = r.OnStatusCloseBalance(rec.RecordId)
		case StatusMvmnt:
			status, act, err = r.OnStatusMovement(rec.RecordId)
		case StatusMvmntDetail:
			status, act, err = r.OnStatusMovementDetail(rec.RecordId)
		case StatusCashOnHand:
			status, act, err = r.OnStatusCashInHand(rec.RecordId)
		}

		switch act {
		case ActionPushBack:
			r.backTrackRecord = rec
			return stmt, nil
		case ActionConsume:
			if IsRecIdMovementDetail(rec.RecordId) {
				stmt.movements[len(stmt.movements)-1].details = append(stmt.movements[len(stmt.movements)-1].details, MovementDetailRecord(rec))
			} else {
				switch rec.RecordId {
				case Rh61OpenBalRecId:
					stmt.openingBalance = OpeningBalanceRecord(rec)
				case Rh62MvmntRecId:
					stmt.movements = append(stmt.movements, MovementRecord{Record: rec})
				case Rh64ClosBalRecId:
					stmt.closingBalance = ClosingBalanceRecord(rec)
				case Rh65CashOnHandRecId:
					stmt.cashInHand = CashInHandRecord(rec)
				}
			}
		case ActionReturn:
			return stmt, nil
		}
	}

	return stmt, nil
}

func (r *RNDRHReader0010605) OnStatusStart(recId string) (string, string, error) {
	var err error
	var status string
	var action = ActionConsume
	if recId == Rh61OpenBalRecId {
		status = StatusOpenBalance
	} else if recId == RhEfRecId {
		action = ActionPushBack
		status = StatusEnd
	} else {
		err = fmt.Errorf("unexpected record %s in status %s", recId, status)
	}

	return status, action, err
}

func (r *RNDRHReader0010605) OnStatusOpenBalance(recId string) (string, string, error) {
	var err error
	var status string
	var action = ActionConsume

	switch recId {
	case Rh64ClosBalRecId:
		status = StatusCloseBalance
	case Rh62MvmntRecId:
		status = StatusMvmnt
	default:
		err = fmt.Errorf("unexpected record %s in status %s", recId, status)
	}

	return status, action, err
}

func (r *RNDRHReader0010605) OnStatusCloseBalance(recId string) (string, string, error) {
	var err error
	var status string
	var action = ActionConsume

	switch recId {
	case Rh65CashOnHandRecId:
		status = StatusCashOnHand
	case Rh61OpenBalRecId:
		status = StatusEnd
		action = ActionPushBack
	case RhEfRecId:
		status = StatusEnd
		action = ActionPushBack
	default:
		err = fmt.Errorf("unexpected record %s in status %s", recId, status)
	}
	return status, action, err
}

func (r *RNDRHReader0010605) OnStatusMovement(recId string) (string, string, error) {
	var err error
	var status string
	var action = ActionConsume

	switch recId {
	case Rh64ClosBalRecId:
		status = StatusCloseBalance
	case Rh62MvmntRecId:
		status = StatusMvmnt
	default:
		if IsRecIdMovementDetail(recId) {
			status = StatusMvmntDetail
		} else {
			err = fmt.Errorf("unexpected record %s in status %s", recId, status)
		}
	}
	return status, action, err
}

func (r *RNDRHReader0010605) OnStatusMovementDetail(recId string) (string, string, error) {
	var err error
	var status string
	var action = ActionConsume

	switch recId {
	case Rh64ClosBalRecId:
		status = StatusCloseBalance
	case Rh62MvmntRecId:
		status = StatusMvmnt
	default:
		if IsRecIdMovementDetail(recId) {
			status = StatusMvmntDetail
		} else {
			err = fmt.Errorf("unexpected record %s in status %s", recId, status)
		}
	}

	return status, action, err
}

func (r *RNDRHReader0010605) OnStatusCashInHand(recId string) (string, string, error) {
	var err error
	var status string
	var action = ActionConsume

	switch recId {
	case Rh61OpenBalRecId:
		status = StatusEnd
		action = ActionPushBack
	case RhEfRecId:
		status = StatusEnd
		action = ActionPushBack
	default:
		err = fmt.Errorf("unexpected record %s in status %s", recId, status)
	}
	return status, action, err
}

func IsRecIdMovementDetail(recId string) bool {
	if recId == Rh63MvmntDetKKKRecId ||
		recId == Rh63MvmntDetYYYRecId ||
		recId == Rh63MvmntDetYY2RecId ||
		recId == Rh63MvmntDetZZ1RecId ||
		recId == Rh63MvmntDetZZ2RecId ||
		recId == Rh63MvmntDetZZ3RecId ||
		recId == Rh63MvmntDetID1RecId ||
		recId == Rh63MvmntDetRI1RecId ||
		recId == Rh63RI2MvmntRecId ||
		recId == Rh63MvmntDetElseRecId {
		return true
	}

	return false
}

func (r *RNDRHReader0010605) Close() {
	if r.fileReader != nil {
		r.fileReader.Close()
	}

	r.fileReader = nil
}
