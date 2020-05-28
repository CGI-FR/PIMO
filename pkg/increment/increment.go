package increment

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a struct to create incremental int
type MaskEngine struct {
	Value     *int
	Increment int
}

// NewMask create an IncrementalMask
func NewMask(start, incr int) MaskEngine {
	value := &start
	return MaskEngine{value, incr}
}

// Mask masks a value with an incremental int
func (incr MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	output := *incr.Value
	*incr.Value += incr.Increment
	return output, nil
}

func NewMaskFromConfig(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.Incremental.Increment != 0 {
		return NewMask(conf.Mask.Incremental.Start, conf.Mask.Incremental.Increment), true, nil
	}
	return nil, false, nil
}
