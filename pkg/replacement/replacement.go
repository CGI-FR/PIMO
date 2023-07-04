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

package replacement

import (
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// ReplacementMask is to mask with a value from another field
type MaskEngine struct {
	Field model.Selector
}

// NewMask return a mask containing another field of the dictionary
func NewMask(field string) MaskEngine {
	return MaskEngine{model.NewPackedPathSelector(field)}
}

// Mask masks a value with another field of the json
func (remp MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask replacement")
	v, _ := remp.Field.Read(context[0])

	return v, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Replacement) != 0 {
		return NewMask(conf.Masking.Mask.Replacement), true, nil
	}
	return nil, false, nil
}
