package model

import (
	"errors"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
)

var maskContextFactories []MaskContextFactory
var maskFactories []MaskFactory

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
