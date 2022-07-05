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

package dateparser

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceDateStringByDate(t *testing.T) {
	layoutISO := "2006-01-02"
	dateMask := NewMask(layoutISO, "")
	data := "2000-01-01"
	resulttime, err := dateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Nanosecond(), resulttime.(time.Time).Nanosecond(), "Should return the same time")
}

func TestMaskingShouldReplaceDateByDateString(t *testing.T) {
	layoutISO := "2006-01-02"
	dateMask := NewMask("", layoutISO)
	data := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	result, err := dateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	expected := "2020-01-01"
	assert.Equal(t, expected, result, "Should return the same time")
}

func TestMaskingShouldReplaceDateStringByDateString(t *testing.T) {
	layoutIn := "2006-01-02"
	layoutOut := "01/02/06"
	dateMask := NewMask(layoutIn, layoutOut)
	data := "2000-01-01"
	result, err := dateMask.Mask(data)
	expected := "01/01/00"
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, expected, result, "Should return the right time format")
}

func TestMaskingShouldReplaceDateStringByDateStringInTwoMasks(t *testing.T) {
	layoutIn := "2006-01-02"
	layoutOut := "01/02/06"
	firstDateMask := NewMask(layoutIn, "")
	secondDateMask := NewMask("", layoutOut)
	data := "2000-01-01"
	firstResult, err := firstDateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	result, err := secondDateMask.Mask(firstResult)
	expected := "01/01/00"
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, expected, result, "Should return the right time format")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{DateParser: model.DateParserType{InputFormat: "2006-02-01", OutputFormat: "01/02/06"}}}
	factoryConfig := model.NewMaskFactoryConfiguration(maskingConfig, 0, nil)
	config, present, err := Factory(factoryConfig)
	mask := NewMask("2006-02-01", "01/02/06")
	assert.Equal(t, mask, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")

	maskingConfig = model.Masking{Mask: model.MaskType{DateParser: model.DateParserType{InputFormat: "2006-02-01", OutputFormat: ""}}}
	factoryConfig = model.NewMaskFactoryConfiguration(maskingConfig, 0, nil)
	config, present, err = Factory(factoryConfig)
	mask = NewMask("2006-02-01", "")
	assert.Equal(t, mask, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")

	maskingConfig = model.Masking{Mask: model.MaskType{DateParser: model.DateParserType{InputFormat: "", OutputFormat: "01/02/06"}}}
	factoryConfig = model.NewMaskFactoryConfiguration(maskingConfig, 0, nil)
	config, present, err = Factory(factoryConfig)
	mask = NewMask("", "01/02/06")
	assert.Equal(t, mask, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestMaskingShouldReplaceUnixEpochByDateString(t *testing.T) {
	outputFormat := "02/01/06"
	dateMask := NewMask("unixEpoch", outputFormat)
	data := model.CleanTypes(json.Number("1647512434"))
	resulttime, err := dateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, "17/03/22", resulttime, "Should return the same time")
}

func TestMaskingShouldReplaceDateStringByUnixEpoch(t *testing.T) {
	inputFormat := "02/01/06"
	dateMask := NewMask(inputFormat, "unixEpoch")
	data := "17/03/22"
	result, err := dateMask.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, int64(1647475200), result, "Should return the same time")
}
