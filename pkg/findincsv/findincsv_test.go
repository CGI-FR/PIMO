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
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/stretchr/testify/assert"
)

func TestFactoryShouldCreateAMaskWith2Templates(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}+123",
		Entry: "{{.nom}}+456",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons.csv",
		ExactMatch: exactMatch,
		Expected:   "only-one",
		Header:     false,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	tempMaskCsv, err := templatemask.NewMask(factoryConfig.Masking.Mask.FindInCSV.ExactMatch.CSV, tmpl.FuncMap{}, 0, "")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	tempMaskEntry, err := templatemask.NewMask(factoryConfig.Masking.Mask.FindInCSV.ExactMatch.Entry, tmpl.FuncMap{}, 0, "")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	data := model.NewDictionary().With("nom", "Jean").With("last_name", "Bonbeur").With("mail", "jean44@outlook.com").Pack()
	resultCsv, err := tempMaskCsv.Mask("jean44@outlook.com", data)
	resultEntry, err := tempMaskEntry.Mask("jean44@outlook.com", data)
	assert.Equal(t, nil, err, "error should be nil")
	waitedCsv := "Bonbeur+123"
	assert.Equal(t, waitedCsv, resultCsv, "Should create the right mail")
	waitedEntry := "Jean+456"
	assert.Equal(t, waitedEntry, resultEntry, "Should create the right mail")
}

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
	assert.Equal(t, masked, "{\"first_name\":\"Luce\",\"last_name\":\"Vidal\",\"email\":\"luce.vidal@yopmail.fr\"}", "should be equal")
	assert.Nil(t, err, "error should be nil")
}
