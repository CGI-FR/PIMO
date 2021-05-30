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
	"time"
)

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry, ...Dictionary) (Entry, error)
}

// MaskContextEngine is a masking algorithm for dictionary
type MaskContextEngine interface {
	MaskContext(Dictionary, string, ...Dictionary) (Dictionary, error)
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	Function func(Entry, ...Dictionary) (Entry, error)
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	return fme.Function(e, context...)
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
	InputFormat  string `yaml:"inputFormat"`
	OutputFormat string `yaml:"outputFormat"`
}

type FF1Type struct {
	KeyFromEnv string `yaml:"keyFromEnv"`
	TweakField string `yaml:"tweakField"`
	Radix      uint
	Decrypt    bool
}

type PipeType struct {
	Masking        []Masking
	InjectParent   string `yaml:"injectParent"`
	InjectRoot     string `yaml:"injectRoot"`
	DefinitionFile string `yaml:"file"`
}

type MaskType struct {
	Add               Entry                `yaml:"add"`
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
	RangeMask         int                  `yaml:"range"`
	RandomDuration    RandomDurationType   `yaml:"randomDuration"`
	FluxURI           string               `yaml:"fluxUri"`
	RandomDecimal     RandomDecimalType    `yaml:"randomDecimal"`
	DateParser        DateParserType       `yaml:"dateParser"`
	FromCache         string               `yaml:"fromCache"`
	FF1               FF1Type              `yaml:"ff1"`
	Pipe              PipeType             `yaml:"pipe"`
}

type Masking struct {
	Selector SelectorType `yaml:"selector"`
	Mask     MaskType     `yaml:"mask"`
	Cache    string       `yaml:"cache"`
}

type CacheDefinition struct {
	Unique bool `yaml:"unique"`
}

type Definition struct {
	Version string                     `yaml:"version"`
	Seed    int64                      `yaml:"seed"`
	Masking []Masking                  `yaml:"masking"`
	Caches  map[string]CacheDefinition `yaml:"caches"`
}

// Processor process Dictionary and produce none, one or many element
type Processor interface {
	Open() error
	ProcessDictionary(Dictionary, Collector) error
}

// CoProcessor process many Dictionary from source and produce none, one or many element
type CoProcessor interface {
	// Open CoProcessor before process.
	Open(Source) error
	// Next return true if CoProcessor have more output.
	Next() bool
	// Value is the last output of CoProcessor.
	Value() Dictionary
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
	// Process transform each Dictionary from the source with a Processor
	Process(Processor) Pipeline
	// CoProcess delegat the source processing to a co processor
	CoProcess(CoProcessor) Pipeline
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
		out.Collect(dictionary)
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
	sink.cache.Put(dictionary.Get("key"), dictionary.Get("value"))
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

func (pipeline SimplePipeline) CoProcess(coProcess CoProcessor) Pipeline {
	return NewCoProcessPipeline(pipeline.source, coProcess)
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

func (p *ProcessPipeline) Open() (err error) {
	err = p.source.Open()
	if err != nil {
		return
	}
	return p.Processor.Open()
}

func (p *ProcessPipeline) Value() Dictionary {
	return CopyDictionary(p.collector.Value())
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

func (p *ProcessPipeline) CoProcess(process CoProcessor) Pipeline {
	return NewCoProcessPipeline(p, process)
}

func (p *ProcessPipeline) WithSource(source Source) Pipeline {
	if s, ok := p.source.(*ProcessPipeline); ok {
		return &ProcessPipeline{NewCollector(), s.WithSource(source).(Source), p.Processor, nil}
	}
	return &ProcessPipeline{NewCollector(), source, p.Processor, nil}
}

func NewCoProcessPipeline(source Source, coProcess CoProcessor) Pipeline {
	return &CoProcessPipeline{source, coProcess, nil}
}

type CoProcessPipeline struct {
	source    Source
	coProcess CoProcessor
	err       error
}

func (p *CoProcessPipeline) Open() error {
	err := p.coProcess.Open(p.source)
	if err != nil {
		return err
	}

	return nil
}

func (p *CoProcessPipeline) Next() bool {
	return p.coProcess.Next()
}

func (p *CoProcessPipeline) Value() Dictionary {
	return CopyDictionary(p.coProcess.Value())
}

func (p *CoProcessPipeline) Err() error {
	return p.err
}

func (p *CoProcessPipeline) AddSink(sink SinkProcess) SinkedPipeline {
	return SimpleSinkedPipeline{p, sink}
}

func (p *CoProcessPipeline) Process(process Processor) Pipeline {
	return NewProcessPipeline(p, process)
}

func (p *CoProcessPipeline) CoProcess(process CoProcessor) Pipeline {
	return NewCoProcessPipeline(p, process)
}

func (p *CoProcessPipeline) WithSource(source Source) Pipeline {
	if s, ok := p.source.(*ProcessPipeline); ok {
		return &CoProcessPipeline{s.WithSource(source).(Source), p.coProcess, nil}
	}
	return &CoProcessPipeline{source, p.coProcess, nil}
}

func (pipeline SimpleSinkedPipeline) Run() (err error) {
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
