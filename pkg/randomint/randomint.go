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

package randomint

import (
	"fmt"
	"hash/fnv"
	"math/rand"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of number to mask randomly
type MaskEngine struct {
	rand   *rand.Rand
	min    int
	max    int
	seeder model.Seeder
}

// NewMask create a MaskEngine with a seed
func NewMask(min int, max int, seed int64, seeder model.Seeder) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), min, max, seeder}
}

// Mask choose a mask int randomly within boundary
func (rim MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask randomInt")
	if len(context) > 0 {
		seed, ok, err := rim.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			rim.rand.Seed(seed)
		}
	}
	return rim.rand.Intn(rim.max+1-rim.min) + rim.min, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if conf.Mask.RandomInt.Min != 0 || conf.Mask.RandomInt.Max != 0 {
		var seeder model.Seeder
		if jpath := conf.Seed.Field; jpath != "" {
			sel := model.NewPathSelector(jpath)
			h := fnv.New64a()
			seeder = func(context model.Dictionary) (int64, bool, error) {
				e, ok := sel.Read(context)
				if !ok {
					return 0, ok, nil
				}
				h.Reset()
				_, err := h.Write([]byte(fmt.Sprintf("%v", e)))
				return int64(h.Sum64()), true, err
			}
		} else {
			// set differents seeds for differents jsonpath
			h := fnv.New64a()
			h.Write([]byte(conf.Selector.Jsonpath))
			seed += int64(h.Sum64())
			seeder = func(context model.Dictionary) (int64, bool, error) {
				return 0, false, nil
			}
		}
		return NewMask(conf.Mask.RandomInt.Min, conf.Mask.RandomInt.Max, seed, seeder), true, nil
	}
	return nil, false, nil
}
