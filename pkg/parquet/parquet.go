// Copyright (C) 2021 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package parquet

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	over "github.com/adrienaury/zeromdc"
	"github.com/apache/arrow/go/v12/parquet/file"
	"github.com/apache/arrow/go/v12/parquet/pqarrow"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

const BufferSize = 4

// NewSource creates a new Source.
func NewSource(path string) (*Source, error) {
	reader, err := file.OpenParquetFile(path, true)
	if err != nil {
		return nil, err
	}

	fileReader, err := pqarrow.NewFileReader(reader, pqarrow.ArrowReadProperties{BatchSize: 10}, nil)
	if err != nil {
		return nil, err
	}

	return &Source{fileReader, []model.Dictionary{}, 0, nil, false, nil}, nil
}

// NewPackedSource creates a new packed Source.
func NewPackedSource(path string) (model.Source, error) {
	log.Trace().Msg("NewPackedSource")

	result, err := NewSource(path)
	if err != nil {
		return nil, err
	}

	result.packed = true
	return result, err
}

// Source export line to JSON format.
type Source struct {
	fileReader   *pqarrow.FileReader
	batch        []model.Dictionary
	index        int
	err          error
	packed       bool
	recordReader pqarrow.RecordReader
}

func (s *Source) Open() error {
	log.Trace().Msg("open parquet file")
	var err error
	s.recordReader, err = s.fileReader.GetRecordReader(context.Background(), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// NextBatch read next batch in parquet file
func (s *Source) NextBatch() bool {
	log.Trace().Msg("read batch of rows")

	if next := s.recordReader.Next(); !next {
		return false
	}

	record := s.recordReader.Record()

	buf, err := record.MarshalJSON()
	if err != nil {
		s.err = err
		return false
	}

	s.batch, s.err = JSONToArray(buf)

	s.index = 0
	return s.err == nil
}

func (s *Source) Next() bool {
	s.index++
	if s.index < len(s.batch) {
		return true
	}

	return s.NextBatch()
}

func (s *Source) Value() model.Entry {
	return s.batch[s.index]
}

func (s *Source) Err() error {
	return s.err
}

// NewSink creates a new Sink.
func NewSink(file io.Writer) model.SinkProcess {
	return Sink{file, ""}
}

// NewSinkWithContext creates a new Sink.
func NewSinkWithContext(file io.Writer, counter string) model.SinkProcess {
	over.MDC().Set(counter, 1)
	return Sink{file, counter}
}

type Sink struct {
	file    io.Writer
	counter string
}

func (s Sink) Open() error {
	return nil
}

func (s Sink) ProcessDictionary(dictionary model.Entry) error {
	log.Trace().Msg("write to parquet")

	_, ok := dictionary.(model.Dictionary)

	if !ok {
		return errors.New("Can't write entry that not dictonary")
	}

	return nil
}

// JSONToArray return a array of model.Dictionary from a jsonline
func JSONToArray(buffer []byte) ([]model.Dictionary, error) {
	result := []model.Dictionary{}

	var rows []json.RawMessage
	log.Debug().Str("jsonline", string(buffer)).Msg("decode")
	err := json.Unmarshal(buffer, &rows)
	if err != nil {
		return result, err
	}

	for _, dict := range rows {
		cleanDict, err := jsonline.JSONToPackedDictionary(dict)
		if err != nil {
			return nil, err
		}
		result = append(result, cleanDict)
	}

	return result, nil
}

// JSONToPackedDictionary return a packed model.Dictionary from a jsonline
func JSONToPackedDictionary(jsonline []byte) (model.Dictionary, error) {
	dict := model.NewDictionary()

	err := json.Unmarshal(jsonline, &dict)
	if err != nil {
		return model.NewDictionary(), err
	}

	// packer
	root := dict.Pack().With("original", string(jsonline))

	return model.CleanDictionary(root), nil
}
