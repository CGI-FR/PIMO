package add

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is a value that always mask the same way
type MaskEngine struct {
	value model.Entry
}

// NewMask return a ConstantMask from a value
func NewMask(value model.Entry) MaskEngine {
	return MaskEngine{value}
}

// MaskContext remove the field
func (am MaskEngine) MaskContext(context model.Dictionary, key string) (model.Dictionary, error) {
	_, present := context[key]
	if !present {
		context[key] = am.value
	}

	return context, nil
}

func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.Add != nil {
		return config.WithContextEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.Add)), true, nil
	}
	return nil, false, nil
}
