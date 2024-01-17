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

package xml

import (
	"testing"
	"time"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/randdate"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/stretchr/testify/assert"
)

func TestFactoryShouldCreateAMask(t *testing.T) {
	xPath := "content"
	masking := []model.Masking{
		{
			Selector: model.SelectorType{Jsonpath: "@author"},
			Mask:     model.MaskType{Template: "new name"},
		},
	}
	model.InjectMaskFactories([]model.MaskFactory{templatemask.Factory})
	maskingConfig := model.Masking{Mask: model.MaskType{XML: model.XMLType{XPath: xPath, InjectParent: "", Masking: masking}}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	config, present, err := Factory(factoryConfig)
	assert.NotNil(t, config, "config shouldn't be nil")
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

func TestMaskShouldReplaceTargetAttributeValue(t *testing.T) {
	masking := []model.Masking{
		{
			Selector: model.SelectorType{Jsonpath: "@author"},
			Mask:     model.MaskType{Template: "new name"},
		},
	}
	config := model.XMLType{
		XPath:        "note",
		InjectParent: "",
		Masking:      masking,
	}
	model.InjectMaskFactories([]model.MaskFactory{templatemask.Factory})
	maskingConfig := model.Masking{Mask: model.MaskType{XML: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 42}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := "<note author='John Doe'><date>10/10/2023</date>This is a note of my blog....</note>"
	masked, err := mask.Mask(data, model.Dictionary{})

	assert.Equal(t,
		"<note author='new name'><date>10/10/2023</date>This is a note of my blog....</note>",
		masked,
	)
	assert.Nil(t, err, "error should be nil")
}

func TestMaskWithMultiMasks(t *testing.T) {
	masking := []model.Masking{
		{
			Selector: model.SelectorType{Jsonpath: "@author"},
			Mask:     model.MaskType{Template: "New name"},
		},
		{
			Selector: model.SelectorType{Jsonpath: "date"},
			Masks: []model.MaskType{
				{RandDate: model.RandDateType{
					DateMin: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
					DateMax: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				}},
				{Template: "{{index . \"date\"}}"},
			},
		},
	}
	config := model.XMLType{
		XPath:   "note",
		Masking: masking,
	}

	model.InjectMaskFactories([]model.MaskFactory{templatemask.Factory, randdate.Factory})
	maskingConfig := model.Masking{Mask: model.MaskType{XML: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 42}
	mask, present, err := Factory(factoryConfig)

	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")

	data := "<note author='John Doe'><date>10/10/2023</date>This is a note of my blog....</note>"
	masked, err := mask.Mask(data, model.Dictionary{})

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t,
		"<note author='New name'><date>1981-05-22 08:40:23 +0000 UTC</date>This is a note of my blog....</note>",
		masked,
	)
}

func TestMaskWithoutXmlValueShouldReturnErrorAndStop(t *testing.T) {
	masking := []model.Masking{
		{
			Selector: model.SelectorType{Jsonpath: "@author"},
			Mask:     model.MaskType{Template: "new name"},
		},
	}
	config := model.XMLType{
		XPath:   "note",
		Masking: masking,
	}

	model.InjectMaskFactories([]model.MaskFactory{templatemask.Factory, randdate.Factory})
	maskingConfig := model.Masking{Mask: model.MaskType{XML: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 42}
	mask, present, err := Factory(factoryConfig)

	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")

	data := "Not a XML"
	masked, err := mask.Mask(data, model.Dictionary{})

	assert.Nil(t, masked)
	assert.Equal(t, err.Error(), "Jsonpath content is not a valid XML")
}

func TestMaskShouldReplaceTargetAttributeValueWithInjectParent(t *testing.T) {
	masking := []model.Masking{
		{
			Selector: model.SelectorType{Jsonpath: "@author"},
			Mask:     model.MaskType{Template: "{{._.title}}"},
		},
		{
			Selector: model.SelectorType{Jsonpath: "date"},
			Mask:     model.MaskType{Template: "{{._.city}}"},
		},
	}
	config := model.XMLType{
		XPath:        "note",
		InjectParent: "_",
		Masking:      masking,
	}
	model.InjectMaskFactories([]model.MaskFactory{templatemask.Factory})
	maskingConfig := model.Masking{Mask: model.MaskType{XML: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 42}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := "<note author='John Doe'><date>10/10/2023</date>This is a note of my blog....</note>"
	dict := model.NewDictionary().With("title", "this is my blog title").With("content", data).With("city", "Nantes").Pack()
	masked, err := mask.Mask(data, dict)

	assert.Equal(t,
		"<note author='this is my blog title'><date>Nantes</date>This is a note of my blog....</note>",
		masked,
	)
	assert.Nil(t, err, "error should be nil")
}
