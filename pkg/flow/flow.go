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

package flow

import (
	"strconv"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
)

func Export(masking model.Definition) (string, error) {
	maskingDef := masking.Masking
	if len(maskingDef) == 0 {
		return "", nil
	}
	res := `flowchart LR
    `
	for i := 0; i < len(maskingDef); i++ {
		if maskingDef[i].Masks != nil {
			for _, v := range maskingDef[i].Masks {
				res += exportMask(maskingDef[i], v) + "\n    "
			}
			return res, nil
		}
		res += exportMask(maskingDef[i], maskingDef[i].Mask) + "\n    "
	}
	return res, nil
}

func exportMask(masking model.Masking, mask model.MaskType) string {
	if mask.Add != nil {
		return mask.Add.(string) + " -->|Add| " + masking.Selector.Jsonpath
	}
	if mask.AddTransient != nil {
		return mask.AddTransient.(string) + " -->|AddTransient| " + masking.Selector.Jsonpath
	}
	if mask.Constant != nil {
		return mask.Constant.(string) + " -->|Constant| " + masking.Selector.Jsonpath
	}
	if mask.RandomChoice != nil {
		return "RandomChoice[[" + flattenChoices(mask) + "]] -->|RandomChoice| " + masking.Selector.Jsonpath
	}
	if mask.RandomChoiceInURI != "" {
		return mask.RandomChoiceInURI + " -->|RandomChoiceInURI| " + masking.Selector.Jsonpath
	}
	if mask.Command != "" {
		return mask.Command + " -->|Command| " + masking.Selector.Jsonpath
	}
	if mask.RandomInt != (model.RandIntType{}) {
		return "RandomInt[" + strconv.Itoa(mask.RandomInt.Min) + "," + strconv.Itoa(mask.RandomInt.Max) + "] -->|RandomInt| " + masking.Selector.Jsonpath
	}
	if len(mask.WeightedChoice) > 0 {
		return "WeightedChoice[[" + flattenWeightedChoices(mask) + "]] -->|WeightedChoice| " + masking.Selector.Jsonpath
	}
	if mask.Regex != "" {
		return "Regex[" + mask.Regex + "] -->|Regex| " + masking.Selector.Jsonpath
	}
	if mask.Hash != nil {
		return "Hash[[" + flattenHash(mask) + "]] -->|Hash| " + masking.Selector.Jsonpath
	}
	if mask.HashInURI != "" {
		return mask.HashInURI + " -->|HashInURI| " + masking.Selector.Jsonpath
	}
	if mask.RandDate != (model.RandDateType{}) {
		return "RandDate[DateMin: " + mask.RandDate.DateMin.String() + ", DateMax: " + mask.RandDate.DateMax.String() + "] -->|RandDate| " + masking.Selector.Jsonpath
	}
	if mask.Incremental != (model.IncrementalType{}) {
		return "Incremental[Start: " + strconv.Itoa(mask.Incremental.Start) + ", Increment: " + strconv.Itoa(mask.Incremental.Increment) + "] -->|Incremental| " + masking.Selector.Jsonpath
	}
	if mask.Replacement != "" {
		return mask.Replacement + " -->|Replacement| " + masking.Selector.Jsonpath
	}
	if mask.Template != "" {
		return mask.Template + " -->|Template| " + masking.Selector.Jsonpath
	}
	if mask.TemplateEach != (model.TemplateEachType{}) {
		return "TemplateEach[Item: " + mask.TemplateEach.Item + ", Index: " + mask.TemplateEach.Index + ", Template: " + mask.TemplateEach.Template + "] -->|TemplateEach| " + masking.Selector.Jsonpath
	}
	if mask.Duration != "" {
		return mask.Duration + " -->|Duration| " + masking.Selector.Jsonpath
	}
	if mask.Remove {
		return masking.Selector.Jsonpath + " -->|Remove| Trash[(Trash)]"
	}
	if mask.RangeMask != 0 {
		return strconv.Itoa(mask.RangeMask) + " -->|RangeMask| " + masking.Selector.Jsonpath
	}
	if mask.RandomDuration != (model.RandomDurationType{}) {
		return "RandomDuration[Min: " + mask.RandomDuration.Min + ", Max: " + mask.RandomDuration.Max + "] -->|RandomDuration| " + masking.Selector.Jsonpath
	}
	if mask.FluxURI != "" {
		return mask.FluxURI + " -->|FluxURI| " + masking.Selector.Jsonpath
	}
	if mask.RandomDecimal != (model.RandomDecimalType{}) {
		min := strconv.FormatFloat(mask.RandomDecimal.Min, 'E', mask.RandomDecimal.Precision, 64)
		max := strconv.FormatFloat(mask.RandomDecimal.Max, 'E', mask.RandomDecimal.Precision, 64)
		precision := strconv.Itoa(mask.RandomDecimal.Precision)
		return "RandomDecimal[Min: " + min + ", Max: " + max + ", Precision: " + precision + "] -->|RandomDecimal| " + masking.Selector.Jsonpath
	}
	if mask.DateParser != (model.DateParserType{}) {
		return "DateParser[InputFormat: " + mask.DateParser.InputFormat + ", OutputFormat: " + mask.DateParser.OutputFormat + "] -->|DateParser| " + masking.Selector.Jsonpath
	}
	if mask.FromCache != "" {
		return mask.FromCache + " -->|FromCache| " + masking.Selector.Jsonpath
	}
	if mask.FF1 != (model.FF1Type{}) {
		return "FF1[KeyFromEnv: " + mask.FF1.KeyFromEnv + ", TweakField: " + mask.FF1.TweakField + ", Radix: " + strconv.FormatUint(uint64(mask.FF1.Radix), 10) + ", Decrypt: " + strconv.FormatBool(mask.FF1.Decrypt) + "] -->|FF1| " + masking.Selector.Jsonpath
	}
	if mask.Pipe.Masking != nil {
		str := "Pipe[DefinitionFile: " + mask.Pipe.DefinitionFile + ", InjectParent: " + mask.Pipe.InjectParent + ", InjectRoot: " + mask.Pipe.InjectRoot + "] -->|Pipe| " + masking.Selector.Jsonpath + "\n    "
		for _, v := range mask.Pipe.Masking {
			str += masking.Selector.Jsonpath + " --> " + exportMask(v, v.Mask) + "\n    "
		}
		return str
	}
	if mask.FromJSON != "" {
		return mask.FromJSON + " -->|FromJSON| " + masking.Selector.Jsonpath
	}
	if mask.Luhn != (&model.LuhnType{}) {
		return mask.Luhn.Universe + " -->|Luhn| " + masking.Selector.Jsonpath
	}

	return ""
}

func flattenChoices(mask model.MaskType) string {
	choices := make([]string, len(mask.RandomChoice))
	for i, v := range mask.RandomChoice {
		choices[i] = v.(string)
	}
	return strings.Join(choices, ",")
}

func flattenWeightedChoices(mask model.MaskType) string {
	choices := make([]string, len(mask.WeightedChoice))
	for i, v := range mask.WeightedChoice {
		weight := uint64(v.Weight)
		weightStr := strconv.FormatUint(weight, 10)
		choices[i] = "{Choice: " + v.Choice.(string) + ", Weight: " + weightStr + "}"
	}
	return strings.Join(choices, ",")
}

func flattenHash(mask model.MaskType) string {
	choices := make([]string, len(mask.Hash))
	for i, v := range mask.Hash {
		choices[i] = v.(string)
	}
	return strings.Join(choices, ",")
}
