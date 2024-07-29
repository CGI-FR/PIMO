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
	"encoding/binary"
	"fmt"
	"math"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/kenshaw/baseconv"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/sha3"
)

type MaskEngine struct {
	length int
	domain string
	salt   []byte
	seeder model.Seeder
}

func NewMask(length int, resistance int, domain string, seed int64, seeder model.Seeder) MaskEngine {
	if len(domain) < 2 {
		domain = "0123456789abcdef"
	}
	if resistance > 0 {
		length = lengthWithResistance(resistance)
	}
	salt := make([]byte, 0, 16)
	salt = binary.LittleEndian.AppendUint64(salt, uint64(seed))
	return MaskEngine{
		length: length,
		domain: domain,
		salt:   salt,
		seeder: seeder,
	}
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask sha3")

	if e == nil {
		return e, nil
	}

	str := fmt.Sprintf("%v", e)

	salt := me.salt
	if len(context) > 0 {
		seed, ok, err := me.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			salt = binary.LittleEndian.AppendUint64(me.salt, uint64(seed))
		}
	}

	h := make([]byte, me.length)
	cshake := sha3.NewCShake256([]byte{}, salt)
	if _, err := cshake.Write([]byte(str)); err != nil {
		return e, err
	}
	if _, err := cshake.Read(h); err != nil {
		return e, err
	}

	if me.domain != "0123456789abcdef" {
		return baseconv.Convert(fmt.Sprintf("%x", h), "0123456789abcdef", me.domain)
	} else {
		return fmt.Sprintf("%x", h), nil
	}
}

func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.Sha3.Length > 0 || conf.Masking.Mask.Sha3.Resistance > 0 {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

		return NewMask(conf.Masking.Mask.Sha3.Length, conf.Masking.Mask.Sha3.Resistance, conf.Masking.Mask.Sha3.Domain, conf.Seed, seeder), true, nil
	}
	return nil, false, nil
}

const (
	BASE2 = 2
	BASE8 = 8
)

func lengthWithResistance(resistance int) int {
	power := 1

	for {
		if math.Pow(BASE2, float64(power)) > float64(resistance) {
			break
		}
		power++
	}

	return int(math.Ceil(float64(power) * BASE2 / BASE8))
}
