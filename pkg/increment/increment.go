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

package increment

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a struct to create incremental int
type MaskEngine struct {
	Value     *int
	Increment int
}

// NewMask create an IncrementalMask
func NewMask(start, incr int) MaskEngine {
	value := &start
	return MaskEngine{value, incr}
}

// Mask masks a value with an incremental int
func (incr MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	output := *incr.Value
	*incr.Value += incr.Increment
	return output, nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.Incremental.Increment != 0 {
		return NewMask(conf.Mask.Incremental.Start, conf.Mask.Incremental.Increment), true, nil
	}
	return nil, false, nil
}
