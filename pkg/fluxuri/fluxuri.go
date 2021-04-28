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

package fluxuri

import (
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of value to mask in order from a list
type MaskEngine struct {
	List    []model.Entry
	LenList int
	Actual  *int
}

// NewMask returns a MaskEngine from a file
func NewMask(name string) (MaskEngine, error) {
	list, err := uri.Read(name)
	if err != nil {
		return MaskEngine{}, err
	}
	length := len(list)
	var actual = 0
	return MaskEngine{list, length, &actual}, nil
}

// MaskContext add the field if not existing or replace the value if existing
func (me MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask fluxuri")
	if *me.Actual < me.LenList {
		context[key] = me.List[*me.Actual]
		*me.Actual++
	}
	return context, nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskContextEngine, bool, error) {
	if len(conf.Mask.FluxURI) != 0 {
		mask, err := NewMask(conf.Mask.FluxURI)
		return mask, true, err
	}
	return nil, false, nil
}
