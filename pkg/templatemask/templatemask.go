package templatemask

import (
	"bytes"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is to mask a value thanks to a template
type MaskEngine struct {
	template *template.Template
}

// rmAcc removes accents from string
// Function derived from: http://blog.golang.org/normalization
func rmAcc(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// NewMask create a MaskEngine
func NewMask(text string) (MaskEngine, error) {
	funcMap := template.FuncMap{
		"ToUpper":  strings.ToUpper,
		"ToLower":  strings.ToLower,
		"NoAccent": rmAcc,
	}
	temp, err := template.New("template").Funcs(funcMap).Parse(text)
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
