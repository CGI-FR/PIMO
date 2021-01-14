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
	config := model.NewMaskConfiguration().
		WithEntry("date", dateMask)
	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"date": "2000-01-01"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	resultmap := result.(map[string]model.Entry)
	resulttime := resultmap["date"].(time.Time)
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Nanosecond(), resulttime.Nanosecond(), "Should return the same time")
}

func TestMaskingShouldReplaceDateByDateString(t *testing.T) {
	layoutISO := "2006-01-02"
	dateMask := NewMask("", layoutISO)
	config := model.NewMaskConfiguration().
		WithEntry("date", dateMask)
	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"date": time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	expected := model.Dictionary{"date": "2020-01-01"}
	assert.Equal(t, expected, result, "Should return the same time")
}

func TestMaskingShouldReplaceDateStringByDateString(t *testing.T) {
	layoutIn := "2006-01-02"
	layoutOut := "01/02/06"
	dateMask := NewMask(layoutIn, layoutOut)
	config := model.NewMaskConfiguration().
		WithEntry("date", dateMask)
	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"date": "2000-01-01"}
	result, err := maskingEngine.Mask(data)
	expected := model.Dictionary{"date": "01/01/00"}
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, expected, result, "Should return the right time format")
}

func TestMaskingShouldReplaceDateStringByDateStringInTwoMasks(t *testing.T) {
	layoutIn := "2006-01-02"
	layoutOut := "01/02/06"
	firstDateMask := NewMask(layoutIn, "")
	secondDateMask := NewMask("", layoutOut)
	config := model.NewMaskConfiguration().
		WithEntry("date", firstDateMask).
		WithEntry("date", secondDateMask)
	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"date": "2000-01-01"}
	result, err := maskingEngine.Mask(data)
	expected := model.Dictionary{"date": "01/01/00"}
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
