package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nameMasking = FunctionMaskEngine{func(name Entry) Entry { return "Toto" }}

func TestMaskingShouldReturnEmptyWhenInputISEmpty(t *testing.T) {

	maskingEngine := NewMaskConfiguration().AsEngine()
	data := Dictionary{}
	result := maskingEngine.Mask(data)

	assert.Equal(t, data, result, "should be empty")

}

func TestMaskingShouldNoReplaceInsensitiveValue(t *testing.T) {
	maskingEngine := NewMaskConfiguration().AsEngine()

	data := Dictionary{"project": "ER456"}
	result := maskingEngine.Mask(data)

	assert.Equal(t, data, result, "should be equal")

}

func TestMaskingShouldReplaceSensitiveValue(t *testing.T) {

	config := NewMaskConfiguration().
		WithEntry("name", nameMasking)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	t.Log(result)

	assert.NotEqual(t, data, result, "should be masked")

}

func TestMaskingShouldReplaceValueInNestedDictionary(t *testing.T) {

	config := NewMaskConfiguration().
		WithEntry("customer", NewMaskConfiguration().
			WithEntry("name", nameMasking).AsEngine(),
		)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"customer": Dictionary{"name": "Benjamin"}, "project": "MyProject"}
	result := maskingEngine.Mask(data)
	t.Log(result)
	want := Dictionary{"customer": Dictionary{"name": "Toto"}, "project": "MyProject"}
	assert.Equal(t, want, result, "should be masked")

}

func TestWithEntryShouldBuildNestedConfigurationWhenKeyContainsDot(t *testing.T) {

	config := NewMaskConfiguration().
		WithEntry("customer.name", nameMasking)
	_, ok := config.GetMaskingEngine("customer")
	assert.True(t, ok, "should be same configuration")

	_, okbis := config.GetMaskingEngine("customer.name")
	assert.False(t, okbis, "should be same configuration")
}
