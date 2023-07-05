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
	"errors"
	"fmt"
	"os"
	"strings"
	tmpl "text/template"
	"time"

	"github.com/cgi-fr/pimo/pkg/functions"
	"github.com/goccy/go-yaml"
)

// nolint: gochecknoglobals
var (
	maskContextFactories []MaskContextFactory
	maskFactories        []MaskFactory
	skipLineOnError      bool
	skipFieldOnError     bool
	skipLogFile          string
)

func InjectMaskContextFactories(factories []MaskContextFactory) {
	maskContextFactories = factories
}

func InjectMaskFactories(factories []MaskFactory) {
	maskFactories = factories
}

func InjectConfig(skipLineOnErrorValue bool, skipFieldOnErrorValue bool, skipLogFileValue string) {
	skipLineOnError = skipLineOnErrorValue
	skipFieldOnError = skipFieldOnErrorValue
	skipLogFile = skipLogFileValue
}

func BuildCaches(caches map[string]CacheDefinition, existing map[string]Cache) map[string]Cache {
	if existing == nil {
		existing = map[string]Cache{}
	}
	for name, conf := range caches {
		if _, exist := existing[name]; !exist {
			if conf.Unique {
				existing[name] = NewUniqueMemCache()
			} else {
				existing[name] = NewMemCache()
			}
		}
	}
	return existing
}

func BuildFuncMap(funcs map[string]Function, fns tmpl.FuncMap) (tmpl.FuncMap, error) {
	b := functions.Builder{
		Definitions: map[string]functions.Definition{},
	}

	for name, f := range funcs {
		params := []functions.ParamDefinition{}
		for _, p := range f.Params {
			params = append(params, functions.ParamDefinition{Name: p.Name})
		}

		b.Definitions[name] = functions.Definition{
			Params: params,
			Body:   f.Body,
		}
	}

	return b.Build(fns)
}

func BuildPipeline(pipeline Pipeline, conf Definition, caches map[string]Cache, functions tmpl.FuncMap) (Pipeline, map[string]Cache, error) {
	caches = BuildCaches(conf.Caches, caches)
	functions, err := BuildFuncMap(conf.Functions, functions)
	if err != nil {
		return nil, nil, err
	}
	cleaners := []Processor{}

	// pack data inside a root container
	pipeline = pipeline.Process(NewPackProcess())

	for _, masking := range conf.Masking {
		allSelectors := masking.Selectors
		if sel := masking.Selector; sel.Jsonpath != "" {
			allSelectors = append(allSelectors, sel)
		}

		for _, sel := range allSelectors {
			nbArg := 0

			if cache, ok := caches[masking.Cache]; ok {
				pipeline = pipeline.Process(&LookupValueSetter{cache, NewPackedPathSelector(sel.Jsonpath)})
			}

			allMasksDefinition := append([]MaskType{masking.Mask}, masking.Masks...)

			for _, maskDefinition := range allMasksDefinition {
				virtualMask := Masking{
					Selector:  sel,
					Selectors: nil,
					Mask:      maskDefinition,
					Masks:     nil,
					Cache:     masking.Cache,
					Preserve:  masking.Preserve,
					Seed:      masking.Seed,
				}

				if virtualMask.Mask.FromCache != "" {
					cache, ok := caches[virtualMask.Mask.FromCache]
					if !ok {
						return nil, nil, errors.New("Cache '" + virtualMask.Cache + "' not found for '" + virtualMask.Selector.Jsonpath + "'")
					}
					pipeline = pipeline.Process(NewFromCacheProcess(NewPackedPathSelector(virtualMask.Selector.Jsonpath), cache, virtualMask.Preserve))
					nbArg++
				}

				configuration := MaskFactoryConfiguration{Masking: virtualMask, Seed: conf.Seed, Cache: caches, Functions: functions}

				for _, factory := range maskFactories {
					mask, present, err := factory(configuration)
					if err != nil {
						return nil, nil, errors.New(err.Error() + " for " + virtualMask.Selector.Jsonpath)
					}
					if present {
						if virtualMask.Cache != "" {
							cache, ok := caches[virtualMask.Cache]
							if !ok {
								return nil, nil, errors.New("Cache '" + virtualMask.Cache + "' not found for '" + virtualMask.Selector.Jsonpath + "'")
							}
							switch typedCache := cache.(type) {
							case UniqueCache:
								mask = NewUniqueMaskCacheEngine(typedCache, mask)
							default:
								mask = NewMaskCacheEngine(typedCache, mask)
							}
						}
						pipeline = pipeline.Process(NewMaskEngineProcess(NewPackedPathSelector(virtualMask.Selector.Jsonpath), mask, virtualMask.Preserve, skipLogFile))
						nbArg++
					}
				}

				for _, factory := range maskContextFactories {
					mask, present, err := factory(configuration)
					if err != nil {
						return nil, nil, errors.New(err.Error() + " for " + virtualMask.Selector.Jsonpath)
					}
					if present {
						if virtualMask.Cache != "" {
							cache, ok := caches[virtualMask.Cache]
							if !ok {
								return nil, nil, errors.New("Cache '" + virtualMask.Cache + "' not found for '" + virtualMask.Selector.Jsonpath + "'")
							}
							switch typedCache := cache.(type) {
							case UniqueCache:
								mask = NewUniqueMaskContextCacheEngine(typedCache, mask)
							default:
								mask = NewMaskContextCacheEngine(typedCache, mask)
							}
						}
						pipeline = pipeline.Process(NewMaskContextEngineProcess(NewPackedPathSelector(virtualMask.Selector.Jsonpath), mask, skipLogFile))
						nbArg++
						if i, hasCleaner := mask.(HasCleaner); hasCleaner {
							cleaners = append(cleaners, NewMaskContextEngineProcess(NewPackedPathSelector(virtualMask.Selector.Jsonpath), i.GetCleaner(), skipLogFile))
						}
					}
				}
			}

			if cache, ok := caches[masking.Cache]; ok {
				pipeline = pipeline.Process(&LookupValueDiscarder{cache})
			}

			if nbArg == 0 {
				return pipeline, nil, errors.New("No masks defined for " + masking.Selector.Jsonpath)
			}
		}
	}
	for _, cleaner := range cleaners {
		pipeline = pipeline.Process(cleaner)
	}

	// unpack data from the container
	pipeline = pipeline.Process(NewUnpackProcess())

	return pipeline, caches, nil
}

func LoadPipelineDefinitionFromYAML(source []byte) (Definition, error) {
	var conf Definition
	err := yaml.Unmarshal(source, &conf)
	if err != nil {
		return conf, err
	}
	(&conf).SetSeed(conf.Seed)
	return conf, nil
}

func LoadPipelineDefinitionFromFile(filename string) (Definition, error) {
	source, err := os.ReadFile(filename)
	if err != nil {
		return Definition{}, err
	}
	return LoadPipelineDefinitionFromYAML(source)
}

func LoadPipelineDefintionFromOneLiner(oneLine []string) (Definition, error) {
	var conf Definition
	conf.Masking = []Masking{}
	for _, value := range oneLine {
		i := strings.Index(value, "=")
		if i == -1 {
			return conf, fmt.Errorf("%s is not of the form 'jsonpath={mask}'", value)
		}
		jsonpath, maskString := value[:i], value[(i+1):]

		masking := Masking{Selector: SelectorType{Jsonpath: jsonpath}}

		if err := yaml.Unmarshal([]byte(maskString), &masking.Mask); err != nil {
			if err2 := yaml.Unmarshal([]byte(maskString), &masking.Masks); err2 != nil {
				return conf, err
			}
		}

		conf.Masking = append(conf.Masking, masking)
	}
	(&conf).SetSeed(conf.Seed)
	return conf, nil
}

// setSeed compute seed from current nano second if seedValue is 0
func (conf *Definition) SetSeed(seedValue int64) {
	switch {
	case seedValue > 0:
		conf.Seed = seedValue
	case seedValue == 0:
		conf.Seed = time.Now().UnixNano()
	}
}
