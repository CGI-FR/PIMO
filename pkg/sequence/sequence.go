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
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
)

type MaskEngineState struct {
	counter int
}

func (mes MaskEngineState) toCounters(length int, base int) []int {
	result := make([]int, length)

	counter := mes.counter
	for counter > 0 {
		result = append([]int{counter % base}, result...)
		counter = counter / base
	}

	return result
}

type MaskEngine struct {
	format   []rune
	varying  []rune
	state    *MaskEngineState
	base     int
	length   int
	counters []int
}

func NewMask(format string, varying string, state *MaskEngineState) (MaskEngine, error) {
	if len(varying) == 0 {
		varying = "0123456789"
	}

	if state == nil {
		state = &MaskEngineState{counter: 0}
	}

	formatRunes := []rune(format)
	varyingRunes := []rune(varying)

	varyingLength := 0
	for _, char := range formatRunes {
		if slices.Contains(varyingRunes, char) {
			varyingLength++
		}
	}

	length := len(formatRunes)
	base := len(varyingRunes)

	return MaskEngine{
		format:   formatRunes,
		varying:  varyingRunes,
		state:    state,
		base:     base,
		length:   length,
		counters: state.toCounters(varyingLength, base),
	}, nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask sequence")

	result := make([]rune, me.length)

	counterIdx := 0
	for i := range me.format {
		char := me.format[me.length-i-1]
		if slices.Contains(me.varying, char) {
			result[me.length-i-1] = me.varying[me.counters[counterIdx]]
			counterIdx++
		} else {
			result[me.length-i-1] = char
		}
	}

	me.increment()

	return string(result), nil
}

func (me MaskEngine) increment() {
	me.state.counter++
	for i := range me.counters {
		if me.counters[i] < me.base-1 {
			me.counters[i]++
			return
		} else {
			me.counters[i] = 0
		}
	}
	me.state.counter = 0
}

func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Sequence.Format) > 0 {
		mask, err := NewMask(conf.Masking.Mask.Sequence.Format, conf.Masking.Mask.Sequence.Varying, nil)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	state := &MaskEngineState{counter: 0}
	return func(format string, varying string) (model.Entry, error) {
		mask, err := NewMask(format, varying, state)
		if err != nil {
			return nil, err
		}
		return mask.Mask(nil)
	}
}
