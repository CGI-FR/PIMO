package replacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByAnotherField(t *testing.T) {
	remplacement := NewMask("name1")

	data := model.Dictionary{"name1": "Dupont", "name2": "Dupond"}
	result, err := remplacement.Mask("Dupond", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "Dupont"
	assert.NotEqual(t, result, data, "Should be different")
	assert.Equal(t, waited, result, "Should be masked with the other value")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Replacement: "name"}}
	config, present, err := Factory(maskingConfig, 0)
	waitedConfig := NewMask("name")
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
