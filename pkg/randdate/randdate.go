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

package randdate

import (
	"hash/fnv"
	"math/rand"
	"time"

	"github.com/cgi-fr/pimo/pkg/model"
)

// MaskEngine is a struct to mask the date
type MaskEngine struct {
	rand    *rand.Rand
	DateMin time.Time
	DateMax time.Time
}

// NewMask return a MaskEngine from 2 dates
func NewMask(min, max time.Time, seed int64) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), min, max}
}

// Mask choose a mask date randomly
func (dateRange MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	delta := dateRange.DateMax.Unix() - dateRange.DateMin.Unix()
	sec := time.Unix(dateRange.rand.Int63n(delta)+dateRange.DateMin.Unix(), 0)
	return sec, nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if conf.Mask.RandDate.DateMin != conf.Mask.RandDate.DateMax {
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Selector.Jsonpath))
		seed += int64(h.Sum64())
		return NewMask(conf.Mask.RandDate.DateMin, conf.Mask.RandDate.DateMax, seed), true, nil
	}
	return nil, false, nil
}
