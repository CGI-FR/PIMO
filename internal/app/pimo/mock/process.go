package mock

import (
	"sync"

	"github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/model"
)

type BlackHoleSink struct{}

func (s BlackHoleSink) Open() error {
	return nil
}

func (s BlackHoleSink) ProcessDictionary(value model.Entry) error {
	return nil
}

type Processor struct {
	mutex    *sync.Mutex
	source   *model.CallableMapSource
	pipeline model.SinkedPipeline
}

func NewProcessor(maskingFile string, globalSeed *int64) (*Processor, error) {
	pdef, err := model.LoadPipelineDefinitionFromFile(maskingFile)
	if err != nil {
		return nil, err
	}

	if globalSeed != nil {
		pdef.SetSeed(*globalSeed)
	}

	source := model.NewCallableMapSource()

	pipeline, _, err := model.BuildPipeline(model.NewPipeline(source), pdef, nil, nil, "", "")
	if err != nil {
		return nil, err
	}

	return &Processor{
		mutex:    &sync.Mutex{},
		source:   source,
		pipeline: pipeline.AddSink(BlackHoleSink{}),
	}, nil
}

func (p *Processor) Process(dict model.Dictionary) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.source.SetValue(dict)

	err := p.pipeline.Run()

	zeromdc.ClearGlobalFields()

	return err
}
