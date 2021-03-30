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
	"fmt"
	"math/rand"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/uri"
)

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand *rand.Rand
	list []model.Entry
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(list []model.Entry, seed int64) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), list}
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return mrl.list[mrl.rand.Intn(len(mrl.list))], nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.RandomChoice) != 0 && len(conf.Mask.RandomChoiceInURI) != 0 {
		return nil, false, fmt.Errorf("2 different random choices")
	}
	if len(conf.Mask.RandomChoice) != 0 {
		return NewMask(conf.Mask.RandomChoice, seed), true, nil
	}
	if len(conf.Mask.RandomChoiceInURI) != 0 {
		list, err := uri.Read(conf.Mask.RandomChoiceInURI)
		return NewMask(list, seed), true, err
	}
	return nil, false, nil
}
