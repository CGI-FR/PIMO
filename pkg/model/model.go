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
	MaskContext(Dictionary, string, ...Dictionary) (Dictionary, error)
}

// NewMaskConfiguration build new configuration
func NewMaskConfiguration() MaskConfiguration {
	return MapMaskConfiguration{[]MaskKey{}}
}

// MaskConfiguration is a configuration to mask dictionaries content
type MaskConfiguration interface {
	GetMaskingEngine(string) (MaskContextEngine, bool)
	WithEntry(string, MaskEngine) MaskConfiguration
	WithContextEntry(string, MaskContextEngine) MaskConfiguration
	AsEngine() MaskEngine
	AsContextEngine() MaskContextEngine
	Entries() []MaskKey
}

// MaskKeyStruct is a structure containing a maskengine and the key which need the mask
type MaskKey struct {
	maskCont MaskContextEngine
	Key      string
}

// MapMaskConfiguration Implements MaskConfiguration with a map
type MapMaskConfiguration struct {
	config []MaskKey
}

// WithEntry append engine for entry in the configuration and return modified configuration
func (mmc MapMaskConfiguration) WithEntry(key string, engine MaskEngine) MaskConfiguration {
	return mmc.WithContextEntry(key, ContextWrapper{engine})
}

func (mmc MapMaskConfiguration) WithContextEntry(key string, engine MaskContextEngine) MaskConfiguration {
	mainEntry := strings.SplitN(key, ".", 2)
	if len(mainEntry) == 2 {
		mmc.config = append(mmc.config, MaskKey{NewMaskConfiguration().WithContextEntry(mainEntry[1], engine).AsContextEngine(), mainEntry[0]})
	} else {
		mmc.config = append(mmc.config, MaskKey{engine, key})
	}

	return mmc
}

// GetMaskingEngine return the MaskEngine configured for that string
func (mmc MapMaskConfiguration) GetMaskingEngine(key string) (MaskContextEngine, bool) {
	for _, k := range mmc.config {
		if k.Key == key {
			return k.maskCont, true
		}
	}
	return nil, false
}

// AsEngine return engine with configuration
func (mmc MapMaskConfiguration) AsEngine() MaskEngine {
	return MaskingEngineFactory(mmc, false)
}

// AsEngine return engine with configuration
func (mmc MapMaskConfiguration) AsContextEngine() MaskContextEngine {
	return ContextWrapper{mmc.AsEngine()}
}

// Entries list every mask in mmc in order
func (mmc MapMaskConfiguration) Entries() []MaskKey {
	return mmc.config
}

// MaskingEngineFactory return Masking function data without private information
func MaskingEngineFactory(config MaskConfiguration, root bool) FunctionMaskEngine {
	return FunctionMaskEngine{func(input Entry, contexts ...Dictionary) (Entry, error) {
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
		for _, maskKey := range config.Entries() {
			engine := maskKey.maskCont
			var err error
			if root {
				output, err = engine.MaskContext(output, maskKey.Key, output)
			} else {
				output, err = engine.MaskContext(output, maskKey.Key, contexts...)
			}
			if err != nil {
				errMask = err
				delete(output, maskKey.Key)
			}
		}
		return output, errMask
	}}
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	Function func(Entry, ...Dictionary) (Entry, error)
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	return fme.Function(e, context...)
}

type ContextWrapper struct {
	engine MaskEngine
}

func NewContextWrapper(engine MaskEngine) MaskContextEngine {
	return ContextWrapper{engine}
}

func (cw ContextWrapper) MaskContext(context Dictionary, key string, contexts ...Dictionary) (Dictionary, error) {
	result := Dictionary{}
	var err error
	for k, v := range context {
		if k == key {
			switch j := context[k].(type) {
			case []Dictionary:
				var tab []Dictionary
				for i := range j {
					var e Entry
					e, erreur := cw.engine.Mask(j[i], contexts...)
					if erreur != nil {
						err = erreur
						tab = append(tab, j[i])
						continue
					}
					tab = append(tab, e.(Dictionary))
				}
				result[k] = tab

			case []Entry:
				var tab []Entry
				for i := range j {
					var e Entry
					e, erreur := cw.engine.Mask(j[i], contexts...)
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
				e, erreur := cw.engine.Mask(context[k], contexts...)
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
}

type Masking struct {
	Selector SelectorType `yaml:"selector"`
	Mask     MaskType     `yaml:"mask"`
	Cache    string       `yaml:"cache"`
}

/***************
 * REFACTORING *
 ***************/

// IProcess process Dictionary and none, one or many element
type IProcess interface {
	Open() error
	ProcessDictionary(Dictionary, ICollector) error
}

// ICollector collect Dictionary generate by IProcess
type ICollector interface {
	Collect(Dictionary)
}

// ISinkProcess send Dictionary process by Pipeline to an output
type ISinkProcess interface {
	Open() error
	ProcessDictionary(Dictionary) error
}

type IPipeline interface {
	Process(IProcess) IPipeline
	AddSink(ISinkProcess) ISinkedPipeline
}

type ISinkedPipeline interface {
	Run() error
}

// ISource is an iterator over Dictionary
type ISource interface {
	Open() error
	Next() bool
	Value() Dictionary
	Err() error
}

// ISelector acces to a specific data into a dictonary
type ISelector interface {
	Read(Dictionary) Entry
	Write(Dictionary, Entry) Dictionary
	Delete(Dictionary) Dictionary
}

type Mapper func(Dictionary) (Dictionary, error)

/******************
 * IMPLEMENTATION *
 ******************/

func NewPipelineFromSlice(dictionaries []Dictionary) IPipeline {
	return Pipeline{source: &SourceFromSlice{dictionaries: dictionaries, offset: 0}}
}

func NewSourceFromSlice(dictionaries []Dictionary) ISource {
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

func NewRepeaterProcess(times int) IProcess {
	return RepeaterProcess{times}
}

type RepeaterProcess struct {
	times int
}

func (p RepeaterProcess) Open() error {
	return nil
}

func (p RepeaterProcess) ProcessDictionary(dictionary Dictionary, out ICollector) error {
	for i := 0; i < p.times; i++ {
		out.Collect(dictionary)
	}
	return nil
}

func NewSimplePathSelector(path string) ISelector {
	return SimplePathSelector{path}
}

type SimplePathSelector struct {
	Path string
}

func (s SimplePathSelector) Read(dictionary Dictionary) Entry {
	return dictionary[s.Path]
}

func (s SimplePathSelector) Write(dictionary Dictionary, entry Entry) Dictionary {
	result := Dictionary{}

	for k, v := range dictionary {
		result[k] = v
	}
	result[s.Path] = entry
	return result
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

func NewDeleteMaskEngineProcess(selector ISelector) IProcess {
	return &DeleteMaskEngineProcess{selector: selector}
}

type DeleteMaskEngineProcess struct {
	selector ISelector
}

func (dp *DeleteMaskEngineProcess) Open() (err error) {
	return err
}

func (dp *DeleteMaskEngineProcess) ProcessDictionary(dictionary Dictionary, out ICollector) error {
	out.Collect(dp.selector.Delete(dictionary))
	return nil
}

func NewMaskEngineProcess(selector ISelector, mask MaskEngine) IProcess {
	return &MaskEngineProcess{selector, mask}
}

type MaskEngineProcess struct {
	selector ISelector
	mask     MaskEngine
}

func (mp *MaskEngineProcess) Open() (err error) {
	return err
}

func (mp *MaskEngineProcess) ProcessDictionary(dictionary Dictionary, out ICollector) error {
	match := mp.selector.Read(dictionary)

	masked, err := mp.mask.Mask(match, dictionary)
	if err != nil {
		return err
	}

	out.Collect(mp.selector.Write(dictionary, masked))

	return nil
}

func NewMapProcess(mapper Mapper) IProcess {
	return MapProcess{mapper: mapper}
}

type MapProcess struct {
	mapper Mapper
}

func (mp MapProcess) Open() error { return nil }

func (mp MapProcess) ProcessDictionary(dictionary Dictionary, out ICollector) error {
	mappedValue, err := mp.mapper(dictionary)
	if err != nil {
		return err
	}
	out.Collect(mappedValue)
	return nil
}

func NewSinkToSlice(dictionaries *[]Dictionary) ISinkProcess {
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

type SinkedPipeline struct {
	source ISource
	sink   ISinkProcess
}

type Pipeline struct {
	source ISource
}

func NewPipepline(source ISource) IPipeline {
	return Pipeline{source: source}
}

func (pipeline Pipeline) Process(process IProcess) IPipeline {
	return NewProcessPipeline(pipeline.source, process)
}

func (pipeline Pipeline) AddSink(sink ISinkProcess) ISinkedPipeline {
	return SinkedPipeline{pipeline, sink}
}

func (pipeline Pipeline) Next() bool {
	return pipeline.source.Next()
}

func (pipeline Pipeline) Value() Dictionary {
	return pipeline.source.Value()
}

func (pipeline Pipeline) Err() error {
	return pipeline.source.Err()
}

func (pipeline Pipeline) Open() error {
	return pipeline.source.Open()
}

func NewCollector() *Collector {
	return &Collector{[]Dictionary{}, nil}
}

type Collector struct {
	queue []Dictionary
	value Dictionary
}

func (c *Collector) Collect(dictionary Dictionary) {
	c.queue = append(c.queue, dictionary)
}

func (c *Collector) Next() bool {
	if len(c.queue) > 0 {
		c.value = c.queue[0]
		c.queue = c.queue[1:]
		return true
	}
	return false
}

func (c *Collector) Value() Dictionary {
	return c.value
}

func NewProcessPipeline(source ISource, process IProcess) IPipeline {
	return ProcessPipeline{NewCollector(), source, process, nil}
}

type ProcessPipeline struct {
	collector *Collector
	source    ISource
	IProcess
	err error
}

func (p ProcessPipeline) Next() bool {
	if p.collector.Next() {
		return true
	}
	for p.source.Next() {
		if p.source.Err() != nil {
			return false
		}
		p.err = p.ProcessDictionary(p.source.Value(), p.collector)
		if p.err != nil {
			return false
		}
		if p.collector.Next() {
			return true
		}
	}
	return false
}

func (p ProcessPipeline) Value() Dictionary {
	return p.collector.Value()
}

func (p ProcessPipeline) Err() error {
	return p.source.Err()
}

func (p ProcessPipeline) AddSink(sink ISinkProcess) ISinkedPipeline {
	return SinkedPipeline{p, sink}
}

func (p ProcessPipeline) Process(process IProcess) IPipeline {
	return NewProcessPipeline(p, process)
}

func (pipeline SinkedPipeline) Run() (err error) {
	err = pipeline.source.Open()
	if err != nil {
		return err
	}
	err = pipeline.sink.Open()
	if err != nil {
		return err
	}

	for pipeline.source.Next() {
		if pipeline.source.Err() != nil {
			return pipeline.source.Err()
		}
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
