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

package regex

import (
	"regexp"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByRegex(t *testing.T) {
	regex := "0[1-7]( ([0-9]){2}){4}"
	regmask, err := NewMask(regex, 0, nil)
	if err != nil {
		assert.Fail(t, "could not initialise regexep")
	}

	data := "00 00 00 00 00"
	result, err := regmask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	match, _ := regexp.MatchString(regex, result.(string))
	assert.NotEqual(t, data, result, "should be masked")
	assert.True(t, match, "should match the regexp")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Regex: "[A-Z]oto([a-z]){3}"}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	_, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
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
