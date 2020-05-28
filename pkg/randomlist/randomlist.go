package randomlist

import (
	"math/rand"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand *rand.Rand
	list []model.Entry
}

// NewMaskSeeded create a MaskRandomList with a seed
func NewMask(list []model.Entry, seed int64) MaskEngine {
	return MaskEngine{rand.New(rand.NewSource(seed)), list}
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return mrl.list[mrl.rand.Intn(len(mrl.list))], nil
}

// NewMaskFromConfig create a mask from a yaml config
func NewMaskFromConfig(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.RandomChoice) != 0 {
		return NewMask(conf.Mask.RandomChoice, seed), true, nil
	}
	return nil, false, nil
}
