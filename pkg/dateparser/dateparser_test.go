package dateparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceDateStringByDate(t *testing.T) {
	layoutISO := "2006-01-02"
	dateMask := NewMask(layoutISO, "")
	data := "2000-01-01"
	resulttime, err := dateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Nanosecond(), resulttime.(time.Time).Nanosecond(), "Should return the same time")
}

func TestMaskingShouldReplaceDateByDateString(t *testing.T) {
	layoutISO := "2006-01-02"
	dateMask := NewMask("", layoutISO)
	data := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	result, err := dateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	expected := "2020-01-01"
	assert.Equal(t, expected, result, "Should return the same time")
}

func TestMaskingShouldReplaceDateStringByDateString(t *testing.T) {
	layoutIn := "2006-01-02"
	layoutOut := "01/02/06"
	dateMask := NewMask(layoutIn, layoutOut)
	data := "2000-01-01"
	result, err := dateMask.Mask(data)
	expected := "01/01/00"
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, expected, result, "Should return the right time format")
}

func TestMaskingShouldReplaceDateStringByDateStringInTwoMasks(t *testing.T) {
	layoutIn := "2006-01-02"
	layoutOut := "01/02/06"
	firstDateMask := NewMask(layoutIn, "")
	secondDateMask := NewMask("", layoutOut)
	data := "2000-01-01"
	firstResult, err := firstDateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	result, err := secondDateMask.Mask(firstResult)
	expected := "01/01/00"
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, expected, result, "Should return the right time format")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{DateParser: model.DateParserType{InputFormat: "2006-02-01", OutputFormat: "01/02/06"}}}
	config, present, err := Factory(maskingConfig, 0)
	mask := NewMask("2006-02-01", "01/02/06")
	assert.Equal(t, mask, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")

	maskingConfig = model.Masking{Mask: model.MaskType{DateParser: model.DateParserType{InputFormat: "2006-02-01", OutputFormat: ""}}}
	config, present, err = Factory(maskingConfig, 0)
	mask = NewMask("2006-02-01", "")
	assert.Equal(t, mask, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")

	maskingConfig = model.Masking{Mask: model.MaskType{DateParser: model.DateParserType{InputFormat: "", OutputFormat: "01/02/06"}}}
	config, present, err = Factory(maskingConfig, 0)
	mask = NewMask("", "01/02/06")
	assert.Equal(t, mask, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
