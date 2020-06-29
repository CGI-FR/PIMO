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

// RegistryMaskToConfiguration create a mask from a yaml config
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if len(conf.Mask.Template) != 0 {
		mask, err := NewMask(conf.Mask.Template)
		if err != nil {
			return nil, false, err
		}
		return config.WithEntry(conf.Selector.Jsonpath, mask), true, nil
	}
	return nil, false, nil
}
