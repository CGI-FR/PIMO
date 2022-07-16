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

package randomuri

import (
	"bytes"
	"hash/fnv"
	"math/rand"
	"text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand     *rand.Rand
	seeder   model.Seeder
	template *template.Template
	cache    map[string][]model.Entry
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(templateSource string, seed int64, seeder model.Seeder) (MaskEngine, error) {
	template, err := template.New("template-randomInUri").Parse(templateSource)
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, template, map[string][]model.Entry{}}, err
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask randomChoiceInUri")

	if len(context) > 0 {
		seed, ok, err := mrl.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			mrl.rand.Seed(seed)
		}
	}

	var output bytes.Buffer
	if len(context) == 0 {
		context = []model.Dictionary{model.NewDictionary()}
	}
	if err := mrl.template.Execute(&output, context[0].Unordered()); err != nil {
		return nil, err
	}
	filename := output.String()
	var list []model.Entry
	if listFromCache, ok := mrl.cache[filename]; ok {
		list = listFromCache
	} else {
		listFromFile, err := uri.Read(filename)
		if err != nil {
			return nil, err
		}
		mrl.cache[filename] = listFromFile
		list = listFromFile
	}

	return list[mrl.rand.Intn(len(list))], nil
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	// set differents seeds for differents jsonpath
	h := fnv.New64a()
	h.Write([]byte(conf.Masking.Selector.Jsonpath))
	conf.Seed += int64(h.Sum64())

	if len(conf.Masking.Mask.RandomChoiceInURI) != 0 {
		mask, err := NewMask(conf.Masking.Mask.RandomChoiceInURI, conf.Seed, model.NewSeeder(conf.Masking.Seed.Field, conf.Seed))
		return mask, true, err
	}
	return nil, false, nil
}
