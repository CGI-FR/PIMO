package rangemask

import (
	"strconv"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is a value that always mask by replacing with a scale
type MaskEngine struct {
	rangeScale int
}

// NewMask return a MaskEngine from a value
func NewMask(scale int) MaskEngine {
	return MaskEngine{scale}
}

// Mask return a range from a MaskEngine
func (rm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	scaledValue := int(e.(float64)) / rm.rangeScale * rm.rangeScale
	rangedValue := "[" + strconv.Itoa(scaledValue) + ";" + strconv.Itoa(scaledValue+rm.rangeScale-1) + "]"
	return rangedValue, nil
}

// Create a mask from a configuration
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.RangeMask != 0 {
		return config.WithEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.RangeMask)), true, nil
	}
	return nil, false, nil
}
