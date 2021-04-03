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

package pipe

import (
	"github.com/cgi-fr/pimo/pkg/model"
)

type source struct {
	value *model.Dictionary
	read  *bool
}

func (s *source) Open() error {
	return nil
}

func (s *source) Err() error {
	return nil
}

func (s *source) Next() bool {
	return *s.read
}

func (s *source) Value() model.Dictionary {
	defer func() { *s.read = false }()
	return *s.value
}

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	seed         int64
	source       *source
	pipeline     model.Pipeline
	injectParent string
	injectRoot   string
}

// NewMask return a MaskEngine from a value
func NewMask(seed int64, injectParent string, injectRoot string, caches map[string]model.Cache, masking ...model.Masking) MaskEngine {
	read := false
	source := source{nil, &read}
	definition := model.Definition{Seed: seed, Masking: masking}
	pipeline := model.NewPipeline(&source)
	pipeline, _, _ = model.BuildPipeline(pipeline, definition, caches)
	return MaskEngine{seed, &source, pipeline, injectParent, injectRoot}
}

// Mask return a Constant from a MaskEngine
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	var result []model.Dictionary
	value := InterfaceToDictionaryEntry(e)
	me.source.value = &value
	*me.source.read = true
	err := me.pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	if err != nil {
		return nil, err
	}
	return result[0], err
}

// Factory create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Pipe.Masking) > 0 {
		return NewMask(seed, conf.Mask.Pipe.InjectParent, conf.Mask.Pipe.InjectRoot, caches, conf.Mask.Pipe.Masking...), true, nil
	}
	return nil, false, nil
}

// InterfaceToDictionary returns a model.Dictionary from an interface
func InterfaceToDictionaryEntry(inter interface{}) model.Dictionary {
	dic := make(map[string]model.Entry)
	mapint := inter.(model.Dictionary)

	for k, v := range mapint {
		switch typedValue := v.(type) {
		case map[string]interface{}:
			dic[k] = model.InterfaceToDictionary(v)
		case []interface{}:
			tab := []model.Entry{}
			for _, item := range typedValue {
				_, dico := item.(map[string]interface{})

				if dico {
					tab = append(tab, model.InterfaceToDictionary(item))
				} else {
					tab = append(tab, item)
				}
			}
			dic[k] = tab
		default:
			dic[k] = v
		}
	}

	return dic
}
