package ff1

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/capitalone/fpe/ff1"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	key        string
	tweakField string
	radix      uint
}

// NewMask return a MaskEngine from a value
func NewMask(key string, tweak string, radix uint) MaskEngine {
	return MaskEngine{key, tweak, radix}
}

// Mask return a Constant from a MaskEngine
func (ff1m MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	// Extract tweak from the Dictionary (context)
	tweak := context[0][ff1m.tweakField].(string)
	if tweak == "" {
		return nil, fmt.Errorf("Le champ \"%s\" doit être présent et doit contenir le tweak pour le chiffrement", ff1m.tweakField)
	}
	// Get encryption key as byte array
	decodedkey, err := decodingKey(ff1m.key)
	if err != nil {
		return nil, err
	}
	// Create cipher based on radix, key and the given tweak
	FF1, err := ff1.NewCipher(int(ff1m.radix), len(tweak), decodedkey, []byte(tweak))
	if err != nil {
		return nil, err
	}
	// Encode targeted string
	ciphertext, err := FF1.Encrypt(e.(string))
	if err != nil {
		return nil, err
	}
	context[0]["name"] = ciphertext
	return context[0], nil
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
		return nil, fmt.Errorf("La clé de chiffrement n'a pas le format adapté: length=%d", keyLength)
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
	if conf.Mask.FF1.Key != "" && conf.Mask.FF1.Radix > 0 && conf.Mask.FF1.TweakField != "" {
		return NewMask(conf.Mask.FF1.Key, conf.Mask.FF1.TweakField, conf.Mask.FF1.Radix), true, nil
	}
	return nil, false, nil
}
