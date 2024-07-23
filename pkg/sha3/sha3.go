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

package sha3

import (
	"fmt"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/kenshaw/baseconv"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/sha3"
)

type MaskEngine struct {
	length int
	domain string
}

func NewMask(length int, domain string) MaskEngine {
	if len(domain) < 2 {
		domain = "0123456789abcdef"
	}
	return MaskEngine{
		length: length,
		domain: domain,
	}
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask sha3")

	if e == nil {
		return e, nil
	}

	if str, ok := e.(string); ok {
		h := make([]byte, me.length)
		sha3.ShakeSum256(h, []byte(str))

		if me.domain != "0123456789abcdef" {
			return baseconv.Convert(fmt.Sprintf("%x", h), "0123456789abcdef", me.domain)
		}
	}

	return e, nil
}

func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.Sha3.Length > 0 {
		return NewMask(conf.Masking.Mask.Sha3.Length, conf.Masking.Mask.Sha3.Domain), true, nil
	}
	return nil, false, nil
}
