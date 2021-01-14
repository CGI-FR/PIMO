package remove

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldRemoveField(t *testing.T) {
	removeMask := NewMask()
	data := model.Dictionary{"field": "SomeInformation", "otherfield": "SomeOtherInformation"}
	result, err := removeMask.MaskContext(data, "field")
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"otherfield": "SomeOtherInformation"}
	assert.Equal(t, waited, result, "should be Toto")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "field"}, Mask: model.MaskType{Remove: true}}
	mask, present, err := Factory(maskingConfig, 0)
	waitedMask := NewMask()
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
