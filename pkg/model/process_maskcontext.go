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
	"github.com/cgi-fr/pimo/pkg/selector"
)

func NewMaskContextEngineProcess(selector Selector, mask MaskContextEngine) Processor {
	return &MaskContextEngineProcess{selector, mask}
}

type MaskContextEngineProcess struct {
	selector Selector
	mask     MaskContextEngine
}

func (mcep *MaskContextEngineProcess) Open() error {
	return nil
}

func (mcep *MaskContextEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) (ret error) {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	mcep.selector.ApplyContext(result, func(rootContext, parentContext Dictionary, key string, _ Entry) (Action, Entry) {
		masked, err := mcep.mask.MaskContext(parentContext, key, rootContext, parentContext)
		if err != nil {
			ret = err
			return selector.NOTHING, nil
		}
		value, ok := masked[key]
		if !ok {
			return selector.NOTHING, nil
		}
		return selector.WRITE, value
	})

	if ret == nil {
		out.Collect(result)
	}

	return
}
