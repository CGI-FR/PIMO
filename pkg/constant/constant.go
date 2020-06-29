package constant

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is a value that always mask the same way
type MaskEngine struct {
	constValue model.Entry
}

// NewMask return a MaskEngine from a value
func NewMask(data model.Entry) MaskEngine {
	return MaskEngine{data}
}

// Mask return a Constant from a MaskEngine
func (cm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return cm.constValue, nil
}

// Create a mask from a configuration
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.Constant != nil {
		return config.WithEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.Constant)), true, nil
	}
	return nil, false, nil
}
