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
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func BenchmarkPimoRun(b *testing.B) {
	definition := model.Definition{
		Version: "1",
		Functions: map[string]model.Function{
			"addTest": {
				Params: []model.Param{
					{Name: "x"},
					{Name: "y"},
				},
				Body: "return x + y",
			},
		},
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "addTransientFunction"},
				Mask:     model.MaskType{AddTransient: "{{addTest 4 5}}"},
			},
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
	_, err = ctx.Execute(io.Discard)
	if err != nil {
		b.FailNow()
	}
}

func BenchmarkPimoRunLarge(b *testing.B) {
	dict, err := LoadJsonLineFromDocument("large.json")
	assert.NoError(b, err)

	os.Setenv("FF1_ENCRYPTION_KEY", "70NZ2NWAqk9/A21vBPxqlA==")

	definition := model.Definition{
		Version: "1",
		Functions: map[string]model.Function{
			"addTest": {
				Params: []model.Param{
					{Name: "x"},
					{Name: "y"},
				},
				Body: "return x + y",
			},
		},
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "add"},
				Mask:     model.MaskType{Add: "done"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "addTransient"},
				Mask:     model.MaskType{AddTransient: "also added"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "addTransientFunction"},
				Mask:     model.MaskType{AddTransient: "{{addTest 4 5}}"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "state"},
				Mask:     model.MaskType{Command: "pwd"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "number"},
				Mask:     model.MaskType{Constant: "{\"words\":[\"hello\", \"world\"]}"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "updated_at"},
				Mask: model.MaskType{DateParser: model.DateParserType{
					// Mon Jan 2 15:04:05 -0700 MST 2006
					InputFormat:  "2006-01-02T15:04:05Z",
					OutputFormat: "unixEpoch",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "created_at"},
				Mask:     model.MaskType{Duration: "-P2D"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "user.login"},
				Mask: model.MaskType{FF1: model.FF1Type{
					Radix:      36,
					KeyFromEnv: "FF1_ENCRYPTION_KEY",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "head.label"},
				Mask:     model.MaskType{FluxURI: "pimo://nameFR"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "head.label"},
				Mask:     model.MaskType{FluxURI: "pimo://nameFR"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "head.user.login"},
				Mask:     model.MaskType{FromJSON: "number"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "base.sha"},
				Mask: model.MaskType{Hash: []model.Entry{
					"Emerald City",
					"Ruby City",
					"Sapphire City",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "active_lock_reason"},
				Mask: model.MaskType{Incremental: model.IncrementalType{
					Start:     10,
					Increment: 6,
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "merge_commit_sha"},
				Mask: model.MaskType{Luhn: &model.LuhnType{
					Universe: "0123456789abcdef",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "_links.self.href"},
				Mask: model.MaskType{Markov: model.MarkovType{
					MaxSize:   20,
					Separator: "",
					Order:     2,
					Sample:    "pimo://nameFR",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "head.repo.topics"},
				Mask: model.MaskType{RandDate: model.RandDateType{
					DateMin: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
					DateMax: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "created_at"},
				Mask: model.MaskType{RandomDuration: model.RandomDurationType{
					Min: "-P2D",
					Max: "-P20D",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "number"},
				Mask: model.MaskType{RandomDecimal: model.RandomDecimalType{
					Min:       0.0,
					Max:       10.0,
					Precision: 2,
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "id"},
				Mask: model.MaskType{RandomInt: model.RandIntType{
					Min: 0,
					Max: 10,
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "milestone"},
				Mask: model.MaskType{RandomChoice: []model.Entry{
					"Mickael",
					"Mathieu",
					"Marcelle",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "base.ref"},
				Mask:     model.MaskType{RandomChoiceInURI: "pimo://surnameFR"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "base.repo.stargazers_count"},
				Mask:     model.MaskType{RangeMask: 6},
			},
			{
				Selector: model.SelectorType{Jsonpath: "auto_merge"},
				Mask:     model.MaskType{Regex: "0[1-7]( ([0-9]){2}){4}"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "add"},
				Mask:     model.MaskType{Remove: true},
			},
			{
				Selector: model.SelectorType{Jsonpath: "head.repo.owner.events_url"},
				Mask:     model.MaskType{Replacement: "_links.statuses.href"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "body"},
				Mask:     model.MaskType{Template: "ID = {{.id}}"},
			},
			{
				Selector: model.SelectorType{Jsonpath: "head.repo.topics"},
				Mask: model.MaskType{TemplateEach: model.TemplateEachType{
					Item:     "topic",
					Index:    "topic_idx",
					Template: "the topic {{.topic}} is in position {{.topic_idx}}",
				}},
			},
			{
				Selector: model.SelectorType{Jsonpath: "milestone"},
				Mask: model.MaskType{WeightedChoice: []model.WeightedChoiceType{
					{Choice: "Mickael", Weight: 10},
					{Choice: "Mathieu", Weight: 1},
					{Choice: "Marcelle", Weight: 76},
				}},
			},
		},
	}

	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	ctx := pimo.NewContext(definition)
	cfg := pimo.Config{
		Iteration:   b.N,
		SingleInput: &dict,
	}

	if err := ctx.Configure(cfg); err != nil {
		b.FailNow()
	}

	b.ResetTimer()
	if _, err := ctx.Execute(io.Discard); err != nil {
		b.FailNow()
	}
}

func LoadJsonLineFromDocument(filename string) (model.Dictionary, error) {
	jsonBytes, err := os.ReadFile("testdata/" + filename)
	if err != nil {
		return model.NewDictionary(), fmt.Errorf(": %w", err)
	}

	return jsonline.JSONToDictionary(jsonBytes)

	// compactLine := bytes.NewBuffer([]byte{})

	// if err := json.Compact(compactLine, jsonBytes); err != nil {
	// 	return model.NewDictionary(), fmt.Errorf(": %w", err)
	// }

	// return jsonline.JSONToDictionary(compactLine.Bytes())
}
