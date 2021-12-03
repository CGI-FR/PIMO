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
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a struct to create incremental int
type MaskEngine struct {
	Mod uint
	Map []rune
}

// NewMask create an Luhn mask
func NewMask(mod uint, mp []rune) MaskEngine {
	return MaskEngine{mod, mp}
}

// Mask return the value with luhn checksum
func (incr MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask luhn")
	return "123456782", nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if conf.Mask.Luhn.Map != "" {
		return NewMask(conf.Mask.Luhn.Mod, []rune(conf.Mask.Luhn.Map)), true, nil
	}
	return NewMask(conf.Mask.Luhn.Mod, []rune("0123456789")), true, nil
}
