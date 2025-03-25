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

//go:build !wasm && !wasi && !(linux && 386) && !(windows && 386)
// +build !wasm
// +build !wasi
// +build !linux !386
// +build !windows !386

package parquet

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"

	over "github.com/adrienaury/zeromdc"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
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
func NewPackedSource(path string) (*Source, error) {
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

func (s *Source) Schema() (*arrow.Schema, error) {
	return s.fileReader.Schema()
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

func (s *Source) Close() error {
	log.Trace().Msg("close parquet file")
	return s.fileReader.ParquetReader().Close()
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

func (s *Source) Values() iter.Seq2[model.Dictionary, error] {
	return func(yield func(model.Dictionary, error) bool) {
		for s.Next() {
			if !yield(s.Value().(model.Dictionary), nil) {
				return
			}
		}
	}
}

// NewSink creates a new Sink.
func NewSink(file io.Writer, schema *arrow.Schema) (Sink, error) {
	mem := memory.NewCheckedAllocator(memory.DefaultAllocator)
	writer, err := pqarrow.NewFileWriter(schema, file, nil, pqarrow.NewArrowWriterProperties(pqarrow.WithAllocator(mem)))
	if err != nil {
		return Sink{}, err
	}

	return Sink{writer, schema, ""}, nil
}

// NewSinkWithContext creates a new Sink.
func NewSinkWithContext(file io.Writer, schema *arrow.Schema, counter string) (Sink, error) {
	over.MDC().Set(counter, 1)
	result, err := NewSink(file, schema)
	if err != nil {
		return result, err
	}

	result.counter = counter
	return result, nil
}

type Sink struct {
	file    *pqarrow.FileWriter
	schema  *arrow.Schema
	counter string
}

func (s Sink) Open() error {
	return nil
}

func (s Sink) Close() error {
	return s.file.Close()
}

func (s Sink) ProcessDictionary(dictionary model.Entry) error {
	log.Trace().Msg("write to parquet")

	dico, ok := dictionary.(model.Dictionary)

	if !ok {
		return fmt.Errorf("Can't write entry that not dictonary %v", dictionary)
	}

	jsonline, err := json.Marshal(dico)
	if err != nil {
		return err
	}

	jsonReader := bytes.NewReader(jsonline)

	reader := array.NewJSONReader(jsonReader, s.schema)

	if !reader.Next() {
		return fmt.Errorf("can't convert row to arrow schema")
	}

	err = s.file.WriteBuffered(reader.Record())
	if err != nil {
		return err
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
