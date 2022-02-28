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

package markov

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	chain   Chain
	maxSize int
}

// NewMask create a MaskEngine with a seed
func NewMask(seed int64, name, separator string, maxSize, order int) (MaskEngine, error) {
	list, err := uri.Read(name)
	if err != nil {
		return MaskEngine{}, err
	}
	if order == 0 {
		order = 2
	}
	//nolint: gosec
	chain := NewChain(order, rand.New(rand.NewSource(seed)), separator)
	for _, el := range list {
		s := strings.Split(el.(string), separator)
		chain.Add(s)
	}
	if err != nil {
		fmt.Println(err)
	}
	if maxSize == 0 {
		maxSize = 20
	}
	return MaskEngine{chain, maxSize}, nil
}

func (mm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask markov")
	tokens := make([]string, 0)
	for i := 0; i < mm.chain.order; i++ {
		tokens = append(tokens, StartToken)
	}
	for tokens[len(tokens)-1] != EndToken && len(tokens) <= mm.maxSize {
		next, _ := mm.chain.Generate(tokens[(len(tokens) - mm.chain.order):])
		tokens = append(tokens, next)
	}

	return strings.Join(tokens[mm.chain.order:len(tokens)-1], mm.chain.separator), nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Markov.Sample) != 0 {
		h := fnv.New64a()
		h.Write([]byte(conf.Selector.Jsonpath))
		seed += int64(h.Sum64())
		mask, err := NewMask(seed, conf.Mask.Markov.Sample,
			conf.Mask.Markov.Separator,
			conf.Mask.Markov.MaxSize,
			conf.Mask.Markov.Order)
		return mask, true, err
	}
	return nil, false, nil
}
