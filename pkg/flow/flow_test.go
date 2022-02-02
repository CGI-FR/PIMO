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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Template({{.TOTO}})"| name_1
        !input[(input)] --> TOTO
        TOTO -->|"Template({{.TOTO}})"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Constant(1)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Regex([A-Z]oto([a-z]){3})"| name_1
    end
    name_1 --> !output>Output]
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
    !add[/Add/] --> name
    subgraph name_sg
        name -->|"Add(Toto)"| name_1
    end
    name_1 --> !output>Output]
    `

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportTemplateAdd(t *testing.T) {
	definition := model.Definition{
		Version: "test",
		Masking: []model.Masking{
			{
				Selector: model.SelectorType{Jsonpath: "name"},
				Mask: model.MaskType{
					Add: "{{.surname}}",
				},
			},
		},
	}

	wanted := `flowchart LR
    !add[/Add/] --> name
    subgraph name_sg
        name -->|"Add({{.surname}})"| name_1
        !input[(input)] --> surname
        surname -->|"Add({{.surname}})"| name_1
    end
    name_1 --> !output>Output]
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
    !add[/Add/] --> name
    subgraph name_sg
        name -->|"AddTransient(Toto)"| name_1
    end
    name_1 --> !remove[\Remove\]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RandomChoice(Toto,Tata)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RandomChoiceInURI(file://../names.txt)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Command(cat file.json)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RandomInt(Min: 10, Max: 25)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"WeightedChoice(toto @ 8,tutu @ 1)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Hash(Toto,Tutu)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"HashInURI(file://../names.txt)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RandDate(DateMin: 2022-01-14 16:14:00 +0000 UTC, DateMax: 2022-01-15 16:14:00 +0000 UTC)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Incremental(Start: 10, Increment: 10)"| name_1
    end
    name_1 --> !output>Output]
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
					Replacement: "surname",
				},
			},
		},
	}

	wanted := `flowchart LR
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Replacement(surname)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"TemplateEach(Item: val, Index: idx, Template: {{title .val}} {{.idx}})"| name_1
        !input[(input)] --> idx
        idx -->|"TemplateEach({{title .val}} {{.idx}})"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Duration(-P2D)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    name --> !remove[\Remove\]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RangeMask(5)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RandomDuration(Min: -P2D, Max: -P27D)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"FluxURI(file://../names.txt)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"RandomDecimal(Min: 0.000E+00, Max: 9.000E+00, Precision: 3)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"DateParser(InputFormat: 2006-01-02, OutputFormat: 01/02/06)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"FromCache(surname)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"FF1(KeyFromEnv: FF1_ENCRYPTION_KEY, TweakField: tweak, Radix: 62, Decrypt: false)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Pipe(DefinitionFile: , InjectParent: , InjectRoot: )"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"FromJSON(nom)"| name_1
    end
    name_1 --> !output>Output]
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
    !input[(input)] --> name
    subgraph name_sg
        name -->|"Luhn(abcdef)"| name_1
    end
    name_1 --> !output>Output]
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
    !add[/Add/] --> name
    subgraph name_sg
        name -->|"Add(abcdef)"| name_1
        name_1 -->|"FromJSON(nom)"| name_2
    end
    name_2 --> !output>Output]
    `

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}

func TestExportMasking(t *testing.T) {
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

	wanted := `flowchart LR
    !input[(input)] --> name
    subgraph name_sg
        name -->|"HashInURI(pimo://nameFR)"| name_1
    end
    name_1 --> !output>Output]
    !input[(input)] --> familyName
    subgraph familyName_sg
        familyName -->|"HashInURI(pimo://surnameFR)"| familyName_1
    end
    familyName_1 --> !output>Output]
    !add[/Add/] --> domaine
    subgraph domaine_sg
        domaine -->|"RandomChoice(gmail.com,msn.com)"| domaine_1
    end
    domaine_1 --> !remove[\Remove\]
    !input[(input)] --> email
    subgraph email_sg
        email -->|"Template({{.name}}.{{.familyName}}@{{.domaine}})"| email_1
        name_1 -->|"Template({{.name}}.{{.familyName}}@{{.domaine}})"| email_1
        familyName_1 -->|"Template({{.name}}.{{.familyName}}@{{.domaine}})"| email_1
        domaine_1 -->|"Template({{.name}}.{{.familyName}}@{{.domaine}})"| email_1
    end
    email_1 --> !output>Output]
    `

	exportedResult, err := flow.Export(definition)
	assert.Nil(t, err)

	assert.Equal(t, wanted, exportedResult)
}
