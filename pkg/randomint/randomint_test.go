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

package randomint

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByRandomNumber(t *testing.T) {
	min := 7
	max := 77
	ageMask := NewMask(min, max, 0, nil)

	result, err := ageMask.Mask(83)
	assert.Equal(t, nil, err, "error should be nil")

	assert.NotEqual(t, 83, result, "Should be masked")
	assert.True(t, result.(int) >= min, "Should be more than min")
	assert.True(t, result.(int) <= max, "Should be less than max")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomInt: model.RandIntType{Min: 18, Max: 25}}}
	mask, present, err := Factory(maskingConfig, 0, nil)
	assert.NotNil(t, mask, "shouldn't be nil")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0, nil)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
