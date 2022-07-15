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

package add

import (
	"testing"
	"text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldAddField(t *testing.T) {
	addMask, err := NewMask("newvalue", template.FuncMap{})
	assert.NoError(t, err, "error should be nil")
	data := model.NewDictionary().With("field", "SomeInformation")
	result, err := addMask.MaskContext(data, "newfield", data)
	assert.NoError(t, err, "error should be nil")
	expected := model.NewDictionary().With("field", "SomeInformation").With("newfield", "newvalue")
	assert.Equal(t, expected, result, "field should be added")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "field"}, Mask: model.MaskType{Add: "value"}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0, Functions: template.FuncMap{}}
	_, present, err := Factory(factoryConfig)
	assert.NoError(t, err, "error should be nil")
	assert.True(t, present, "should be true")
}
