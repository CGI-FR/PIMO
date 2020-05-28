package hash

import (
	"hash/fnv"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is a list of masking value for hash masking
type MaskEngine struct {
	List []model.Entry
}

// Mask choose a mask value by hash
func (hm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(e.(string)))
	return hm.List[int(h.Sum32())%len(hm.List)], err
}

func NewMaskFromConfig(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Hash) != 0 {
		var maskHash MaskEngine
		maskHash.List = append(maskHash.List, conf.Mask.Hash...)
		return maskHash, true, nil
	}
	return nil, false, nil
}
