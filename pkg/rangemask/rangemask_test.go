package rangemask

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRangedValue(t *testing.T) {
	rangeMask := NewMask(10)
	config := model.NewMaskConfiguration().
		WithEntry("age", rangeMask)
	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"age": float64(25)}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"age": "[20;29]"}
	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RangeMask: 15}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedMask := model.NewMaskConfiguration().WithEntry("", NewMask(15))
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
