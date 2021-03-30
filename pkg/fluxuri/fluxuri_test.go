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

func TestNewMaskShouldCreateTheRightMask(t *testing.T) {
	fileName := "file://../../test/csvvalues.csv"
	mask, err := NewMask(fileName)
	assert.Nil(t, err, "error should be nil")
	list := []model.Entry{1623, 1512, 905}
	length := 3
	actual := 0
	expectedMask := MaskEngine{list, length, &actual}
	assert.Equal(t, mask, expectedMask, "Should create the right mask")
}

func TestMaskingShouldMaskAsExpected(t *testing.T) {
	fileName := "file://../../test/csvvalues.csv"
	mask, err := NewMask(fileName)
	assert.Nil(t, err, "error should be nil")

	// Creating the field
	firstData := model.Dictionary{"field": "thing", "field2": "thing"}
	firstMasked, err := mask.MaskContext(firstData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	firstWaited := model.Dictionary{"id": 1623, "field": "thing", "field2": "thing"}
	assert.Equal(t, firstMasked, firstWaited, "First id masking should be equal")

	// Creating the field with the second value
	secondData := model.Dictionary{"field": "thing", "field2": "thing"}
	secondMasked, err := mask.MaskContext(secondData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	secondWaited := model.Dictionary{"id": 1512, "field": "thing", "field2": "thing"}
	assert.Equal(t, secondMasked, secondWaited, "Second id masking should be equal")

	// Replacing the existing field
	thirdData := model.Dictionary{"id": 25, "field": "thing", "field2": "thing"}
	thirdMasked, err := mask.MaskContext(thirdData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	thirdWaited := model.Dictionary{"id": 905, "field": "thing", "field2": "thing"}
	assert.Equal(t, thirdMasked, thirdWaited, "Third id masking should be equal")

	// Not creating field if every data is used
	fourthData := model.Dictionary{"field": "thing", "field2": "thing"}
	fourthMasked, err := mask.MaskContext(fourthData, "id")
	assert.Equal(t, nil, err, "error should be nil")
	fourthWaited := model.Dictionary{"field": "thing", "field2": "thing"}
	assert.Equal(t, fourthMasked, fourthWaited, "Third id masking should be equal")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{FluxURI: "file://../../test/csvvalues.csv"}}
	conf, present, err := Factory(maskingConfig, 0)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
	waitedMask, err := NewMask("file://../../test/csvvalues.csv")
	assert.Equal(t, conf, waitedMask, "should be equal")
	assert.Nil(t, err, "error should be nil")
}
