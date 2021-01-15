package add

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a value that will be the initialisation of the field when it's created
type MaskEngine struct {
	value model.Entry
}

// NewMask return a MaskEngine from a value
func NewMask(value model.Entry) MaskEngine {
	return MaskEngine{value}
}

// MaskContext add the field
func (am MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	_, present := context[key]
	if !present {
		context[key] = am.value
	}

	return context, nil
}

// Create a mask from a configuration
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.Add != nil {
		return config.WithContextEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.Add)), true, nil
	}
	return nil, false, nil
}
