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

	data := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	result, err := durationMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	expectedDate := time.Date(2020, 3, 16, 0, 0, 0, 0, time.UTC)
	assert.WithinDurationf(t, expectedDate, (result).(time.Time), 15*24*time.Hour, "Should be in March")
}

func TestMaskingShouldReplaceDateByNegatifIncrement(t *testing.T) {
	firstNegaDuration := "-P60D"
	secondNegaDuration := "-P90D"
	negaMaskingEngine, err := NewMask(firstNegaDuration, secondNegaDuration, 0)

	assert.Equal(t, nil, err, "error is not nil")

	negaData := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	negaResult, err := negaMaskingEngine.Mask(negaData)
	assert.Equal(t, nil, err, "error should be nil")
	negaExpectedDate := time.Date(2019, 10, 18, 0, 0, 0, 0, time.UTC)
	assert.WithinDurationf(t, negaExpectedDate, (negaResult).(time.Time), 15*24*time.Hour, "Should be around October")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomDuration: model.RandomDurationType{Min: "-P60D", Max: "-P90D"}}}
	config, present, err := Factory(maskingConfig, 0)
	mask, _ := NewMask("-P60D", "-P90D", 0)
	waitedConfig := mask
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
