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

package model

import (
	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/rs/zerolog/log"
)

func NewMaskEngineProcess(selector Selector, mask MaskEngine, preserve string) Processor {
	return &MaskEngineProcess{selector, mask, preserve}
}

type MaskEngineProcess struct {
	selector Selector
	mask     MaskEngine
	preserve string
}

func (mep *MaskEngineProcess) Open() error {
	return nil
}

func (mep *MaskEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) (ret error) {
	over.AddGlobalFields("path")
	over.MDC().Set("path", mep.selector)
	defer func() { over.MDC().Remove("path") }()
	result := dictionary
	applied := mep.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		switch {
		case value == nil && (mep.preserve == "null" || mep.preserve == "blank"):
			log.Trace().Msgf("Preserve %s value, skip masking", mep.preserve)
			return NOTHING, nil
		case value == "" && (mep.preserve == "empty" || mep.preserve == "blank"):
			log.Trace().Msgf("Preserve %s value, skip masking", mep.preserve)
			return NOTHING, nil
		default:
			masked, err := mep.mask.Mask(value, rootContext.Get(".").(Dictionary), parentContext)
			if err != nil {
				ret = err
				return NOTHING, nil
			}
			return WRITE, masked
		}
	})

	if !applied {
		statistics.IncIgnoredPathsCount()
		log.Warn().Msg("Path not found")
	}

	if ret == nil {
		out.Collect(result)
		return
	}

	if ret != nil && skipLineOnError {
		log.Warn().AnErr("error", ret).Msg("Line skipped")
		statistics.IncIgnoredLinesCount()
		return nil
	}

	if ret != nil && skipFieldOnError {
		log.Warn().AnErr("error", ret).Msg("Field skipped")
		statistics.IncIgnoredFieldsCount()
		mep.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
			return DELETE, nil
		})
		out.Collect(result)
		return nil
	}

	return ret
}
