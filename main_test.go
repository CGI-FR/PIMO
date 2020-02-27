package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReturnEmptyWhenInputISEmpty(t *testing.T) {

	maskingEngine := MaskingFactory(MapMaskConfiguration{})
	data := Dictionary{}
	result := maskingEngine.Mask(data)

	assert.Equal(t, data, result, "should be empty")

}

func TestMaskingShouldNoReplaceInsensitiveValue(t *testing.T) {
	maskingEngine := MaskingFactory(MapMaskConfiguration{})

	data := Dictionary{"project": "ER456"}
	result := maskingEngine.Mask(data)

	assert.Equal(t, data, result, "should be equal")

}

func TestMaskingShouldReplaceSensitiveValue(t *testing.T) {
	nameMasking := FunctionMaskEngine{func(name Entry) Entry { return "Toto" }}

	maskingEngine := MaskingFactory(MapMaskConfiguration{map[string]MaskEngine{"name": nameMasking}})

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	t.Log(result)

	assert.NotEqual(t, data, result, "should be masked")

}

func TestMaskingShouldReplaceValueInNestedDictionary(t *testing.T) {
	nameMasking := FunctionMaskEngine{func(name Entry) Entry { return "Toto" }}

	customerMaskingEngine := MaskingFactory(MapMaskConfiguration{map[string]MaskEngine{"name": nameMasking}})
	maskingEngine := MaskingFactory(MapMaskConfiguration{map[string]MaskEngine{"customer": customerMaskingEngine}})

	data := Dictionary{"customer": Dictionary{"name": "Benjamin"}, "project": "MyProject"}
	result := maskingEngine.Mask(data)
	t.Log(result)
	want := Dictionary{"customer": Dictionary{"name": "Toto"}, "project": "MyProject"}
	assert.Equal(t, want, result, "should be masked")

}
