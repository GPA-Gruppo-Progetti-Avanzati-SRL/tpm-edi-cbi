package cbirnd0010605

import (
	"bytes"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fixedlengthfile/reader"
	"github.com/rs/zerolog/log"
)

type RNDReader0010605 struct {
	isError    bool
	fileReader reader.Reader
}

func NewReader(dataContent []byte) (*RNDReader0010605, error) {
	const semLogContext = "cbi-rnd-001.06.05::new-reader"

	strReader := bytes.NewReader(dataContent)
	flr, err := reader.NewReader(fixedLengthReaderConfig, reader.WithIoReader(strReader))
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	r := RNDReader0010605{fileReader: flr}
	return &r, nil
}

func (r *RNDReader0010605) Read() (reader.Record, error) {
	const semLogContext = "cbi-rnd-001.06.05::read"

	if r.isError {
		return reader.Record{}, errors.New("cbi-rnd-001.06.05 already in error state")
	}

	rec, err := r.fileReader.Read()
	if err != nil {
		r.isError = true
		return reader.Record{}, err
	}

	return rec, nil
}

func (r *RNDReader0010605) Close() {
	if r.fileReader != nil {
		r.fileReader.Close()
	}

	r.fileReader = nil
}
