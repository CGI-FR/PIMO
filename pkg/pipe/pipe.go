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
	"fmt"
	"hash/fnv"
	"regexp"

	over "github.com/Trendyol/overlog"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	source       string
	pipeline     model.Pipeline
	injectParent string
	injectRoot   string
}

// NewMask return a MaskEngine from a value
func NewMask(seed int64, injectParent string, injectRoot string, caches map[string]model.Cache, filename string, masking ...model.Masking) (MaskEngine, error) {
	var definition model.Definition
	var err error
	if len(filename) > 0 {
		definition, err = model.LoadPipelineDefinitionFromFile(filename)
		if err != nil {
			return MaskEngine{filename, nil, injectParent, injectRoot}, err
		}
		// merge the current seed with the seed provided by configuration on the pipe
		definition.Seed += seed
	} else {
		definition = model.Definition{Seed: seed + 1, Masking: masking}
	}
	pipeline := model.NewPipeline(nil)
	pipeline, _, err = model.BuildPipeline(pipeline, definition, caches)
	return MaskEngine{"", pipeline, injectParent, injectRoot}, err
}

func (me MaskEngine) MaskContext(e model.Dictionary, key string, context ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask pipe")
	var result []model.Dictionary
	input := []model.Dictionary{}

	copy := model.CopyDictionary(e)

	value, ok := e.GetValue(key)
	if !ok || value == nil {
		return copy, nil
	}

	for _, elemInput := range model.CleanDictionarySlice(e.Get(key)) {
		if len(me.injectParent) > 0 {
			elemInput.Set(me.injectParent, copy)
		}
		if len(me.injectRoot) > 0 {
			rootcopy := model.CopyDictionary(context[0])
			elemInput.Set(me.injectRoot, rootcopy)
		}
		input = append(input, elemInput)
	}
	saveConfig, _ := over.MDC().Get("config")
	savePath, _ := over.MDC().Get("path")
	saveContext, _ := over.MDC().Get("context")
	over.MDC().Set("context", fmt.Sprintf("%s/%s[1]", saveContext, savePath))
	if len(me.source) > 0 {
		over.MDC().Set("config", me.source)
	}
	// TODO: possible refactoring
	// model.NewSourceFromSlice(input).
	//			Process(model.NewCounterProcessWithCallback("input-line", 1, updateContext)).
	//			Process(me.pipeline).
	//			AddSink(model.NewSinkToSlice(&result)).
	//			Run()
	err := me.pipeline.
		WithSource(model.NewSourceFromSlice(input)).
		Process(model.NewCounterProcessWithCallback("internal", 1, updateContext)).
		AddSink(model.NewSinkToSlice(&result)).
		Run()
	over.MDC().Set("config", saveConfig)
	over.MDC().Set("path", savePath)
	over.MDC().Set("context", saveContext)
	if err != nil {
		return model.NewDictionary(), err
	}
	for _, dict := range result {
		if len(me.injectParent) > 0 {
			dict.Delete(me.injectParent)
		}
		if len(me.injectRoot) > 0 {
			dict.Delete(me.injectRoot)
		}
	}
	copy.Set(key, result)
	return copy, nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskContextEngine, bool, error) {
	if len(conf.Masking.Mask.Pipe.Masking) > 0 || len(conf.Masking.Mask.Pipe.DefinitionFile) > 0 {
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		mask, err := NewMask(conf.Seed, conf.Masking.Mask.Pipe.InjectParent, conf.Masking.Mask.Pipe.InjectRoot, conf.Cache, conf.Masking.Mask.Pipe.DefinitionFile, conf.Masking.Mask.Pipe.Masking...)
		if err != nil {
			return mask, true, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}

var re = regexp.MustCompile(`(\[\d*\])?$`)

func updateContext(counter int) {
	context := over.MDC().GetString("context")
	over.MDC().Set("context", re.ReplaceAllString(context, fmt.Sprintf("[%d]", counter)))
}
