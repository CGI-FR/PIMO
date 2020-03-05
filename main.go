package main

import (
	"math/rand"
	"os/exec"
	"strings"
	"time"

	wr "github.com/mroth/weightedrand"
)

// Dictionary is a Map with string as key and Entry as value
type Dictionary = map[string]Entry

// Entry is a dictionary value
type Entry interface{}

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry) Entry
}

// NewMaskConfiguration build new configuration
func NewMaskConfiguration() MaskConfiguration {
	return MapMaskConfiguration{map[string]MaskEngine{}}
}

// MaskConfiguration is a configuration to mask dictionaries content
type MaskConfiguration interface {
	GetMaskingEngine(string) (MaskEngine, bool)
	WithEntry(string, MaskEngine) MaskConfiguration
	AsEngine() MaskEngine
}

// MapMaskConfiguration Implements MaskConfiguration with a map
type MapMaskConfiguration struct {
	config map[string]MaskEngine
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
	} else {
		mmc.config[key] = engine
	}

	return mmc
}

// AsEngine return engine with configuration
func (mmc MapMaskConfiguration) AsEngine() MaskEngine {
	return MaskingEngineFactory(mmc)
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	function func(Entry) Entry
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry) Entry { return fme.function(e) }

// MaskingEngineFactory return Masking function data without private information
func MaskingEngineFactory(config MaskConfiguration) MaskEngine {
	return FunctionMaskEngine{func(input Entry) Entry {
		output := Dictionary{}

		dico, ok := input.(Dictionary)
		if !ok {
			return input
		}

		for k, v := range dico {
			engine, ok := config.GetMaskingEngine(k)
			if ok {
				output[k] = engine.Mask(v)
			} else {
				output[k] = v
			}
		}
		return output
	}}
}

// CommandMaskEngine implements MaskEngine with a console command
type CommandMaskEngine struct {
	cmd string
}

// Mask delegate mask algorithm to an external program
func (cme CommandMaskEngine) Mask(e Entry) Entry {
	splitCommand := strings.Split(cme.cmd, " ")
	/* #nosec */
	out, err := exec.Command(splitCommand[0], splitCommand[1:]...).Output()

	resulting := strings.Trim(string(out), "\n")
	if err != nil {
		return "ERROR"
	}
	return resulting
}

// MaskList is a list of masking value
type MaskList struct {
	list []string
}

// Mask choose a mask value randomly
func (ml MaskList) Mask(e Entry) Entry {
	rand.Seed(time.Now().UnixNano())
	return ml.list[rand.Intn(len(ml.list))]
}

// WeightedChoice is a tuple of choice and weight for WeightedMaskList
type WeightedChoice struct {
	data   Entry
	weight uint
}

// WeightedMaskList is a list of masking value with weight for random
type WeightedMaskList struct {
	cs wr.Chooser
}

// NewWeightedMaskList return a WeightedMaskList from slice of entry and weights
func NewWeightedMaskList(list []WeightedChoice) WeightedMaskList {
	var cs []wr.Choice
	for k := range list {
		cs = append(cs, wr.Choice{Item: list[k].data, Weight: list[k].weight})
	}
	return WeightedMaskList{wr.NewChooser(cs...)}
}

// Mask choose a mask value randomly with weight for choice
func (wml WeightedMaskList) Mask(e Entry) Entry {
	return wml.cs.Pick()
}
