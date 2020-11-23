package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceValueInArrayDictionary(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry) (Entry, error) { return "Toto", nil }}

	config := NewMaskConfiguration().
		WithEntry("customer", NewMaskConfiguration().
			WithEntry("name", nameMasking).AsEngine(),
		)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"customer": []Dictionary{{"name": "Benjamin"}}}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	want := Dictionary{"customer": []Dictionary{{"name": "Toto"}}}
	assert.Equal(t, want, result, "should be masked")
}

func TestMaskingShouldReplaceValueInArrayString(t *testing.T) {
	var nameMasking = FunctionMaskEngine{Function: func(name Entry) (Entry, error) { return "Toto", nil }}

	config := NewMaskConfiguration().
		WithEntry("customer", nameMasking)

	maskingEngine := MaskingEngineFactory(config)
	data := Dictionary{"customer": []Entry{"Benjamin"}}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	want := Dictionary{"customer": []Entry{"Toto"}}
	assert.Equal(t, want, result, "should be masked")
}
