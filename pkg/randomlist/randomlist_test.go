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

package randomlist

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByRandomInList(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}

	data := model.NewDictionary().With("name", "Benjamin")
	result, err := NewMask(nameList, 0, nil).Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.NotEqual(t, data, result, "should be masked")

	assert.Contains(t, nameList, result, "Should be in the list")
}

func TestMaskingShouldReplaceSensitiveValueByRandomAndDifferent(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}

	mask := NewMask(nameList, 0, nil)

	diff := 0
	for i := 0; i < 1000; i++ {
		result, err := mask.Mask("Benjamin")
		assert.Equal(t, nil, err, "error should be nil")
		resultBis, err := mask.Mask("Benjamin")
		assert.Equal(t, nil, err, "error should be nil")
		if result != resultBis {
			diff++
		}
	}
	assert.True(t, diff >= 750, "Should be the same less than 250 times")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoice: []model.Entry{"Michael", "Paul", "Marc"}}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	_, present, err := Factory(factoryConfig)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
