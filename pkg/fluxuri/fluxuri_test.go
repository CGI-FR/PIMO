package fluxuri

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestNewMaskShouldCreateTheRightMask(t *testing.T) {
	fileName := "file://../../test/csvvalues.csv"
	mask, err := NewMask(fileName)
	assert.Nil(t, err, "error should be nil")
	list := []model.Entry{1623, 1512, 905}
	length := 3
	actual := 0
	expectedMask := MaskEngine{list, length, &actual}
	assert.Equal(t, mask, expectedMask, "Should create the right mask")
}

func TestMaskingShouldMaskAsExpected(t *testing.T) {
	fileName := "file://../../test/csvvalues.csv"
	mask, err := NewMask(fileName)
	assert.Nil(t, err, "error should be nil")

	// Creating the field
	firstData := model.Dictionary{"field": "thing", "field2": "thing"}
	firstMasked, err := mask.MaskContext(firstData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	firstWaited := model.Dictionary{"id": 1623, "field": "thing", "field2": "thing"}
	assert.Equal(t, firstMasked, firstWaited, "First id masking should be equal")

	// Creating the field with the second value
	secondData := model.Dictionary{"field": "thing", "field2": "thing"}
	secondMasked, err := mask.MaskContext(secondData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	secondWaited := model.Dictionary{"id": 1512, "field": "thing", "field2": "thing"}
	assert.Equal(t, secondMasked, secondWaited, "Second id masking should be equal")

	// Replacing the existing field
	thirdData := model.Dictionary{"id": 25, "field": "thing", "field2": "thing"}
	thirdMasked, err := mask.MaskContext(thirdData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	thirdWaited := model.Dictionary{"id": 905, "field": "thing", "field2": "thing"}
	assert.Equal(t, thirdMasked, thirdWaited, "Third id masking should be equal")

	// Not creating field if every data is used
	fourthData := model.Dictionary{"field": "thing", "field2": "thing"}
	fourthMasked, err := mask.MaskContext(fourthData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	fourthWaited := model.Dictionary{"field": "thing", "field2": "thing"}
	assert.Equal(t, fourthMasked, fourthWaited, "Third id masking should be equal")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{FluxURI: "file://../../test/csvvalues.csv"}}
	conf, present, err := Factory(maskingConfig, 0)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
	waitedMask, err := NewMask("file://../../test/csvvalues.csv")
	assert.Equal(t, conf, waitedMask, "should be equal")
	assert.Nil(t, err, "error should be nil")
}
