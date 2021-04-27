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

package dateparser

import (
	"fmt"
	"time"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a type to change a date format
type MaskEngine struct {
	inputFormat  string
	outputFormat string
}

// NewMask create a MaskEngine
func NewMask(input, output string) MaskEngine {
	return MaskEngine{input, output}
}

// Mask change a time format
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Debug().Interface("data", e).Msg("Mask dateparser")
	var t time.Time
	var err error
	if me.inputFormat != "" {
		timestring := fmt.Sprintf("%v", e)
		t, err = time.Parse(me.inputFormat, timestring)
		if err != nil {
			return nil, err
		}
	} else {
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
	}
	if me.outputFormat != "" {
		return t.Format(me.outputFormat), nil
	}
	return t, nil
}

// Factory Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.DateParser.InputFormat) != 0 || len(conf.Mask.DateParser.OutputFormat) != 0 {
		mask := NewMask(conf.Mask.DateParser.InputFormat, conf.Mask.DateParser.OutputFormat)
		return mask, true, nil
	}
	return nil, false, nil
}
