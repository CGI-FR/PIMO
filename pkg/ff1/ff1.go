package ff1

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/capitalone/fpe/ff1"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	keyFromEnv string
	tweakField string
	radix      uint
	decrypt    bool
}

// NewMask return a MaskEngine from a value
func NewMask(key string, tweak string, radix uint, decrypt bool) MaskEngine {
	return MaskEngine{key, tweak, radix, decrypt}
}

// Mask return a Constant from a MaskEngine
func (ff1m MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	// Extract tweak from the Dictionary (context)
	var tweak string
	if context[0][ff1m.tweakField] == nil {
		tweak = ""
	} else {
		tweak = context[0][ff1m.tweakField].(string)
	}
	// Get encryption key as byte array
	envKey := os.Getenv(ff1m.keyFromEnv)
	if envKey == "" {
		return nil, fmt.Errorf("Environment variable named FF1_ENCRYPTION_KEY should be defined")
	}
	decodedKey, err := decodingKey(envKey)
	if err != nil {
		return nil, err
	}
	// Create cipher based on radix, key and the given tweak
	FF1, err := ff1.NewCipher(int(ff1m.radix), len(tweak), decodedKey, []byte(tweak))
	if err != nil {
		return nil, err
	}
	var ciphertext string
	if ff1m.decrypt {
		// Decrypt targeted string
		ciphertext, err = FF1.Decrypt(e.(string))
	} else {
		// Encrypt targeted string
		ciphertext, err = FF1.Encrypt(e.(string))
	}
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// Checking encryption key length then converting it
func decodingKey(key string) ([]byte, error) {
	// Base64 Key length authorized
	validLength := map[int]bool{
		16: true,
		24: true,
		32: true,
	}
	// Checking key length
	keyLength := b64.StdEncoding.EncodedLen(len(key))
	if !validLength[keyLength] {
		return nil, fmt.Errorf("Encryption Key's length not valid (accepted: 16, 24, 32); actual length is %d", keyLength)
	}
	// Decoding base64 key to byte array
	decodedkey, err := b64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return decodedkey, nil
}

// Factory create a mask from a configuration
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.FF1.KeyFromEnv != "" && conf.Mask.FF1.Radix > 0 {
		return NewMask(conf.Mask.FF1.KeyFromEnv, conf.Mask.FF1.TweakField, conf.Mask.FF1.Radix, conf.Mask.FF1.Decrypt), true, nil
	}
	return nil, false, nil
}
