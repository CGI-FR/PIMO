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
	"io"
	"os"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/parquet-go/parquet-go"
	"github.com/rs/zerolog/log"
)

const BufferSize = 4

// NewSource creates a new Source.
func NewSource(path string) model.Source {
	return &Source{path, nil, make([]parquet.Row, BufferSize), 0, model.NewDictionary(), nil, false, 0}
}

// NewPackedSource creates a new packed Source.
func NewPackedSource(path string) model.Source {
	log.Trace().Msg("NewPackedSource")

	return &Source{path, nil, make([]parquet.Row, BufferSize), 0, model.NewDictionary(), nil, true, 0}
}

// Source export line to JSON format.
type Source struct {
	path       string
	file       *parquet.File
	buffer     []parquet.Row
	groupIndex int
	value      model.Dictionary
	err        error
	packed     bool
	size       int
}

func (s *Source) Open() error {
	log.Trace().Msg("open parquet file")

	f, err := os.Open(s.path)
	print(err)
	if err != nil {
		return err
	}

	stat, err := os.Stat(s.path)
	if err != nil {
		return err
	}
	s.file, err = parquet.OpenFile(f, stat.Size())

	print(s.file.Schema().String())

	if err != nil {
		return err
	}

	s.groupIndex = 0
	log.Printf("%v", s.file.RowGroups())

	s.size, err = s.file.RowGroups()[0].Rows().ReadRows(s.buffer)

	if err != nil && errors.Is(err, io.EOF) {
		return err
	}

	return nil
}

func (s *Source) readColumn() {
	dict := model.NewDictionary()
	s.buffer[s.groupIndex].Range(func(columnIndex int, columnValues []parquet.Value) bool {
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
		s.groupIndex++
		return true
	})
}

// Next convert next line to model.Dictionary
func (s *Source) Next() bool {
	// read next buffer
	if s.groupIndex < BufferSize {
		s.readColumn()
		return true
	}

	return false
	for _, rowGroup := range s.file.RowGroups() {
		rows := rowGroup.Rows()
		_, err := rows.ReadRows(s.buffer)
		if err != nil {
			s.err = err
			return false
		}
	}
	s.buffer[0].Range(func(columnIndex int, columnValues []parquet.Value) bool {
		print(columnValues[0].GoString())
		return true
	})
	return false // s.err == nil
}

func (s *Source) Value() model.Entry {
	return s.value
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
	if len(s.counter) > 0 {
		value, exists := over.MDC().Get(s.counter)
		if !exists {
			return nil
		}

		if counter, ok := value.(int); ok {
			over.MDC().Set(s.counter, counter+1)
		}
	}

	return nil
}

// JSONToDictionary return a model.Dictionary from a jsonline
func JSONToDictionary(jsonline []byte) (model.Dictionary, error) {
	dict := model.NewDictionary()

	return model.CleanDictionary(dict), nil
}

// JSONToPackedDictionary return a packed model.Dictionary from a jsonline
func JSONToPackedDictionary(jsonline []byte) (model.Dictionary, error) {
	dict := model.NewDictionary()

	// packer
	root := dict.Pack().With("original", string(jsonline))

	return model.CleanDictionary(root), nil
}
