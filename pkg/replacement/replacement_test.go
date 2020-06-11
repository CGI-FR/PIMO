package replacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/constant"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByAnotherField(t *testing.T) {
	remplacement := NewMask("name1")
	config := model.NewMaskConfiguration().
		WithEntry("name2", remplacement)
	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name1": "Dupont", "name2": "Dupond"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"name1": "Dupont", "name2": "Dupont"}
	assert.NotEqual(t, result, data, "Should be different")
	assert.Equal(t, waited, result, "Should be masked with the other value")
}

func TestMaskingShouldReplaceSensitiveValueByAMaskedField(t *testing.T) {
	remplacement := NewMask("name1")
	consta := constant.NewMask("Toto")
	config := model.NewMaskConfiguration().
		WithEntry("name1", consta).
		WithEntry("name2", remplacement)
	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name1": "Dupont", "name2": "Dupond"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"name1": "Toto", "name2": "Toto"}
	assert.Equal(t, waited, result, "Should be masked with the constant value")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Replacement: "name"}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedConfig := model.NewMaskConfiguration().WithEntry("", NewMask("name"))
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
