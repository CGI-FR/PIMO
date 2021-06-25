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

package templateeach

import (
	"bytes"
	"strings"
	"text/template"
	"unicode"

	sprig "github.com/Masterminds/sprig/v3"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// MaskEngine is to mask a value thanks to a template
type MaskEngine struct {
	template  *template.Template
	itemName  string
	indexName string
}

// rmAcc removes accents from string
// Function derived from: http://blog.golang.org/normalization
func rmAcc(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// NewMask create a MaskEngine
func NewMask(text string, itemName string, indexName string) (MaskEngine, error) {
	funcMap := template.FuncMap{
		"ToUpper":  strings.ToUpper,
		"ToLower":  strings.ToLower,
		"NoAccent": rmAcc,
	}
	temp, err := template.New("template-each").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(text)
	return MaskEngine{temp, itemName, indexName}, err
}

// Mask masks a value with a template
func (tmpl MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask template-each")
	// var output bytes.Buffer
	// err := tmpl.template.Execute(&output, context[0].Unordered())
	// return output.String(), err

	log.Trace().Interface("input", context).Str("key", key).Interface("contexts", contexts).Msg("Enter MaskContext")

	copy := model.CopyDictionary(context)
	tmplctx := contexts[0].Unordered()

	value, ok := copy.GetValue(key)
	if ok {
		switch typedVal := model.CleanTypes(value).(type) {
		case []model.Entry:
			result := []model.Entry{}

			for idx, item := range typedVal {
				tmplctx[tmpl.itemName] = item
				if len(tmpl.indexName) > 0 {
					tmplctx[tmpl.indexName] = idx + 1
				}

				var output bytes.Buffer
				if err := tmpl.template.Execute(&output, tmplctx); err != nil {
					return copy, err
				}

				result = append(result, output.String())
			}

			copy.Set(key, result)
		default:
			log.Debug().Str("key", key).Msg("Mask template-each - value is not an array, ignored masking")
		}
	} else {
		log.Debug().Str("key", key).Msg("Mask template-each - key is not present in context, ignored masking")
	}

	log.Trace().Interface("output", copy).Str("key", key).Msg("Exit MaskContext")

	return copy, nil
}

// Factory create a mask from a yaml config
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskContextEngine, bool, error) {
	if len(conf.Mask.TemplateEach.Template) != 0 {
		mask, err := NewMask(conf.Mask.TemplateEach.Template, conf.Mask.TemplateEach.Item, conf.Mask.TemplateEach.Index)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
