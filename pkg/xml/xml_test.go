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

package pimo

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFactoryShouldCreateAMask(t *testing.T) {
	xPath := "content"
	subMasking := `
	- selector:
		jsonpath: "@author"
  	mask:
		template: "New Author Name"`
	maskingConfig := model.Masking{Mask: model.MaskType{XML: model.XMLType{XPath: xPath, InjectParent: "", Masking: subMasking}}}
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
	subMasking := `
    - selector:
          jsonpath: "@author"
      mask:
        template: "new name"
`

	config := model.XMLType{
		XPath:        "note",
		InjectParent: "",
		Masking:      subMasking,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{XML: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 42}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().
		With("title", "my blog note").
		With("content", `<note author="John Doe"><date>10/10/2023</date>This is a note of my blog....</note>`).
		Pack()
	masked, err := mask.Mask("content", data)

	assert.Equal(t,
		model.NewDictionary().
			With("title", "my blog note").
			With("content", `<note author="new name"><date>10/10/2023</date>This is a note of my blog....</note>`).
			Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
	assert.Nil(t, err, "error should be nil")
}

func TestMaskShouldRemoveTargetAttribute(t *testing.T) {
	subMasking := `
    - selector:
          jsonpath: "@author"
      mask:
        remove: true
`

	config := model.XMLType{
		XPath:        "note",
		InjectParent: "",
		Masking:      subMasking,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{XML: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 42}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().
		With("title", "my blog note").
		With("content", `<note author="John Doe"><date>10/10/2023</date>This is a note of my blog....</note>`).
		Pack()
	masked, err := mask.Mask("content", data)

	assert.Equal(t,
		model.NewDictionary().
			With("title", "my blog note").
			With("content", `<note author="*remove"><date>10/10/2023</date>This is a note of my blog....</note>`).
			Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
	assert.Nil(t, err, "error should be nil")
}
