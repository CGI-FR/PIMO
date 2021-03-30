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

package constant

import (
	"github.com/cgi-fr/pimo/pkg/model"
)

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	constValue model.Entry
}

// NewMask return a MaskEngine from a value
func NewMask(data model.Entry) MaskEngine {
	return MaskEngine{data}
}

// Mask return a Constant from a MaskEngine
func (cm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return cm.constValue, nil
}

// Factory create a mask from a configuration
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.Constant != nil {
		return NewMask(conf.Mask.Constant), true, nil
	}
	return nil, false, nil
}
