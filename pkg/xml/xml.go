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

package xml

import (
	"hash/fnv"
	"math/rand"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a struct to mask the date
type MaskEngine struct {
	rand         *rand.Rand
	seeder       model.Seeder
	xPath        string
	injectParent string
	masking      []byte
}

// NewMask return a MaskEngine from 2 dates
func NewMask(xPath, injectParent, subMasking string, seed int64, seeder model.Seeder) MaskEngine {
	maskingConfig := []byte(subMasking)
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, xPath, injectParent, maskingConfig}
}

// Mask choose a mask date randomly
func (engine MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask XML")
	var transformedXML string
	return transformedXML, nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.XML.XPath) != 0 || len(conf.Masking.Mask.XML.Masking) != 0 {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)
		if len(conf.Masking.Mask.XML.InjectParent) == 0 {
			conf.Masking.Mask.XML.InjectParent = ""
		}
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		return NewMask(conf.Masking.Mask.XML.XPath,
			conf.Masking.Mask.XML.InjectParent,
			conf.Masking.Mask.XML.Masking,
			conf.Seed, seeder), true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(xPath string, injectParent string, masking string) (model.Entry, error) {
		mask := NewMask(xPath, injectParent, masking, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		callnumber++
		return mask.Mask(nil)
	}
}
