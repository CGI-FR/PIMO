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

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a type to mask with a random float
type MaskEngine struct {
	rand      *rand.Rand
	seeder    model.Seeder
	precision int
	min       float64
	max       float64
}

// NewMask create a MaskEngine with a seed
func NewMask(min float64, max float64, precision int, seed int64, seeder model.Seeder) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, precision, min, max}
}

// Mask choose a mask int randomly within boundary
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask randomDecimal")

	if len(context) > 0 {
		seed, ok, err := me.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			me.rand.Seed(seed)
		}
	}

	// Generate the random number
	r := me.min + me.rand.Float64()*(me.max-me.min)

	// Round the number to the precision
	rounded := math.Round(r*math.Pow(10, float64(me.precision))) / math.Pow(10, float64(me.precision))
	return rounded, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.RandomDecimal.Precision != 0 {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64()) //nolint:gosec
		return NewMask(conf.Masking.Mask.RandomDecimal.Min, conf.Masking.Mask.RandomDecimal.Max, conf.Masking.Mask.RandomDecimal.Precision, conf.Seed, seeder), true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(min float64, max float64, precision int) (model.Entry, error) {
		mask := NewMask(min, max, precision, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		callnumber++
		return mask.Mask(nil)
	}
}
