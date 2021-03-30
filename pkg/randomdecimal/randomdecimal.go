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

package randomdecimal

import (
	"hash/fnv"
	"math"
	"math/rand"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a type to mask with a random float
type MaskEngine struct {
	rand      *rand.Rand
	precision int
	min       float64
	max       float64
}

// NewMask create a MaskEngine with a seed
func NewMask(min float64, max float64, precision int, seed int64) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), precision, min, max}
}

// Mask choose a mask int randomly within boundary
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	// Generate the random number
	r := me.min + me.rand.Float64()*(me.max-me.min)

	// Round the number to the precision
	rounded := math.Round(r*math.Pow(10, float64(me.precision))) / math.Pow(10, float64(me.precision))
	return rounded, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.RandomDecimal.Precision != 0 {
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Selector.Jsonpath))
		seed += int64(h.Sum64())
		return NewMask(conf.Mask.RandomDecimal.Min, conf.Mask.RandomDecimal.Max, conf.Mask.RandomDecimal.Precision, seed), true, nil
	}
	return nil, false, nil
}
