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

package hash

import (
	"fmt"
	"hash/fnv"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of masking value for hash masking
type MaskEngine struct {
	List []model.Entry
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(list []model.Entry) MaskEngine {
	// nolint: gosec
	return MaskEngine{model.CleanTypes(list).([]model.Entry)}
}

// Mask choose a mask value by hash
func (hm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask hash")
	if e == nil {
		return e, nil
	}
	h := fnv.New32a()
	_, err := h.Write([]byte(e.(string)))
	return hm.List[int(h.Sum32())%len(hm.List)], err
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Hash) != 0 && len(conf.Masking.Mask.HashInURI) != 0 {
		return nil, false, fmt.Errorf("2 different hash choices")
	}
	if len(conf.Masking.Mask.Hash) != 0 {
		var maskHash MaskEngine
		maskHash.List = append(maskHash.List, conf.Masking.Mask.Hash...)
		return maskHash, true, nil
	}
	if len(conf.Masking.Mask.HashInURI) != 0 {
		list, err := uri.Read(conf.Masking.Mask.HashInURI)
		return MaskEngine{list.Collect()}, true, err
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	return func(choices_ []interface{}, input model.Entry) (model.Entry, error) {
		choices := make([]model.Entry, len(choices_))
		for i, c := range choices_ {
			choices[i] = c
		}
		mask := NewMask(choices)
		return mask.Mask(input)
	}
}

func FuncInUri(seed int64, seedField string) interface{} {
	return func(uristr string, input model.Entry) (model.Entry, error) {
		log.Warn().Msg("Using MaskHashInUri from a template can cause performance issues")
		list, err := uri.Read(uristr)
		if err != nil {
			return nil, err
		}
		mask := NewMask(list.Collect())
		return mask.Mask(input)
	}
}
