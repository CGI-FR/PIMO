package model

import (
	"github.com/barkimedes/go-deepcopy"
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
	// fmt.Printf("ProcessDictionary(%v\n\tdictionary = %#v\n\tout = %#v\n)\n", mep, dictionary, out)
	copy, err := deepcopy.Anything(dictionary)
	if err != nil {
		return err
	}
	result := copy.(Dictionary)
	// defer func() {
	// 	fmt.Printf("return ProcessDictionary(%v)\n", result)
	// }()
	mep.selector.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		masked, err := mep.mask.Mask(value, rootContext)
		if err != nil {
			ret = err
			return selector.NOTHING, nil
		}
		return selector.WRITE, masked
	})

	if err == nil {
		out.Collect(result)
	}

	return
}
