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

import "github.com/cgi-fr/pimo/pkg/selector"

func NewDeleteMaskEngineProcess(selector Selector) Processor {
	return &DeleteMaskEngineProcess{selector: selector}
}

type DeleteMaskEngineProcess struct {
	selector Selector
}

func (dp *DeleteMaskEngineProcess) Open() (err error) {
	return err
}

func (dp *DeleteMaskEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	dp.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		return selector.DELETE, nil
	})
	out.Collect(result)
	return nil
}

func NewMapProcess(mapper Mapper) Processor {
	return MapProcess{mapper: mapper}
}
