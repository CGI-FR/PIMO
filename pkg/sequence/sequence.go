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
)

type MaskEngine struct {
	format  string
	varying string
}

func NewMask(format string, varying string) (MaskEngine, error) {
	if len(varying) == 0 {
		varying = "0123456789"
	}

	return MaskEngine{
		format:  format,
		varying: varying,
	}, nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask sequence")

	return e, nil
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
