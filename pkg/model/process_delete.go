package model

import "github.com/cgi-fr/pimo/pkg/selector"

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
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	dp.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		return selector.DELETE, nil
	})
	out.Collect(result)
	return nil
}

func NewMapProcess(mapper Mapper) Processor {
	return MapProcess{mapper: mapper}
}
