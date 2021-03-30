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

package remove

import (
	"github.com/cgi-fr/pimo/pkg/model"
)

// MaskEngine is an empty mask to link to a field to remove
type MaskEngine struct {
}

// NewMask return an empty mask
func NewMask() MaskEngine {
	return MaskEngine{}
}

// MaskContext remove the field
func (rm MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	delete(context, key)

	return context, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64) (model.MaskContextEngine, bool, error) {
	if conf.Mask.Remove {
		return NewMask(), true, nil
	}
	return nil, false, nil
}
