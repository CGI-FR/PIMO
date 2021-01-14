package constant

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a value that always mask the same way
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

// Factory create a mask from a configuration
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.Constant != nil {
		return NewMask(conf.Mask.Constant), true, nil
	}
	return nil, false, nil
}
