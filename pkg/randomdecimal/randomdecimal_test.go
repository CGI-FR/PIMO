// Copyright (C) 2021 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package randomdecimal

import (
	"strconv"
	"testing"
	"time"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

// TODO: this test will fail randomly 1/10th of the time
func TestMaskingShouldReplaceSensitiveValueByRandomNumber(t *testing.T) {
	min := float64(0)
	max := float64(10)
	precision := 2
	randomMask := NewMask(min, max, precision, time.Now().UnixNano())

	result, err := randomMask.Mask(20)
	assert.Equal(t, nil, err, "error should be nil")
	amount := result.(float64)
	assert.NotEqual(t, 20, amount, "Should be masked")
	assert.True(t, amount >= min, "Should be more than min")
	assert.True(t, amount <= max, "Should be less than max")
	assert.Equal(t, 4, len(strconv.FormatFloat(amount, 'f', -1, 64)), "Should be of length 4")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomDecimal: model.RandomDecimalType{Min: float64(0), Max: float64(10), Precision: 2}}}
	_, present, err := Factory(maskingConfig, 0, nil)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
