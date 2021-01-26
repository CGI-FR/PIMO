package model

import (
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

func TestPipelineWithAddMaskEngine(t *testing.T) {
	var cityMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { return "Nantes", nil }}

	mySlice := []Dictionary{{"name": "Bob"}}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("city"), cityMasking)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{{"name": "Bob", "city": "Nantes"}}
	assert.Equal(t, wanted, result)

	assert.NotEqual(t, wanted, mySlice)
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
