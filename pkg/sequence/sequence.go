// Copyright (C) 2024 CGI France
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

package sequence

import (
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	format  string
	varying string

	runesCount int
	runes      []rune
	varunes    []rune
	counters   []int
}

func NewMask(format string, varying string) (MaskEngine, error) {
	if len(varying) == 0 {
		varying = "0123456789"
	}

	runes := []rune(format)
	runesCount := len(runes)

	lenVarying := 0
	for _, c := range runes {
		if strings.ContainsRune(varying, c) {
			lenVarying++
		}
	}

	counters := make([]int, lenVarying)

	return MaskEngine{
		format:     format,
		varying:    varying,
		runesCount: runesCount,
		runes:      runes,
		varunes:    []rune(varying),
		counters:   counters,
	}, nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask sequence")

	result := make([]rune, me.runesCount)

	counterIdx := 0
	for i := range me.runes {
		char := me.runes[me.runesCount-i-1]
		if strings.ContainsRune(me.varying, char) {
			result[me.runesCount-i-1] = me.varunes[me.counters[counterIdx]]
			counterIdx++
		} else {
			result[me.runesCount-i-1] = char
		}
	}

	me.increment()

	return string(result), nil
}

func (me MaskEngine) increment() {
	for i := range me.counters {
		if me.counters[i] < len(me.varunes)-1 {
			me.counters[i]++
			break
		} else {
			me.counters[i] = 0
		}
	}
}

func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Sequence.Format) > 0 {
		mask, err := NewMask(conf.Masking.Mask.Sequence.Format, conf.Masking.Mask.Sequence.Varying)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	return func(format string, varying string, input model.Entry) (model.Entry, error) {
		mask, err := NewMask(format, varying)
		if err != nil {
			return nil, err
		}
		return mask.Mask(input)
	}
}
