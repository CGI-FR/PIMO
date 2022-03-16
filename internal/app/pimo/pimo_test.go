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

package pimo_test

import (
	"io/ioutil"
	"testing"

	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog"
)

func BenchmarkPimoRun(b *testing.B) {
	definition := model.Definition{
		Version: "1",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					HashInURI: "pimo://nameFR",
				},
			},
			{
				Selector: model.SelectorType{Jsonpath: "familyName"},
				Mask: model.MaskType{
					HashInURI: "pimo://surnameFR",
				},
			},
			{
				Selector: model.SelectorType{Jsonpath: "domaine"},
				Masks: []model.MaskType{
					{Add: ""},
					{RandomChoice: []model.Entry{"gmail.com", "msn.com"}},
				},
			},
			{
				Selector: model.SelectorType{Jsonpath: "email"},
				Mask: model.MaskType{
					Template: "{{.name}}.{{.familyName}}@{{.domaine}}",
				},
			},
			{
				Selector: model.SelectorType{Jsonpath: "domaine"},
				Mask: model.MaskType{
					Remove: true,
				},
			},
		},
	}

	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	ctx := pimo.NewContext(definition)

	data := model.NewDictionary().With("name", "John").With("familyName", "Doe").With("email", "john.doe@gmail.com")

	cfg := pimo.Config{
		EmptyInput:  false,
		Iteration:   b.N,
		SingleInput: &data,
	}
	err := ctx.Configure(cfg)
	if err != nil {
		b.FailNow()
	}

	b.ResetTimer()
	_, err = ctx.Execute(ioutil.Discard)
	if err != nil {
		b.FailNow()
	}
}
