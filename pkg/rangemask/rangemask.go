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

package rangemask

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that always mask by replacing with a scale
type MaskEngine struct {
	rangeScale int
}

// NewMask return a MaskEngine from a value
func NewMask(scale int) MaskEngine {
	return MaskEngine{scale}
}

// Mask return a range from a MaskEngine
func (rm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask range")
	if e == nil {
		return e, nil
	}
	val := tryConvertToFloat64(e)
	i, ok := val.(float64)
	if !ok {
		return e, fmt.Errorf("%v is not a number", e)
	}
	scaledValue := int(i) / rm.rangeScale * rm.rangeScale
	rangedValue := "[" + strconv.Itoa(scaledValue) + ";" + strconv.Itoa(scaledValue+rm.rangeScale-1) + "]"
	return rangedValue, nil
}

// Factory Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.RangeMask != 0 {
		return NewMask(conf.Masking.Mask.RangeMask), true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	return func(maybescale interface{}, input model.Entry) (model.Entry, error) {
		var err error
		var scale int
		switch typedscale := maybescale.(type) {
		case string:
			scale, err = strconv.Atoi(typedscale)
			if err != nil {
				return nil, err
			}
		case int:
			scale = typedscale
		case int64:
			if typedscale > math.MaxInt || typedscale < math.MinInt {
				return nil, errors.New("scale is out of range")
			}
			scale = int(typedscale)
		case int32:
			scale = int(typedscale)
		case int16:
			scale = int(typedscale)
		case int8:
			scale = int(typedscale)
		case uint:
			scale = int(typedscale)
		case uint64:
			if typedscale > math.MaxInt {
				return nil, errors.New("scale is out of range")
			}
			scale = int(typedscale)
		case uint32:
			scale = int(typedscale)
		case uint16:
			scale = int(typedscale)
		case uint8:
			scale = int(typedscale)
		}
		mask := NewMask(scale)
		return mask.Mask(input)
	}
}

func tryConvertToFloat64(input model.Entry) model.Entry {
	switch typedinput := input.(type) {
	case string:
		if v, err := strconv.ParseFloat(typedinput, 64); err == nil {
			return v
		} else {
			log.Warn().Err(err).Msg("")
		}
	case int:
		return float64(typedinput)
	case int64:
		return float64(typedinput)
	case int32:
		return float64(typedinput)
	case int16:
		return float64(typedinput)
	case int8:
		return float64(typedinput)
	case uint:
		return float64(typedinput)
	case uint64:
		return float64(typedinput)
	case uint32:
		return float64(typedinput)
	case uint16:
		return float64(typedinput)
	case uint8:
		return float64(typedinput)
	case float32:
		return float64(typedinput)
	}
	return input
}
