package model

func NewMaskContextEngineProcessV1(selector Selector, mask MaskContextEngine) Processor {
	return &MaskContextEngineProcess{selector, mask}
}

type MaskContextEngineProcessV1 struct {
	selector Selector
	mask     MaskContextEngine
}

func (mp *MaskContextEngineProcessV1) Open() (err error) {
	return err
}

func (mp *MaskContextEngineProcessV1) ProcessDictionary(dictionary Dictionary, out Collector) error {
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
