package randomlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRandomInList(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}

	data := model.Dictionary{"name": "Benjamin"}
	result, err := NewMask(nameList, 0).Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.NotEqual(t, data, result, "should be masked")

	assert.Contains(t, nameList, result, "Should be in the list")
}

func TestMaskingShouldReplaceSensitiveValueByRandomAndDifferent(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}

	mask := NewMask(nameList, 0)

	diff := 0
	for i := 0; i < 1000; i++ {
		result, err := mask.Mask("Benjamin")
		assert.Equal(t, nil, err, "error should be nil")
		resultBis, err := mask.Mask("Benjamin")
		assert.Equal(t, nil, err, "error should be nil")
		if result != resultBis {
			diff++
		}
	}
	assert.True(t, diff >= 750, "Should be the same less than 250 times")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoice: []model.Entry{"Michael", "Paul", "Marc"}}}
	config, present, err := Factory(maskingConfig, 0)
	waitedConfig := NewMask([]model.Entry{"Michael", "Paul", "Marc"}, 0)
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldCreateAMaskFromAList(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoiceInURI: "file://../../test/names.txt"}}
	config, present, err := Factory(maskingConfig, 0)
	waitedConfig := NewMask([]model.Entry{"Mickael", "Marc", "Benjamin"}, 0)
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

func TestFactoryShouldReturnErrorFromADoubleConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoice: []model.Entry{"Michael", "Paul", "Marc"}, RandomChoiceInURI: "file://../../test/names.txt"}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.NotNil(t, err, "error should be nil")
}
