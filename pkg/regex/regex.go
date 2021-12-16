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

package regex

import (
	"hash/fnv"
	"math/rand"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
	regen "github.com/zach-klippenstein/goregen"
)

// MaskEngine is a value that mask thanks to a regular expression
type MaskEngine struct {
	generator regen.Generator
}

// NewMask return a RegexMask from a regexp
func NewMask(exp string, seed int64) (MaskEngine, error) {
	generator, err := regen.NewGenerator(exp, &regen.GeneratorArgs{RngSource: rand.NewSource(seed)})
	return MaskEngine{generator}, err
}

// Mask returns a string thanks to a regular expression
func (rm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask regex")
	out := rm.generator.Generate()
	return out, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Regex) != 0 {
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Selector.Jsonpath))
		seed += int64(h.Sum64())

		mask, err := NewMask(conf.Mask.Regex, seed)
		if err != nil {
			return nil, true, err
		}
		return mask, true, err
	}
	return nil, false, nil
}

func Func(seed int64) func(regex string) (model.Entry, error) {
	return func(regex string) (model.Entry, error) {
		mask, err := NewMask(regex, seed)
		if err != nil {
			return "", err
		}
		return mask.Mask(nil)
	}
}
