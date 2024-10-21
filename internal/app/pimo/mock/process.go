package mock

import "github.com/cgi-fr/pimo/pkg/model"

type Processor struct {
	process model.SinkProcess
	sink    model.QueueCollector
}

func (p Processor) Process(dict model.Dictionary) (model.Dictionary, error) {
	if err := p.process.ProcessDictionary(dict); err != nil {
		return dict, err
	}

	p.sink.Next()

	return p.sink.Value().(model.Dictionary), p.sink.Err()
}
