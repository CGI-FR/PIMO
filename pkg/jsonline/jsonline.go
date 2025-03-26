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

package jsonline

import (
	"bufio"
	"io"
	"iter"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/goccy/go-json"
)

// NewSource creates a new Source.
func NewSource(file io.Reader) model.Source {
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024*100) // increase buffer up to 100 MB
	return &Source{scanner, model.NewDictionary(), nil, false}
}

// NewPackedSource creates a new packed Source.
func NewPackedSource(file io.Reader) model.Source {
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024*100) // increase buffer up to 100 MB
	return &Source{scanner, model.NewDictionary(), nil, true}
}

// Source export line to JSON format.
type Source struct {
	fscanner *bufio.Scanner
	value    model.Dictionary
	err      error
	packed   bool
}

func (s *Source) Open() error {
	return nil
}

func (s *Source) Close() error {
	return nil
}

func (s *Source) Values() iter.Seq2[model.Dictionary, error] {
	return func(yield func(model.Dictionary, error) bool) {
		for s.fscanner.Scan() {
			line := s.fscanner.Bytes()

			var dict model.Dictionary
			var err error

			if s.packed {
				dict, err = JSONToPackedDictionary(line)
			} else {
				dict, err = JSONToDictionary(line)
			}

			if err != nil {
				yield(dict, err)
				return
			}

			if !yield(dict, nil) {
				return
			}
		}

		if s.fscanner.Err() != nil {
			yield(model.NewDictionary(), s.fscanner.Err())
		}
	}
}

// Next convert next line to model.Dictionary
func (s *Source) Next() bool {
	if !s.fscanner.Scan() {
		s.err = s.fscanner.Err()
		return false
	}
	line := s.fscanner.Bytes()
	if s.packed {
		s.value, s.err = JSONToPackedDictionary(line)
	} else {
		s.value, s.err = JSONToDictionary(line)
	}
	return s.err == nil
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
	jsonline, err := json.Marshal(dictionary)
	if err != nil {
		return err
	}
	jsonline = append(jsonline, "\n"...)

	_, err = s.file.Write(jsonline)
	if err != nil {
		return err
	}

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
	err := dict.UnmarshalJSON(jsonline)
	return model.CleanDictionary(dict), err
}

// JSONToPackedDictionary return a packed model.Dictionary from a jsonline
func JSONToPackedDictionary(jsonline []byte) (model.Dictionary, error) {
	dict, err := JSONToDictionary(jsonline)
	return dict.Pack().With("original", string(jsonline)), err
}
