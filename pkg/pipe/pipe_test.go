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

package pipe_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/pipe"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/stretchr/testify/assert"
)

func TestMaskEngineShouldMaskNestedArray(t *testing.T) {
	mySlice := []model.Dictionary{
		model.NewDictionary().With("address", model.NewDictionary().With("city", []model.Dictionary{
			model.NewDictionary().With("name", "Nantes"),
			// model.NewDictionary().With("name", "Rennes"),
			// model.NewDictionary().With("name", "Grenoble"),
		})),
	}

	var result []model.Entry

	model.InjectMaskFactories([]model.MaskFactory{templatemask.Factory})
	pipe, err := pipe.NewMask(42, "parent", "root", map[string]model.Cache{}, nil, "",
		model.Masking{
			Selector: model.SelectorType{Jsonpath: "name"},
			Mask: model.MaskType{
				Template: "{{.name}}+",
			},
		})
	assert.Nil(t, err)

	pipeline := model.NewPipelineFromSlice(mySlice).
		Process(model.NewRepeaterProcess(2)).
		Process(model.NewMaskContextEngineProcess(model.NewPathSelector("address.city"), pipe, "")).
		AddSink(model.NewSinkToSlice(&result))

	err = pipeline.Run()

	assert.Nil(t, err)

	wanted := []model.Dictionary{
		model.NewDictionary().With("address", model.NewDictionary().With("city", []model.Dictionary{
			model.NewDictionary().With("name", "Nantes+"),
		})),
		model.NewDictionary().With("address", model.NewDictionary().With("city", []model.Dictionary{
			model.NewDictionary().With("name", "Nantes+"),
		})),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func DictionariesToJSONLine(dictionaries []model.Dictionary) string {
	var jsonline []byte
	for _, dictionary := range dictionaries {
		result, _ := json.Marshal(dictionary)
		jsonline = append(jsonline, result...)
		jsonline = append(jsonline, "\n"...)
	}
	return strings.TrimSuffix(string(jsonline), "\n")
}

func EntriesToJSONLine(dictionaries []model.Entry) string {
	var jsonline []byte
	for _, dictionary := range dictionaries {
		result, _ := json.Marshal(dictionary)
		jsonline = append(jsonline, result...)
		jsonline = append(jsonline, "\n"...)
	}
	return strings.TrimSuffix(string(jsonline), "\n")
}
