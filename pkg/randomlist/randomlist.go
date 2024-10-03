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

package randomlist

import (
	"hash/fnv"
	"math/rand"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand   *rand.Rand
	seeder model.Seeder
	list   []model.Entry
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(list []model.Entry, seed int64, seeder model.Seeder) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, model.CleanTypes(list).([]model.Entry)}
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask randomChoice")

	if len(context) > 0 {
		seed, ok, err := mrl.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			mrl.rand.Seed(seed)
		}
	}

	return mrl.list[mrl.rand.Intn(len(mrl.list))], nil
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

	// set differents seeds for differents jsonpath
	h := fnv.New64a()
	h.Write([]byte(conf.Masking.Selector.Jsonpath))
	conf.Seed += int64(h.Sum64()) //nolint:gosec

	if len(conf.Masking.Mask.RandomChoice) != 0 {
		return NewMask(conf.Masking.Mask.RandomChoice, conf.Seed, seeder), true, nil
	}

	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(choices_ []interface{}) (model.Entry, error) {
		choices := make([]model.Entry, len(choices_))
		for i, c := range choices_ {
			choices[i] = c
		}
		mask := NewMask(choices, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		callnumber++
		return mask.Mask(nil)
	}
}
