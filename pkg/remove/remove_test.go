package remove

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByConstantValue(t *testing.T) {
	removeMask := NewMask()
	config := model.NewMaskConfiguration().
		WithContextEntry("field", removeMask)
	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"field": "SomeInformation", "otherfield": "SomeOtherInformation"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"otherfield": "SomeOtherInformation"}
	assert.NotEqual(t, data, result, "field should be removed")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "field"}, Mask: model.MaskType{Remove: true}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedMask := model.NewMaskConfiguration().WithContextEntry("field", NewMask())
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
