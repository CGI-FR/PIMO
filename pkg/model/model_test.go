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

package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelineSource(t *testing.T) {
	mySlice := []Dictionary{{"v": 1}, {"v": 2}, {"v": 3}, {"v": 4}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)
	assert.Equal(t, mySlice, result)
}

func TestPipelineWithProcessorSource(t *testing.T) {
	mySlice := []Dictionary{{"v": 1}, {"v": 2}, {"v": 3}, {"v": 4}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d["v"].(int)
			return Dictionary{"v": value + 1}, nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"v": 2}, {"v": 3}, {"v": 4}, {"v": 5}}
	assert.Equal(t, wanted, result)
}

func TestPipelineWithChainedProcessorSource(t *testing.T) {
	mySlice := []Dictionary{{"v": 1}, {"v": 2}, {"v": 3}, {"v": 4}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d["v"].(int)
			return Dictionary{"v": value + 1}, nil
		})).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d["v"].(int)
			return Dictionary{"v": value * value}, nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"v": 4}, {"v": 9}, {"v": 16}, {"v": 25}}
	assert.Equal(t, wanted, result)
}

func TestPipelineWithRepeaterProcessor(t *testing.T) {
	mySlice := []Dictionary{{"v": 1}, {"v": 2}, {"v": 3}, {"v": 4}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewRepeaterProcess(2)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"v": 1}, {"v": 1}, {"v": 2}, {"v": 2}, {"v": 3}, {"v": 3}, {"v": 4}, {"v": 4}}
	assert.Equal(t, wanted, result)
}

func TestPipelineWithRepeaterAndMapChainedProcessor(t *testing.T) {
	mySlice := []Dictionary{{"v": 1}, {"v": 2}, {"v": 3}, {"v": 4}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewRepeaterProcess(2)).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d["v"].(int)
			return Dictionary{"v": value * value}, nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"v": 1}, {"v": 1}, {"v": 4}, {"v": 4}, {"v": 9}, {"v": 9}, {"v": 16}, {"v": 16}}
	assert.Equal(t, wanted, result)
}

func TestPipelineWithMaskEngine(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Toto", nil }}

	mySlice := []Dictionary{{"name": "Bob"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("name"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"name": "Toto"}}
	assert.Equal(t, wanted, result)
}

func TestPipelineWithDeleteMaskEngine(t *testing.T) {
	mySlice := []Dictionary{{"name": "Bob", "city": "Nantes"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewDeleteMaskEngineProcess(NewSimplePathSelector("name"))).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"city": "Nantes"}}
	assert.Equal(t, wanted, result)

	assert.NotEqual(t, wanted, mySlice)
}

func TestMaskEngineShouldNotCreateField(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Toto", nil }}

	mySlice := []Dictionary{{"city": "Nantes"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("name"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"city": "Nantes"}}
	assert.Equal(t, wanted, result)
}

// MaskEngine is a value that will be the initialisation of the field when it's created
type TestAddMaskEngine struct {
	value Entry
}

// MaskContext add the field
func (am TestAddMaskEngine) MaskContext(context Dictionary, key string, contexts ...Dictionary) (Dictionary, error) {
	_, present := context[key]
	if !present {
		context[key] = am.value
	}

	return context, nil
}

func TestMaskEngineWithContext(t *testing.T) {
	mySlice := []Dictionary{{"city": "Nantes"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskContextEngineProcess(NewSimplePathSelector("name"), TestAddMaskEngine{"Toto"})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"city": "Nantes", "name": "Toto"}}
	assert.Equal(t, wanted, result)
}

func TestMaskEngineShouldMaskAllEntriesInArray(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{{"city": []Entry{"Nantes", "Rennes", "Grenoble"}}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("city"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"city": []Entry{"Paris", "Paris", "Paris"}}}
	assert.Equal(t, wanted, result)
}

func TestMaskEngineShouldMaskNestedEntry(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{{"address": map[string]Entry{"city": "Nantes"}}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("address.city"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"address": map[string]Entry{"city": "Paris"}}}
	assert.Equal(t, wanted, result)
}

func TestMaskEngineShouldMaskNestedDictionariesArray(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{
		{"address": []Entry{
			Dictionary{"city": "Nantes"},
			Dictionary{"city": "Rennes"},
			Dictionary{"city": "Grenoble"},
		}},
	}

	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("address.city"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"address": []Dictionary{
			{"city": "Paris"},
			{"city": "Paris"},
			{"city": "Paris"},
		}},
	}
	assert.Equal(t, wanted, result)
}

func TestMaskEngineShouldMaskNestedArray(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{
		{"address": Dictionary{"city": []Entry{
			"Nantes",
			"Rennes",
			"Grenoble",
		}}},
	}

	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("address.city"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"address": Dictionary{"city": []Entry{
			"Paris",
			"Paris",
			"Paris",
		}},
		},
	}
	assert.Equal(t, wanted, result)
}

func jsonlineToDictionaries(jsl string) []Dictionary {
	result := []Dictionary{}
	jsons := strings.Split(jsl, "\n")
	for _, js := range jsons {
		var inter interface{}
		err := json.Unmarshal(([]byte)(js), &inter)
		if err != nil {
			return nil
		}
		result = append(result, InterfaceToDictionary(inter))
	}
	return result
}

func dictionariesToJSONLine(dictionaries []Dictionary) string {
	var jsonline []byte
	for _, dictionary := range dictionaries {
		result, _ := json.Marshal(dictionary)
		jsonline = append(jsonline, result...)
		jsonline = append(jsonline, "\n"...)
	}
	return strings.TrimSuffix(string(jsonline), "\n")
}

func TestMaskEngineShouldMaskNestedArrays(t *testing.T) {
	var i = 0
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}

	iput := `{"persons":[{"phonenumber":"027123456"},{"phonenumber":"028123456"}]}
			 {"persons":[{"phonenumber":"027123456"}]}
			 {"persons":[]}`
	oput := `{"persons":[{"phonenumber":"1"},{"phonenumber":"2"}]}
			 {"persons":[{"phonenumber":"3"}]}
			 {"persons":[]}`

	mySlice := jsonlineToDictionaries(iput)

	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("persons.phonenumber"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	expected := dictionariesToJSONLine(jsonlineToDictionaries(oput))
	actual := dictionariesToJSONLine(result)
	assert.Equal(t, expected, actual)
}

func TestMaskEngineShouldMaskNestedNestedArrays(t *testing.T) {
	var i = 0
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}

	iput := `{"elements":[{"persons":[{"phonenumber":"027123456"},{"phonenumber":"028123456"}]},{"persons":[{"phonenumber":"029123456"},{"phonenumber":"020123456"}]}]}`
	oput := `{"elements":[{"persons":[{"phonenumber":"1"},{"phonenumber":"2"}]},{"persons":[{"phonenumber":"3"},{"phonenumber":"4"}]}]}`

	mySlice := jsonlineToDictionaries(iput)

	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("elements.persons.phonenumber"), nameMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	expected := dictionariesToJSONLine(jsonlineToDictionaries(oput))
	actual := dictionariesToJSONLine(result)
	assert.Equal(t, expected, actual)
}

func TestMaskEngineShouldMaskMultipleNestedNestedArrays(t *testing.T) {
	var i = 0
	var j = 0
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}
	var emailMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		j++
		return fmt.Sprintf("email.%d@company.com", j), nil
	}}

	iput := `{"elements":[{"persons":[{"phonenumber":"027123456","email":"person1@company.com"},{"phonenumber":"028123456","email":"person2@company.com"}]},{"persons":[{"phonenumber":"029123456","email":"person3@company.com"},{"phonenumber":"020123456","email":"person4@company.com"}]}]}`
	oput := `{"elements":[{"persons":[{"phonenumber":"1","email":"email.1@company.com"},{"phonenumber":"2","email":"email.2@company.com"}]},{"persons":[{"phonenumber":"3","email":"email.3@company.com"},{"phonenumber":"4","email":"email.4@company.com"}]}]}`

	mySlice := jsonlineToDictionaries(iput)

	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("elements.persons.phonenumber"), nameMasking)).
		Process(NewMaskEngineProcess(NewPathSelector("elements.persons.email"), emailMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	expected := dictionariesToJSONLine(jsonlineToDictionaries(oput))
	actual := dictionariesToJSONLine(result)
	assert.Equal(t, expected, actual)
}

func TestMaskEngineShouldReturnError(t *testing.T) {
	var errorMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "", fmt.Errorf("Test error") }}

	mySlice := []Dictionary{{"city": "Nantes"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("city"), errorMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.NotNil(t, err)
}

func TestPipelineShouldReturnError(t *testing.T) {
	var errorMap = func(Dictionary) (Dictionary, error) {
		return nil, fmt.Errorf("Test error")
	}

	mySlice := []Dictionary{{"city": "Nantes"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMapProcess(errorMap)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.NotNil(t, err)
}

type ErrorSource struct {
}

func (s ErrorSource) Err() error {
	return fmt.Errorf("Test error")
}
func (s ErrorSource) Next() bool {
	return false
}

func (s ErrorSource) Value() Dictionary {
	return nil
}

func (s ErrorSource) Open() error {
	return nil
}

func TestPipelineShouldReturnErrorFromSource(t *testing.T) {
	var result []Dictionary

	pipeline := NewPipeline(ErrorSource{}).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.NotNil(t, err)
}

func TestCacheShouldProvide(t *testing.T) {
	i := 0

	var errorMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		i++
		return fmt.Sprintf("%s - %d", name, i), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{{"city": "Nantes"}, {"city": "Grenoble"}, {"city": "Nantes"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("city"), NewMaskCacheEngine(cache, errorMasking))).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"city": "Nantes - 1"}, {"city": "Grenoble - 2"}, {"city": "Nantes - 1"}}

	assert.Equal(t, wanted, result)
}
