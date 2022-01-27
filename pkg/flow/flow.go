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
	"regexp"
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
	input := make(map[string]struct{})
	variables := make(map[string]interface{})
	for i := 0; i < len(maskingDef); i++ {
		if maskingDef[i].Masks != nil {
			for _, v := range maskingDef[i].Masks {
				res += exportMask(maskingDef[i], v, input, variables) + "\n    "
			}
		}
		mask := exportMask(maskingDef[i], maskingDef[i].Mask, input, variables)
		if mask != "" {
			res += mask + "\n    "
		}
	}
	res += addInputValues(input)
	res += addOutputValues(variables)
	return res, nil
}

func exportMask(masking model.Masking, mask model.MaskType, input map[string]struct{}, variables map[string]interface{}) string {
	if mask.Add != nil {
		str := "!add --> " + masking.Selector.Jsonpath
		key := masking.Selector.Jsonpath + "_1"
		variables[key] = mask.Add
		if mask.Add.(string) != "" {
			return str + "\n    " + mask.Add.(string) + " --> " + masking.Selector.Jsonpath
		}
		return str
	}
	if mask.AddTransient != nil {
		return mask.AddTransient.(string) + " -->|AddTransient| " + masking.Selector.Jsonpath
	}
	if mask.Constant != nil {
		return mask.Constant.(string) + " -->|Constant| " + masking.Selector.Jsonpath
	}
	if mask.RandomChoice != nil {
		return masking.Selector.Jsonpath + " -->|\"RandomChoice(" + flattenChoices(mask) + ")\"| " + masking.Selector.Jsonpath + "_1"
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
		input[masking.Selector.Jsonpath] = struct{}{}
		return masking.Selector.Jsonpath + " -->|\"HashInURI(" + mask.HashInURI + ")\"| " + masking.Selector.Jsonpath + "_1"
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
		input[masking.Selector.Jsonpath] = struct{}{}
		key := masking.Selector.Jsonpath + "_1"
		variables[key] = ""
		definitionString := " -->|\"Template(" + mask.Template + ")\"| " + masking.Selector.Jsonpath + "_1\n    "
		res := masking.Selector.Jsonpath + definitionString
		res += unescapeTemplateValues(mask.Template, definitionString, variables)
		return res
	}
	if mask.TemplateEach != (model.TemplateEachType{}) {
		return "TemplateEach[Item: " + mask.TemplateEach.Item + ", Index: " + mask.TemplateEach.Index + ", Template: " + mask.TemplateEach.Template + "] -->|TemplateEach| " + masking.Selector.Jsonpath
	}
	if mask.Duration != "" {
		return mask.Duration + " -->|Duration| " + masking.Selector.Jsonpath
	}
	if mask.Remove {
		return checkValueToRemove(masking, input, variables)
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
		str += flattenPipe(mask.Pipe, masking, input, variables)
		return str
	}
	if mask.FromJSON != "" {
		return mask.FromJSON + " -->|FromJSON| " + masking.Selector.Jsonpath
	}
	if mask.Luhn != nil {
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

func unescapeTemplateValues(templateValue, definitionString string, variables map[string]interface{}) string {
	res := ""
	regex := regexp.MustCompile(`(?:{{\.)([0-z]+)(?:}})`)
	splittedTemplate := regex.FindAllString(templateValue, -1)
	for i := range splittedTemplate {
		value := splittedTemplate[i][3:len(splittedTemplate[i])-2] + "_1"
		variables[value] = ""
		res += value + definitionString
	}
	return res[:len(res)-5]
}

func flattenPipe(mask model.PipeType, masking model.Masking, input map[string]struct{}, variables map[string]interface{}) string {
	str := make([]string, len(mask.Masking))
	for i, v := range mask.Masking {
		str[i] = masking.Selector.Jsonpath + " --> " + exportMask(v, v.Mask, input, variables) + "\n    "
	}
	return strings.Join(str, "\n    ")
}

func checkValueToRemove(masking model.Masking, input map[string]struct{}, variables map[string]interface{}) string {
	_, ok := input[masking.Selector.Jsonpath]
	keyVar := masking.Selector.Jsonpath + "_1"
	_, okVar := variables[keyVar]
	if ok {
		delete(input, masking.Selector.Jsonpath)
	}
	if okVar {
		delete(variables, keyVar)
		return keyVar + " --> !remove"
	}
	return masking.Selector.Jsonpath + " --> !remove"
}

func addInputValues(input map[string]struct{}) string {
	res := ""
	for k, _ := range input {
		res += "input[(input)] --> " + k + "\n    "
	}
	return res
}

func addOutputValues(variables map[string]interface{}) string {
	res := ""
	for k, _ := range variables {
		res += k + " --> output>output]\n    "
	}
	return res
}
