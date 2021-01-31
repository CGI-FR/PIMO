package rangemask

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRangedValue(t *testing.T) {
	rangeMask := NewMask(10)
	result, err := rangeMask.Mask(float64(25))
	assert.Equal(t, nil, err, "error should be nil")
	waited := "[20;29]"
	assert.NotEqual(t, float64(25), result, "should be masked")
	assert.Equal(t, waited, result, "should be [20;29]")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RangeMask: 15}}
	mask, present, err := Factory(maskingConfig, 0)
	waitedMask := NewMask(15)
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
