// Copyright (C) 2021 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package templatemask

import (
	"bytes"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig"
	"github.com/cgi-fr/pimo/pkg/model"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	temp, err := template.New("template").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(text)
	return MaskEngine{temp}, err
}

// Mask masks a value with a template
func (tmpl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	var output bytes.Buffer
	err := tmpl.template.Execute(&output, context[0])
	return output.String(), err
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Template) != 0 {
		mask, err := NewMask(conf.Mask.Template)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
