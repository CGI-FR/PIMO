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

package add

import (
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that will be the initialisation of the field when it's created
type MaskEngine struct {
	value model.Entry
}

// NewMask return a MaskEngine from a value
func NewMask(value model.Entry) MaskEngine {
	return MaskEngine{value}
}

// MaskContext add the field
func (am MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask add")
	_, present := context.GetValue(key)
	if !present {
		context.Set(key, am.value)
	}

	return context, nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskContextEngine, bool, error) {
	if conf.Mask.Add != nil {
		return NewMask(conf.Mask.Add), true, nil
	}
	return nil, false, nil
}
