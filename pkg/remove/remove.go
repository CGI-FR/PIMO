package remove

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is an empty mask to link to a field to remove
type MaskEngine struct {
}

// NewMask return an empty mask
func NewMask() MaskEngine {
	return MaskEngine{}
}

// MaskContext remove the field
func (rm MaskEngine) MaskContext(context model.Dictionary, key string) (model.Dictionary, error) {
	delete(context, key)

	return context, nil
}

// RegistryMaskToConfiguration create a mask from a yaml config
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.Remove {
		return config.WithContextEntry(conf.Selector.Jsonpath, NewMask()), true, nil
	}
	return nil, false, nil
}
