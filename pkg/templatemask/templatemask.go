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
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/rs/zerolog/log"
)

// MaskEngine is to mask a value thanks to a template
type MaskEngine struct {
	template *template.Engine
}

// NewMask create a MaskEngine
func NewMask(text string, funcs tmpl.FuncMap) (MaskEngine, error) {
	temp, err := template.NewEngine(text, funcs)
	return MaskEngine{temp}, err
}

// Mask masks a value with a template
func (tmpl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask template")
	var output bytes.Buffer
	err := tmpl.template.Execute(&output, context[0].Unordered())
	return output.String(), err
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Template) != 0 {
		mask, err := NewMask(conf.Masking.Mask.Template, conf.Functions)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
