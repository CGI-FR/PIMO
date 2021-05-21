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

package rangemask

import (
	"encoding/json"
	"strconv"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that always mask by replacing with a scale
type MaskEngine struct {
	rangeScale int
}

// NewMask return a MaskEngine from a value
func NewMask(scale int) MaskEngine {
	return MaskEngine{scale}
}

// Mask return a range from a MaskEngine
func (rm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask range")
	i, err := e.(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	scaledValue := int(i) / rm.rangeScale * rm.rangeScale
	rangedValue := "[" + strconv.Itoa(scaledValue) + ";" + strconv.Itoa(scaledValue+rm.rangeScale-1) + "]"
	return rangedValue, nil
}

// Factory Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if conf.Mask.RangeMask != 0 {
		return NewMask(conf.Mask.RangeMask), true, nil
	}
	return nil, false, nil
}
