package randdate

import (
	"math/rand"
	"time"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a struct to mask the date
type MaskEngine struct {
	rand    *rand.Rand
	DateMin time.Time
	DateMax time.Time
}

// NewMask return a MaskEngine from 2 dates
func NewMask(min, max time.Time, seed int64) MaskEngine {
	return MaskEngine{rand.New(rand.NewSource(seed)), min, max}
}

// Mask choose a mask date randomly
func (dateRange MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	delta := dateRange.DateMax.Unix() - dateRange.DateMin.Unix()
	sec := time.Unix(dateRange.rand.Int63n(delta)+dateRange.DateMin.Unix(), 0)
	return sec, nil
}

func NewMaskFromConfig(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if conf.Mask.RandDate.DateMin != conf.Mask.RandDate.DateMax {
		return NewMask(conf.Mask.RandDate.DateMin, conf.Mask.RandDate.DateMax, seed), true, nil
	}
	return nil, false, nil
}
