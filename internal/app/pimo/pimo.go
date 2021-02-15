package pimo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/jsonline"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// YAMLStructure of the file
type YAMLStructure struct {
	Version string               `yaml:"version"`
	Seed    int64                `yaml:"seed"`
	Masking []model.Masking      `yaml:"masking"`
	Caches  map[string]YAMLCache `yaml:"caches"`
}

type MaskingV2 struct {
	Selector SelectorTypeV2         `yaml:"selector"`
	Mask     map[string]interface{} `yaml:"mask"`
	Cache    string                 `yaml:"cache"`
}

type SelectorTypeV2 struct {
	JsonPath string `yaml:"jsonpath"`
}

type YAMLStructureV2 struct {
	Version string               `yaml:"version"`
	Seed    int64                `yaml:"seed"`
	Masking []MaskingV2          `yaml:"masking"`
	Caches  map[string]YAMLCache `yaml:"caches"`
}

type YAMLCache struct {
	Unique bool `yaml:"unique"`
}

type CachedMaskEngineFactories func(model.MaskEngine) model.MaskEngine

func buildCache(caches map[string]YAMLCache) map[string]model.Cache {
	result := map[string]model.Cache{}

	for name, conf := range caches {
		if conf.Unique {
			result[name] = model.NewUniqueMemCache()
		} else {
			result[name] = model.NewMemCache()
		}
	}
	return result
}

// YamlConfig create a MaskConfiguration from a file
func YamlPipeline(pipeline model.Pipeline, filename string, maskFactories []model.MaskFactory, maskContextFactories []model.MaskContextFactory) (model.Pipeline, map[string]model.Cache, error) {
	conf, err := ParseYAML(filename)
	caches := buildCache(conf.Caches)

	if err != nil {
		return nil, nil, err
	}
	for _, v := range conf.Masking {
		nbArg := 0

		if v.Mask.FromCache != "" {
			cache, ok := caches[v.Mask.FromCache]
			if !ok {
				return nil, nil, errors.New("Cache '" + v.Cache + "' not found for '" + v.Selector.Jsonpath + "'")
			}
			pipeline = pipeline.Process(model.NewFromCacheProcess(model.NewPathSelector(v.Selector.Jsonpath), cache))
			nbArg++
		}

		for _, factory := range maskFactories {
			mask, present, err := factory(v, conf.Seed)
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
					case model.UniqueCache:
						mask = model.NewUniqueMaskCacheEngine(typedCache, mask)
					default:
						mask = model.NewMaskCacheEngine(typedCache, mask)
					}
				}
				pipeline = pipeline.Process(model.NewMaskEngineProcess(model.NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}

		for _, factory := range maskContextFactories {
			mask, present, err := factory(v, conf.Seed)
			if err != nil {
				return nil, nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				pipeline = pipeline.Process(model.NewMaskContextEngineProcess(model.NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}
		if nbArg != 1 {
			return pipeline, nil, errors.New("Not the right number of argument for " + v.Selector.Jsonpath + ". There should be 1 and there is " + strconv.Itoa(nbArg))
		}
	}
	return pipeline, caches, nil
}

func ParseYAML(filename string) (YAMLStructure, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return YAMLStructure{}, err
	}
	var conf YAMLStructure
	err = yaml.Unmarshal(source, &conf)
	if err != nil {
		return conf, err
	}
	if conf.Seed == 0 {
		conf.Seed = time.Now().UnixNano()
	}
	return conf, nil
}

func DumpCache(name string, cache model.Cache, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Cache %s not dump : %s", name, err.Error())
	}

	defer file.Close()
	err = model.NewPipeline(cache.Iterate()).AddSink(jsonline.NewSink(file)).Run()
	if err != nil {
		return fmt.Errorf("Cache %s not dump : %s", name, err.Error())
	}

	return nil
}

func LoadCache(name string, cache model.Cache, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	defer file.Close()
	err = model.NewPipeline(jsonline.NewSource(file)).AddSink(model.NewSinkToCache(cache)).Run()
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	return nil
}
