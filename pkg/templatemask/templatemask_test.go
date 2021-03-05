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

package templatemask

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByTemplate(t *testing.T) {
	template := "{{.name}}.{{.surname}}@gmail.com"
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.Dictionary{"name": "Jean", "surname": "Bonbeur", "mail": "jean44@outlook.com"}
	result, err := tempMask.Mask("jean44@outlook.com", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "Jean.Bonbeur@gmail.com"
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestMaskingShouldReplaceSensitiveValueByTemplateInNested(t *testing.T) {
	template := "{{.customer.identity.name}}.{{.customer.identity.surname}}@gmail.com"
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.Dictionary{"customer": model.Dictionary{"identity": model.Dictionary{"name": "Jean", "surname": "Bonbeur"}}, "mail": "jean44@outlook.com"}
	result, err := tempMask.Mask("jean44@outlook.com", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "Jean.Bonbeur@gmail.com"
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "mail"}, Mask: model.MaskType{Template: "{{.name}}.{{.surname}}@gmail.com"}}
	config, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, err, "error should be nil")
	maskingEngine, _ := NewMask("{{.name}}.{{.surname}}@gmail.com")
	assert.IsType(t, maskingEngine, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
	data := model.Dictionary{"name": "Toto", "surname": "Tata", "mail": ""}
	result, err := maskingEngine.Mask(data, data)
	assert.Nil(t, err, "error should be nil")
	waitedResult := "Toto.Tata@gmail.com"
	assert.Equal(t, waitedResult, result, "Should create a right mail")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldReturnAnErrorInWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Template: "{{.name}.{{.surname}}@gmail.com"}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.NotNil(t, err, "error shouldn't be nil")
}

func TestMaskingTemplateShouldFormat(t *testing.T) {
	template := `{{"hello!" | upper | repeat 2}}`
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.Dictionary{"field": "anything"}
	result, err := tempMask.Mask("anything", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "HELLO!HELLO!"
	assert.Equal(t, waited, result, "Should create the right field")
}

func TestRFactoryShouldCreateANestedContextMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "foo.bar"}, Mask: model.MaskType{Template: "{{.baz}}"}}
	maskEngine, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, err, "should be nil")
	assert.True(t, present, "should be true")
	data := model.Dictionary{"baz": "BAZ", "foo": model.Dictionary{"bar": "BAR"}}

	masked, _ := maskEngine.Mask("bar", data)

	waited := "BAZ"
	assert.Equal(t, waited, masked, "Should replace foo.bar with BAZ")
}
