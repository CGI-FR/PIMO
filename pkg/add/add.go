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

package add

import (
	"bytes"
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that will be the initialisation of the field when it's created
type MaskEngine struct {
	value    model.Entry
	template *template.Engine
}

// NewMask return a MaskEngine from a value
func NewMask(value model.Entry) (MaskEngine, error) {
	if tmplstr, ok := value.(string); ok {
		// TODO: on passe tmpl.FuncMap{} ou model.MaskFactoryConfiguration.Functions à NewEgine
		temp, err := template.NewEngine(tmplstr, tmpl.FuncMap{})
		return MaskEngine{value, temp}, err
	}
	return MaskEngine{value, nil}, nil
}

// MaskContext add the field
func (am MaskEngine) MaskContext(context model.Dictionary, key string, contexts ...model.Dictionary) (model.Dictionary, error) {
	log.Info().Msg("Mask add")
	_, present := context.GetValue(key)
	if !present {
		if am.template != nil {
			var output bytes.Buffer
			if err := am.template.Execute(&output, contexts[0].Unordered()); err != nil {
				return context, err
			}
			context.Set(key, output.String())
		} else {
			context.Set(key, am.value)
		}
	}

	return context, nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskContextEngine, bool, error) {
	if conf.Masking.Mask.Add != nil {
		mask, err := NewMask(conf.Masking.Mask.Add)
		return mask, true, err
	}
	return nil, false, nil
}
