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
	"text/template"

	over "github.com/adrienaury/zeromdc"
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
func NewMask(seed int64, injectParent string, injectRoot string, caches map[string]model.Cache, fns template.FuncMap, filename string, masking ...model.Masking) (MaskEngine, error) {
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
	pipeline, _, err = model.BuildPipeline(pipeline, definition, caches, fns)
	return MaskEngine{"", pipeline, injectParent, injectRoot}, err
}

func (me MaskEngine) MaskContext(entry model.Dictionary, key string, context ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask pipe")
	var result []model.Entry
	input := []model.Dictionary{}

	value, ok := entry.GetValue(key)
	if !ok || value == nil {
		return entry, nil
	}

	for _, elemInput := range model.CleanDictionarySlice(entry.Get(key)) {
		if len(me.injectParent) > 0 {
			elemInput.Set(me.injectParent, entry)
		}
		if len(me.injectRoot) > 0 {
			elemInput.Set(me.injectRoot, model.CopyDictionary(context[0]))
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
	// possible refactoring
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
		return entry, err
	}
	for _, dict := range result {
		if len(me.injectParent) > 0 {
			dict.(model.Dictionary).Delete(me.injectParent)
		}
		if len(me.injectRoot) > 0 {
			dict.(model.Dictionary).Delete(me.injectRoot)
		}
	}
	entry.Set(key, result)
	return entry, nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskContextEngine, bool, error) {
	if len(conf.Masking.Mask.Pipe.Masking) > 0 || len(conf.Masking.Mask.Pipe.DefinitionFile) > 0 {
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64()) //nolint:gosec
		mask, err := NewMask(conf.Seed, conf.Masking.Mask.Pipe.InjectParent, conf.Masking.Mask.Pipe.InjectRoot, conf.Cache, conf.Functions, conf.Masking.Mask.Pipe.DefinitionFile, conf.Masking.Mask.Pipe.Masking...)
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
