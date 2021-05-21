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

package command

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByCommand(t *testing.T) {
	nameProgramMasking := NewMask("echo Toto")
	data := "Benjamin"
	result, err := nameProgramMasking.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "Toto"
	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestMaskingShouldReturnAnErrorInCaseOfWrongCommand(t *testing.T) {
	nameCommandMasking := NewMask("WrongCommand")

	data := model.NewDictionaryFromMap(map[string]model.Entry{"name": "Benjamin"})
	result, err := nameCommandMasking.Mask(data)
	resultmap := result.(model.Dictionary)
	assert.Equal(t, "Benjamin", resultmap.Get("name"), "Result should unchanged")
	assert.NotEqual(t, nil, err, "Error should not be nil")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Command: "echo Toto"}}
	config, present, err := Factory(maskingConfig, 0, nil)
	waitedConfig := NewMask("echo Toto")
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0, nil)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
