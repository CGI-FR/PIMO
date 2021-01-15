package duration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestParseDurationShouldCreateTimeDuration(t *testing.T) {
	dur, err := ParseDuration("P2Y6M5DT12H35M30S")
	assert.Nil(t, err, "error is not nil")
	waited := time.Duration(2*24*365*int64(time.Hour) + 6*30*24*int64(time.Hour) + 5*24*int64(time.Hour) + 12*int64(time.Hour) + 35*int64(time.Minute) + 30*int64(time.Second))
	assert.Equal(t, waited, dur, "should return the good duration")
	negadur, err := ParseDuration("-PT2H")
	assert.Equal(t, nil, err, "error is not nil")
	negawaited := time.Duration(-2 * int64(time.Hour))
	assert.Equal(t, negawaited, negadur, "Should return a negative duration")
}

func TestParseDurationShouldReturnErrorIfNotIso8601(t *testing.T) {
	dur, err := ParseDuration("FalseDuration")
	assert.Equal(t, time.Duration(0), dur, "Should return 0 for duration")
	assert.NotNil(t, err, "error should not be nil")
}

func TestMaskingShouldReplaceDateByIncrement(t *testing.T) {
	duration := "P60D"
	durationMask, err := NewMask(duration)

	assert.Equal(t, nil, err, "error is not nil")

	config := model.NewMaskConfiguration().
		WithEntry("date", durationMask)

	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"date": time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"date": time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)}
	assert.Equal(t, result, waited, "Should change the date according to mask")
	duration = "-P60D"
	durationMask, err = NewMask(duration)

	assert.Equal(t, nil, err, "error is not nil")

	config = model.NewMaskConfiguration().
		WithEntry("date", durationMask)

	maskingEngine = model.MaskingEngineFactory(config, true)
	data = model.Dictionary{"date": time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}
	result, err = maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited = model.Dictionary{"date": time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC)}
	assert.Equal(t, result, waited, "Should change the date according to mask")
}

func TestParseInt64ShouldReturn0IfNotANumber(t *testing.T) {
	number := ParseInt64("number")
	assert.Equal(t, int64(0), number, "should be 0")
}

func TestMaskingShouldReturnAnErrorIfNotATime(t *testing.T) {
	duration := "P2D"
	durationMask, _ := NewMask(duration)

	data := model.Dictionary{"date": "SomeText"}
	_, err := durationMask.Mask(data)
	assert.NotNil(t, err, "Error shouldn't Be Nil")

	secondData := model.Dictionary{"date": 12}
	result, err := durationMask.Mask(secondData)
	waitedResult := model.Dictionary{"date": 12}
	assert.Equal(t, waitedResult, result, "Shouldn't have change")
	assert.NotNil(t, err, "err should not be nil")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Duration: "P2D"}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	mask, _ := NewMask("P2D")
	waitedConfig := model.NewMaskConfiguration().WithEntry("", mask)
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldReturnAnErrorInWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Duration: "NotADuration"}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.NotNil(t, err, "error shouldn't be nil")
}
