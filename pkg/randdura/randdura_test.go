package randdura

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceDateByIncrement(t *testing.T) {
	firstDuration := "P60D"
	secondDuration := "P90D"
	durationMask, err := NewMask(firstDuration, secondDuration, 0)

	assert.Equal(t, nil, err, "error is not nil")

	config := model.NewMaskConfiguration().
		WithEntry("date", durationMask)

	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"date": time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}
	result, err := maskingEngine.Mask(data)
	resultMap := result.(map[string]model.Entry)
	assert.Equal(t, nil, err, "error should be nil")
	expectedDate := time.Date(2020, 3, 16, 0, 0, 0, 0, time.UTC)
	assert.WithinDurationf(t, expectedDate, (resultMap["date"]).(time.Time), 15*24*time.Hour, "Should be in March")

	firstNegaDuration := "-P60D"
	secondNegaDuration := "-P90D"
	negaDurationMask, err := NewMask(firstNegaDuration, secondNegaDuration, 0)

	assert.Equal(t, nil, err, "error is not nil")

	negaConfig := model.NewMaskConfiguration().
		WithEntry("date", negaDurationMask)

	negaMaskingEngine := model.MaskingEngineFactory(negaConfig)
	negaData := model.Dictionary{"date": time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}
	negaResult, err := negaMaskingEngine.Mask(negaData)
	negaResultMap := negaResult.(map[string]model.Entry)
	assert.Equal(t, nil, err, "error should be nil")
	negaExpectedDate := time.Date(2019, 10, 18, 0, 0, 0, 0, time.UTC)
	assert.WithinDurationf(t, negaExpectedDate, (negaResultMap["date"]).(time.Time), 15*24*time.Hour, "Should be around October")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomDuration: model.RandomDurationType{Min: "-P60D", Max: "-P90D"}}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	mask, _ := NewMask("-P60D", "-P90D", 0)
	waitedConfig := model.NewMaskConfiguration().WithEntry("", mask)
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
