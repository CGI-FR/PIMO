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
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestSourceReturnDictionary(t *testing.T) {
	jsonline := []byte(`{"personne": [{"name": "Benjamin", "age" : 35}, {"name": "Nicolas", "age" : 38}]}`)
	pipeline := model.NewPipeline(NewSource(bytes.NewReader(jsonline)))
	var result []model.Dictionary
	err := pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	waited := model.Dictionary{"personne": []model.Entry{
		model.Dictionary{"name": "Benjamin", "age": float64(35)},
		model.Dictionary{"name": "Nicolas", "age": float64(38)},
	}}
	assert.Equal(t, waited, result[0], "Should create the right model.Dictionary")
}

func TestSourceReturnError(t *testing.T) {
	jsonline := []byte(`{"personne" [{"name": "Benjamin", "age" : 35}, {"name": "Nicolas", "age" : 38}]}`)
	pipeline := model.NewPipeline(NewSource(bytes.NewReader(jsonline)))
	var result []model.Dictionary
	err := pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid character '[' after object key")
}

func TestSinkWriteDictionary(t *testing.T) {
	source := model.Dictionary{"personne": []model.Entry{
		model.Dictionary{"name": "Benjamin", "age": float64(35)},
		model.Dictionary{"name": "Nicolas", "age": float64(38)},
	}}

	result := bytes.Buffer{}

	err := model.NewPipelineFromSlice([]model.Dictionary{source}).AddSink(NewSink(&result)).Run()
	jsonline := []byte(`{"personne":[{"age":35,"name":"Benjamin"},{"age":38,"name":"Nicolas"}]}
`)

	assert.Nil(t, err)
	assert.Equal(t, string(jsonline), result.String(), "Should create the right model.Dictionary")
}
