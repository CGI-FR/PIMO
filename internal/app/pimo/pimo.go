package pimo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/goccy/go-yaml"
	wr "github.com/mroth/weightedrand"
	regen "github.com/zach-klippenstein/goregen"
)

// Dictionary is a Map with string as key and Entry as value
type Dictionary = map[string]Entry

// Entry is a dictionary value
type Entry interface{}

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry, ...Dictionary) Entry
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

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	function func(Entry) Entry
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry, context ...Dictionary) Entry { return fme.function(e) }

// MaskingEngineFactory return Masking function data without private information
func MaskingEngineFactory(config MaskConfiguration) MaskEngine {
	return FunctionMaskEngine{func(input Entry) Entry {
		output := Dictionary{}
		dico, ok := input.(Dictionary)
		if !ok {
			return input
		}

		for k, v := range dico {
			output[k] = v
		}

		for _, key := range config.Entries() {
			_, present := output[key]
			if present {
				engine, ok := config.GetMaskingEngine(key)
				if ok {
					output[key] = engine.Mask(output[key], output)
				}
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
func (cme CommandMaskEngine) Mask(e Entry, context ...Dictionary) Entry {
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

// NewMaskRandomList create a MaskRandomList with no seed
func NewMaskRandomList(list []Entry) MaskRandomList {
	return NewMaskRandomListSeeded(list, time.Now().UnixNano())
}

// NewMaskRandomListSeeded create a MaskRandomList with a seed
func NewMaskRandomListSeeded(list []Entry, seed int64) MaskRandomList {
	return MaskRandomList{rand.New(rand.NewSource(seed)), list}
}

// Mask choose a mask value randomly
func (mrl MaskRandomList) Mask(e Entry, context ...Dictionary) Entry {
	return mrl.list[mrl.rand.Intn(len(mrl.list))]
}

// MaskHashList is a list of masking value for hash masking
type MaskHashList struct {
	list []Entry
}

// Mask choose a mask value by hash
func (mhl MaskHashList) Mask(e Entry, context ...Dictionary) Entry {
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
func (wml WeightedMaskList) Mask(e Entry, context ...Dictionary) Entry {
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
func (cm ConstMask) Mask(e Entry, context ...Dictionary) Entry {
	return cm.constValue
}

// RandomIntMask is a list of number to mask randomly
type RandomIntMask struct {
	min int
	max int
}

// Mask choose a mask int randomly within boundary
func (rim RandomIntMask) Mask(e Entry, context ...Dictionary) Entry {
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
func (rm RegexMask) Mask(e Entry, context ...Dictionary) Entry {
	out := rm.generator.Generate()
	return out
}

// DateMask is a struct to mask the date
type DateMask struct {
	dateMin time.Time
	dateMax time.Time
}

// NewDateMask create a date mask
func NewDateMask(min, max time.Time) DateMask {
	t1 := min
	t2 := max
	return DateMask{t1, t2}
}

// Mask choose a mask date randomly
func (dateRange DateMask) Mask(e Entry, context ...Dictionary) Entry {
	delta := dateRange.dateMax.Unix() - dateRange.dateMin.Unix()
	sec := time.Unix(rand.Int63n(delta)+dateRange.dateMin.Unix(), 0)
	return sec
}

//IncrementalMask is a struct to create incremental int
type IncrementalMask struct {
	value     *int
	increment int
}

// NewIncrementalMask create an IncrementalMask
func NewIncrementalMask(start, incr int) IncrementalMask {
	value := &start
	return IncrementalMask{value, incr}
}

// Mask masks a value with an incremental int
func (incr IncrementalMask) Mask(e Entry, context ...Dictionary) Entry {
	output := *incr.value
	*incr.value += incr.increment
	return output
}

// ReplacementMask is to mask a value with another field
type ReplacementMask struct {
	field string
}

// Mask masks a value with another field of the json
func (remp ReplacementMask) Mask(e Entry, context ...Dictionary) Entry {
	return context[0][remp.field]
}

// TemplateMask is to mask a value thanks to a template
type TemplateMask struct {
	template *template.Template
}

// NewTemplateMask create an TemplateMask
func NewTemplateMask(text string) TemplateMask {
	temp, _ := template.New("template").Parse(text)
	return TemplateMask{temp}
}

// Mask masks a value with a template
func (tmpl TemplateMask) Mask(e Entry, context ...Dictionary) Entry {
	var output bytes.Buffer
	err := tmpl.template.Execute(&output, context[0])
	if err != nil {
		return "ERROR"
	}
	return output.String()
}

// YAMLStructure of the file
type YAMLStructure struct {
	Version string `yaml:"version"`
	Seed    int64  `yaml:"seed"`
	Masking []struct {
		Selector struct {
			Jsonpath string `yaml:"jsonpath"`
		} `yaml:"selector"`
		Mask struct {
			Constant     Entry   `yaml:"constant"`
			RandomChoice []Entry `yaml:"randomChoice"`
			Command      string  `yaml:"command"`
			RandomInt    struct {
				Min int `yaml:"min"`
				Max int `yaml:"max"`
			} `yaml:"randomInt"`
			WeightedChoice []struct {
				Choice Entry `yaml:"choice"`
				Weight uint  `yaml:"weight"`
			} `yaml:"weightedChoice"`
			Regex    string  `yaml:"regex"`
			Hash     []Entry `yaml:"hash"`
			RandDate struct {
				DateMin time.Time `yaml:"dateMin"`
				DateMax time.Time `yaml:"dateMax"`
			} `yaml:"randDate"`
			Incremental struct {
				Start     int `yaml:"start"`
				Increment int `yaml:"increment"`
			} `yaml:"incremental"`
			Replacement string `yaml:"replacement"`
			Template    string `yaml:"template"`
		} `yaml:"mask"`
	} `yaml:"masking"`
}

// YamlConfig create a MaskConfiguration from a file
func YamlConfig(filename string) (MaskConfiguration, error) {
	source, _ := ioutil.ReadFile(filename)
	var conf YAMLStructure
	err := yaml.Unmarshal(source, &conf)
	if err != nil {
		return NewMaskConfiguration(), err
	}
	if conf.Seed == 0 {
		conf.Seed = time.Now().UnixNano()
	}
	config := NewMaskConfiguration()
	for _, v := range conf.Masking {
		nbArg := 0
		if v.Mask.Constant != nil {
			config = config.WithEntry(v.Selector.Jsonpath, NewConstMask(v.Mask.Constant))
			nbArg++
		}
		if len(v.Mask.RandomChoice) != 0 {
			config = config.WithEntry(v.Selector.Jsonpath, NewMaskRandomListSeeded(v.Mask.RandomChoice, conf.Seed))
			nbArg++
		}
		if len(v.Mask.Command) != 0 {
			config = config.WithEntry(v.Selector.Jsonpath, CommandMaskEngine{v.Mask.Command})
			nbArg++
		}
		if v.Mask.RandomInt.Max != 0 || v.Mask.RandomInt.Min != 0 {
			config = config.WithEntry(v.Selector.Jsonpath, RandomIntMask{v.Mask.RandomInt.Min, v.Mask.RandomInt.Max})
			nbArg++
		}
		if len(v.Mask.WeightedChoice) != 0 {
			var maskWeight []WeightedChoice
			for _, v := range v.Mask.WeightedChoice {
				maskWeight = append(maskWeight, WeightedChoice{v.Choice, v.Weight})
			}
			config = config.WithEntry(v.Selector.Jsonpath, NewWeightedMaskList(maskWeight))
			nbArg++
		}
		if len(v.Mask.Regex) != 0 {
			config = config.WithEntry(v.Selector.Jsonpath, NewRegexMask(v.Mask.Regex))
			nbArg++
		}
		if len(v.Mask.Hash) != 0 {
			var maskHash MaskHashList
			maskHash.list = append(maskHash.list, v.Mask.Hash...)
			config = config.WithEntry(v.Selector.Jsonpath, maskHash)
			nbArg++
		}
		if v.Mask.RandDate.DateMin != v.Mask.RandDate.DateMax {
			config = config.WithEntry(v.Selector.Jsonpath, NewDateMask(v.Mask.RandDate.DateMin, v.Mask.RandDate.DateMax))
			nbArg++
		}
		if v.Mask.Incremental.Increment != 0 {
			config = config.WithEntry(v.Selector.Jsonpath, NewIncrementalMask(v.Mask.Incremental.Start, v.Mask.Incremental.Increment))
			nbArg++
		}
		if len(v.Mask.Replacement) != 0 {
			config.WithEntry(v.Selector.Jsonpath, ReplacementMask{v.Mask.Replacement})
			nbArg++
		}
		if len(v.Mask.Template) != 0 {
			config.WithEntry(v.Selector.Jsonpath, NewTemplateMask(v.Mask.Template))
			nbArg++
		}
		if nbArg == 0 || nbArg >= 2 {
			return NewMaskConfiguration(), errors.New("Not the right number of argument")
		}
	}
	return config, nil
}

//InterfaceToDictionary returns a dictionary from an interface
func InterfaceToDictionary(inter interface{}) Dictionary {
	dic := make(map[string]Entry)
	mapint := inter.(map[string]interface{})
	for k, v := range mapint {
		switch v.(type) {
		case map[string]interface{}:
			dic[k] = InterfaceToDictionary(v)
		default:
			dic[k] = v
		}
	}
	return dic
}

//JSONToDictionary return a dictionary from a jsonline
func JSONToDictionary(jsonline []byte) (Dictionary, error) {
	var inter interface{}
	err := json.Unmarshal(jsonline, &inter)
	if err != nil {
		return nil, err
	}
	dic := InterfaceToDictionary(inter)
	return dic, nil
}

//StopIteratorError is an error for le JSONLineIterator
type StopIteratorError struct{}

func (e StopIteratorError) Error() string {
	return fmt.Sprintf("Iterator couldn't find any value")
}

// JSONLineIterator export line to JSON format.
type JSONLineIterator struct {
	fscanner *bufio.Scanner
}

// NewJSONLineIterator creates a new JSONLineIterator.
func NewJSONLineIterator(file io.Reader) JSONLineIterator {
	return JSONLineIterator{bufio.NewScanner(file)}
}

// Next convert next line to dictionary
func (jli *JSONLineIterator) Next() (Dictionary, error) {
	if !jli.fscanner.Scan() {
		return nil, StopIteratorError{}
	}
	line := jli.fscanner.Bytes()
	return JSONToDictionary(line)
}

//DictionaryToJSON create a jsonline from a dictionary
func DictionaryToJSON(dic Dictionary) ([]byte, error) {
	jsonline, err := json.Marshal(dic)
	if err != nil {
		return nil, err
	}
	jsonline = append(jsonline, "\n"...)
	return jsonline, nil
}
