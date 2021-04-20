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

package increment

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldCreateIncrementalInt(t *testing.T) {
	incrementMask := NewMask(1, 1)

	firstMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, firstMasked, 1, "First  id masking should be equal")
	secondMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, secondMasked, 2, "Second id masking should be equal")
}

func TestMaskingShouldCreateIncrementalIntwithOffset(t *testing.T) {
	incrementMask := NewMask(5, 2)

	firstMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, firstMasked, 5, "First  id masking should be equal")
	secondMasked, err := incrementMask.Mask(0)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, secondMasked, 7, "Second id masking should be equal")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Incremental: model.IncrementalType{Start: 1, Increment: 3}}}
	actual, present, err := Factory(maskingConfig, 0, nil)
	waited := NewMask(1, 3)
	assert.Equal(t, waited, actual, "should be equal")
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
