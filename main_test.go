package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReturnEmptyWhenInputISEmpty(t *testing.T) {

	masking := MaskingFactory(MaskConfiguration{})
	data := Dictionnary{}
	result := masking(data)

	assert.Equal(t, data, result, "should be empty")

}

func TestMaskingShouldNoReplaceInsensitiveValue(t *testing.T) {
	masking := MaskingFactory(MaskConfiguration{})

	data := Dictionnary{"project": "ER456"}
	result := masking(data)

	assert.Equal(t, data, result, "should be equal")

}

func TestMaskingShouldReplaceSensitiveValue(t *testing.T) {
	nameMasking := func(name interface{}) interface{} { return "Toto" }

	masking := MaskingFactory(MaskConfiguration{"name": nameMasking})

	data := Dictionnary{"name": "Benjamin"}
	result := masking(data)
	t.Log(result)

	assert.NotEqual(t, data, result, "should be masked")

}
