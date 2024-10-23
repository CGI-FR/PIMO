package mock

import (
	"sync"

	"github.com/cgi-fr/pimo/pkg/model"
)

type SynchonizedSink struct {
	value model.Entry
}

func (s *SynchonizedSink) Open() error {
	return nil
}

func (s *SynchonizedSink) ProcessDictionary(value model.Entry) error {
	s.value = value
	return nil
}

func (s *SynchonizedSink) Dict() model.Dictionary {
	return s.value.(model.Dictionary)
}

type Processor struct {
	mutex    *sync.Mutex
	source   *model.CallableMapSource
	sink     *SynchonizedSink
	pipeline model.SinkedPipeline
}

func NewProcessor(maskingFile string) (*Processor, error) {
	pdef, err := model.LoadPipelineDefinitionFromFile(maskingFile)
	if err != nil {
		return nil, err
	}

	source := model.NewCallableMapSource()

	pipeline, _, err := model.BuildPipeline(model.NewPipeline(source), pdef, nil, nil, "", "")
	if err != nil {
		return nil, err
	}

	sink := &SynchonizedSink{}

	pipeline.AddSink(sink)

	return &Processor{
		mutex:    &sync.Mutex{},
		source:   source,
		sink:     sink,
		pipeline: pipeline.AddSink(sink),
	}, nil
}

func (p *Processor) Process(dict model.Dictionary) (model.Dictionary, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.source.SetValue(dict)

	err := p.pipeline.Run()

	return p.sink.Dict(), err
}
