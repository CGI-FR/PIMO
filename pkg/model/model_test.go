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
	mySlice := []Entry{1, 2, 3, 4}
	var result []Entry

	pipeline := NewSourceFromSlice(mySlice).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)
	assert.Equal(t, mySlice, result)
}

func TestPipelineWithProcessorSource(t *testing.T) {
	mySlice := []Entry{1, 2, 3, 4}
	var result []Entry

	pipeline := NewSourceFromSlice(mySlice).
		Process(NewMapProcess(func(e Entry) (Entry, error) {
			value := e.(int)
			return value + 1, nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Entry{2, 3, 4, 5}
	assert.Equal(t, wanted, result)
}

func TestPipelineWithChainedProcessorSource(t *testing.T) {
	mySlice := []Entry{1, 2, 3, 4}
	var result []Entry

	pipeline := NewSourceFromSlice(mySlice).
		Process(NewMapProcess(func(e Entry) (Entry, error) {
			value := e.(int)
			return value + 1, nil
		})).
		Process(NewMapProcess(func(e Entry) (Entry, error) {
			value := e.(int)
			return value * value, nil
		})).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Entry{4, 9, 16, 25}
	assert.Equal(t, wanted, result)
}
