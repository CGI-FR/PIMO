package pimo

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
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

func buildCachedMaskEngineFactories(caches map[string]YAMLCache) map[string]CachedMaskEngineFactories {
	result := map[string]CachedMaskEngineFactories{}

	for name, conf := range caches {
		if conf.Unique {
			cache := model.NewMemCache()
			result[name] = func(orignale model.MaskEngine) model.MaskEngine {
				return model.NewUniqueMaskCacheEngine(cache, orignale)
			}
		} else {
			cache := model.NewMemCache()
			result[name] = func(orignale model.MaskEngine) model.MaskEngine {
				return model.NewMaskCacheEngine(cache, orignale)
			}
		}
	}
	return result
}

// YamlConfig create a MaskConfiguration from a file
func YamlPipeline(pipeline model.IPipeline, filename string, maskFactories []model.MaskFactory, maskContextFactories []model.MaskContextFactory) (model.IPipeline, error) {
	conf, err := ParseYAML(filename)
	cacheFactories := buildCachedMaskEngineFactories(conf.Caches)

	if err != nil {
		return nil, err
	}
	for _, v := range conf.Masking {
		nbArg := 0
		for _, factory := range maskFactories {
			mask, present, err := factory(v, conf.Seed)
			if err != nil {
				return nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				if v.Cache != "" {
					cacheFactory, ok := cacheFactories[v.Cache]
					if !ok {
						return nil, errors.New("Cache '" + v.Cache + "' not found for '" + v.Selector.Jsonpath + "'")
					}
					mask = cacheFactory(mask)
				}

				pipeline = pipeline.Process(model.NewMaskEngineProcess(model.NewPathSelector(v.Selector.Jsonpath), mask))
				nbArg++
			}
		}

		for _, factory := range maskContextFactories {
			mask, present, err := factory(v, conf.Seed)
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

// YamlConfig create a MaskConfiguration from a file
func YamlConfig(filename string, maskFactories []model.MaskFactory, maskContextFactories []model.MaskContextFactory) (model.MaskConfiguration, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return model.NewMaskConfiguration(), err
	}
	var conf YAMLStructure
	err = yaml.Unmarshal(source, &conf)
	if err != nil {
		return model.NewMaskConfiguration(), err
	}
	if conf.Seed == 0 {
		conf.Seed = time.Now().UnixNano()
	}
	config := model.NewMaskConfiguration()
	cacheFactories := buildCachedMaskEngineFactories(conf.Caches)

	for _, v := range conf.Masking {
		nbArg := 0
		for _, factory := range maskFactories {
			mask, present, err := factory(v, conf.Seed)

			if err != nil {
				return nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present && v.Cache != "" {
				cacheFactory, ok := cacheFactories[v.Cache]
				if !ok {
					return nil, errors.New("cache " + v.Cache + " not found for " + v.Selector.Jsonpath)
				}

				mask = cacheFactory(mask)
			}

			if present {
				config = config.WithEntry(v.Selector.Jsonpath, mask)

				nbArg++
			}
		}
		for _, factory := range maskContextFactories {
			mask, present, err := factory(v, conf.Seed)

			if err != nil {
				return nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				config = config.WithContextEntry(v.Selector.Jsonpath, mask)

				nbArg++
			}
		}
		if nbArg != 1 {
			return config, errors.New("Not the right number of argument for " + v.Selector.Jsonpath + ". There should be 1 and there is " + strconv.Itoa(nbArg))
		}
	}
	return config, nil
}

// InterfaceToDictionary returns a model.Dictionary from an interface
func InterfaceToDictionary(inter interface{}) model.Dictionary {
	dic := make(map[string]model.Entry)
	mapint := inter.(map[string]interface{})

	for k, v := range mapint {
		switch typedValue := v.(type) {
		case map[string]interface{}:
			dic[k] = InterfaceToDictionary(v)
		case []interface{}:
			tab := []model.Entry{}
			for _, item := range typedValue {
				_, dico := item.(map[string]interface{})

				if dico {
					tab = append(tab, InterfaceToDictionary(item))
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

// JSONToDictionary return a model.Dictionary from a jsonline
func JSONToDictionary(jsonline []byte) (model.Dictionary, error) {
	var inter interface{}
	err := json.Unmarshal(jsonline, &inter)
	if err != nil {
		return nil, err
	}
	dic := InterfaceToDictionary(inter)
	return dic, nil
}

// StopIteratorError is an error for le JSONLineIterator
type StopIteratorError struct{}

func (e StopIteratorError) Error() string {
	return "Iterator couldn't find any value"
}

// JSONLineIterator export line to JSON format.
type JSONLineIterator struct {
	fscanner *bufio.Scanner
}

// NewJSONLineIterator creates a new JSONLineIterator.
func NewJSONLineIterator(file io.Reader) JSONLineIterator {
	return JSONLineIterator{bufio.NewScanner(file)}
}

// Next convert next line to model.Dictionary
func (jli *JSONLineIterator) Next() (model.Dictionary, error) {
	if !jli.fscanner.Scan() {
		return nil, StopIteratorError{}
	}
	line := jli.fscanner.Bytes()
	return JSONToDictionary(line)
}

// DictionaryToJSON create a jsonline from a model.Dictionary
func DictionaryToJSON(dic model.Dictionary) ([]byte, error) {
	jsonline, err := json.Marshal(dic)
	if err != nil {
		return nil, err
	}
	jsonline = append(jsonline, "\n"...)
	return jsonline, nil
}
