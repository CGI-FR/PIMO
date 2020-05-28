package templatemask

import (
	"bytes"
	"text/template"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is to mask a value thanks to a template
type MaskEngine struct {
	template *template.Template
}

// NewMask create a MaskEngine
func NewMask(text string) (MaskEngine, error) {
	temp, err := template.New("template").Parse(text)
	return MaskEngine{temp}, err
}

// Mask masks a value with a template
func (tmpl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	var output bytes.Buffer
	err := tmpl.template.Execute(&output, context[0])
	return output.String(), err
}

func NewMaskFromConfig(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Template) != 0 {
		mask, err := NewMask(conf.Mask.Template)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
