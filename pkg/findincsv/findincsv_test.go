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

func TestFactoryShouldCreateAMask(t *testing.T) {
	exactMatch := model.ExactMatchType{
		CSV:   "{{.last_name}}",
		Entry: "Blanc",
	}
	config := model.FindInCSVType{
		URI:        "file://../../test/persons.csv",
		ExactMatch: exactMatch,
		Header:     true,
	}
	maskingConfig := model.Masking{Mask: model.MaskType{FindInCSV: config}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	_, present, err := Factory(factoryConfig)
	assert.Nil(t, err, "error should be nil")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}
