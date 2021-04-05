package model

import (
	"github.com/cgi-fr/pimo/pkg/selector"
)

type Action = selector.Action

func NewMaskEngineProcess(selector Selector, mask MaskEngine) Processor {
	return &MaskEngineProcess{selector, mask}
}

type MaskEngineProcess struct {
	selector Selector
	mask     MaskEngine
}

func (mep *MaskEngineProcess) Open() error {
	return nil
}

func (mep *MaskEngineProcess) ProcessDictionary(dictionary Dictionary, out Collector) (ret error) {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	mep.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		masked, err := mep.mask.Mask(value, rootContext, parentContext)
		if err != nil {
			ret = err
			return selector.NOTHING, nil
		}
		return selector.WRITE, masked
	})

	if ret == nil {
		out.Collect(result)
	}

	return
}
