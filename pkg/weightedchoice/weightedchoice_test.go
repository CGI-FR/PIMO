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

package weightedchoice

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByWeightedRandom(t *testing.T) {
	NameWeighted := []model.WeightedChoiceType{{Choice: "Michel", Weight: 4}, {Choice: "Marc", Weight: 1}}
	weightMask := NewMask(NameWeighted, 0, nil)

	wait := "Michel"
	equal := 0
	data := "Benjamin"
	for i := 0; i < 1000; i++ {
		result, err := weightMask.Mask(data)
		assert.Equal(t, nil, err, "error should be nil")
		if result == wait {
			equal++
		}
	}

	assert.True(t, equal >= 750, "Should be more than 750 Michel")
	assert.True(t, equal <= 850, "Should be less than 150 Marc")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingChoice := []model.WeightedChoiceType{{Choice: "Dupont", Weight: 9}, {Choice: "Dupond", Weight: 1}}
	maskingConfig := model.Masking{Mask: model.MaskType{WeightedChoice: maskingChoice}}
	_, present, err := Factory(maskingConfig, 0, nil)
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
