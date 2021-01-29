package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceValueInArrayDictionary(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Toto", nil }}

	config := NewMaskConfiguration().
		WithEntry("customer", NewMaskConfiguration().
			WithEntry("name", nameMasking).AsEngine(),
		)

	maskingEngine := MaskingEngineFactory(config, true)

	data := Dictionary{"customer": []Dictionary{{"name": "Benjamin"}}}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	want := Dictionary{"customer": []Dictionary{{"name": "Toto"}}}
	assert.Equal(t, want, result, "should be masked")
}

func TestMaskingShouldReplaceValueInArrayString(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Toto", nil }}

	config := NewMaskConfiguration().
		WithEntry("customer", nameMasking)

	maskingEngine := MaskingEngineFactory(config, true)
	data := Dictionary{"customer": []Entry{"Benjamin"}}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	want := Dictionary{"customer": []Entry{"Toto"}}
	assert.Equal(t, want, result, "should be masked")
}

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
