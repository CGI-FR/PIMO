// Copyright (C) 2022 CGI France
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

package flow_test

import (
	"testing"
	"time"

	"github.com/cgi-fr/pimo/pkg/flow"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestExportTemplate(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Template: "{{.TOTO}}",
				},
			},
		},
	}

	wanted := `flowchart LR
	{{.TOTO}} -->|Template| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportConstant(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Constant: "1",
				},
			},
		},
	}

	wanted := `flowchart LR
	1 -->|Constant| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRegex(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Regex: "[A-Z]oto([a-z]){3}",
				},
			},
		},
	}

	wanted := `flowchart LR
	Regex[[A-Z]oto([a-z]){3}] -->|Regex| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportAdd(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Add: "Toto",
				},
			},
		},
	}

	wanted := `flowchart LR
	Toto -->|Add| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportAddTransient(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					AddTransient: "Toto",
				},
			},
		},
	}

	wanted := `flowchart LR
	Toto -->|AddTransient| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRandomChoice(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RandomChoice: []model.Entry{"Toto", "Tata"},
				},
			},
		},
	}

	wanted := `flowchart LR
	RandomChoice[[Toto,Tata]] -->|RandomChoice| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRandomChoiceInURI(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RandomChoiceInURI: "file://../names.txt",
				},
			},
		},
	}

	wanted := `flowchart LR
	file://../names.txt -->|RandomChoiceInURI| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportCommand(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Command: "cat file.json",
				},
			},
		},
	}

	wanted := `flowchart LR
	cat file.json -->|Command| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRandomInt(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RandomInt: model.RandIntType{
						Min: 10,
						Max: 25,
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	RandomInt[10,25] -->|RandomInt| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportWeightedChoice(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					WeightedChoice: []model.WeightedChoiceType{
						{
							Choice: "toto",
							Weight: 8,
						},
						{
							Choice: "tutu",
							Weight: 1,
						},
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	WeightedChoice[[{Choice: toto, Weight: 8},{Choice: tutu, Weight: 1}]] -->|WeightedChoice| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportHash(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Hash: []model.Entry{"Toto", "Tutu"},
				},
			},
		},
	}

	wanted := `flowchart LR
	Hash[[Toto,Tutu]] -->|Hash| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportHashInURI(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					HashInURI: "file://../names.txt",
				},
			},
		},
	}

	wanted := `flowchart LR
	file://../names.txt -->|HashInURI| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRandDate(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RandDate: model.RandDateType{
						DateMin: time.Date(2022, time.January, 14, 16, 14, 0, 0, time.UTC),
						DateMax: time.Date(2022, time.January, 15, 16, 14, 0, 0, time.UTC),
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	RandDate[DateMin: 2022-01-14 16:14:00 +0000 UTC, DateMax: 2022-01-15 16:14:00 +0000 UTC] -->|RandDate| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportIncremental(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Incremental: model.IncrementalType{
						Start:     10,
						Increment: 10,
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	Incremental[Start: 10, Increment: 10] -->|Incremental| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportReplacement(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Replacement: "name2",
				},
			},
		},
	}

	wanted := `flowchart LR
	name2 -->|Replacement| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportTemplateEach(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					TemplateEach: model.TemplateEachType{
						Item:     "val",
						Index:    "idx",
						Template: "{{title .val}} {{.idx}}",
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	TemplateEach[Item: val, Index: idx, Template: {{title .val}} {{.idx}}] -->|TemplateEach| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportDuration(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Duration: "-P2D",
				},
			},
		},
	}

	wanted := `flowchart LR
	-P2D -->|Duration| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRemove(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Remove: true,
				},
			},
		},
	}

	wanted := `flowchart LR
	name -->|Remove| Trash[(Trash)]
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRangeMask(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RangeMask: 5,
				},
			},
		},
	}

	wanted := `flowchart LR
	5 -->|RangeMask| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRandomDuration(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RandomDuration: model.RandomDurationType{
						Min: "-P2D",
						Max: "-P27D",
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	RandomDuration[Min: -P2D, Max: -P27D] -->|RandomDuration| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportFluxURI(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					FluxURI: "file://../names.txt",
				},
			},
		},
	}

	wanted := `flowchart LR
	file://../names.txt -->|FluxURI| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportRandomDecimal(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					RandomDecimal: model.RandomDecimalType{
						Min:       0,
						Max:       9,
						Precision: 3,
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	RandomDecimal[Min: 0.000E+00, Max: 9.000E+00, Precision: 3] -->|RandomDecimal| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportDateParser(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					DateParser: model.DateParserType{
						InputFormat:  "2006-01-02",
						OutputFormat: "01/02/06",
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	DateParser[InputFormat: 2006-01-02, OutputFormat: 01/02/06] -->|DateParser| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportFromCache(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					FromCache: "surname",
				},
			},
		},
	}

	wanted := `flowchart LR
	surname -->|FromCache| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportFF1(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					FF1: model.FF1Type{
						KeyFromEnv: "FF1_ENCRYPTION_KEY",
						TweakField: "tweak",
						Radix:      62,
						Decrypt:    false,
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	FF1[KeyFromEnv: FF1_ENCRYPTION_KEY, TweakField: tweak, Radix: 62, Decrypt: false] -->|FF1| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportPipe(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Pipe: model.PipeType{
						Masking: []model.Masking{
							{
								Selector: model.SelectorType{Jsonpath: "surname"},
								Mask: model.MaskType{
									Add: "Tutu",
								},
							},
						},
						InjectParent:   "",
						InjectRoot:     "",
						DefinitionFile: "",
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	Pipe[DefinitionFile: , InjectParent: , InjectRoot: ] -->|Pipe| name
	name --> Tutu -->|Add| surname
	
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportFromJSON(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					FromJSON: "nom",
				},
			},
		},
	}

	wanted := `flowchart LR
	nom -->|FromJSON| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportLuhn(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Luhn: &model.LuhnType{
						Universe: "abcdef",
					},
				},
			},
		},
	}

	wanted := `flowchart LR
	abcdef -->|Luhn| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportMasks(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Masks: []model.MaskType{
					{Add: "abcdef"},
					{FromJSON: "nom"},
				},
			},
		},
	}

	wanted := `flowchart LR
	abcdef -->|Add| name
	nom -->|FromJSON| name
	`

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}
