package randomint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRandomNumber(t *testing.T) {
	min := 7
	max := 77
	ageMask := NewMask(min, max, 0)
	config := model.NewMaskConfiguration().
		WithEntry("age", ageMask)

	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"age": 83}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	resultmap := result.(map[string]model.Entry)

	assert.NotEqual(t, data, result, "Should be masked")
	assert.True(t, resultmap["age"].(int) >= min, "Should be more than min")
	assert.True(t, resultmap["age"].(int) <= max, "Should be less than max")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomInt: model.RandIntType{Min: 18, Max: 25}}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.NotNil(t, mask, "shouldn't be nil")
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
