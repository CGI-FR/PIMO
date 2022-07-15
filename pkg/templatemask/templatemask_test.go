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
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldReplaceSensitiveValueByTemplate(t *testing.T) {
	template := "{{.name}}.{{.surname}}@gmail.com"
	tempMask, err := NewMask(template, tmpl.FuncMap{}, 0)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.NewDictionary().With("name", "Jean").With("surname", "Bonbeur").With("mail", "jean44@outlook.com")
	result, err := tempMask.Mask("jean44@outlook.com", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "Jean.Bonbeur@gmail.com"
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestMaskingShouldReplaceSensitiveValueByTemplateInNested(t *testing.T) {
	template := "{{.customer.identity.name}}.{{.customer.identity.surname}}@gmail.com"
	tempMask, err := NewMask(template, tmpl.FuncMap{}, 0)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.NewDictionary().
		With("customer", model.NewDictionary().
			With("identity", model.NewDictionary().
				With("name", "Jean").
				With("surname", "Bonbeur"),
			).
			With("mail", "jean44@outlook.com"),
		)
	result, err := tempMask.Mask("jean44@outlook.com", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "Jean.Bonbeur@gmail.com"
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "mail"}, Mask: model.MaskType{Template: "{{.name}}.{{.surname}}@gmail.com"}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0, Functions: make(tmpl.FuncMap)}
	config, present, err := Factory(factoryConfig)
  assert.Nil(t, err, "error should be nil")
	maskingEngine, _ := NewMask("{{.name}}.{{.surname}}@gmail.com", tmpl.FuncMap{}, 0)
	assert.IsType(t, maskingEngine, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
	data := model.NewDictionary().With("name", "Toto").With("surname", "Tata").With("mail", "")
	result, err := maskingEngine.Mask(data, data)
	assert.Nil(t, err, "error should be nil")
	waitedResult := "Toto.Tata@gmail.com"
	assert.Equal(t, waitedResult, result, "Should create a right mail")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0, Functions: make(tmpl.FuncMap)}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldReturnAnErrorInWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Template: "{{.name}.{{.surname}}@gmail.com"}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0, Functions: make(tmpl.FuncMap)}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.NotNil(t, err, "error shouldn't be nil")
}

func TestMaskingTemplateShouldFormat(t *testing.T) {
	template := `{{"hello!" | upper | repeat 2}}`
	tempMask, err := NewMask(template, tmpl.FuncMap{}, 0)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.NewDictionary().With("field", "anything")
	result, err := tempMask.Mask("anything", data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := "HELLO!HELLO!"
	assert.Equal(t, waited, result, "Should create the right field")
}

func TestRFactoryShouldCreateANestedContextMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "foo.bar"}, Mask: model.MaskType{Template: "{{.baz}}"}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0, Functions: make(tmpl.FuncMap)}
	maskEngine, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().
		With("baz", "BAZ").
		With("foo", model.NewDictionary().With("bar", "BAR"))

	masked, _ := maskEngine.Mask("bar", data)

	waited := "BAZ"
	assert.Equal(t, waited, masked, "Should replace foo.bar with BAZ")
}

func TestMaskingTemplateShouldIterOverContextArray(t *testing.T) {
	template := `{{- range $index, $rel := .REL_PERMIS -}}{{.ID_PERMIS}}{{- end -}}`
	tempMask, err := NewMask(template, tmpl.FuncMap{}, 0)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	data := model.NewDictionary().With("REL_PERMIS", []model.Dictionary{model.NewDictionary().With("ID_PERMIS", 1)})

	result, err := tempMask.Mask("anything", data)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waited := "1"
	assert.Equal(t, waited, result, "Should create the right field")
}
