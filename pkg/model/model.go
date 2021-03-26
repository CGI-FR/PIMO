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
	"log"
	"reflect"
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

type MaskFactory func(Masking, int64) (MaskEngine, bool, error)

type MaskContextFactory func(Masking, int64) (MaskContextEngine, bool, error)

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
}

type Masking struct {
	Selector SelectorType `yaml:"selector"`
	Mask     MaskType     `yaml:"mask"`
	Cache    string       `yaml:"cache"`
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

// Selector acces to a specific data into a dictonary
type Selector interface {
	Read(Dictionary) (Entry, bool)
	Write(Dictionary, Entry) Dictionary
	Delete(Dictionary) Dictionary
	ReadContext(Dictionary) (Dictionary, string, bool)
	WriteContext(Dictionary, Entry) Dictionary
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

func NewPathSelector(path string) Selector {
	mainEntry := strings.SplitN(path, ".", 2)
	if len(mainEntry) == 2 {
		return ComplexePathSelector{mainEntry[0], NewPathSelector(mainEntry[1])}
	} else {
		return NewSimplePathSelector(mainEntry[0])
	}
}

type ComplexePathSelector struct {
	path        string
	subSelector Selector
}

func (s ComplexePathSelector) Read(dictionary Dictionary) (entry Entry, ok bool) {
	entry, ok = dictionary[s.path]
	if !ok {
		return
	}

	switch typedMatch := entry.(type) {
	case []Entry:
		entry = []Entry{}
		for _, subEntry := range typedMatch {
			var subResult Entry
			subResult, ok = s.subSelector.Read(subEntry.(Dictionary))
			if !ok {
				return
			}

			entry = append(entry.([]Entry), subResult.(Entry))
		}
		return
	case Dictionary:
		return s.subSelector.Read(typedMatch)
	default:
		return nil, false
	}
}

func (s ComplexePathSelector) ReadContext(dictionary Dictionary) (Dictionary, string, bool) {
	entry, ok := dictionary[s.path]
	if !ok {
		return dictionary, "", false
	}
	subEntry, ok := entry.(Dictionary)
	if !ok {
		return dictionary, "", false
	}
	return s.subSelector.ReadContext(subEntry)
}

var deep int = 0

func (s ComplexePathSelector) Write(dictionary Dictionary, entry Entry) Dictionary {
	defer func() { deep++ }()
	result := Dictionary{}

	deep++
	log.Println(deep, "Write(", dictionary, ",", entry, "(", reflect.TypeOf(entry), ")")

	for k, v := range dictionary {
		result[k] = v
	}
	log.Println(deep, "result=", result, "(", reflect.TypeOf(result), ")")
	v := reflect.ValueOf(entry)
	log.Println(deep, "entry is", v.Kind())
	switch v.Kind() {
	case reflect.Slice:
		log.Println(deep, "case: entry is a slice")
		subTargetArray, isArray := result[s.path].([]Entry)
		if !isArray {
			result[s.path] = s.subSelector.Write(result[s.path].(Dictionary), entry)
		} else {
			subArray := []Dictionary{}
			for i := 0; i < v.Len(); i++ {
				subArray = append(subArray, s.subSelector.Write(subTargetArray[i].(Dictionary), v.Index(i).Interface()))
			}
			result[s.path] = subArray
		}

	default:
		log.Println(deep, "case: entry is something else")
		result[s.path] = s.subSelector.Write(result[s.path].(Dictionary), entry)
	}

	return result
}

func (s ComplexePathSelector) WriteContext(dictionary Dictionary, entry Entry) Dictionary {
	return dictionary
}

func (s ComplexePathSelector) Delete(dictionary Dictionary) Dictionary {
	result := Dictionary{}

	for k, v := range dictionary {
		if k == s.path {
			result[k] = s.subSelector.Delete(v.(Dictionary))
		}
	}
	return result
}

func NewSimplePathSelector(path string) Selector {
	return SimplePathSelector{path}
}

type SimplePathSelector struct {
	Path string
}

func (s SimplePathSelector) Read(dictionary Dictionary) (entry Entry, ok bool) {
	entry, ok = dictionary[s.Path]
	return
}

func (s SimplePathSelector) ReadContext(dictionary Dictionary) (Dictionary, string, bool) {
	return dictionary, s.Path, true
}

func (s SimplePathSelector) Write(dictionary Dictionary, entry Entry) Dictionary {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	result[s.Path] = entry
	return result
}

func (s SimplePathSelector) WriteContext(dictionary Dictionary, entry Entry) Dictionary {
	return dictionary
}

func (s SimplePathSelector) Delete(dictionary Dictionary) Dictionary {
	result := Dictionary{}

	for k, v := range dictionary {
		if k != s.Path {
			result[k] = v
		}
	}
	return result
}

func NewDeleteMaskEngineProcess(selector Selector) Processor {
	return &DeleteMaskEngineProcess{selector: selector}
}

type DeleteMaskEngineProcess struct {
	selector Selector
}

func (dp *DeleteMaskEngineProcess) Open() (err error) {
	return err
}

func (dp *DeleteMaskEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	out.Collect(dp.selector.Delete(dictionary))
	return nil
}

func NewMaskContextEngineProcess(selector Selector, mask MaskContextEngine) Processor {
	return &MaskContextEngineProcess{selector, mask}
}

type MaskContextEngineProcess struct {
	selector Selector
	mask     MaskContextEngine
}

func (mp *MaskContextEngineProcess) Open() (err error) {
	return err
}

func (mp *MaskContextEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	subDictionary, key, ok := mp.selector.ReadContext(dictionary)
	if !ok {
		out.Collect(dictionary)

		return nil
	}

	masked, err := mp.mask.MaskContext(subDictionary, key, dictionary)
	if err != nil {
		return err
	}

	out.Collect(mp.selector.WriteContext(dictionary, masked))

	return nil
}

func NewMaskEngineProcess(selector Selector, mask MaskEngine) Processor {
	return &MaskEngineProcess{selector, mask}
}

type MaskEngineProcess struct {
	selector Selector
	mask     MaskEngine
}

func (mp *MaskEngineProcess) Open() (err error) {
	return err
}

func (mp *MaskEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	log.Println("ProcessDictionary(", dictionary, ")")
	match, ok := mp.selector.Read(dictionary)
	if !ok {
		log.Println("mp.selector.Read returned no match")
		out.Collect(dictionary)
		return nil
	}

	log.Println("mp.selector.Read returned", match, "(", reflect.TypeOf(match), ")")

	var masked Entry
	var err error

	switch typedMatch := match.(type) {
	case []Entry:
		masked = []Entry{}
		var maskedEntry Entry

		for _, matchEntry := range typedMatch {
			maskedEntry, err = mp.mask.Mask(matchEntry, dictionary)
			if err != nil {
				return err
			}
			masked = append(masked.([]Entry), maskedEntry)
		}

	default:
		masked, err = mp.mask.Mask(typedMatch, dictionary)
		if err != nil {
			return err
		}
	}

	log.Println("masking produced", masked, "(", reflect.TypeOf(masked), ")")
	// log.Println("!!! replacing with correct value [[1 2][3 4]] !!!")
	// masked = [][]Entry{
	// 	{"1", "2"},
	// 	{"3", "4"},
	// }

	out.Collect(mp.selector.Write(dictionary, masked))

	return nil
}

func NewMapProcess(mapper Mapper) Processor {
	return MapProcess{mapper: mapper}
}

type MapProcess struct {
	mapper Mapper
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
	sink.cache.Put(dictionary["key"], dictionary["value"])
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
	return &QueueCollector{[]Dictionary{}, nil}
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
	result := Dictionary{}
	for k, v := range p.collector.Value() {
		result[k] = v
	}
	return result
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

func (pipeline SimpleSinkedPipeline) Run() (err error) {
	err = pipeline.source.Open()
	if err != nil {
		return err
	}
	err = pipeline.sink.Open()
	if err != nil {
		return err
	}

	log.Println("==============================================")
	log.Println("BEGIN PROCESSING OF PIPELINE")
	log.Println("==============================================")

	for pipeline.source.Next() {
		err := pipeline.sink.ProcessDictionary(pipeline.source.Value())
		if err != nil {
			return err
		}
	}
	return pipeline.source.Err()
}

// InterfaceToDictionary returns a model.Dictionary from an interface
func InterfaceToDictionary(inter interface{}) Dictionary {
	dic := make(map[string]Entry)
	mapint := inter.(map[string]interface{})

	for k, v := range mapint {
		switch typedValue := v.(type) {
		case map[string]interface{}:
			dic[k] = InterfaceToDictionary(v)
		case []interface{}:
			tab := []Entry{}
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
