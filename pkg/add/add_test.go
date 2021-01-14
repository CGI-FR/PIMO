package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldAddField(t *testing.T) {
	addMask := NewMask("newvalue")
	data := model.Dictionary{"field": "SomeInformation"}
	result, err := addMask.MaskContext(data, "newfield")
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"field": "SomeInformation", "newfield": "newvalue"}
	assert.Equal(t, waited, result, "field should be added")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "field"}, Mask: model.MaskType{Add: "value"}}
	mask, present, err := Factory(maskingConfig, 0)
	waitedMask := NewMask("value")
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
