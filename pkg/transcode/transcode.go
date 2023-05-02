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

package transcode

import (
	"hash/fnv"
	"math/rand"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand    *rand.Rand
	Classes []model.Class
	seeder  model.Seeder
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(list []model.Class, seed int64, seeder model.Seeder) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), list, seeder}
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask transcode")

	if len(context) > 0 {
		seed, ok, err := mrl.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			mrl.rand.Seed(seed)
		}
	}

	switch eString := e.(type) {
	case string:
		eSplit := strings.Split(eString, "")
		for i, c := range eSplit {
			for _, class := range mrl.Classes {
				if strings.Contains(class.Input, c) {
					eSplit[i] = pick(class.Output, mrl.rand)
				}
			}
		}
		return strings.Join(eSplit, ""), nil
	default:
		return context, nil
	}
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	// set differents seeds for differents jsonpath
	if conf.Masking.Mask.Transcode != nil {
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed, conf.SeedFromClock)
		if classes := conf.Masking.Mask.Transcode.Classes; len(classes) > 0 {
			return NewMask(classes, conf.Seed, seeder), true, nil
		}
		return NewMask(defaultClasses(), conf.Seed, seeder), true, nil
	}

	return nil, false, nil
}

func pick(output string, rand *rand.Rand) string {
	return strings.Split(output, "")[rand.Intn(len(output))]
}

func defaultClasses() []model.Class {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	return []model.Class{{Input: lower, Output: lower}, {Input: upper, Output: upper}, {Input: digits, Output: digits}}
}

func Func(seed int64, seedField string, seedFromClock bool) interface{} {
	var callnumber int64
	return func(input model.Entry) (model.Entry, error) {
		mask := NewMask(defaultClasses(), seed+callnumber, model.NewSeeder(seedField, seed+callnumber, seedFromClock))
		callnumber++
		return mask.Mask(input)
	}
}
