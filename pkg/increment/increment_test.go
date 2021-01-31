package increment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldCreateIncrementalInt(t *testing.T) {
	incrementMask := NewMask(1, 1)

	firstMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, firstMasked, 1, "First  id masking should be equal")
	secondMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, secondMasked, 2, "Second id masking should be equal")
}

func TestMaskingShouldCreateIncrementalIntwithOffset(t *testing.T) {
	incrementMask := NewMask(5, 2)

	firstMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, firstMasked, 5, "First  id masking should be equal")
	secondMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, secondMasked, 7, "Second id masking should be equal")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Incremental: model.IncrementalType{Start: 1, Increment: 3}}}
	actual, present, err := Factory(maskingConfig, 0)
	waited := NewMask(1, 3)
	assert.Equal(t, waited, actual, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
