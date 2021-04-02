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
	"errors"
	"strconv"

	"github.com/cgi-fr/pimo/pkg/add"
	"github.com/cgi-fr/pimo/pkg/command"
	"github.com/cgi-fr/pimo/pkg/constant"
	"github.com/cgi-fr/pimo/pkg/dateparser"
	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/ff1"
	"github.com/cgi-fr/pimo/pkg/fluxuri"
	"github.com/cgi-fr/pimo/pkg/hash"
	"github.com/cgi-fr/pimo/pkg/increment"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/randdate"
	"github.com/cgi-fr/pimo/pkg/randdura"
	"github.com/cgi-fr/pimo/pkg/randomdecimal"
	"github.com/cgi-fr/pimo/pkg/randomint"
	"github.com/cgi-fr/pimo/pkg/randomlist"
	"github.com/cgi-fr/pimo/pkg/rangemask"
	"github.com/cgi-fr/pimo/pkg/regex"
	"github.com/cgi-fr/pimo/pkg/remove"
	"github.com/cgi-fr/pimo/pkg/replacement"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/cgi-fr/pimo/pkg/weightedchoice"
)

type source struct {
	value model.Dictionary
	read  *bool
}

func (s source) Open() error {
	return nil
}

func (s source) Err() error {
	return nil
}

func (s source) Next() bool {
	return !*s.read
}

func (s source) Value() model.Dictionary {
	defer func() { *s.read = true }()
	return s.value
}

func injectMaskFactories() []model.MaskFactory {
	return []model.MaskFactory{

		constant.Factory,
		command.Factory,
		randomlist.Factory,
		randomint.Factory,
		weightedchoice.Factory,
		regex.Factory,
		hash.Factory,
		randdate.Factory,
		increment.Factory,
		replacement.Factory,
		duration.Factory,
		templatemask.Factory,

		rangemask.Factory,
		randdura.Factory,
		randomdecimal.Factory,
		dateparser.Factory,
		ff1.Factory,
		Factory,
	}
}

func injectMaskContextFactories() []model.MaskContextFactory {
	return []model.MaskContextFactory{
		fluxuri.Factory,
		add.Factory,
		remove.Factory,
	}
}

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	masking      []model.Masking
	injectParent string
	injectRoot   string
}

func BuildPipeline(pipeline model.Pipeline, masking []model.Masking) (model.Pipeline, error) {
	maskFactories := injectMaskFactories()
	maskContextFactories := injectMaskContextFactories()
	for _, v := range masking {
		nbArg := 0

		for _, factory := range maskFactories {
			mask, present, err := factory(v, 0)
			if err != nil {
				return nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				pipeline = pipeline.Process(model.NewMaskEngineProcess(model.NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}

		for _, factory := range maskContextFactories {
			mask, present, err := factory(v, 0)
			if err != nil {
				return nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				pipeline = pipeline.Process(model.NewMaskContextEngineProcess(model.NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}
		if nbArg != 1 {
			return pipeline, errors.New("Not the right number of argument for " + v.Selector.Jsonpath + ". There should be 1 and there is " + strconv.Itoa(nbArg))
		}
	}
	return pipeline, nil
}

// NewMask return a MaskEngine from a value
func NewMask(injectParent string, injectRoot string, masking ...model.Masking) MaskEngine {
	return MaskEngine{masking, injectParent, injectRoot}
}

// Mask return a Constant from a MaskEngine
func (pipem MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	read := false
	source := source{InterfaceToDictionaryEntry(e), &read}
	pipeline := model.NewPipeline(source)
	pipeline, _ = BuildPipeline(pipeline, pipem.masking)
	var result []model.Dictionary
	pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	return result[0], nil
}

// Factory create a mask from a configuration
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Pipe.Masking) > 0 {
		return NewMask(conf.Mask.Pipe.InjectParent, "nil", conf.Mask.Pipe.Masking...), true, nil
	}
	return nil, false, nil
}

// InterfaceToDictionary returns a model.Dictionary from an interface
func InterfaceToDictionaryEntry(inter interface{}) model.Dictionary {
	dic := make(map[string]model.Entry)
	mapint := inter.(map[string]model.Entry)

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
