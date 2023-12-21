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
	"bytes"
	"fmt"
	"hash/fnv"
	tmpl "text/template"
	"time"

	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry, ...Dictionary) (Entry, error)
}

// MaskContextEngine is a masking algorithm for dictionary
type MaskContextEngine interface {
	MaskContext(Dictionary, string, ...Dictionary) (Dictionary, error)
}

// HasCleaner interface provides a function to apply on cleanup
type HasCleaner interface {
	GetCleaner() FunctionMaskContextEngine
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	Function func(Entry, ...Dictionary) (Entry, error)
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	return fme.Function(e, context...)
}

// FunctionMaskContextEngine implements MaskContextEngine with a simple function
type FunctionMaskContextEngine struct {
	Function func(Dictionary, string, ...Dictionary) (Dictionary, error)
}

// MaskContext delegate mask algorithm to the function
func (fme FunctionMaskContextEngine) MaskContext(e Dictionary, key string, context ...Dictionary) (Dictionary, error) {
	return fme.Function(e, key, context...)
}

type MaskFactoryConfiguration struct {
	Masking   Masking
	Seed      int64
	Cache     map[string]Cache
	Functions tmpl.FuncMap
}

type MaskFactory func(MaskFactoryConfiguration) (MaskEngine, bool, error)

type MaskContextFactory func(MaskFactoryConfiguration) (MaskContextEngine, bool, error)

type SelectorType struct {
	Jsonpath string `yaml:"jsonpath" json:"jsonpath" jsonschema_description:"Path of the target value to mask in the JSON input"`
}

type IncrementalType struct {
	Start     int `yaml:"start" json:"start" jsonschema_description:"First value in the sequence"`
	Increment int `yaml:"increment" json:"increment" jsonschema_description:"Increment to add to reach the next value in the sequence"`
}

type RandDateType struct {
	DateMin time.Time `yaml:"dateMin" json:"dateMin" jsonschema_description:"Lower bound of the date range"`
	DateMax time.Time `yaml:"dateMax" json:"dateMax" jsonschema_description:"Higher bound of the date range"`
}

type RandIntType struct {
	Min int `yaml:"min" json:"min" jsonschema_description:"Lower bound of the integer range"`
	Max int `yaml:"max" json:"max" jsonschema_description:"Lower bound of the integer range"`
}

type WeightedChoiceType struct {
	Choice Entry `yaml:"choice" json:"choice" jsonschema_description:"Value for this choice"`
	Weight uint  `yaml:"weight" json:"weight" jsonschema_description:"Weight of this choice, higher weights will be selected more frequently"`
}

type RandomDurationType struct {
	Min string `yaml:"min" json:"min" jsonschema_description:"Lower bound of the duration range (ISO 8601 notation)"`
	Max string `yaml:"max" json:"max" jsonschema_description:"Higher bound of the duration range (ISO 8601 notation)"`
}

type RandomDecimalType struct {
	Min       float64 `yaml:"min" json:"min" jsonschema_description:"Lower bound of the decimal range"`
	Max       float64 `yaml:"max" json:"max" jsonschema_description:"Lower bound of the decimal range"`
	Precision int     `yaml:"precision" json:"precision" jsonschema_description:"Precision of the generated value"`
}

type DateParserType struct {
	InputFormat  string `yaml:"inputFormat,omitempty" json:"inputFormat,omitempty" jsonschema_description:"Format of the input datetime, it should always display the following date : Mon Jan 2 15:04:05 -0700 MST 2006 or the constant value 'unixEpoch'"`
	OutputFormat string `yaml:"outputFormat,omitempty" json:"outputFormat,omitempty" jsonschema_description:"Format of the output datetime, it should always display the following date : Mon Jan 2 15:04:05 -0700 MST 2006 or the constant value 'unixEpoch'"`
}

type FF1Type struct {
	KeyFromEnv string  `yaml:"keyFromEnv" json:"keyFromEnv" jsonschema_description:"Name of the system environment variable that contains the private key"`
	TweakField string  `yaml:"tweakField,omitempty" json:"tweakField,omitempty" jsonschema_description:"Name of the field to use as 'tweak' value : reduce the attack surface by using a varying value on each record, it can be considered as an extension of the secret key that change on each record"`
	Radix      uint    `yaml:"radix,omitempty" json:"radix,omitempty" jsonschema_description:"determine which part of the fixed FF1 domain definition will actually be used, for example 10 will use the first 10 characters of 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	Domain     string  `yaml:"domain,omitempty" json:"domain,omitempty" jsonschema_description:"allowed characters domain that will be encrypted, do not use radix and domain at the same time"`
	Preserve   bool    `yaml:"preserve,omitempty" json:"preserve,omitempty" jsonschema_description:"preserve characters that are not part of the allowed domain"`
	OnError    *string `yaml:"onError,omitempty" json:"onError,omitempty" jsonschema_description:"template to execute if there is an error while encrypting value"`
	Decrypt    bool    `yaml:"decrypt,omitempty" json:"decrypt,omitempty" jsonschema_description:"Decrypt the value instead of encrypt"`
}

type PipeType struct {
	Masking        []Masking `yaml:"masking,omitempty" json:"masking,omitempty" jsonschema_description:"Define a list of selector/mask couple to apply on the JSON stream, in this order"`
	InjectParent   string    `yaml:"injectParent,omitempty" json:"injectParent,omitempty" jsonschema_description:"Used in conjunction with the 'pipe' mask, inject the parent object with the given field name"`
	InjectRoot     string    `yaml:"injectRoot,omitempty" json:"injectRoot,omitempty" jsonschema_description:"Used in conjunction with the 'pipe' mask, inject the root object with the given field name"`
	DefinitionFile string    `yaml:"file,omitempty" json:"file,omitempty" jsonschema_description:"URI to an external resource to read the pipeline definition"`
}

type TemplateEachType struct {
	Item     string `yaml:"item,omitempty" json:"item,omitempty" jsonschema_description:"Inject the current element value under the given field name"`
	Index    string `yaml:"index,omitempty" json:"index,omitempty" jsonschema_description:"Inject the current element index under the given field name"`
	Template string `yaml:"template,omitempty" json:"template,omitempty" jsonschema_description:"Replace the current value with the result of executing this Golang template"`
}

type LuhnType struct {
	Universe string `yaml:"universe,omitempty" json:"universe,omitempty" jsonschema_description:"All possible characters that can be encountered as input value"`
}

type MarkovType struct {
	MaxSize   int    `yaml:"max-size,omitempty" json:"max-size,omitempty" jsonschema_description:"Maximum length for the generated text"`
	Sample    string `yaml:"sample,omitempty" json:"sample,omitempty" jsonschema_description:"URI to an external resource to train the Markiv model"`
	Separator string `yaml:"separator,omitempty" json:"separator,omitempty" jsonschema_description:"Separator to use to read tokens, leave empty to treat each character as a token"`
	Order     int    `yaml:"order,omitempty" json:"order,omitempty" jsonschema_description:"Number of tokens to consider, a higher value = more similar to sample, too high and the generated text will be completly similar to sample"`
}

type Class struct {
	Input  string `yaml:"input" json:"input" jsonschema_description:"Characters to replace in the input value"`
	Output string `yaml:"output" json:"output" jsonschema_description:"Characters to use to generate the output value"`
}

type TranscodeType struct {
	Classes []Class `yaml:"classes,omitempty" json:"classes,omitempty" jsonschema_description:"Each class will define a rule to replace a set of characters by another"`
}

type ChoiceInCSVType struct {
	URI             string `yaml:"uri" json:"uri" jsonschema_description:"URI of the CSV resource"`
	Header          bool   `yaml:"header,omitempty" json:"header,omitempty" jsonschema_description:"Does the CSV resource contains a header line"`
	Separator       string `yaml:"separator,omitempty" json:"separator,omitempty" jsonschema_description:"Separator character"`
	Comment         string `yaml:"comment,omitempty" json:"comment,omitempty" jsonschema_description:"Lines beginning with the comment character without preceding whitespace are ignored"`
	FieldsPerRecord int    `yaml:"fieldsPerRecord,omitempty" json:"fieldsPerRecord,omitempty" jsonschema_description:"FieldsPerRecord is the number of expected fields per record, 0 means the number of fields in the first record"`
	TrimSpace       bool   `yaml:"trim,omitempty" json:"trim,omitempty" jsonschema_description:"If true leading white space in a field is ignored"`
}

type ExactMatchType struct {
	CSV   string `yaml:"csv" json:"csv" jsonschema_description:"Characters exact to match in the csv file"`
	Entry string `yaml:"entry" json:"entry" jsonschema_description:"Characters exact to find in the entry file"`
}

type FindInCSVType struct {
	URI             string         `yaml:"uri" json:"uri" jsonschema_description:"URI of the CSV resource"`
	ExactMatch      ExactMatchType `yaml:"exactMatch,omitempty" json:"exactMatch,omitempty" jsonschema_description:"Compare csv data and entry data, find the exected matched csv line"`
	JaccardMatch    ExactMatchType `yaml:"jaccard,omitempty" json:"jaccard,omitempty" jsonschema_description:"Compare csv data and entry data with jaccard, find the similarity matched csv line"`
	Expected        string         `yaml:"expected" json:"expected" jsonschema_description:"How much result return, 3 modes availables: only-one, at-least-one, many"`
	Header          bool           `yaml:"header,omitempty" json:"header,omitempty" jsonschema_description:"Does the CSV resource contains a header line"`
	Separator       string         `yaml:"separator,omitempty" json:"separator,omitempty" jsonschema_description:"Separator character"`
	Comment         string         `yaml:"comment,omitempty" json:"comment,omitempty" jsonschema_description:"Lines beginning with the comment character without preceding whitespace are ignored"`
	FieldsPerRecord int            `yaml:"fieldsPerRecord,omitempty" json:"fieldsPerRecord,omitempty" jsonschema_description:"FieldsPerRecord is the number of expected fields per record, 0 means the number of fields in the first record"`
	TrimSpace       bool           `yaml:"trim,omitempty" json:"trim,omitempty" jsonschema_description:"If true leading white space in a field is ignored"`
}

type MaskType struct {
	Add               Entry                `yaml:"add,omitempty" json:"add,omitempty" jsonschema:"oneof_required=Add,title=Add Mask,description=Add a new field in the JSON stream"`
	AddTransient      Entry                `yaml:"add-transient,omitempty" json:"add-transient,omitempty" jsonschema:"oneof_required=AddTransient,title=Add Transient Mask" jsonschema_description:"Add a new temporary field, that will not show in the JSON output"`
	Constant          Entry                `yaml:"constant,omitempty" json:"constant,omitempty" jsonschema:"oneof_required=Constant,title=Constant Mask" jsonschema_description:"Replace the input value with a constant field"`
	RandomChoice      []Entry              `yaml:"randomChoice,omitempty" json:"randomChoice,omitempty" jsonschema:"oneof_required=RandomChoice,title=Random Choice Mask" jsonschema_description:"Replace the input value with a value chosen randomly from a constant list"`
	RandomChoiceInURI string               `yaml:"randomChoiceInUri,omitempty" json:"randomChoiceInUri,omitempty" jsonschema:"oneof_required=RandomChoiceInURI,title=Random Choice in URI" jsonschema_description:"Replace the input value with a value chosen randomly from an external resource (1 line = 1 value)"`
	RandomChoiceInCSV ChoiceInCSVType      `yaml:"randomChoiceInCSV,omitempty" json:"randomChoiceInCSV,omitempty" jsonschema:"oneof_required=RandomChoiceInCSV,title=Random Choice in CSV" jsonschema_description:"Replace the input value with a record chosen randomly from an external CSV resource"`
	Command           string               `yaml:"command,omitempty" json:"command,omitempty" jsonschema:"oneof_required=Command,title=Command Mask" jsonschema_description:"Replace the input value with the result of the given system command"`
	RandomInt         RandIntType          `yaml:"randomInt,omitempty" json:"randomInt,omitempty" jsonschema:"oneof_required=RandomInt,title=Random Integer Mask" jsonschema_description:"Replace the input value with a value chosen randomly from an integer range"`
	WeightedChoice    []WeightedChoiceType `yaml:"weightedChoice,omitempty" json:"weightedChoice,omitempty" jsonschema:"oneof_required=WeightedChoice,title=Weighted Choice Mask" jsonschema_description:"Replace the input value with a value chosen randomly from a constant list, each value is given a weight (higher weight value has higher chance to be selected)"`
	Regex             string               `yaml:"regex,omitempty" json:"regex,omitempty" jsonschema:"oneof_required=Regex,title=Regex Mask" jsonschema_description:"Replace the input value with a random generated value that match the given regular expression"`
	Hash              []Entry              `yaml:"hash,omitempty" json:"hash,omitempty" jsonschema:"oneof_required=Hash,title=Hash Mask" jsonschema_description:"Replace the input value with a value chosen deterministically from a constant list, the same input will always be replaced by the same output"`
	HashInURI         string               `yaml:"hashInUri,omitempty" json:"hashInUri,omitempty" jsonschema:"oneof_required=HashInURI,title=Hash in URI Mask" jsonschema_description:"Replace the input value with a value chosen deterministically from an external resource (1 line = 1 value), the same input will always be replaced by the same output"`
	HashInCSV         ChoiceInCSVType      `yaml:"hashInCSV,omitempty" json:"hashInCSV,omitempty" jsonschema:"oneof_required=HashInCSV,title=Hash in CSV" jsonschema_description:"Replace the input value with a record chosen deterministically from an external CSV resource, the same input will always be replaced by the same output"`
	RandDate          RandDateType         `yaml:"randDate,omitempty" json:"randDate,omitempty" jsonschema:"oneof_required=RandDate,title=Random Date Mask" jsonschema_description:"Replace the input value with a value chosen randomly from an date range"`
	Incremental       IncrementalType      `yaml:"incremental,omitempty" json:"incremental,omitempty" jsonschema:"oneof_required=Incremental,title=Incremental Mask" jsonschema_description:"Replace the input value with an incrementing integer sequence"`
	Replacement       string               `yaml:"replacement,omitempty" json:"replacement,omitempty" jsonschema:"oneof_required=Replacement,title=Replacement Mask" jsonschema_description:"Replace the input value with the value of another field"`
	Template          string               `yaml:"template,omitempty" json:"template,omitempty" jsonschema:"oneof_required=Template,title=Template Mask" jsonschema_description:"Replace the input value with the result of executing the given Golang template"`
	TemplateEach      TemplateEachType     `yaml:"template-each,omitempty" json:"template-each,omitempty" jsonschema:"oneof_required=TemplateEach,title=Template Each Mask" jsonschema_description:"Replace all input values (selector must be an array field) with the result of executing the given Golang template on each value"`
	Duration          string               `yaml:"duration,omitempty" json:"duration,omitempty" jsonschema:"oneof_required=Duration,title=Duration Mask" jsonschema_description:"Modify the input value (selector must be a date field) increasing or decreasing by the given amount of time"`
	Remove            bool                 `yaml:"remove,omitempty" json:"remove,omitempty" jsonschema:"oneof_required=Remove,title=Remove Mask" jsonschema_description:"Remove the field from the JSON stream"`
	RangeMask         int                  `yaml:"range,omitempty" json:"range,omitempty" jsonschema:"oneof_required=RangeMask,title=Range Mask" jsonschema_description:"Replace the integer value with a range of the given size"`
	RandomDuration    RandomDurationType   `yaml:"randomDuration,omitempty" json:"randomDuration,omitempty" jsonschema:"oneof_required=RandomDuration,title=Random Duration Mask" jsonschema_description:"Modify the input value (selector must be a date field) increasing or decreasing by a random amount of time"`
	FluxURI           string               `yaml:"fluxUri,omitempty" json:"fluxUri,omitempty" jsonschema:"oneof_required=FluxURI,title=Flux in URI Mask" jsonschema_description:"Replace the input value with the next value in the sequence given by an external resource (1 line = 1 value)"`
	RandomDecimal     RandomDecimalType    `yaml:"randomDecimal,omitempty" json:"randomDecimal,omitempty" jsonschema:"oneof_required=RandomDecimal,title=Random Decimal Mask" jsonschema_description:"Replace the input value with a value chosen randomly from an decimal range"`
	DateParser        DateParserType       `yaml:"dateParser,omitempty" json:"dateParser,omitempty" jsonschema:"oneof_required=DateParser,title=Date Parser Mask" jsonschema_description:"Change the format of the input date"`
	FromCache         string               `yaml:"fromCache,omitempty" json:"fromCache,omitempty" jsonschema:"oneof_required=FromCache,title=From Cache Mask" jsonschema_description:"Replace the input value with the value stored at the corresponding key in the given cache"`
	FF1               FF1Type              `yaml:"ff1,omitempty" json:"ff1,omitempty" jsonschema:"oneof_required=FF1,title=FF1 Mask" jsonschema_description:"Encrypt the input value using the FF1 algorithm (format preserving encryption)"`
	Pipe              PipeType             `yaml:"pipe,omitempty" json:"pipe,omitempty" jsonschema:"oneof_required=Pipe,title=Pipe Mask" jsonschema_description:"If the input value contains an array of object, stream each object with the given masking pipeline definition, this mask exists to handle complex data structures"`
	FromJSON          string               `yaml:"fromjson,omitempty" json:"fromjson,omitempty" jsonschema:"oneof_required=FromJSON,title=From JSON Mask" jsonschema_description:"Parse the input value as raw JSON, and add the resulting structure to the JSON stream"`
	Luhn              *LuhnType            `yaml:"luhn,omitempty" json:"luhn,omitempty" jsonschema:"oneof_required=Luhn,title=Luhn Mask" jsonschema_description:"Concatenate a checksum key to the input value computed with the luhn algorithm"`
	Markov            MarkovType           `yaml:"markov,omitempty" json:"markov,omitempty" jsonschema:"oneof_required=Markov,title=Markov Mask" jsonschema_description:"Produces pseudo text based on sample text"`
	Transcode         *TranscodeType       `yaml:"transcode,omitempty" json:"transcode,omitempty" jsonschema:"oneof_required=Transcode,title=Transcode Mask" jsonschema_description:"Produce a random string by preserving character classes from the original value"`
	FindInCSV         FindInCSVType        `yaml:"findInCSV,omitempty" json:"findInCSV,omitempty" jsonschema:"oneof_required=FindInCSV,title=Find in CSV Mask" jsonschema_description:"find matched values in a CSV file based on input json file and save the matched csv line as type objet"`
}

type Masking struct {
	// Masking requires at least one Selector and one Mask definition.
	// Case1: One selector, One mask
	// Case2: One selector, Multiple masks
	// Case3: Multiple selectors, One mask
	// Case4: Multiple selectors, Multiple masks
	Selector  SelectorType   `yaml:"selector,omitempty" json:"selector,omitempty" jsonschema:"oneof_required=case1,oneof_required=case2" jsonschema_description:"A selector defines on which field the mask will be applied"`
	Selectors []SelectorType `yaml:"selectors,omitempty" json:"selectors,omitempty" jsonschema:"oneof_required=case3,oneof_required=case4" jsonschema_description:"Defines on which fields the mask will be applied"`
	Mask      MaskType       `yaml:"mask,omitempty" json:"mask,omitempty" jsonschema:"oneof_required=case1,oneof_required=case3" jsonschema_description:"Defines how the selected value(s) will be masked"`
	Masks     []MaskType     `yaml:"masks,omitempty" json:"masks,omitempty" jsonschema:"oneof_required=case2,oneof_required=case4" jsonschema_description:"Defines how the selected value(s) will be masked"`
	Cache     string         `yaml:"cache,omitempty" json:"cache,omitempty" jsonschema_description:"Use an in-memory cache to preserve coherence between original/masked values"`
	Preserve  string         `yaml:"preserve,omitempty" json:"preserve,omitempty" jsonschema:"enum=null,enum=empty,enum=blank,enum=notInCache" jsonschema_description:"Preserve (do not mask) some values : null = preserve null value, empty = preserve empty strings, blank = preserve both null and empty values, notInCache = preserve value even if not present in cache (fromCache mask)"`
	Seed      SeedType       `yaml:"seed,omitempty" json:"seed,omitempty" jsonschema_description:"Initialize the Pseaudo-Random-Generator with the value given field"`
}

type SeedType struct {
	Field string `yaml:"field,omitempty" json:"field,omitempty" jsonschema_description:"Initialize the Pseaudo-Random-Generator with the given field value, a Golang Template can be used here"`
}

type CacheDefinition struct {
	Unique  bool `yaml:"unique,omitempty" json:"unique,omitempty" jsonschema_description:"The cache will not allow a masked value to be used multiple times, the mask will be reapplied until a unique value is generated"`
	Reverse bool `yaml:"reverse,omitempty" json:"reverse,omitempty" jsonschema_description:"Reverse the cache, keys will be used as values, and values will be used as keys"`
}
type Function struct {
	Params []Param `yaml:"params" json:"params" jsonschema_description:"Declare parameters function"`
	Body   string  `yaml:"body" json:"body" jsonschema_description:"Declare body function"`
}

type Param struct {
	Name string `yaml:"name" json:"name" jsonschema_description:"Declare name parameters"`
}

type Definition struct {
	Version   string                     `yaml:"version" json:"version" jsonschema_description:"Version of the pipeline definition, use the value 1"`
	Seed      int64                      `yaml:"seed,omitempty" json:"seed,omitempty" jsonschema_description:"Initialize the Pseaudo-Random-Generator with the given value"`
	Functions map[string]Function        `yaml:"functions,omitempty" json:"functions,omitempty" jsonschema_description:"Declare functions to be used in the masking"`
	Masking   []Masking                  `yaml:"masking" json:"masking" jsonschema_description:"Masking pipeline definition"`
	Caches    map[string]CacheDefinition `yaml:"caches,omitempty" json:"caches,omitempty" jsonschema_description:"Declare in-memory caches"`
}

/***************
 * REFACTORING *
 ***************/

// Processor process Dictionary and none, one or many element
type Processor interface {
	Open() error
	ProcessDictionary(Dictionary, Collector) error
}

// Collector collect Dictionary generate by Process
type Collector interface {
	Collect(Entry)
}

// SinkProcess send Dictionary process by Pipeline to an output
type SinkProcess interface {
	Open() error
	ProcessDictionary(Entry) error
}

type Pipeline interface {
	Source() Source
	Process(Processor) Pipeline
	WithSource(Source) Pipeline
	AddSink(SinkProcess) SinkedPipeline
}

type SinkedPipeline interface {
	Run() error
}

// Source is an iterator over Dictionary
type Source interface {
	Open() error
	Next() bool
	Value() Entry
	Err() error
}

type Mapper func(Dictionary) (Dictionary, error)

/******************
 * IMPLEMENTATION *
 ******************/

func NewPipelineFromSlice(dictionaries []Dictionary) Pipeline {
	return SimplePipeline{source: &SourceFromSlice{dictionaries: dictionaries, offset: 0}}
}

func NewSourceFromSlice(dictionaries []Dictionary) Source {
	return &SourceFromSlice{dictionaries: dictionaries, offset: 0}
}

type SourceFromSlice struct {
	dictionaries []Dictionary
	offset       int
}

func (source *SourceFromSlice) Next() bool {
	result := source.offset < len(source.dictionaries)
	source.offset++
	return result
}

func (source *SourceFromSlice) Value() Entry {
	return source.dictionaries[source.offset-1]
}

func (source *SourceFromSlice) Err() error {
	return nil
}

func (source *SourceFromSlice) Open() error {
	source.offset = 0
	return nil
}

func NewRepeaterUntilProcess(source *TempSource, text, mode string, skipLogFile string) (Processor, error) {
	eng, err := template.NewEngine(text, tmpl.FuncMap{}, 0, "")
	var errlogger *MsgLogger
	if len(skipLogFile) > 0 {
		errlogger = NewMsgLogger(skipLogFile)
	}
	return RepeaterUntilProcess{eng, source, mode, errlogger}, err
}

type RepeaterUntilProcess struct {
	tmpl      *template.Engine
	tmp       *TempSource
	mode      string
	errlogger *MsgLogger
}

func (p RepeaterUntilProcess) Open() error {
	return nil
}

func (p RepeaterUntilProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	var output bytes.Buffer
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Cannot execute template, error: %v", r)
		}
	}()
	err = p.tmpl.Execute(&output, dictionary.UnpackAsDict().Untyped())

	if err != nil && skipLineOnError {
		log.Warn().AnErr("error", err).Msg("Line skipped")
		statistics.IncIgnoredLinesCount()
		if p.errlogger != nil {
			if msg, ok := dictionary.GetValue("original"); !ok {
				return nil
			} else if msgstr, ok := msg.(string); !ok {
				return nil
			} else if err := p.errlogger.Log(msgstr); err != nil {
				log.Err(err)
			}
		}
		return nil
	}

	if err == nil {
		b := output.String()
		switch p.mode {
		case "while":
			p.tmp.repeat = b == "true"
			if p.tmp.repeat {
				out.Collect(dictionary)
			}
		case "until":
			p.tmp.repeat = b == "false"
			out.Collect(dictionary)
		default:
			p.tmp.repeat = false
			out.Collect(dictionary)
		}
	}

	return err
}

func NewTempSource(sourceValue Source) *TempSource {
	return &TempSource{repeat: false, source: sourceValue, value: NewPackedDictionary()}
}

type TempSource struct {
	repeat bool
	value  Entry
	source Source
}

func (s *TempSource) Open() error { return s.source.Open() }

func (s *TempSource) Next() bool {
	if s.repeat {
		return true
	}
	if s.source.Next() {
		s.value = s.source.Value()
		return true
	}
	return false
}

func (s *TempSource) Err() error { return s.source.Err() }

func (s *TempSource) Value() Entry {
	return s.value
}

func NewRepeaterProcess(times int) Processor {
	return RepeaterProcess{times}
}

type RepeaterProcess struct {
	times int
}

func (p RepeaterProcess) Open() error {
	return nil
}

func (p RepeaterProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	for i := 0; i < p.times; i++ {
		out.Collect(dictionary.Copy())
	}
	return nil
}

type MapProcess struct {
	mapper Mapper
}

func NewMapProcess(mapper Mapper) Processor {
	return MapProcess{mapper: mapper}
}

func (mp MapProcess) Open() error { return nil }

func (mp MapProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	mappedValue, err := mp.mapper(dictionary)
	if err != nil {
		return err
	}
	out.Collect(mappedValue)
	return nil
}

func NewSinkToSlice(dictionaries *[]Entry) SinkProcess {
	return &SinkToSlice{dictionaries}
}

type SinkToSlice struct {
	dictionaries *[]Entry
}

func (sink *SinkToSlice) Open() error {
	*sink.dictionaries = []Entry{}
	return nil
}

func (sink *SinkToSlice) ProcessDictionary(dictionary Entry) error {
	*sink.dictionaries = append(*sink.dictionaries, dictionary)
	return nil
}

func NewSinkToCache(cache Cache) SinkProcess {
	return &SinkToCache{cache}
}

type SinkToCache struct {
	cache Cache
}

func (sink *SinkToCache) Open() error {
	return nil
}

func (sink *SinkToCache) ProcessDictionary(entry Entry) error {
	dictionary := entry.(Dictionary)
	sink.cache.Put(CleanTypes(dictionary.Get("key")), CleanTypes(dictionary.Get("value")))
	return nil
}

type SimpleSinkedPipeline struct {
	source Source
	sink   SinkProcess
}

type SimplePipeline struct {
	source Source
}

func NewPipeline(source Source) Pipeline {
	return SimplePipeline{source: source}
}

func (pipeline SimplePipeline) Source() Source {
	return pipeline.source
}

func (pipeline SimplePipeline) WithSource(source Source) Pipeline {
	return SimplePipeline{source: source}
}

func (pipeline SimplePipeline) Process(process Processor) Pipeline {
	return NewProcessPipeline(pipeline.source, process)
}

func (pipeline SimplePipeline) AddSink(sink SinkProcess) SinkedPipeline {
	return SimpleSinkedPipeline{pipeline, sink}
}

func (pipeline SimplePipeline) Next() bool {
	return pipeline.source.Next()
}

func (pipeline SimplePipeline) Value() Entry {
	return pipeline.source.Value()
}

func (pipeline SimplePipeline) Err() error {
	return pipeline.source.Err()
}

func (pipeline SimplePipeline) Open() error {
	return pipeline.source.Open()
}

func NewCollector() *QueueCollector {
	return &QueueCollector{[]Entry{}, NewDictionary()}
}

type QueueCollector struct {
	queue []Entry
	value Entry
}

func (c *QueueCollector) Err() error {
	return nil
}

func (c *QueueCollector) Open() error {
	return nil
}

func (c *QueueCollector) Collect(dictionary Entry) {
	c.queue = append(c.queue, dictionary)
}

func (c *QueueCollector) Next() bool {
	if len(c.queue) > 0 {
		c.value = c.queue[0]
		c.queue = c.queue[1:]
		return true
	}
	return false
}

func (c *QueueCollector) Value() Entry {
	return c.value
}

func NewProcessPipeline(source Source, process Processor) Pipeline {
	return &ProcessPipeline{NewCollector(), source, process, nil}
}

type ProcessPipeline struct {
	collector *QueueCollector
	source    Source
	Processor
	err error
}

func (p *ProcessPipeline) Next() bool {
	if p.collector.Next() {
		return true
	}
	for p.source.Next() {
		p.err = p.ProcessDictionary(p.source.Value().(Dictionary), p.collector)
		if p.err != nil {
			return false
		}
		if p.collector.Next() {
			return true
		}
	}
	p.err = p.source.Err()
	return false
}

func (p *ProcessPipeline) Value() Entry {
	return p.collector.Value()
}

func (p *ProcessPipeline) Err() error {
	return p.err
}

func (p *ProcessPipeline) AddSink(sink SinkProcess) SinkedPipeline {
	return SimpleSinkedPipeline{p, sink}
}

func (p *ProcessPipeline) Process(process Processor) Pipeline {
	return NewProcessPipeline(p, process)
}

func (p *ProcessPipeline) Source() Source {
	return p.source
}

func (p *ProcessPipeline) WithSource(source Source) Pipeline {
	if s, ok := p.source.(*ProcessPipeline); ok {
		return &ProcessPipeline{NewCollector(), s.WithSource(source).(Source), p.Processor, nil}
	}
	return &ProcessPipeline{NewCollector(), source, p.Processor, nil}
}

func (pipeline SimpleSinkedPipeline) Run() (err error) {
	log.Trace().Msg("Enter SimpleSinkedPipeline.Run")
	defer log.Trace().Err(err).Msg("Exit SimpleSinkedPipeline.Run")

	err = pipeline.source.Open()
	if err != nil {
		return err
	}
	err = pipeline.sink.Open()
	if err != nil {
		return err
	}

	for pipeline.source.Next() {
		err := pipeline.sink.ProcessDictionary(pipeline.source.Value())
		if err != nil {
			return err
		}
	}
	return pipeline.source.Err()
}

type Seeder func(Dictionary) (int64, bool, error)

func NewSeeder(sourceField string, seed int64) Seeder {
	var seeder Seeder

	if jpath := sourceField; jpath != "" {
		sel := NewPackedPathSelector(jpath)
		hash := fnv.New64a()
		seeder = func(context Dictionary) (int64, bool, error) {
			e, ok := sel.Read(context)
			if !ok {
				return 0, ok, nil
			}
			hash.Reset()
			_, err := hash.Write([]byte(fmt.Sprintf("%v", e)))
			return int64(hash.Sum64()) + seed, true, err
		}
	} else {
		seeder = func(context Dictionary) (int64, bool, error) {
			return seed, false, nil
		}
	}
	return seeder
}

func NewCallableMapSource() *CallableMapSource {
	return &CallableMapSource{}
}

type CallableMapSource struct {
	value     Entry
	nextValue Entry
}

func (source *CallableMapSource) Open() error {
	return nil
}

func (source *CallableMapSource) Next() bool {
	source.value = source.nextValue
	source.nextValue = nil
	return source.value != nil
}

func (source *CallableMapSource) SetValue(value Entry) {
	source.nextValue = value
}

func (source *CallableMapSource) Value() Entry {
	return source.value
}

func (source *CallableMapSource) Err() error {
	return nil
}
