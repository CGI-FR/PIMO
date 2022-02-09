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

type class struct {
	Input  []string
	Output []string
}

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand    *rand.Rand
	Classes []class
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(list []class, seed int64) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), list}
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask transcode")
	switch eString := e.(type) {
	case string:
		eSplit := strings.Split(eString, "")
		for i, c := range eSplit {
			for _, class := range mrl.Classes {
				if in(class.Input, c) {
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
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	// set differents seeds for differents jsonpath
	h := fnv.New64a()
	h.Write([]byte(conf.Selector.Jsonpath))
	seed += int64(h.Sum64())
	classes := []class{}
	for _, c := range conf.Mask.Transcode {
		input := strings.Split(c.Input, "")
		output := strings.Split(c.Output, "")
		classes = append(classes, class{input, output})
	}
	if len(classes) > 0 {
		return NewMask(classes, seed), true, nil
	}
	return nil, false, nil
}

func pick(output []string, rand *rand.Rand) string {
	return output[rand.Intn(len(output))]
}

func in(input []string, e string) bool {
	for _, i := range input {
		if e == i {
			return true
		}
	}
	return false
}
