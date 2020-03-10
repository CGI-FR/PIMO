package main

import (
	"hash/fnv"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	wr "github.com/mroth/weightedrand"
	regen "github.com/zach-klippenstein/goregen"
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

// MaskRandomList is a list of masking value and a rand init to mask
type MaskRandomList struct {
	rand *rand.Rand
	list []Entry
}

// NewMaskRandomList create a MaskRandomList
func NewMaskRandomList(list []Entry) MaskRandomList {
	return MaskRandomList{rand.New(rand.NewSource(time.Now().UnixNano())), list}
}

// Mask choose a mask value randomly
func (mrl MaskRandomList) Mask(e Entry) Entry {
	return mrl.list[mrl.rand.Intn(len(mrl.list))]
}

// MaskHashList is a list of masking value for hash masking
type MaskHashList struct {
	list []Entry
}

// Mask choose a mask value by hash
func (mhl MaskHashList) Mask(e Entry) Entry {
	h := fnv.New32a()
	_, err := h.Write([]byte(e.(string)))
	if err != nil {
		return "ERROR"
	}
	return mhl.list[int(h.Sum32())%len(mhl.list)]
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

// NewWeightedMaskList returns a WeightedMaskList from slice of entry and weights
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

//ConstMask is a value that always mask the same way
type ConstMask struct {
	constValue Entry
}

// NewConstMask return a ConstantMask from a value
func NewConstMask(data Entry) ConstMask {
	return ConstMask{data}
}

// Mask return a Constant from a ConstMask
func (cm ConstMask) Mask(e Entry) Entry {
	return cm.constValue
}

// RandomIntMask is a list of number to mask randomly
type RandomIntMask struct {
	min int
	max int
}

// Mask choose a mask int randomly within boundary
func (rim RandomIntMask) Mask(e Entry) Entry {
	return rand.Intn(rim.max+1-rim.min) + rim.min
}

//RegexMask is a value that mask thanks to a regular expression
type RegexMask struct {
	generator regen.Generator
}

//NewRegexMask return a RegexMask from a regexp
func NewRegexMask(exp string) RegexMask {
	generator, _ := regen.NewGenerator(exp, &regen.GeneratorArgs{RngSource: rand.NewSource(time.Now().UnixNano())})
	return RegexMask{generator}
}

//Mask returns a string thanks to a regular expression
func (rm RegexMask) Mask(e Entry) Entry {
	out := rm.generator.Generate()
	return out
}
