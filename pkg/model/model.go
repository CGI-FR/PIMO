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

type MaskFactory func(Masking, int64, map[string]Cache) (MaskEngine, bool, error)

type MaskContextFactory func(Masking, int64, map[string]Cache) (MaskContextEngine, bool, error)

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

type RandomDurationType struct {
	Min string
	Max string
}

type RandomDecimalType struct {
	Min       float64
	Max       float64
	Precision int
}

type DateParserType struct {
	InputFormat  string `yaml:"inputFormat,omitempty"`
	OutputFormat string `yaml:"outputFormat,omitempty"`
}

type FF1Type struct {
	KeyFromEnv string `yaml:"keyFromEnv"`
	TweakField string `yaml:"tweakField,omitempty"`
	Radix      uint   `yaml:"radix,omitempty"`
	Decrypt    bool   `yaml:"decrypt,omitempty"`
}

type PipeType struct {
	Masking        []Masking `yaml:"masking,omitempty"`
	InjectParent   string    `yaml:"injectParent,omitempty"`
	InjectRoot     string    `yaml:"injectRoot,omitempty"`
	DefinitionFile string    `yaml:"file,omitempty"`
}

type TemplateEachType struct {
	Item     string `yaml:"item,omitempty"`
	Index    string `yaml:"index,omitempty"`
	Template string `yaml:"template,omitempty"`
}

type LuhnType struct {
	Universe string `yaml:"universe,omitempty"`
}

type MarkovType struct {
	MaxSize   int    `yaml:"max-size,omitempty"`
	Sample    string `yaml:"sample,omitempty"`
	Separator string `yaml:"separator,omitempty"`
	Order     int    `yaml:"order,omitempty"`
}

type Class struct {
	Input  string `yaml:"input"`
	Output string `yaml:"output"`
}

type TranscodeType struct {
	Classes []Class `yaml:"classes,omitempty"`
}

type MaskType struct {
	Add               Entry                `yaml:"add,omitempty" jsonschema:"oneof_required=Add"`
	AddTransient      Entry                `yaml:"add-transient,omitempty" jsonschema:"oneof_required=AddTransient"`
	Constant          Entry                `yaml:"constant,omitempty" jsonschema:"oneof_required=Constant"`
	RandomChoice      []Entry              `yaml:"randomChoice,omitempty" jsonschema:"oneof_required=RandomChoice"`
	RandomChoiceInURI string               `yaml:"randomChoiceInUri,omitempty" jsonschema:"oneof_required=RandomChoiceInURI"`
	Command           string               `yaml:"command,omitempty" jsonschema:"oneof_required=Command"`
	RandomInt         RandIntType          `yaml:"randomInt,omitempty" jsonschema:"oneof_required=RandomInt"`
	WeightedChoice    []WeightedChoiceType `yaml:"weightedChoice,omitempty" jsonschema:"oneof_required=WeightedChoice"`
	Regex             string               `yaml:"regex,omitempty" jsonschema:"oneof_required=Regex"`
	Hash              []Entry              `yaml:"hash,omitempty" jsonschema:"oneof_required=Hash"`
	HashInURI         string               `yaml:"hashInUri,omitempty" jsonschema:"oneof_required=HashInURI"`
	RandDate          RandDateType         `yaml:"randDate,omitempty" jsonschema:"oneof_required=RandDate"`
	Incremental       IncrementalType      `yaml:"incremental,omitempty" jsonschema:"oneof_required=Incremental"`
	Replacement       string               `yaml:"replacement,omitempty" jsonschema:"oneof_required=Replacement"`
	Template          string               `yaml:"template,omitempty" jsonschema:"oneof_required=Template"`
	TemplateEach      TemplateEachType     `yaml:"template-each,omitempty" jsonschema:"oneof_required=TemplateEach"`
	Duration          string               `yaml:"duration,omitempty" jsonschema:"oneof_required=Duration"`
	Remove            bool                 `yaml:"remove,omitempty" jsonschema:"oneof_required=Remove"`
	RangeMask         int                  `yaml:"range,omitempty" jsonschema:"oneof_required=RangeMask"`
	RandomDuration    RandomDurationType   `yaml:"randomDuration,omitempty" jsonschema:"oneof_required=RandomDuration"`
	FluxURI           string               `yaml:"fluxUri,omitempty" jsonschema:"oneof_required=FluxURI"`
	RandomDecimal     RandomDecimalType    `yaml:"randomDecimal,omitempty" jsonschema:"oneof_required=RandomDecimal"`
	DateParser        DateParserType       `yaml:"dateParser,omitempty" jsonschema:"oneof_required=DateParser"`
	FromCache         string               `yaml:"fromCache,omitempty" jsonschema:"oneof_required=FromCache"`
	FF1               FF1Type              `yaml:"ff1,omitempty" jsonschema:"oneof_required=FF1"`
	Pipe              PipeType             `yaml:"pipe,omitempty" jsonschema:"oneof_required=Pipe"`
	FromJSON          string               `yaml:"fromjson,omitempty" jsonschema:"oneof_required=FromJSON"`
	Luhn              *LuhnType            `yaml:"luhn,omitempty" jsonschema:"oneof_required=Luhn"`
	Markov            MarkovType           `yaml:"markov,omitempty" jsonschema:"oneof_required=Markov"`
	Transcode         *TranscodeType       `yaml:"transcode,omitempty" jsonschema:"oneof_required=Transcode"`
}

type Masking struct {
	// Masking requires at least one Selector and one Mask definition.
	// Case1: One selector, One mask
	// Case2: One selector, Multiple masks
	// Case3: Multiple selectors, One mask
	// Case4: Multiple selectors, Multiple masks
	Selector  SelectorType   `yaml:"selector,omitempty" jsonschema:"oneof_required=case1,oneof_required=case2"`
	Selectors []SelectorType `yaml:"selectors,omitempty" jsonschema:"oneof_required=case3,oneof_required=case4"`
	Mask      MaskType       `yaml:"mask,omitempty" jsonschema:"oneof_required=case1,oneof_required=case3"`
	Masks     []MaskType     `yaml:"masks,omitempty" jsonschema:"oneof_required=case2,oneof_required=case4"`
	Cache     string         `yaml:"cache,omitempty"`
	Preserve  string         `yaml:"preserve,omitempty"`
	Seed      SeedType       `yaml:"seed,omitempty"`
}

type SeedType struct {
	Field string `yaml:"field,omitempty"`
}

type CacheDefinition struct {
	Unique  bool `yaml:"unique,omitempty"`
	Reverse bool `yaml:"reverse,omitempty"`
}

type Definition struct {
	Version string                     `yaml:"version"`
	Seed    int64                      `yaml:"seed,omitempty"`
	Masking []Masking                  `yaml:"masking"`
	Caches  map[string]CacheDefinition `yaml:"caches,omitempty"`
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
	Collect(Dictionary)
}

// SinkProcess send Dictionary process by Pipeline to an output
type SinkProcess interface {
	Open() error
	ProcessDictionary(Dictionary) error
}

type Pipeline interface {
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
	Value() Dictionary
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

func (source *SourceFromSlice) Value() Dictionary {
	return source.dictionaries[source.offset-1]
}

func (source *SourceFromSlice) Err() error {
	return nil
}

func (source *SourceFromSlice) Open() error {
	source.offset = 0
	return nil
}

func NewRepeaterUntilProcess(source *TempSource, text, mode string) (Processor, error) {
	eng, err := template.NewEngine(text)

	return RepeaterUntilProcess{eng, source, mode}, err
}

type RepeaterUntilProcess struct {
	tmpl *template.Engine
	tmp  *TempSource
	mode string
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
	err = p.tmpl.Execute(&output, dictionary.Untyped())

	if err != nil && skipLineOnError {
		log.Warn().AnErr("error", err).Msg("Line skipped")
		statistics.IncIgnoredLinesCount()
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

func NewTempSource(sourceValue Source) Source {
	return &TempSource{repeat: false, source: sourceValue, value: NewDictionary()}
}

type TempSource struct {
	repeat bool
	value  Dictionary
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

func (s *TempSource) Value() Dictionary {
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

func NewSinkToSlice(dictionaries *[]Dictionary) SinkProcess {
	return &SinkToSlice{dictionaries}
}

type SinkToSlice struct {
	dictionaries *[]Dictionary
}

func (sink *SinkToSlice) Open() error {
	*sink.dictionaries = []Dictionary{}
	return nil
}

func (sink *SinkToSlice) ProcessDictionary(dictionary Dictionary) error {
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

func (sink *SinkToCache) ProcessDictionary(dictionary Dictionary) error {
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

func (pipeline SimplePipeline) Value() Dictionary {
	return pipeline.source.Value()
}

func (pipeline SimplePipeline) Err() error {
	return pipeline.source.Err()
}

func (pipeline SimplePipeline) Open() error {
	return pipeline.source.Open()
}

func NewCollector() *QueueCollector {
	return &QueueCollector{[]Dictionary{}, NewDictionary()}
}

type QueueCollector struct {
	queue []Dictionary
	value Dictionary
}

func (c *QueueCollector) Err() error {
	return nil
}

func (c *QueueCollector) Open() error {
	return nil
}

func (c *QueueCollector) Collect(dictionary Dictionary) {
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

func (c *QueueCollector) Value() Dictionary {
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
		p.err = p.ProcessDictionary(p.source.Value(), p.collector)
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

func (p *ProcessPipeline) Value() Dictionary {
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

func NewSeeder(conf Masking, seed int64) Seeder {
	var seeder Seeder
	if jpath := conf.Seed.Field; jpath != "" {
		sel := NewPathSelector(jpath)
		h := fnv.New64a()
		seeder = func(context Dictionary) (int64, bool, error) {
			e, ok := sel.Read(context)
			if !ok {
				return 0, ok, nil
			}
			h.Reset()
			_, err := h.Write([]byte(fmt.Sprintf("%v", e)))
			return int64(h.Sum64()), true, err
		}
	} else {
		seeder = func(context Dictionary) (int64, bool, error) {
			return seed, false, nil
		}
	}
	return seeder
}
