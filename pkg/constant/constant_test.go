package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByConstantValue(t *testing.T) {
	maskingConstant := "Toto"
	constMask := NewMask(maskingConstant)
	config := model.NewMaskConfiguration().
		WithEntry("name", constMask)
	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"name": "Benjamin"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"name": "Toto"}
	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Constant: "Toto"}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedMask := model.NewMaskConfiguration().WithEntry("", NewMask("Toto"))
	assert.Equal(t, waitedMask, mask, "should be equal")
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
