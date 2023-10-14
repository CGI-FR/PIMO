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
	"bytes"
	"io"

	json "github.com/goccy/go-json"

	"github.com/adrienaury/flexjson"
	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/model"
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
	decoder := json.NewDecoder(bytes.NewReader(jsonline))
	flex := flexjson.NewFlexDecoder(
		decoder,
		func() (model.Dictionary, error) { return model.NewDictionary(), nil },
		func(obj model.Dictionary, key string, value any) (model.Dictionary, error) {
			obj.Set(key, value)
			return obj, nil
		},
		flexjson.StandardArrayMaker(),
		flexjson.StandardArrayAdder(),
	)

	dict, err := flex.DecodeObject()
	if err != nil {
		return model.NewDictionary(), err
	}

	return model.CleanDictionary(dict), nil
}

// JSONToPackedDictionary return a packed model.Dictionary from a jsonline
func JSONToPackedDictionary(jsonline []byte) (model.Dictionary, error) {
	decoder := json.NewDecoder(bytes.NewReader(jsonline))
	flex := flexjson.NewFlexDecoder(
		decoder,
		func() (model.Dictionary, error) { return model.NewDictionary(), nil },
		func(obj model.Dictionary, key string, value any) (model.Dictionary, error) {
			obj.Set(key, value)
			return obj, nil
		},
		flexjson.StandardArrayMaker(),
		flexjson.StandardArrayAdder(),
	)

	dict, err := flex.DecodeObject()
	if err != nil {
		return model.NewDictionary(), err
	}

	// packer
	root := dict.Pack().With("original", string(jsonline))

	return model.CleanDictionary(root), nil
}
