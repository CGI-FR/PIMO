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

// NewMaskConfiguration build new configuration
func NewMaskConfiguration() MaskConfiguration {
	return MapMaskConfiguration{map[string]MaskEngine{}, []string{}}
}

// MaskConfiguration is a configuration to mask dictionaries content
type MaskConfiguration interface {
	GetMaskingEngine(string) (MaskEngine, bool)
	WithEntry(string, MaskEngine) MaskConfiguration
	AsEngine() MaskEngine
	Entries() []string
}

// MapMaskConfiguration Implements MaskConfiguration with a map
type MapMaskConfiguration struct {
	config     map[string]MaskEngine
	components []string
}

// GetMaskingEngine return the MaskEngine configured for that string
func (mmc MapMaskConfiguration) GetMaskingEngine(key string) (MaskEngine, bool) {
	engine, test := mmc.config[key]
	return engine, test
}

// WithEntry append engine for entry in the configuration and return modified configuration
func (mmc MapMaskConfiguration) WithEntry(key string, engine MaskEngine) MaskConfiguration {
	mainEntry := strings.SplitN(key, ".", 2)
	if len(mainEntry) == 2 {
		mmc.config[mainEntry[0]] = NewMaskConfiguration().WithEntry(mainEntry[1], engine).AsEngine()
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

// MaskingEngineFactory return Masking function data without private information
func MaskingEngineFactory(config MaskConfiguration) MaskEngine {
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
					masked, err := engine.Mask(output[key], output)
					if err != nil {
						errMask = err
						delete(output, key)
					}
					output[key] = masked
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
	RandDate          RandDateType         `yaml:"randDate"`
	Incremental       IncrementalType      `yaml:"incremental"`
	Replacement       string               `yaml:"replacement"`
	Template          string               `yaml:"template"`
	Duration          string               `yaml:"duration"`
}
type Masking struct {
	Selector SelectorType `yaml:"selector"`
	Mask     MaskType     `yaml:"mask"`
}
