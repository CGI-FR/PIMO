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
		Expected:   "many",
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

func TestExpactedAtLeastOneShouldReturnFirstResult(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons_same_name.csv",
		ExactMatch: exactMatch,
		Expected:   "at-least-one",
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
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t,
		model.NewDictionary().
			With("last_name", "Vidal").
			With("email", "vincent.vidal@yopmail.fr").
			With("first_name", "Vincent").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
}

func TestExpactedByDefaultShouldReturnFirstResult(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons_same_name.csv",
		ExactMatch: exactMatch,
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
			With("email", "vincent.vidal@yopmail.fr").
			With("first_name", "Vincent").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
	assert.Nil(t, err, "error should be nil")
}

func TestExpactedOnlyOneShouldReturnErrorWhen3Results(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons_same_name.csv",
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
	assert.Equal(t, "Expected one result, but got 3", err.Error())
	assert.Nil(t, masked, "masked should be nil")
}

func TestExpactedByDefaultShouldReturnErrorWhenNoMatch(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons_same_name.csv",
		ExactMatch: exactMatch,
		Header:     true,
		TrimSpace:  true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().With("nom", "Vi").With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)
	assert.Equal(t, "Expected at least one result, but got none", err.Error())
	assert.Nil(t, masked, "masked should be nil")
}

func TestJaccardMatchShouldFindAMatchWithExactSameString(t *testing.T) {
	JaccardMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	config := model.FindInCSVType{
		URI:          "file://../../test/persons.csv",
		JaccardMatch: JaccardMatch,
		Header:       true,
		TrimSpace:    true,
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
			With("first_name", "Luce").
			With("last_name", "Vidal").
			With("email", "luce.vidal@yopmail.fr").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
}

func TestJaccardMatchShouldFindMostSimilarMatch(t *testing.T) {
	JaccardMatch := model.ExactMatchType{
		CSV:   "{{.email}}",
		Entry: "{{.email}}",
	}
	config := model.FindInCSVType{
		URI:          "file://../../test/persons_same_name.csv",
		JaccardMatch: JaccardMatch,
		Header:       true,
		TrimSpace:    true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().With("email", "luc.vidal@yopmail.fr").With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t,
		model.NewDictionary().
			With("first_name", "Luce").
			With("last_name", "Vidal").
			With("email", "luce.vidal@yopmail.fr").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
}

func TestJaccardMatchShouldFindMostSimilarMatchWithoutHeader(t *testing.T) {
	JaccardMatch := model.ExactMatchType{
		CSV:   "{{index . \"2\"}}",
		Entry: "{{.email}}",
	}
	config := model.FindInCSVType{
		URI:          "file://../../test/persons_same_name.csv",
		JaccardMatch: JaccardMatch,
		Header:       false,
		TrimSpace:    true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().With("email", "luc.vidal@yopmail.fr").With("info_personne", "").Pack()
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

// Should find {Luce,Vidal,luc.vidal@yopmail.fr} not {Luce,Vida,luce.vidal@yopmail.fr}
func TestJaccardMatchWithExactMatchShouldFindMostSimilarMatch(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+123",
	}
	jaccardMatch := model.ExactMatchType{
		CSV:   "{{.email}}+456",
		Entry: "{{.email}}+456",
	}
	config := model.FindInCSVType{
		URI:          "file://../../test/persons_exact_jaccard.csv",
		ExactMatch:   exactMatch,
		JaccardMatch: jaccardMatch,
		Header:       true,
		TrimSpace:    true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	data := model.NewDictionary().
		With("nom", "Vidal").
		With("email", "luc.vidal@yopmail.fr").
		With("info_personne", "").Pack()
	masked, err := mask.Mask("info_personne", data)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t,
		model.NewDictionary().
			With("first_name", "Luce").
			With("last_name", "Vidal").
			With("email", "luce.vidal@yopmail.fr").Unordered(),
		masked.(model.Dictionary).Unordered(),
	)
}
