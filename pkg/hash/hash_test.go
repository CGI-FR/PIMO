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

package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByHashing(t *testing.T) {
	nameList := []model.Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}
	mask := MaskEngine{nameList}

	data := "Alexis"
	result, err := mask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	resultBis, err := mask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")

	assert.Equal(t, result, resultBis, "Should be hashed the same way")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Hash: []model.Entry{"Mickael", "Benjamin", "Michel"}}}
	config, present, err := Factory(maskingConfig, 0)
	waitedConfig := MaskEngine{[]model.Entry{"Mickael", "Benjamin", "Michel"}}
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
