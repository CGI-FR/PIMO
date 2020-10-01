package randomdecimal

import (
	"math"
	"math/rand"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a type to mask with a random float
type MaskEngine struct {
	rand      *rand.Rand
	precision int
	min       float64
	max       float64
}

// NewMask create a MaskEngine with a seed
func NewMask(min float64, max float64, precision int, seed int64) MaskEngine {
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), precision, min, max}
}

// Mask choose a mask int randomly within boundary
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	// Generate the random number
	r := me.min + me.rand.Float64()*(me.max-me.min)

	// Round the number to the precision
	rounded := math.Round(r*math.Pow(10, float64(me.precision))) / math.Pow(10, float64(me.precision))
	return rounded, nil
}

// RegistryMaskToConfiguration create a mask from a yaml config
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if conf.Mask.RandomDecimal.Precision != 0 {
		return config.WithEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.RandomDecimal.Min, conf.Mask.RandomDecimal.Max, conf.Mask.RandomDecimal.Precision, seed)), true, nil
	}
	return nil, false, nil
}
