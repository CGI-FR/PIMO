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

package transcode

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldTranscodeValue(t *testing.T) {
	Input := "abcdefghijklmnopqrstuvwxyz"
	Output := "*"
	Class := model.Class{Input: Input, Output: Output}
	transcodeMask := NewMask([]model.Class{Class}, 0, func(context model.Dictionary) (int64, bool, error) {
		return 0, false, nil
	})

	result, err := transcodeMask.Mask("mark_23")

	assert.Nil(t, nil, err, "error should be nil")
	assert.Equal(t, "****_23", result)
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingChoice := model.TranscodeType{Classes: []model.Class{{Input: "01234", Output: "56789"}, {Input: "ABCDE", Output: "abcde"}}}
	maskingConfig := model.Masking{Mask: model.MaskType{Transcode: &maskingChoice}}
	_, present, err := Factory(maskingConfig, 0, nil)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0, nil)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}
