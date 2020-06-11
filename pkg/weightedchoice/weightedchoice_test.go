package weightedchoice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByWeightedRandom(t *testing.T) {
	NameWeighted := []model.WeightedChoiceType{{Choice: "Michel", Weight: 4}, {Choice: "Marc", Weight: 1}}
	weightMask := NewMask(NameWeighted, 0)
	config := model.NewMaskConfiguration().
		WithEntry("name", weightMask)

	maskingEngine := model.MaskingEngineFactory(config)

	wait := model.Dictionary{"name": "Michel"}
	equal := 0
	data := model.Dictionary{"name": "Benjamin"}
	for i := 0; i < 1000; i++ {
		result, err := maskingEngine.Mask(data)
		resultmap := result.(map[string]model.Entry)
		assert.Equal(t, nil, err, "error should be nil")
		if resultmap["name"].(string) == wait["name"] {
			equal++
		}
	}

	assert.True(t, equal >= 750, "Should be more than 750 Michel")
	assert.True(t, equal <= 850, "Should be less than 150 Marc")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingChoice := []model.WeightedChoiceType{{Choice: "Dupont", Weight: 9}, {Choice: "Dupond", Weight: 1}}
	maskingConfig := model.Masking{Mask: model.MaskType{WeightedChoice: maskingChoice}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedConfig := model.NewMaskConfiguration().WithEntry("", NewMask(maskingChoice, 0))
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
