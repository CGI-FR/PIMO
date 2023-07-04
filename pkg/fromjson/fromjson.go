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

package fromjson

import (
	"encoding/json"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that will be the initialisation of the field when it's created
type MaskEngine struct {
	sourcefield model.Selector
}

// NewMask return a MaskEngine from a value
func NewMask(source string) MaskEngine {
	return MaskEngine{model.NewPackedPathSelector(source)}
}

// MaskContext
func (am MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask fromjson")

	value, _ := am.sourcefield.Read(contexts[0])

	switch t := value.(type) {
	case string:
		var result interface{}
		err := json.Unmarshal([]byte(t), &result)
		if err != nil {
			return context, err
		}
		context.Set(key, model.CleanTypes(result))
	default:
		log.Warn().Msg("Invalid type")
	}

	return context, nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskContextEngine, bool, error) {
	if conf.Masking.Mask.FromJSON != "" {
		return NewMask(conf.Masking.Mask.FromJSON), true, nil
	}
	return nil, false, nil
}
