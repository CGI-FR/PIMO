package remove

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is a value that always mask the same way
type MaskEngine struct {
}

// NewMask return a ConstantMask from a value
func NewMask() MaskEngine {
	return MaskEngine{}
}

// MaskContext remove the field
func (rm MaskEngine) MaskContext(context model.Dictionary, key string) (model.Dictionary, error) {
	delete(context, key)

	return context, nil
}

func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.Remove {
		return config.WithContextEntry(conf.Selector.Jsonpath, NewMask()), true, nil
	}
	return nil, false, nil
}
