package randomlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRandomInList(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}
	config := model.NewMaskConfiguration().
		WithEntry("name", NewMask(nameList, 0))

	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name": "Benjamin"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.NotEqual(t, data, result, "should be masked")

	namemap := result.(map[string]model.Entry)
	assert.Contains(t, nameList, namemap["name"], "Should be in the list")
}

func TestMaskingShouldReplaceSensitiveValueByRandomAndDifferent(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}
	config := model.NewMaskConfiguration().
		WithEntry("name", NewMask(nameList, 0))
	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name": "Benjamin"}
	diff := 0
	for i := 0; i < 1000; i++ {
		result, err := maskingEngine.Mask(data)
		assert.Equal(t, nil, err, "error should be nil")
		resultBis, err := maskingEngine.Mask(data)
		assert.Equal(t, nil, err, "error should be nil")
		if result.(map[string]model.Entry)["name"] != resultBis.(map[string]model.Entry)["name"] {
			diff++
		}
	}
	assert.True(t, diff >= 750, "Should be the same less than 250 times")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoice: []model.Entry{"Michael", "Paul", "Marc"}}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedConfig := model.NewMaskConfiguration().WithEntry("", NewMask([]model.Entry{"Michael", "Paul", "Marc"}, 0))
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldCreateAMaskFromAList(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoiceInURI: "file://../../test/names.txt"}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedConfig := model.NewMaskConfiguration().WithEntry("", NewMask([]model.Entry{"Mickael", "Marc", "Benjamin"}, 0))
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

func TestRegistryMaskToConfigurationShouldReturnErrorFromADoubleConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoice: []model.Entry{"Michael", "Paul", "Marc"}, RandomChoiceInURI: "file://../../test/names.txt"}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.NotNil(t, err, "error should be nil")
}
