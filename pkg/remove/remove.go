package remove

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is an empty mask to link to a field to remove
type MaskEngine struct {
}

// NewMask return an empty mask
func NewMask() MaskEngine {
	return MaskEngine{}
}

// MaskContext remove the field
func (rm MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	delete(context, key)

	return context, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64) (model.MaskContextEngine, bool, error) {
	if conf.Mask.Remove {
		return NewMask(), true, nil
	}
	return nil, false, nil
}
