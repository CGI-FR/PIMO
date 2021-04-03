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
)

// MaskEngine is a list of masking value for hash masking
type MaskEngine struct {
	List []model.Entry
}

// Mask choose a mask value by hash
func (hm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(e.(string)))
	return hm.List[int(h.Sum32())%len(hm.List)], err
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Hash) != 0 && len(conf.Mask.HashInURI) != 0 {
		return nil, false, fmt.Errorf("2 different hash choices")
	}
	if len(conf.Mask.Hash) != 0 {
		var maskHash MaskEngine
		maskHash.List = append(maskHash.List, conf.Mask.Hash...)
		return maskHash, true, nil
	}
	if len(conf.Mask.HashInURI) != 0 {
		list, err := uri.Read(conf.Mask.HashInURI)
		return MaskEngine{list}, true, err
	}
	return nil, false, nil
}
