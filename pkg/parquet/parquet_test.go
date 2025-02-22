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
	"os"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSourceReturnDictionary(t *testing.T) {
	source, err := NewSource("testdata/alltypes_plain.parquet")
	assert.Nil(t, err)

	pipeline := model.NewPipeline(source)
	var result []model.Entry
	err = pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	assert.Nil(t, err)

	assert.Equal(t, 8, len(result))
	document := result[0].(model.Dictionary).Get(".").(model.Dictionary)

	assert.Equal(t, true, document.Get("bool_col"))
	assert.Equal(t, float64(0), document.Get("tinyint_col"))
	assert.Equal(t, float64(0), document.Get("smallint_col"))
	assert.Equal(t, float64(0), document.Get("int_col"))
	assert.Equal(t, float64(0), document.Get("bigint_col"))
	assert.Equal(t, float64(0), document.Get("float_col"))
	assert.Equal(t, float64(0), document.Get("double_col"))
	// assert.Equal(t, "03/01/09", document.Get("date_string_col"))
	// assert.Equal(t, "0", document.Get("string_col"))
	assert.Equal(t, "2009-03-01 00:00:00", document.Get("timestamp_col"))
}

func TestSinkWriteDictionary(t *testing.T) {
	source, err := NewSource("testdata/alltypes_plain.parquet")
	assert.Nil(t, err)

	schema, err := source.Schema()
	assert.Nil(t, err)

	fileWriter, err := os.Create("testdata/alltypes_plain_out.parquet")
	assert.Nil(t, err)
	sink, err := NewSink(fileWriter, schema)
	assert.Nil(t, err)
	pipeline := model.NewPipeline(source)
	err = pipeline.AddSink(sink).Run()
	assert.Nil(t, err)

	err = sink.Close()
	assert.Nil(t, err)
}
