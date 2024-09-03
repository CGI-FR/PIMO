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
	mySlice := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)
	assert.Equal(t, DictionariesToJSONLine(mySlice), EntriesToJSONLine(result))
}

func TestPipelineWithProcessorSource(t *testing.T) {
	mySlice := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d.Get("v").(int)
			return NewDictionary().With("v", value+1), nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
		NewDictionary().With("v", 5),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestPipelineWithChainedProcessorSource(t *testing.T) {
	mySlice := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d.Get("v").(int)
			return NewDictionary().With("v", value+1), nil
		})).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d.Get("v").(int)
			return NewDictionary().With("v", value*value), nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("v", 4),
		NewDictionary().With("v", 9),
		NewDictionary().With("v", 16),
		NewDictionary().With("v", 25),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestPipelineWithRepeaterProcessor(t *testing.T) {
	mySlice := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewRepeaterProcess(2)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
		NewDictionary().With("v", 4),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestPipelineWithRepeaterAndMapChainedProcessor(t *testing.T) {
	mySlice := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 2),
		NewDictionary().With("v", 3),
		NewDictionary().With("v", 4),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewRepeaterProcess(2)).
		Process(NewMapProcess(func(d Dictionary) (Dictionary, error) {
			value := d.Get("v").(int)
			return NewDictionary().With("v", value*value), nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 1),
		NewDictionary().With("v", 4),
		NewDictionary().With("v", 4),
		NewDictionary().With("v", 9),
		NewDictionary().With("v", 9),
		NewDictionary().With("v", 16),
		NewDictionary().With("v", 16),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestPipelineWithMaskEngine(t *testing.T) {
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Toto", nil }}

	mySlice := []Dictionary{NewDictionary().With("name", "Bob")}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("name"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{NewDictionary().With("name", "Toto")}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestPipelineWithDeleteMaskEngine(t *testing.T) {
	mySlice := []Dictionary{
		NewDictionary().With("name", "Bob").With("city", "Nantes"),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewDeleteMaskEngineProcess(NewPathSelector("name"))).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("city", "Nantes"),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestMaskEngineShouldNotCreateField(t *testing.T) {
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Toto", nil }}

	mySlice := []Dictionary{NewDictionary().With("city", "Nantes")}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("name"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{NewDictionary().With("city", "Nantes")}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

// MaskEngine is a value that will be the initialisation of the field when it's created
type TestAddMaskEngine struct {
	value Entry
}

// MaskContext add the field
func (am TestAddMaskEngine) MaskContext(context Dictionary, key string, contexts ...Dictionary) (Dictionary, error) {
	_, present := context.GetValue(key)
	if !present {
		context.Set(key, am.value)
	}

	return context, nil
}

func TestMaskEngineWithContext(t *testing.T) {
	mySlice := []Dictionary{NewDictionary().With("city", "Nantes")}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskContextEngineProcess(NewPathSelector("name"), TestAddMaskEngine{"Toto"}, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{NewDictionary().With("city", "Nantes").With("name", "Toto")}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestMaskEngineShouldMaskAllEntriesInArray(t *testing.T) {
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{NewDictionary().With("city", []Entry{"Nantes", "Rennes", "Grenoble"})}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("city"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{NewDictionary().With("city", []Entry{"Paris", "Paris", "Paris"})}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestMaskEngineShouldMaskNestedEntry(t *testing.T) {
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{NewDictionary().With("address", NewDictionary().With("city", "Nantes"))}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("address.city"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{NewDictionary().With("address", NewDictionary().With("city", "Paris"))}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestMaskEngineShouldMaskNestedDictionariesArray(t *testing.T) {
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{
		NewDictionary().With("address", []Entry{
			NewDictionary().With("city", "Nantes"),
			NewDictionary().With("city", "Rennes"),
			NewDictionary().With("city", "Grenoble"),
		}),
	}

	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("address.city"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("address", []Entry{
			NewDictionary().With("city", "Paris"),
			NewDictionary().With("city", "Paris"),
			NewDictionary().With("city", "Paris"),
		}),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func TestMaskEngineShouldMaskNestedArray(t *testing.T) {
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Paris", nil }}

	mySlice := []Dictionary{
		NewDictionary().With("address", NewDictionary().With("city", []Entry{
			"Nantes",
			"Rennes",
			"Grenoble",
		})),
	}

	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("address.city"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With(
			"address", NewDictionary().With("city", []Entry{
				"Paris",
				"Paris",
				"Paris",
			}),
		),
	}
	assert.Equal(t, DictionariesToJSONLine(wanted), EntriesToJSONLine(result))
}

func jsonlineToDictionaries(jsl string) []Dictionary {
	result := []Dictionary{}
	jsons := strings.Split(jsl, "\n")
	for _, js := range jsons {
		dict := NewDictionary()
		err := json.Unmarshal(([]byte)(js), &dict)
		if err != nil {
			return nil
		}
		result = append(result, CleanDictionary(dict))
	}
	return result
}

func DictionariesToJSONLine(dictionaries []Dictionary) string {
	var jsonline []byte
	for _, dictionary := range dictionaries {
		result, _ := json.Marshal(dictionary)
		jsonline = append(jsonline, result...)
		jsonline = append(jsonline, "\n"...)
	}
	return strings.TrimSuffix(string(jsonline), "\n")
}

func EntriesToJSONLine(dictionaries []Entry) string {
	var jsonline []byte
	for _, dictionary := range dictionaries {
		result, _ := json.Marshal(dictionary)
		jsonline = append(jsonline, result...)
		jsonline = append(jsonline, "\n"...)
	}
	return strings.TrimSuffix(string(jsonline), "\n")
}

func TestMaskEngineShouldMaskNestedArrays(t *testing.T) {
	i := 0
	nameMasking := FunctionMaskEngine{
		Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
			i++
			return fmt.Sprintf("%d", i), nil
		},
	}

	iput := `{"persons":[{"phonenumber":"001"},{"phonenumber":"002"}]}
			 {"persons":[{"phonenumber":"003"}]}
			 {"persons":[]}`
	oput := `{"persons":[{"phonenumber":"1"},{"phonenumber":"2"}]}
			 {"persons":[{"phonenumber":"3"}]}
			 {"persons":[]}`

	mySlice := jsonlineToDictionaries(iput)

	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("persons.phonenumber"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	expected := DictionariesToJSONLine(jsonlineToDictionaries(oput))
	actual := EntriesToJSONLine(result)
	assert.Equal(t, expected, actual)
}

func TestMaskEngineShouldMaskNestedNestedArrays(t *testing.T) {
	i := 0
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}

	iput := `{"elements":[{"persons":[{"phonenumber":"027123456"},{"phonenumber":"028123456"}]},{"persons":[{"phonenumber":"029123456"},{"phonenumber":"020123456"}]}]}`
	oput := `{"elements":[{"persons":[{"phonenumber":"1"},{"phonenumber":"2"}]},{"persons":[{"phonenumber":"3"},{"phonenumber":"4"}]}]}`

	mySlice := jsonlineToDictionaries(iput)

	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("elements.persons.phonenumber"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	expected := DictionariesToJSONLine(jsonlineToDictionaries(oput))
	actual := EntriesToJSONLine(result)
	assert.Equal(t, expected, actual)
}

func TestInOutFormat1(t *testing.T) {
	masking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "mask", nil }}
	var odict []Entry
	iput := `{"key": ["mask"]}`
	idict := jsonlineToDictionaries(iput)
	pipeline := NewPipelineFromSlice(idict).
		Process(NewMaskEngineProcess(NewPathSelector("key"), masking, "", nil, "")).
		AddSink(NewSinkToSlice(&odict))
	err := pipeline.Run()

	assert.Nil(t, err)
	assert.Equal(t, DictionariesToJSONLine(idict), EntriesToJSONLine(odict))
}

func TestInOutFormat2(t *testing.T) {
	masking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "mask", nil }}
	var odict []Entry
	iput := `{"key1": [{"key2": "mask"}]}`
	idict := jsonlineToDictionaries(iput)
	pipeline := NewPipelineFromSlice(idict).
		Process(NewMaskEngineProcess(NewPathSelector("key1.key2"), masking, "", nil, "")).
		AddSink(NewSinkToSlice(&odict))
	err := pipeline.Run()

	assert.Nil(t, err)
	assert.Equal(t, DictionariesToJSONLine(idict), EntriesToJSONLine(odict))
}

func TestMaskEngineShouldMaskMultipleNestedNestedArrays(t *testing.T) {
	i := 0
	j := 0
	nameMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}
	emailMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		j++
		return fmt.Sprintf("email.%d@company.com", j), nil
	}}

	iput := `{"elements":[{"persons":[{"phonenumber":"027123456","email":"person1@company.com"},{"phonenumber":"028123456","email":"person2@company.com"}]},{"persons":[{"phonenumber":"029123456","email":"person3@company.com"},{"phonenumber":"020123456","email":"person4@company.com"}]}]}`
	oput := `{"elements":[{"persons":[{"phonenumber":"1","email":"email.1@company.com"},{"phonenumber":"2","email":"email.2@company.com"}]},{"persons":[{"phonenumber":"3","email":"email.3@company.com"},{"phonenumber":"4","email":"email.4@company.com"}]}]}`

	mySlice := jsonlineToDictionaries(iput)

	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("elements.persons.email"), emailMasking, "", nil, "")).
		Process(NewMaskEngineProcess(NewPathSelector("elements.persons.phonenumber"), nameMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	expected := DictionariesToJSONLine(jsonlineToDictionaries(oput))
	actual := EntriesToJSONLine(result)
	assert.Equal(t, expected, actual)
}

func TestMaskEngineShouldReturnError(t *testing.T) {
	errorMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "", fmt.Errorf("Test error") }}

	mySlice := []Dictionary{NewDictionary().With("city", "Nantes")}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("city"), errorMasking, "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.NotNil(t, err)
}

func TestPipelineShouldReturnError(t *testing.T) {
	errorMap := func(Dictionary) (Dictionary, error) {
		return NewDictionary(), fmt.Errorf("Test error")
	}

	mySlice := []Dictionary{NewDictionary().With("city", "Nantes")}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMapProcess(errorMap)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.NotNil(t, err)
}

type ErrorSource struct{}

func (s ErrorSource) Err() error {
	return fmt.Errorf("Test error")
}

func (s ErrorSource) Next() bool {
	return false
}

func (s ErrorSource) Value() Entry {
	return NewDictionary()
}

func (s ErrorSource) Open() error {
	return nil
}

func TestPipelineShouldReturnErrorFromSource(t *testing.T) {
	var result []Entry

	pipeline := NewPipeline(ErrorSource{}).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.NotNil(t, err)
}

func TestCacheShouldProvide(t *testing.T) {
	i := 0

	errorMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		i++
		return fmt.Sprintf("%s - %d", name, i), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		NewDictionary().With("city", "Nantes"),
		NewDictionary().With("city", "Grenoble"),
		NewDictionary().With("city", "Nantes"),
	}
	var result []Entry

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("city"), NewMaskCacheEngine(cache, errorMasking), "", nil, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Entry{
		NewDictionary().With("city", "Nantes - 1"),
		NewDictionary().With("city", "Grenoble - 2"),
		NewDictionary().With("city", "Nantes - 1"),
	}

	assert.Equal(t, wanted, result)
}

func TestCallableMapSource(t *testing.T) {
	source := NewCallableMapSource()
	assert.NotNil(t, source)
}

func TestCallableMapSourceNextShouldReturnFalseBeforeValueIsSetted(t *testing.T) {
	source := NewCallableMapSource()
	assert.False(t, source.Next())
}

func TestCallableMapSourceNextShouldReturnTrueAfterValueIsSetted(t *testing.T) {
	source := NewCallableMapSource()
	source.SetValue(NewDictionary())
	assert.True(t, source.Next())
}

func TestCallableMapSourceNextShouldReturnFalseAfterNextCalledTwice(t *testing.T) {
	source := NewCallableMapSource()
	source.SetValue(NewDictionary())
	assert.True(t, source.Next())
	assert.False(t, source.Next())
}
