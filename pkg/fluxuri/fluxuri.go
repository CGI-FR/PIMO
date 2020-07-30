package fluxuri

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/uri"
)

//MaskEngine is a list of value to mask in order from a list
type MaskEngine struct {
	List    []model.Entry
	LenList int
	Actual  *int
}

// NewMask returns a MaskEngine from a file
func NewMask(name string) (MaskEngine, error) {
	list, err := uri.Read(name)
	if err != nil {
		return MaskEngine{}, err
	}
	length := len(list)
	var actual = 0
	return MaskEngine{list, length, &actual}, nil
}

// MaskContext add the field if not existing or replace the value if existing
func (me MaskEngine) MaskContext(context model.Dictionary, key string) (model.Dictionary, error) {
	if *me.Actual < me.LenList {
		context[key] = me.List[*me.Actual]
		*me.Actual++
	}
	return context, nil
}

// Create a mask from a configuration
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if len(conf.Mask.FluxURI) != 0 {
		mask, err := NewMask(conf.Mask.FluxURI)
		return config.WithContextEntry(conf.Selector.Jsonpath, mask), true, err
	}
	return nil, false, nil
}
