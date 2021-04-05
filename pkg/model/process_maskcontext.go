package model

import (
	"github.com/cgi-fr/pimo/pkg/selector"
)

func NewMaskContextEngineProcess(selector Selector, mask MaskContextEngine) Processor {
	return &MaskContextEngineProcess{selector, mask}
}

type MaskContextEngineProcess struct {
	selector Selector
	mask     MaskContextEngine
}

func (mcep *MaskContextEngineProcess) Open() error {
	return nil
}

func (mcep *MaskContextEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) (ret error) {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	mcep.selector.ApplyContext(result, func(rootContext, parentContext Dictionary, key string, _ Entry) (Action, Entry) {
		masked, err := mcep.mask.MaskContext(parentContext, key, rootContext, parentContext)
		if err != nil {
			ret = err
			return selector.NOTHING, nil
		}
		value, ok := masked[key]
		if !ok {
			return selector.NOTHING, nil
		}
		return selector.WRITE, value
	})

	if ret == nil {
		out.Collect(result)
	}

	return
}
