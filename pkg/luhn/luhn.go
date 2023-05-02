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

package luhn

import (
	"bytes"
	"fmt"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a struct to create incremental int
type MaskEngine struct {
	Universe []byte
}

// NewMask create an Luhn mask
func NewMask(universe []byte) MaskEngine {
	return MaskEngine{universe}
}

// Mask return the value with luhn checksum
func (l MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	if e == nil {
		// Cannot use a nil value so we leave it untouched
		log.Warn().Msg("Mask luhn - ignored null value")
		return e, nil
	}

	log.Info().Msg("Mask luhn")

	factor := 2
	sum := 0
	n := len(l.Universe)
	input := e.(string)

	// Starting from the right and working leftwards is easier since
	// the initial "factor" will always be "2".
	for i := len(input) - 1; i >= 0; i-- {
		codePoint := bytes.IndexByte(l.Universe, input[i])
		addend := factor * codePoint

		// Alternate the "factor" that each "codePoint" is multiplied by
		if factor == 2 {
			factor = 1
		} else {
			factor = 2
		}

		// Sum the digits of the "addend" as expressed in base "n"
		addend = addend/n + (addend % n)
		sum += addend
	}

	// Calculate the number that must be added to the "sum"
	// to make it divisible by "n".
	remainder := sum % n
	checkCodePoint := (n - remainder) % n

	return input + string(l.Universe[checkCodePoint]), nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.Luhn != nil {
		if conf.Masking.Mask.Luhn.Universe != "" {
			if len(conf.Masking.Mask.Luhn.Universe)%2 != 0 {
				return nil, true, fmt.Errorf("luhn universe size must be divisible by 2")
			}
			return NewMask([]byte(conf.Masking.Mask.Luhn.Universe)), true, nil
		}
		return NewMask([]byte("0123456789")), true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string, seedFromClock bool) interface{} {
	return func(input model.Entry) (model.Entry, error) {
		mask := NewMask([]byte("0123456789"))
		return mask.Mask(input)
	}
}
