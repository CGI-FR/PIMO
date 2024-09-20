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
	"errors"
	"fmt"
	"io"
	"os"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/parquet-go/parquet-go"
	"github.com/parquet-go/parquet-go/deprecated"
	"github.com/rs/zerolog/log"
)

const BufferSize = 4

// NewSource creates a new Source.
func NewSource(path string) *Source {
	return &Source{path, nil, make([]parquet.Row, BufferSize), nil, 0, 0, false, model.NewDictionary(), nil, false, 0}
}

// NewPackedSource creates a new packed Source.
func NewPackedSource(path string) model.Source {
	log.Trace().Msg("NewPackedSource")

	return &Source{path, nil, make([]parquet.Row, BufferSize), nil, 0, 0, false, model.NewDictionary(), nil, true, 0}
}

// Source export line to JSON format.
type Source struct {
	path        string
	file        *parquet.File
	buffer      []parquet.Row
	rows        parquet.Rows
	groupIndex  int
	bufferIndex int
	endOfFile   bool
	value       model.Dictionary
	err         error
	packed      bool
	size        int
}

func (s *Source) Open() error {
	log.Trace().Msg("open parquet file")

	f, err := os.Open(s.path)
	if err != nil {
		return err
	}

	stat, err := os.Stat(s.path)
	if err != nil {
		return err
	}
	s.file, err = parquet.OpenFile(f, stat.Size())

	if err != nil {
		return err
	}

	s.groupIndex = 0
	s.readGroup()

	err = s.readRows()

	if err != nil {
		return err
	}

	return nil
}

func (s *Source) readGroup() {
	s.rows = s.file.RowGroups()[s.groupIndex].Rows()
}

func (s *Source) readRows() error {
	log.Trace().Int("groupIndex", s.groupIndex).Msg("read rows")

	size, err := s.rows.ReadRows(s.buffer)
	s.size = size
	s.bufferIndex = 0

	if errors.Is(err, io.EOF) {
		s.endOfFile = true
	} else if err != nil {
		return err
	}

	return nil
}

func (s *Source) readColumns() {
	log.Trace().Int("bufferIndex", s.bufferIndex).Msg("read columns")

	dict := model.NewDictionary()
	s.buffer[s.bufferIndex].Range(func(columnIndex int, columnValues []parquet.Value) bool {
		column := s.file.Schema().Columns()[columnIndex][0]
		switch columnValues[0].Kind() {
		case parquet.ByteArray:
			dict = dict.With(column, columnValues[0].String())
		case parquet.Boolean:
			dict = dict.With(column, columnValues[0].Boolean())
		case parquet.Int32:
			dict = dict.With(column, columnValues[0].Int32())
		case parquet.Int64:
			dict = dict.With(column, columnValues[0].Int64())
		case parquet.Int96:
			dict = dict.With(column, columnValues[0].Int64())
		case parquet.Float:
			dict = dict.With(column, columnValues[0].Float())
		case parquet.Double:
			dict = dict.With(column, columnValues[0].Double())
		}
		s.value = model.CleanDictionary(dict)
		return true
	})
	s.bufferIndex++
}

// Next convert next line to model.Dictionary
func (s *Source) Next() bool {
	if s.err != nil {
		return false
	}

	// read next buffer
	if s.bufferIndex >= s.size && !s.endOfFile {
		s.err = s.readRows()
		if s.err != nil {
			return false
		}
	}

	if s.bufferIndex < s.size {
		s.readColumns()
		return true
	}

	return false
}

func (s *Source) Value() model.Entry {
	return s.value
}

func (s *Source) Err() error {
	return s.err
}

// NewSink creates a new Sink.
func NewSink(file io.Writer, schema parquet.Node) model.SinkProcess {
	return Sink{file, parquet.Writer{}, "", schema}
}

// NewSinkWithContext creates a new Sink.
func NewSinkWithContext(file io.Writer, counter string, schema parquet.Node) model.SinkProcess {
	over.MDC().Set(counter, 1)
	return Sink{file, parquet.Writer{}, counter, schema}
}

type Sink struct {
	file    io.Writer
	writer  parquet.Writer
	counter string
	schema  parquet.Node
}

func (s Sink) Open() error {
	config, err := parquet.NewWriterConfig()
	if err != nil {
		return err
	}

	s.writer = *parquet.NewWriter(s.file, config)
	return nil
}

func (s Sink) ProcessDictionary(dictionary model.Entry) error {
	log.Trace().Msg("write to parquet")

	dico, ok := dictionary.(model.Dictionary)

	if !ok {
		return errors.New("Can't write entry that not dictonary")
	}

	rowBuilder := parquet.NewRowBuilder(s.schema)

	for columnIndex, field := range s.schema.Fields() {
		entry, ok := dico.GetValue(field.Name())

		if !ok {
			return fmt.Errorf("Can't find %s in dictionary", field.Name())
		}

		switch field.Type() {
		case parquet.ByteArrayType:
			rowBuilder.Add(columnIndex, parquet.ByteArrayValue([]byte(entry.(string))))
		case parquet.DoubleType:
			rowBuilder.Add(columnIndex, parquet.DoubleValue(entry.(float64)))
		case parquet.FloatType:
			rowBuilder.Add(columnIndex, parquet.FloatValue(entry.(float32)))
		case parquet.Int96Type:
			rowBuilder.Add(columnIndex, parquet.Int96Value(deprecated.Int64ToInt96(int64(entry.(float64)))))
		case parquet.Int64Type:
			rowBuilder.Add(columnIndex, parquet.Int64Value(int64(entry.(float64))))
		case parquet.Int32Type:
			rowBuilder.Add(columnIndex, parquet.Int32Value(entry.(int32)))
		case parquet.BooleanType:
			rowBuilder.Add(columnIndex, parquet.BooleanValue(entry.(bool)))
		default:
			return fmt.Errorf("can't find how to tranform %s", field.Type())
		}

	}

	writer := parquet.NewWriter(s.file, parquet.NewSchema("test", s.schema))
	row := rowBuilder.Row()
	_, err := writer.WriteRows([]parquet.Row{row})

	return err
}
