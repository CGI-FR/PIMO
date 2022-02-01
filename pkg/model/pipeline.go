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
	"io/ioutil"
	"time"

	"github.com/goccy/go-yaml"
)

// nolint: gochecknoglobals
var (
	maskContextFactories []MaskContextFactory
	maskFactories        []MaskFactory
	skipLineOnError      bool
	skipFieldOnError     bool
)

func InjectMaskContextFactories(factories []MaskContextFactory) {
	maskContextFactories = factories
}

func InjectMaskFactories(factories []MaskFactory) {
	maskFactories = factories
}

func InjectConfig(skipLineOnErrorValue bool, skipFieldOnErrorValue bool) {
	skipLineOnError = skipLineOnErrorValue
	skipFieldOnError = skipFieldOnErrorValue
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

func BuildPipeline(pipeline Pipeline, conf Definition, caches map[string]Cache) (Pipeline, map[string]Cache, error) {
	caches = BuildCaches(conf.Caches, caches)
	cleaners := []Processor{}
	for _, masking := range conf.Masking {
		nbArg := 0

		allMasksDefinition := append([]MaskType{masking.Mask}, masking.Masks...)

		for _, maskDefinition := range allMasksDefinition {
			virtualMask := Masking{
				Selector: masking.Selector,
				Mask:     maskDefinition,
				Masks:    nil,
				Cache:    masking.Cache,
				Preserve: masking.Preserve,
			}

			if virtualMask.Mask.FromCache != "" {
				cache, ok := caches[virtualMask.Mask.FromCache]
				if !ok {
					return nil, nil, errors.New("Cache '" + virtualMask.Cache + "' not found for '" + virtualMask.Selector.Jsonpath + "'")
				}
				pipeline = pipeline.Process(NewFromCacheProcess(NewPathSelector(virtualMask.Selector.Jsonpath), cache))
				nbArg++
			}

			for _, factory := range maskFactories {
				mask, present, err := factory(virtualMask, conf.Seed, caches)
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
					pipeline = pipeline.Process(NewMaskEngineProcess(NewPathSelector(virtualMask.Selector.Jsonpath), mask, virtualMask.Preserve))
					nbArg++
				}
			}

			for _, factory := range maskContextFactories {
				mask, present, err := factory(virtualMask, conf.Seed, caches)
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
					pipeline = pipeline.Process(NewMaskContextEngineProcess(NewPathSelector(virtualMask.Selector.Jsonpath), mask))
					nbArg++
					if i, hasCleaner := mask.(HasCleaner); hasCleaner {
						cleaners = append(cleaners, NewMaskContextEngineProcess(NewPathSelector(virtualMask.Selector.Jsonpath), i.GetCleaner()))
					}
				}
			}
		}
		if nbArg == 0 {
			return pipeline, nil, errors.New("No masks defined for " + masking.Selector.Jsonpath)
		}
	}
	for _, cleaner := range cleaners {
		pipeline = pipeline.Process(cleaner)
	}
	return pipeline, caches, nil
}

func LoadPipelineDefinitionFromYAML(filename string) (Definition, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return Definition{}, err
	}
	var conf Definition
	err = yaml.Unmarshal(source, &conf)
	if err != nil {
		return conf, err
	}
	if conf.Seed == 0 {
		conf.Seed = time.Now().UnixNano()
	}
	return conf, nil
}

func LoadMaskDefintionFromOneLiner(oneLine map[string]string) (Definition, error) {
	var conf Definition
	conf.Masking = []Masking{}
	for key, value := range oneLine {
		masking := Masking{}
		var mask []MaskType
		err := yaml.Unmarshal([]byte(value), &mask)
		if err != nil {
			return conf, err
		}
		if len(mask) == 1 {
			masking.Mask = mask[0]
		} else {
			masking.Masks = mask
		}
		masking.Selector = SelectorType{Jsonpath: key}

		conf.Masking = append(conf.Masking, masking)
	}
	if conf.Seed == 0 {
		conf.Seed = time.Now().UnixNano()
	}
	return conf, nil
}
