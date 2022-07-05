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

package randomuri

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFactoryShouldCreateAMaskFromAList(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoiceInURI: "file://../../test/names.txt"}}
	factoryConfig := model.NewMaskFactoryConfiguration(maskingConfig, 0, nil)
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	masked, err := mask.Mask("")
	assert.Equal(t, masked, "Mickael", "should be equal")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldCreateAMaskFromAListWithTemplate(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomChoiceInURI: "file://../../test/names_{{.gender}}.txt"}}
	factoryConfig := model.NewMaskFactoryConfiguration(maskingConfig, 0, nil)
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	masked, err := mask.Mask("", model.NewDictionary().With("gender", "M"))
	assert.Equal(t, masked, "Mickael", "should be equal")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	factoryConfig := model.NewMaskFactoryConfiguration(maskingConfig, 0, nil)
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
