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
	"bytes"
	"encoding/json"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSourceReturnDictionary(t *testing.T) {
	jsonline := []byte(`{"personne": [{"name": "Benjamin", "age" : 35}, {"name": "Nicolas", "age" : 38}]}`)
	pipeline := model.NewPipeline(NewSource(bytes.NewReader(jsonline)))
	var result []model.Entry
	err := pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	expected := model.NewDictionary().
		With("personne", []model.Entry{
			model.NewDictionary().With("name", "Benjamin").With("age", json.Number("35")),
			model.NewDictionary().With("name", "Nicolas").With("age", json.Number("38")),
		})
	assert.Equal(t, expected, result[0], "Should create the right model.Dictionary")
}

func TestSourceReturnError(t *testing.T) {
	jsonline := []byte(`{"personne: [{"name": "Benjamin", "age" : 35}, {"name": "Nicolas", "age" : 38}]}`)
	pipeline := model.NewPipeline(NewSource(bytes.NewReader(jsonline)))
	var result []model.Entry
	err := pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	assert.NotNil(t, err)
	assert.EqualError(t, err, "json: invalid character a as null")
}

func TestSinkWriteDictionary(t *testing.T) {
	source := model.NewDictionary().
		With("personne", []model.Entry{
			model.NewDictionary().With("name", "Benjamin").With("age", json.Number("35")),
			model.NewDictionary().With("name", "Nicolas").With("age", json.Number("38")),
		})

	result := bytes.Buffer{}

	err := model.NewPipelineFromSlice([]model.Dictionary{source}).AddSink(NewSink(&result)).Run()
	jsonline := []byte(`{"personne":[{"name":"Benjamin","age":35},{"name":"Nicolas","age":38}]}
`)

	assert.Nil(t, err)
	assert.Equal(t, string(jsonline), result.String(), "Should create the right model.Dictionary")
}
