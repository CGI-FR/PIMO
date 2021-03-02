package ff1

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldEncryptStringWithTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	var context = map[string]model.Entry{
		"name":  "Toto",
		"tweak": "mytweak",
	}
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, false)
	line := "Toto"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "nhIy", result, "Should be equal")
}

func TestMaskingShouldEncryptStringWithoutTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	var context = map[string]model.Entry{
		"name": "Toto",
	}
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, false)
	line := "Toto"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "Uaow", result, "Should be equal")
}

func TestMaskingShouldDecryptStringWithTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	var context = map[string]model.Entry{
		"name":  "nhIy",
		"tweak": "mytweak",
	}
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, true)
	line := "nhIy"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "Toto", result, "Should be equal")
}

func TestMaskingShouldDecryptStringWithoutTweak(t *testing.T) {
	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")
	var context = map[string]model.Entry{
		"name": "Uaow",
	}
	ff1Mask := NewMask("FF1_ENCRYPTION_KEY", "tweak", 62, true)
	line := "Uaow"
	result, err := ff1Mask.Mask(line, context)
	assert.Nil(t, err)
	assert.Equal(t, "Toto", result, "Should be equal")
}

func TestDecodingKeyShouldReturnErrorOnWrongKeyLength(t *testing.T) {
	expectedError := fmt.Errorf("Encryption Key's length not valid (accepted: 16, 24, 32); actual length is %d", 64)
	decodedKey, err := decodingKey("aHR0cDovL21hc3Rlcm1pbmRzLmdpdGh1Yi5pby9zcHJpZy8=")
	assert.Equal(t, err, expectedError, "Should be equal")
	assert.Nil(t, decodedKey, "Should be nil")
}

func TestDecodingKeyShouldWork(t *testing.T) {
	original := "cm9nZXJ0aGF0Y2FwdGFpbg=="
	decodedKey, err := decodingKey(original)
	assert.Equal(t, decodedKey, []byte("rogerthatcaptain"), "Should be nil")
	assert.Nil(t, err, "Should be nil")
}

func TestFactoryShouldReturnNilOnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.Nil(t, err, "should be nil")
}

func TestFactoryShouldReturnNilOnWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{FF1: model.FF1Type{KeyFromEnv: "XXXXXX", Radix: 0, TweakField: "tweak"}}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.Nil(t, err, "should be nil")
}
