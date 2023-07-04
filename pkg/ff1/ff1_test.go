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

package ff1

import (
	"os"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMaskingShouldEncryptStringWithTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	context := model.NewDictionary().
		With("name", "Toto").
		With("tweak", "mytweak").Pack()
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, false, "", false, nil)
	line := "Toto"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "nhIy", result, "Should be equal")
}

func TestMaskingShouldEncryptStringWithoutTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	context := model.NewDictionary().
		With("name", "Toto").Pack()
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, false, "", false, nil)
	line := "Toto"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "Uaow", result, "Should be equal")
}

func TestMaskingShouldDecryptStringWithTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	context := model.NewDictionary().
		With("name", "nhIy").
		With("tweak", "mytweak").Pack()
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, true, "", false, nil)
	line := "nhIy"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "Toto", result, "Should be equal")
}

func TestMaskingShouldDecryptStringWithoutTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	context := model.NewDictionary().
		With("name", "Uaow").Pack()
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, true, "", false, nil)
	line := "Uaow"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "Toto", result, "Should be equal")
}

func TestDecodingKeyShouldNotReturnErrorOnWrongKeyLength(t *testing.T) {
	_, err := decodingKey("aHR0cDovL21hc3Rlcm1pbmRzLmdpdGh1Yi5pby9zcHJpZy8=")
	assert.Nil(t, err, "Should be nil")
}

func TestDecodingKeyShouldWork(t *testing.T) {
	original := "cm9nZXJ0aGF0Y2FwdGFpbg=="
	decodedKey, err := decodingKey(original)
	assert.Equal(t, decodedKey, []byte("rogerthatcaptain"), "Should be nil")
	assert.Nil(t, err, "Should be nil")
}

func TestFactoryShouldReturnNilOnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "should be nil")
}

func TestFactoryShouldReturnErrorOnWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{FF1: model.FF1Type{KeyFromEnv: "XXXXXX", Radix: 0, TweakField: "tweak"}}}
	factoryConfig := model.MaskFactoryConfiguration{Masking: maskingConfig, Seed: 0}
	mask, present, err := Factory(factoryConfig)
	assert.Nil(t, mask, "should be nil")
	assert.True(t, present, "should be true")
	assert.EqualErrorf(t, err, "one of the radix or domain attributes should be set", "should be nil")
}
