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

package weightedchoice

import (
	"hash/fnv"
	"math/rand"
	"sort"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// WeightedChoice adapted from github.com/mroth/weightedrand

// Choice is a wrapper of item and its weight
type Choice struct {
	Item   interface{}
	Weight uint
}

// A Chooser is to choose a Choice by weight
type Chooser struct {
	data   []Choice
	totals []int
	max    int
	rand   *rand.Rand
}

// NewChooser creates a Chooser from Choices
func NewChooser(seed int64, cs ...Choice) Chooser {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Weight < cs[j].Weight
	})
	totals := make([]int, len(cs))
	runningTotal := 0
	for i, c := range cs {
		runningTotal += int(c.Weight)
		totals[i] = runningTotal
	}
	// nolint: gosec
	return Chooser{data: cs, totals: totals, max: runningTotal, rand: rand.New(rand.NewSource(seed))}
}

// Pick returns a choice from a Chooser
func (chs Chooser) Pick() interface{} {
	r := chs.rand.Intn(chs.max) + 1
	i := sort.SearchInts(chs.totals, r)
	return chs.data[i].Item
}

// MaskEngine a list of masking value with weight for random
type MaskEngine struct {
	cs     Chooser
	seeder model.Seeder
}

// NewMask returns a WeightedMaskList from slice of model.Entry and weights
func NewMask(list []model.WeightedChoiceType, seed int64, seeder model.Seeder) MaskEngine {
	rand.Seed(seed)
	cs := []Choice{}
	for k := range list {
		cs = append(cs, Choice{Item: list[k].Choice, Weight: list[k].Weight})
	}
	return MaskEngine{NewChooser(seed, cs...), seeder}
}

// Mask choose a mask value randomly
func (wml MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask weightedChoice")

	if len(context) > 0 {
		seed, ok, err := wml.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			wml.cs.rand.Seed(seed)
		}
	}

	return wml.cs.Pick(), nil
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.WeightedChoice) != 0 {
		var maskWeight []model.WeightedChoiceType
		for _, v := range conf.Masking.Mask.WeightedChoice {
			maskWeight = append(maskWeight, model.WeightedChoiceType{Choice: v.Choice, Weight: v.Weight})
		}

		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		return NewMask(maskWeight, conf.Seed, model.NewSeeder(conf.Masking, conf.Seed)), true, nil
	}
	return nil, false, nil
}
