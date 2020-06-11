package weightedchoice

import (
	"math/rand"

	wr "github.com/mroth/weightedrand"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine a list of masking value with weight for random
type MaskEngine struct {
	cs wr.Chooser
}

// NewMask returns a WeightedMaskList from slice of model.Entry and weights
func NewMask(list []model.WeightedChoiceType, seed int64) MaskEngine {
	rand.Seed(seed)
	cs := []wr.Choice{}
	for k := range list {
		cs = append(cs, wr.Choice{Item: list[k].Choice, Weight: list[k].Weight})
	}
	return MaskEngine{wr.NewChooser(cs...)}
}

// Mask choose a mask value randomly
func (wml MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return wml.cs.Pick(), nil
}

// RegistryMaskToConfiguration create a mask from a yaml config
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if len(conf.Mask.WeightedChoice) != 0 {
		var maskWeight []model.WeightedChoiceType
		for _, v := range conf.Mask.WeightedChoice {
			maskWeight = append(maskWeight, model.WeightedChoiceType{Choice: v.Choice, Weight: v.Weight})
		}
		return config.WithEntry(conf.Selector.Jsonpath, NewMask(maskWeight, seed)), true, nil
	}
	return nil, false, nil
}
