package randomlist

import (
	"fmt"
	"math/rand"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/uri"
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
	if len(conf.Mask.RandomChoice) != 0 && len(conf.Mask.RandomChoiceInURI) != 0 {
		return nil, false, fmt.Errorf("2 different random choices")
	}
	if len(conf.Mask.RandomChoice) != 0 {
		return NewMask(conf.Mask.RandomChoice, seed), true, nil
	}
	if len(conf.Mask.RandomChoiceInURI) != 0 {
		list, err := uri.Read(conf.Mask.RandomChoiceInURI)
		return NewMask(list, seed), true, err
	}
	return nil, false, nil
}
