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

package findincsv

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFactoryShouldCreateAMaskAndFindMatchSimple(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons.csv",
		ExactMatch: exactMatch,
		Expected:   "only-one",
		Header:     true,
		TrimSpace:  true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().With("nom", "Vidal").With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)

	assert.Equal(t,
		model.NewDictionary().
			With("last_name", "Vidal").
			With("email", "luce.vidal@yopmail.fr").
			With("first_name", "Luce").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldCreateAMaskAndFindMatchWithoutHeader(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{index . \"1\"}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons.csv",
		ExactMatch: exactMatch,
		Expected:   "only-one",
		Header:     false,
		TrimSpace:  true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().With("nom", "Vidal").With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)
	assert.Nil(t, err, "error should be nil")

	assert.Equal(t,
		model.NewDictionary().
			With("0", "Luce").
			With("1", "Vidal").
			With("2", "luce.vidal@yopmail.fr").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
}

func TestFactoryShouldCreateAMaskAndDontFindMatch(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons.csv",
		ExactMatch: exactMatch,
		Expected:   "only-one",
		Header:     false,
		TrimSpace:  true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().With("nom", "Hello").With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)
	assert.Nil(t, err, "error should be nil")

	assert.Equal(t, []model.Entry{}, masked)
}

func TestFactoryShouldCreateAMaskWithTemplatedUri(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/{{.filename}}.csv",
		ExactMatch: exactMatch,
		Expected:   "only-one",
		Header:     true,
		TrimSpace:  true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().
		With("filename", "persons").
		With("nom", "Vidal").
		With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)

	assert.Equal(t,
		model.NewDictionary().
			With("last_name", "Vidal").
			With("email", "luce.vidal@yopmail.fr").
			With("first_name", "Luce").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
	assert.Nil(t, err, "error should be nil")
}
