package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByHashing(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}
	config := model.NewMaskConfiguration().
		WithEntry("name", MaskEngine{nameList})

	maskingEngine := model.MaskingEngineFactory(config, true)

	data := model.Dictionary{"name": "Alexis"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	resultBis, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")

	assert.Equal(t, result, resultBis, "Should be hashed the same way")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Hash: []model.Entry{"Mickael", "Benjamin", "Michel"}}}
	config, present, err := Factory(maskingConfig, 0)
	waitedConfig := MaskEngine{[]model.Entry{"Mickael", "Benjamin", "Michel"}}
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
