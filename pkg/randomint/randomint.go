package randomint

import (
	"math/rand"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a list of number to mask randomly
type MaskEngine struct {
	rand *rand.Rand
	min  int
	max  int
}

// NewMask create a MaskEngine with a seed
func NewMask(min int, max int, seed int64) MaskEngine {
	return MaskEngine{rand.New(rand.NewSource(seed)), min, max}
}

// Mask choose a mask int randomly within boundary
func (rim MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return rim.rand.Intn(rim.max+1-rim.min) + rim.min, nil
}

// RegistryMaskToConfiguration create a mask from a yaml config
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.RandomInt.Min != 0 || conf.Mask.RandomInt.Max != 0 {
		return config.WithEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.RandomInt.Min, conf.Mask.RandomInt.Max, seed)), true, nil
	}
	return nil, false, nil
}
