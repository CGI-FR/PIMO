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

	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name": "Benjamin"}
	result, err := maskingEngine.Mask(data)

	assert.Equal(t, nil, err, "error should be nil")

	waited := model.Dictionary{"name": "Toto"}
	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestMaskingShouldReturnAnErrorInCaseOfWrongCommand(t *testing.T) {
	nameCommandMasking := NewMask("WrongCommand")
	config := model.NewMaskConfiguration().
		WithEntry("name", nameCommandMasking)

	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name": "Benjamin"}
	result, err := maskingEngine.Mask(data)
	resultmap := result.(map[string]model.Entry)
	assert.Equal(t, nil, resultmap["name"], "Result should be nil")
	assert.NotEqual(t, nil, err, "Error should not be nil")
}

func TestNewMaskFromConfigShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Command: "echo Toto"}}
	mask, present, err := NewMaskFromConfig(maskingConfig, 0)
	waitedMask := NewMask("echo Toto")
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestNewMaskFromConfigShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := NewMaskFromConfig(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
