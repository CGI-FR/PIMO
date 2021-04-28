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

package randdura

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"

	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is to mask a value thanks to 2 durations
type MaskEngine struct {
	Min  time.Duration
	Max  time.Duration
	rand *rand.Rand
}

// NewMask create a MaskEngine with 2 ISO8601 duration strings
func NewMask(minString, maxString string, seed int64) (MaskEngine, error) {
	var durMin time.Duration
	var durMax time.Duration
	var err error
	if minString == "" {
		durMin = 0
	} else {
		durMin, err = duration.ParseDuration(minString)
	}
	if err != nil {
		return MaskEngine{}, err
	}
	if maxString == "" {
		durMax = 0
	} else {
		durMax, err = duration.ParseDuration(maxString)
	}
	// nolint: gosec
	return MaskEngine{durMin, durMax, rand.New(rand.NewSource(seed))}, err
}

// Mask masks a time value with a duration
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask randomDuration")
	var t time.Time
	var err error
	switch v := e.(type) {
	case string:
		t, err = time.Parse(time.RFC3339, v)
	case time.Time:
		t = v
	default:
		return e, fmt.Errorf("Field to mask is not a time nor a string")
	}
	if err != nil {
		return nil, err
	}

	var dura int64

	switch diff := me.Max.Nanoseconds() - me.Min.Nanoseconds(); { // missing expression means "true"
	case diff > 0:
		dura = me.rand.Int63n(diff) + me.Min.Nanoseconds()
	case diff < 0:
		dura = me.Min.Nanoseconds() - me.rand.Int63n(-1*diff)
	default:
		return t, nil
	}

	return t.Add(time.Duration(dura)), nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.RandomDuration.Min) != 0 || len(conf.Mask.RandomDuration.Max) != 0 { // set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Selector.Jsonpath))
		seed += int64(h.Sum64())
		mask, err := NewMask(conf.Mask.RandomDuration.Min, conf.Mask.RandomDuration.Max, seed)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
