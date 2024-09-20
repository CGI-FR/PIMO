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
	"fmt"
	"os"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/parquet-go/parquet-go"
	"github.com/stretchr/testify/assert"
)

func TestSourceReturnDictionary(t *testing.T) {
	pipeline := model.NewPipeline(NewSource("testdata/alltypes_plain.parquet"))
	var result []model.Entry
	err := pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	assert.Nil(t, err)

	assert.Equal(t, 8, len(result))
	document := result[0].(model.Dictionary)

	assert.Equal(t, true, document.Get("bool_col"))
	assert.Equal(t, int32(0), document.Get("tinyint_col"))
	assert.Equal(t, int32(0), document.Get("smallint_col"))
	assert.Equal(t, int32(0), document.Get("int_col"))
	assert.Equal(t, float64(0), document.Get("bigint_col"))
	assert.Equal(t, float32(0), document.Get("float_col"))
	assert.Equal(t, float64(0), document.Get("double_col"))
	assert.Equal(t, "03/01/09", document.Get("date_string_col"))
	assert.Equal(t, "0", document.Get("string_col"))
	// assert.Equal(t, int64(0), document.Get("timestamp_col"))
}

func TestSinkWriteDictionary(t *testing.T) {
	pathIn := "testdata/alltypes_plain.parquet"
	pathOut := "/tmp/test.parquet"

	source := NewSource(pathIn)
	source.Open()
	schema := source.file.Schema()

	source = NewSource(pathIn)

	fileWriter, _ := os.Create(pathOut)

	err := model.
		NewPipeline(source).
		AddSink(NewSink(fileWriter, schema)).
		Run()
	assert.Nil(t, err)

	fileWriter.Close()

	file, err := os.Open(pathOut)
	assert.Nil(t, err)

	stat, err := os.Stat(pathOut)
	assert.Nil(t, err)

	rows, err := parquet.Read[any](file, stat.Size(), schema)
	assert.Nil(t, err)

	assert.Len(t, rows, 1)

	for _, row := range rows {
		fmt.Printf("%v\n", row)
	}
}
