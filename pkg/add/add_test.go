package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldRemoveField(t *testing.T) {
	addMask := NewMask("newvalue")
	config := model.NewMaskConfiguration().
		WithContextEntry("newfield", addMask)
	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"field": "SomeInformation"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"field": "SomeInformation", "newfield": "newvalue"}
	assert.NotEqual(t, data, result, "field should be removed")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "field"}, Mask: model.MaskType{Add: "value"}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedMask := model.NewMaskConfiguration().WithContextEntry("field", NewMask("value"))
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
