package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nameMasking = FunctionMaskEngine{func(name Entry) Entry { return "Toto" }}
var nameProgramMasking = CommandMaskEngine{"echo Toto"}
var nameList = []Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}

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

	assert.NotEqual(t, data, result, "should be masked")
}

func TestMaskingShouldReplaceSensitiveValueByCommand(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", nameProgramMasking)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	waited := Dictionary{"name": "Toto"}

	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestMaskingShouldReplaceSensitiveValueByRandomInList(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", NewMaskRandomList(nameList))

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)

	assert.NotEqual(t, data, result, "should be masked")

	namemap := result.(map[string]Entry)
	assert.Contains(t, nameList, namemap["name"], "Should be in the list")
}

func TestMaskingShouldReplaceSensitiveValueByRandomAndDifferent(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", NewMaskRandomList(nameList))
	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	diff := 0
	for i := 0; i < 1000; i++ {
		result := maskingEngine.Mask(data)
		resultBis := maskingEngine.Mask(data)
		if result.(map[string]Entry)["name"] != resultBis.(map[string]Entry)["name"] {
			diff++
		}
	}
	assert.True(t, diff >= 750, "Should be the same less than 250 times")
}

func TestMaskingShouldReplaceSensitiveValueByHashing(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", MaskHashList{nameList})

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Alexis"}
	result := maskingEngine.Mask(data)
	resultBis := maskingEngine.Mask(data)

	assert.Equal(t, result, resultBis, "Should be hashed the same way")
}

func TestMaskingShouldReplaceValueInNestedDictionary(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("customer", NewMaskConfiguration().
			WithEntry("name", nameMasking).AsEngine(),
		)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"customer": Dictionary{"name": "Benjamin"}, "project": "MyProject"}
	result := maskingEngine.Mask(data)
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

func TestMaskingShouldReplaceSensitiveValueByWeightedRandom(t *testing.T) {
	NameWeighted := []WeightedChoice{{data: "Michel", weight: 4}, {data: "Marc", weight: 1}}
	weightMask := NewWeightedMaskList(NameWeighted)
	config := NewMaskConfiguration().
		WithEntry("name", weightMask)

	maskingEngine := MaskingEngineFactory(config)

	wait := Dictionary{"name": "Michel"}
	equal := 0
	data := Dictionary{"name": "Benjamin"}
	for i := 0; i < 1000; i++ {
		result := maskingEngine.Mask(data).(map[string]Entry)
		if result["name"].(string) == wait["name"] {
			equal++
		}
	}

	assert.True(t, equal >= 750, "Should be more than 750 Michel")
	assert.True(t, equal <= 850, "Should be less than 150 Marc")
}

func TestMaskingShouldReplaceSensitiveValueByConstantValue(t *testing.T) {
	maskingConstant := "Toto"
	constMask := NewConstMask(maskingConstant)
	config := NewMaskConfiguration().
		WithEntry("name", constMask)
	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	waited := Dictionary{"name": "Toto"}

	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}
