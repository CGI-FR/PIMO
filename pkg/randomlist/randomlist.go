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
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, list}
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
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	// set differents seeds for differents jsonpath
	h := fnv.New64a()
	h.Write([]byte(conf.Selector.Jsonpath))
	seed += int64(h.Sum64())

	if len(conf.Mask.RandomChoice) != 0 {
		return NewMask(conf.Mask.RandomChoice, seed, model.NewSeeder(conf, seed)), true, nil
	}

	return nil, false, nil
}
