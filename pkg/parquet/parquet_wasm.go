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

//go:build wasi || wasm
// +build wasi wasm

package parquet

import (
	"io"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cgi-fr/pimo/pkg/model"
)

type Source struct{}

func NewSource(path string) (*Source, error) {
	panic("parquet is not supported on your environment")
}

func NewPackedSource(path string) (*Source, error) {
	panic("parquet is not supported on your environment")
}

func (s *Source) Schema() (*arrow.Schema, error) {
	panic("parquet is not supported on your environment")
}

func (s *Source) Open() error {
	panic("parquet is not supported on your environment")
}

func (s *Source) NextBatch() bool {
	panic("parquet is not supported on your environment")
}

func (s *Source) Next() bool {
	panic("parquet is not supported on your environment")
}

func (s *Source) Value() model.Entry {
	panic("parquet is not supported on your environment")
}

func (s *Source) Err() error {
	panic("parquet is not supported on your environment")
}

func NewSink(file io.Writer, schema *arrow.Schema) (Sink, error) {
	panic("parquet is not supported on your environment")
}

func NewSinkWithContext(file io.Writer, schema *arrow.Schema, counter string) (Sink, error) {
	panic("parquet is not supported on your environment")
}

type Sink struct{}

func (s Sink) Open() error {
	panic("parquet is not supported on your environment")
}

func (s Sink) Close() error {
	panic("parquet is not supported on your environment")
}

func (s Sink) ProcessDictionary(dictionary model.Entry) error {
	panic("parquet is not supported on your environment")
}

func JSONToArray(buffer []byte) ([]model.Dictionary, error) {
	panic("parquet is not supported on your environment")
}

func JSONToPackedDictionary(jsonline []byte) (model.Dictionary, error) {
	panic("parquet is not supported on your environment")
}
