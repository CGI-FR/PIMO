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
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog"
)

// nolint: gochecknoglobals
var (
	log                  zerolog.Logger = zerolog.Nop()
	maskContextFactories []MaskContextFactory
	maskFactories        []MaskFactory
)

func InjectLogger(logger zerolog.Logger) {
	log = logger
}

func InjectMaskContextFactories(factories []MaskContextFactory) {
	maskContextFactories = factories
}

func InjectMaskFactories(factories []MaskFactory) {
	maskFactories = factories
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
	for _, v := range conf.Masking {
		nbArg := 0

		if v.Mask.FromCache != "" {
			cache, ok := caches[v.Mask.FromCache]
			if !ok {
				return nil, nil, errors.New("Cache '" + v.Cache + "' not found for '" + v.Selector.Jsonpath + "'")
			}
			pipeline = pipeline.Process(NewFromCacheProcess(NewPathSelector(v.Selector.Jsonpath), cache))
			nbArg++
		}

		for _, factory := range maskFactories {
			mask, present, err := factory(v, conf.Seed, caches)
			if err != nil {
				return nil, nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				if v.Cache != "" {
					cache, ok := caches[v.Cache]
					if !ok {
						return nil, nil, errors.New("Cache '" + v.Cache + "' not found for '" + v.Selector.Jsonpath + "'")
					}
					switch typedCache := cache.(type) {
					case UniqueCache:
						mask = NewUniqueMaskCacheEngine(typedCache, mask)
					default:
						mask = NewMaskCacheEngine(typedCache, mask)
					}
				}
				pipeline = pipeline.Process(NewMaskEngineProcess(NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}

		for _, factory := range maskContextFactories {
			mask, present, err := factory(v, conf.Seed, caches)
			if err != nil {
				return nil, nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				pipeline = pipeline.Process(NewMaskContextEngineProcess(NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}
		if nbArg != 1 {
			return pipeline, nil, errors.New("Not the right number of argument for " + v.Selector.Jsonpath + ". There should be 1 and there is " + strconv.Itoa(nbArg))
		}
	}
	return pipeline, caches, nil
}

func LoadPipelineDefinitionFromYAML(filename string) (Definition, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Err(err).Str("filename", filename).Msg("Cannot load pipeline definition from file")
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
