package model

import (
	"strings"
	"time"
)

// Dictionary is a Map with string as key and Entry as value
type Dictionary = map[string]Entry

// Entry is a dictionary value
type Entry interface{}

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry, ...Dictionary) (Entry, error)
}

// MaskContextEngine is a masking algorithm for dictionary
type MaskContextEngine interface {
	MaskContext(Dictionary, string) (Dictionary, error)
}

// NewMaskConfiguration build new configuration
func NewMaskConfiguration() MaskConfiguration {
	return MapMaskConfiguration{map[string]MaskContextEngine{}, []string{}}
}

// MaskConfiguration is a configuration to mask dictionaries content
type MaskConfiguration interface {
	GetMaskingEngine(string) (MaskContextEngine, bool)
	WithEntry(string, MaskEngine) MaskConfiguration
	WithContextEntry(string, MaskContextEngine) MaskConfiguration
	AsEngine() MaskEngine
	AsContextEngine() MaskContextEngine
	Entries() []string
}

// MapMaskConfiguration Implements MaskConfiguration with a map
type MapMaskConfiguration struct {
	config     map[string]MaskContextEngine
	components []string
}

// GetMaskingEngine return the MaskEngine configured for that string
func (mmc MapMaskConfiguration) GetMaskingEngine(key string) (MaskContextEngine, bool) {
	engine, test := mmc.config[key]
	return engine, test
}

// WithEntry append engine for entry in the configuration and return modified configuration
func (mmc MapMaskConfiguration) WithEntry(key string, engine MaskEngine) MaskConfiguration {
	return mmc.WithContextEntry(key, ContextWrapper{engine})
}

func (mmc MapMaskConfiguration) WithContextEntry(key string, engine MaskContextEngine) MaskConfiguration {
	mainEntry := strings.SplitN(key, ".", 2)
	if len(mainEntry) == 2 {
		mmc.config[mainEntry[0]] = NewMaskConfiguration().WithContextEntry(mainEntry[1], engine).AsContextEngine()
		mmc.components = append(mmc.components, mainEntry[0])
	} else {
		mmc.config[key] = engine
		mmc.components = append(mmc.components, key)
	}

	return mmc
}

//Entries list every mask in mmc in order
func (mmc MapMaskConfiguration) Entries() []string {
	return mmc.components
}

// AsEngine return engine with configuration
func (mmc MapMaskConfiguration) AsEngine() MaskEngine {
	return MaskingEngineFactory(mmc)
}

// AsEngine return engine with configuration
func (mmc MapMaskConfiguration) AsContextEngine() MaskContextEngine {
	return ContextWrapper{mmc.AsEngine()}
}

// MaskingEngineFactory return Masking function data without private information
func MaskingEngineFactory(config MaskConfiguration) FunctionMaskEngine {
	return FunctionMaskEngine{func(input Entry) (Entry, error) {
		output := Dictionary{}
		dico, ok := input.(Dictionary)
		if !ok {
			return input, nil
		}

		for k, v := range dico {
			output[k] = v
		}
		var errMask error
		errMask = nil
		for _, key := range config.Entries() {
			_, present := output[key]
			if present {
				engine, ok := config.GetMaskingEngine(key)
				if ok {
					var err error
					output, err = engine.MaskContext(output, key)
					if err != nil {
						errMask = err
						delete(output, key)
					}
				}
			}
		}

		return output, errMask
	}}
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	Function func(Entry) (Entry, error)
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	return fme.Function(e)
}

type ContextWrapper struct {
	engine MaskEngine
}

func (cw ContextWrapper) MaskContext(context Dictionary, key string) (Dictionary, error) {
	result := Dictionary{}
	var err error
	for k, v := range context {
		if k == key {
			switch j := context[k].(type) {
			case []interface{}:
				var tab []Entry
				for i := range j {
					var e Entry
					e, erreur := cw.engine.Mask(j[i], context)
					if erreur != nil {
						err = erreur
						tab = append(tab, j[i])
						continue
					}
					tab = append(tab, e)
				}
				result[k] = tab
			case Entry:
				var e Entry
				e, erreur := cw.engine.Mask(context[k], context)
				if erreur != nil {
					err = erreur
					result[k] = v
					continue
				}
				result[k] = e
			}
		} else {
			result[k] = v
		}
	}
	return result, err
}

type SelectorType struct {
	Jsonpath string `yaml:"jsonpath"`
}

type IncrementalType struct {
	Start     int `yaml:"start"`
	Increment int `yaml:"increment"`
}

type RandDateType struct {
	DateMin time.Time `yaml:"dateMin"`
	DateMax time.Time `yaml:"dateMax"`
}

type RandIntType struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

type WeightedChoiceType struct {
	Choice Entry `yaml:"choice"`
	Weight uint  `yaml:"weight"`
}

type MaskType struct {
	Constant          Entry                `yaml:"constant"`
	RandomChoice      []Entry              `yaml:"randomChoice"`
	RandomChoiceInURI string               `yaml:"randomChoiceInUri"`
	Command           string               `yaml:"command"`
	RandomInt         RandIntType          `yaml:"randomInt"`
	WeightedChoice    []WeightedChoiceType `yaml:"weightedChoice"`
	Regex             string               `yaml:"regex"`
	Hash              []Entry              `yaml:"hash"`
	HashInURI         string               `yaml:"hashInUri"`
	RandDate          RandDateType         `yaml:"randDate"`
	Incremental       IncrementalType      `yaml:"incremental"`
	Replacement       string               `yaml:"replacement"`
	Template          string               `yaml:"template"`
	Duration          string               `yaml:"duration"`
	Remove            bool                 `yaml:"remove"`
}
type Masking struct {
	Selector SelectorType `yaml:"selector"`
	Mask     MaskType     `yaml:"mask"`
}
