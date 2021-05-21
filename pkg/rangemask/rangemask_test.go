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

package rangemask

import (
	"encoding/json"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByRangedValue(t *testing.T) {
	rangeMask := NewMask(10)
	result, err := rangeMask.Mask(json.Number("25"))
	assert.Equal(t, nil, err, "error should be nil")
	waited := "[20;29]"
	assert.NotEqual(t, float64(25), result, "should be masked")
	assert.Equal(t, waited, result, "should be [20;29]")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RangeMask: 15}}
	mask, present, err := Factory(maskingConfig, 0, nil)
	waitedMask := NewMask(15)
	assert.Equal(t, waitedMask, mask, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
