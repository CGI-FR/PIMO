package weightedchoice

import (
	"math/rand"
	"sort"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// WeightedChoice adapted from github.com/mroth/weightedrand

// Choice is a wrapper of item and its weight
type Choice struct {
	Item   interface{}
	Weight uint
}

// A Chooser is to choose a Choice by weight
type Chooser struct {
	data   []Choice
	totals []int
	max    int
	rand   *rand.Rand
}

// NewChooser creates a Chooser from Choices
func NewChooser(seed int64, cs ...Choice) Chooser {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Weight < cs[j].Weight
	})
	totals := make([]int, len(cs))
	runningTotal := 0
	for i, c := range cs {
		runningTotal += int(c.Weight)
		totals[i] = runningTotal
	}
	// nolint: gosec
	return Chooser{data: cs, totals: totals, max: runningTotal, rand: rand.New(rand.NewSource(seed))}
}

// Pick returns a choice from a Chooser
func (chs Chooser) Pick() interface{} {
	r := chs.rand.Intn(chs.max) + 1
	i := sort.SearchInts(chs.totals, r)
	return chs.data[i].Item
}

// MaskEngine a list of masking value with weight for random
type MaskEngine struct {
	cs Chooser
}

// NewMask returns a WeightedMaskList from slice of model.Entry and weights
func NewMask(list []model.WeightedChoiceType, seed int64) MaskEngine {
	rand.Seed(seed)
	cs := []Choice{}
	for k := range list {
		cs = append(cs, Choice{Item: list[k].Choice, Weight: list[k].Weight})
	}
	return MaskEngine{NewChooser(seed, cs...)}
}

// Mask choose a mask value randomly
func (wml MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return wml.cs.Pick(), nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.WeightedChoice) != 0 {
		var maskWeight []model.WeightedChoiceType
		for _, v := range conf.Mask.WeightedChoice {
			maskWeight = append(maskWeight, model.WeightedChoiceType{Choice: v.Choice, Weight: v.Weight})
		}
		return NewMask(maskWeight, seed), true, nil
	}
	return nil, false, nil
}
