package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByCommand(t *testing.T) {
	nameProgramMasking := NewMask("echo Toto")
	config := model.NewMaskConfiguration().
		WithEntry("name", nameProgramMasking)

	maskingEngine := model.MaskingEngineFactory(config, true)

	data := model.Dictionary{"name": "Benjamin"}
	result, err := maskingEngine.Mask(data)

	assert.Equal(t, nil, err, "error should be nil")

	waited := model.Dictionary{"name": "Toto"}
	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestMaskingShouldReturnAnErrorInCaseOfWrongCommand(t *testing.T) {
	nameCommandMasking := NewMask("WrongCommand")

	data := model.Dictionary{"name": "Benjamin"}
	result, err := nameCommandMasking.Mask(data)
	resultmap := result.(map[string]model.Entry)
	assert.Equal(t, "Benjamin", resultmap["name"], "Result should unchanged")
	assert.NotEqual(t, nil, err, "Error should not be nil")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Command: "echo Toto"}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	waitedConfig := model.NewMaskConfiguration().WithEntry("", NewMask("echo Toto"))
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
