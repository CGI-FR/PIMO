package increment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldCreateIncrementalInt(t *testing.T) {
	incrementMask := NewMask(1, 1)
	secIncrMask := NewMask(5, 2)
	config := model.NewMaskConfiguration().
		WithEntry("id1", incrementMask).
		WithEntry("id2", secIncrMask)

	maskingEngine := model.MaskingEngineFactory(config)
	firstData := model.Dictionary{"id1": 0, "id2": 0}
	secondData := model.Dictionary{"id1": 0, "id2": 0}
	firstMasked, err := maskingEngine.Mask(firstData)
	assert.Equal(t, nil, err, "error should be nil")
	firstWaited := model.Dictionary{"id1": 1, "id2": 5}
	assert.Equal(t, firstMasked, firstWaited, "First  id masking should be equal")
	secondMasked, err := maskingEngine.Mask(secondData)
	assert.Equal(t, nil, err, "error should be nil")
	secondWaited := model.Dictionary{"id1": 2, "id2": 7}
	assert.Equal(t, secondMasked, secondWaited, "Second id masking should be equal")
}

func TestNewMaskFromConfigShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Incremental: model.IncrementalType{Start: 1, Increment: 3}}}
	mask, present, err := NewMaskFromConfig(maskingConfig, 0)
	incrementMask := mask.(MaskEngine)
	waitedMask := NewMask(1, 3)
	assert.Equal(t, waitedMask.Increment, incrementMask.Increment, "should be equal")
	assert.Equal(t, &waitedMask.Value, &incrementMask.Value, "Should have the right value")
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
