package model

func NewMaskEngineProcessV1(selector Selector, mask MaskEngine) Processor {
	return &MaskEngineProcess{selector, mask}
}

type MaskEngineProcessV1 struct {
	selector Selector
	mask     MaskEngine
}

func (mp *MaskEngineProcessV1) Open() (err error) {
	return err
}

func (mp *MaskEngineProcessV1) ProcessDictionary(dictionary Dictionary, out Collector) error {
	match, ok := mp.selector.Read(dictionary)
	if !ok {
		out.Collect(dictionary)
		return nil
	}

	var masked Entry
	var err error

	switch typedMatch := match.(type) {
	case []Entry:
		masked, err = deepmask(mp.mask, typedMatch, dictionary)
		if err != nil {
			return err
		}

	default:
		masked, err = mp.mask.Mask(typedMatch, dictionary)
		if err != nil {
			return err
		}
	}

	out.Collect(mp.selector.Write(dictionary, masked))

	return nil
}
