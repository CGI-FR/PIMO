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

func NewDeleteMaskEngineProcess(selector Selector) Processor {
	return &DeleteMaskEngineProcess{selector: selector}
}

type DeleteMaskEngineProcess struct {
	selector Selector
}

func (dp *DeleteMaskEngineProcess) Open() (err error) {
	return err
}

func (dp *DeleteMaskEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector, _ SinkProcess) error {
	over.AddGlobalFields("path")
	over.MDC().Set("path", dp.selector)
	defer func() { over.MDC().Remove("path") }()
	result := dictionary
	applied := dp.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		return DELETE, nil
	})

	if !applied {
		statistics.IncIgnoredPathsCount()
		log.Warn().Msg("Path not found")
	}

	out.Collect(result)
	return nil
}
