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
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/capitalone/fpe/ff1"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	keyFromEnv string
	tweakField string
	radix      uint
	decrypt    bool
	domain     string
	preserve   bool
	onError    string
}

// NewMask return a MaskEngine from a value
func NewMask(key string, tweak string, radix uint, decrypt bool, domain string, preserve bool, onError string) MaskEngine {
	return MaskEngine{key, tweak, radix, decrypt, domain, preserve, onError}
}

// Mask return a Constant from a MaskEngine
func (ff1m MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	if e == nil {
		// Cannot encrypt or decrypt a nil value so we leave it untouched
		log.Warn().Msg("Mask ff1 - ignored null value")
		return e, nil
	}

	log.Info().Msg("Mask ff1")

	// Extract tweak from the Dictionary (context)
	var tweak string
	if context[0].Get(ff1m.tweakField) == nil {
		tweak = ""
	} else {
		tweak = context[0].Get(ff1m.tweakField).(string)
	}
	// Get encryption key as byte array
	envKey := os.Getenv(ff1m.keyFromEnv)
	if envKey == "" {
		return nil, fmt.Errorf("Environment variable named '%s' should be defined", ff1m.keyFromEnv)
	}
	decodedKey, err := decodingKey(envKey)
	if err != nil {
		return nil, err
	}

	radix := int(ff1m.radix)
	value := e.(string)
	if len(ff1m.domain) > 0 {
		if value, err = toFF1Domain(value, ff1m.domain); err != nil {
			return nil, err
		}
		radix = len(ff1m.domain)
	}

	// Create cipher based on radix, key and the given tweak
	FF1, err := ff1.NewCipher(radix, len(tweak), decodedKey, []byte(tweak))
	if err != nil {
		return nil, err
	}

	var ciphertext string
	if ff1m.decrypt {
		// Decrypt targeted string
		ciphertext, err = FF1.Decrypt(value)
	} else {
		// Encrypt targeted string
		ciphertext, err = FF1.Encrypt(value)
	}
	if err != nil {
		return nil, err
	}

	if len(ff1m.domain) > 0 {
		ciphertext = fromFF1Domain(ciphertext, ff1m.domain)
	}

	return ciphertext, nil
}

// Checking encryption key length then converting it
func decodingKey(key string) ([]byte, error) {
	// Decoding base64 key to byte array
	decodedkey, err := b64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return decodedkey, nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.FF1.KeyFromEnv != "" || conf.Masking.Mask.FF1.Radix > 0 || conf.Masking.Mask.FF1.Domain != "" {
		if conf.Masking.Mask.FF1.KeyFromEnv == "" {
			return nil, true, fmt.Errorf("keyFromEnv attribut is not optional")
		}
		if conf.Masking.Mask.FF1.Radix == 0 && conf.Masking.Mask.FF1.Domain == "" {
			return nil, true, fmt.Errorf("radix or domain should be given")
		}
		return NewMask(conf.Masking.Mask.FF1.KeyFromEnv, conf.Masking.Mask.FF1.TweakField, conf.Masking.Mask.FF1.Radix, conf.Masking.Mask.FF1.Decrypt, conf.Masking.Mask.FF1.Domain, conf.Masking.Mask.FF1.Preserve, conf.Masking.Mask.FF1.OnError), true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	return func(key string, tweak string, radix uint, decrypt bool, domain string, preserve bool, input model.Entry) (model.Entry, error) {
		context := model.NewDictionary()
		context.Set("tweak", tweak)
		mask := NewMask(key, "tweak", radix, decrypt, domain, preserve, "")
		return mask.Mask(input, context)
	}
}

const ff1domain = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toFF1Domain(value string, domain string) (string, error) {
	var result strings.Builder
	for _, char := range value {
		index := strings.IndexRune(domain, char)
		if index > -1 {
			result.WriteByte(ff1domain[index])
		} else {
			return value, fmt.Errorf("character %c is outside of the domain %s", char, domain)
		}
	}
	return result.String(), nil
}

func fromFF1Domain(value string, domain string) string {
	var result strings.Builder
	for _, char := range value {
		index := strings.IndexRune(ff1domain, char)
		if index > -1 {
			result.WriteByte(domain[index])
		} else {
			panic(fmt.Errorf("character %c is outside of the ff1 domain", char))
		}
	}
	return result.String()
}
