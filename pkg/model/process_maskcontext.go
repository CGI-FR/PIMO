package model

import (
	"github.com/barkimedes/go-deepcopy"
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
	// fmt.Printf("ProcessDictionaryContext(%v\n\tdictionary = %#v\n\tout = %#v\n)\n", mcep, dictionary, out)
	copy, err := deepcopy.Anything(dictionary)
	if err != nil {
		return err
	}
	result := copy.(Dictionary)
	// defer func() {
	// 	fmt.Printf("return ProcessDictionaryContext(%v)\n", result)
	// }()
	mcep.selector.ApplyContext(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		masked, err := mcep.mask.MaskContext(parentContext, key, rootContext)
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
