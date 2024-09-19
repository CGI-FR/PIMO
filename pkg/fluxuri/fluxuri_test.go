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

package fluxuri

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldMaskAsExpected(t *testing.T) {
	fileName := "file://../../test/csvvalues.csv"
	mask, err := NewMask(fileName)
	assert.Nil(t, err, "error should be nil")

	// Creating the field
	firstData := model.NewDictionary().With("field", "thing").With("field2", "thing")
	firstMasked, err := mask.MaskContext(firstData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	firstExpected := model.NewDictionary().With("field", "thing").With("field2", "thing").With("id", 1623)
	assert.Equal(t, firstMasked, firstExpected, "First id masking should be equal")

	// Creating the field with the second value
	secondData := model.NewDictionary().With("field", "thing").With("field2", "thing")
	secondMasked, err := mask.MaskContext(secondData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	secondExpected := model.NewDictionary().With("field", "thing").With("field2", "thing").With("id", 1512)
	assert.Equal(t, secondMasked, secondExpected, "Second id masking should be equal")

	// Replacing the existing field
	thirdData := model.NewDictionary().With("id", 25).With("field", "thing").With("field2", "thing")
	thirdMasked, err := mask.MaskContext(thirdData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	thirdExpected := model.NewDictionary().With("id", 905).With("field", "thing").With("field2", "thing")
	assert.Equal(t, thirdMasked, thirdExpected, "Third id masking should be equal")

	// Not creating field if every data is used
	fourthData := model.NewDictionary().With("field", "thing").With("field2", "thing")
	fourthMasked, err := mask.MaskContext(fourthData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	fourthExpected := model.NewDictionary().With("field", "thing").With("field2", "thing")
	assert.Equal(t, fourthMasked, fourthExpected, "Third id masking should be equal")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{FluxURI: "file://../../test/csvvalues.csv"}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	conf, present, err := Factory(factoryConfig)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
	expectedMask, err := NewMask("file://../../test/csvvalues.csv")
	assert.Equal(t, conf, expectedMask, "should be equal")
	assert.Nil(t, err, "error should be nil")
}
