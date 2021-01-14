package randomdecimal

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRandomNumber(t *testing.T) {
	min := float64(0)
	max := float64(10)
	precision := 2
	randomMask := NewMask(min, max, precision, time.Now().UnixNano())
	config := model.NewMaskConfiguration().
		WithEntry("amount", randomMask)

	maskingEngine := model.MaskingEngineFactory(config, true)
	data := model.Dictionary{"amount": 20}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	resultmap := result.(map[string]model.Entry)
	amount := resultmap["amount"].(float64)

	assert.NotEqual(t, data, result, "Should be masked")
	assert.True(t, amount >= min, "Should be more than min")
	assert.True(t, amount <= max, "Should be less than max")
	assert.Equal(t, 4, len(strconv.FormatFloat(amount, 'f', -1, 64)), "Should be of length 4")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomDecimal: model.RandomDecimalType{Min: float64(0), Max: float64(10), Precision: 2}}}
	conf, present, err := Factory(maskingConfig, 0)
	waitedConf := NewMask(float64(0), float64(10), 2, 0)
	assert.Equal(t, conf, waitedConf, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
