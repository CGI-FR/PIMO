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

type subgraph struct {
	name    string
	masks   []edge
	removed bool
	added   bool
}

type edge struct {
	source      string
	destination string
	mask        string
	param       string
	key         string
}

func Export(masking model.Definition) (string, error) {
	maskingDef := masking.Masking
	if len(maskingDef) == 0 {
		return "", nil
	}
	res := `flowchart LR
    `
	variables := make(map[string]subgraph)
	maskOrder := make([]string, 0, 10)
	for i := 0; i < len(maskingDef); i++ {
		_, ok := variables[maskingDef[i].Selector.Jsonpath]
		if !ok {
			maskOrder = append(maskOrder, maskingDef[i].Selector.Jsonpath)
		}
		if maskingDef[i].Masks != nil {
			for _, v := range maskingDef[i].Masks {
				exportMask(maskingDef[i], v, variables)
			}
		}
		exportMask(maskingDef[i], maskingDef[i].Mask, variables)
	}
	res += printSubgraphs(variables, maskOrder)
	return res, nil
}

func exportMask(masking model.Masking, mask model.MaskType, variables map[string]subgraph) {
	maskSubgraph := subgraph{
		name:    masking.Selector.Jsonpath + "_sg",
		removed: false,
		added:   false,
		masks:   make([]edge, 0, 10),
	}
	edgeToAdd := edge{}
	edgeToAdd.key = masking.Selector.Jsonpath
	edgeToAdd.source = masking.Selector.Jsonpath
	edgeToAdd.destination = masking.Selector.Jsonpath + "_1"
	if elem, ok := variables[masking.Selector.Jsonpath]; ok {
		maskSubgraph = elem
		edgeToAdd = checkSourceAndDestination(edgeToAdd, len(elem.masks), maskSubgraph, masking)
	}

	if mask.Add != nil {
		edgeToAdd.mask = "Add"
		edgeToAdd.param = sanitizeParam(mask.Add.(string))
		maskSubgraph = exportAdd(maskSubgraph, edgeToAdd, mask, masking, variables)
		maskSubgraph.added = true
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.AddTransient != nil {
		edgeToAdd.mask = "AddTransient"
		edgeToAdd.param = sanitizeParam(mask.AddTransient.(string))
		maskSubgraph = exportAddTransient(maskSubgraph, edgeToAdd, mask, masking, variables)
		maskSubgraph.added = true
		maskSubgraph.removed = true
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Constant != nil {
		edgeToAdd.mask = "Constant"
		edgeToAdd.param = sanitizeParam(mask.Constant.(string))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RandomChoice != nil {
		edgeToAdd.mask = "RandomChoice"
		edgeToAdd.param = sanitizeParam(flattenChoices(mask))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RandomChoiceInURI != "" {
		edgeToAdd.mask = "RandomChoiceInURI"
		edgeToAdd.param = sanitizeParam(mask.RandomChoiceInURI)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Command != "" {
		edgeToAdd.mask = "Command"
		edgeToAdd.param = sanitizeParam(mask.Command)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RandomInt != (model.RandIntType{}) {
		edgeToAdd.mask = "RandomInt"
		edgeToAdd.param = sanitizeParam("Min: " + strconv.Itoa(mask.RandomInt.Min) + ", Max: " + strconv.Itoa(mask.RandomInt.Max))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if len(mask.WeightedChoice) > 0 {
		edgeToAdd.mask = "WeightedChoice"
		edgeToAdd.param = sanitizeParam(flattenWeightedChoices(mask))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Regex != "" {
		edgeToAdd.mask = "Regex"
		edgeToAdd.param = sanitizeParam(mask.Regex)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Hash != nil {
		edgeToAdd.mask = "Hash"
		edgeToAdd.param = sanitizeParam(flattenHash(mask))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.HashInURI != "" {
		edgeToAdd.mask = "HashInURI"
		edgeToAdd.param = sanitizeParam(mask.HashInURI)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RandDate != (model.RandDateType{}) {
		edgeToAdd.mask = "RandDate"
		edgeToAdd.param = sanitizeParam("DateMin: " + mask.RandDate.DateMin.String() + ", DateMax: " + mask.RandDate.DateMax.String())
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Incremental != (model.IncrementalType{}) {
		edgeToAdd.mask = "Incremental"
		edgeToAdd.param = sanitizeParam("Start: " + strconv.Itoa(mask.Incremental.Start) + ", Increment: " + strconv.Itoa(mask.Incremental.Increment))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Replacement != "" {
		edgeToAdd.mask = "Replacement"
		edgeToAdd.param = sanitizeParam(mask.Replacement)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Template != "" {
		edgeToAdd.mask = "Template"
		edgeToAdd.param = sanitizeParam(mask.Template)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		maskSubgraph = unescapeTemplateValues(mask.Template, "Template", masking.Selector.Jsonpath, variables, maskSubgraph)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.TemplateEach != (model.TemplateEachType{}) {
		edgeToAdd.mask = "TemplateEach"
		edgeToAdd.param = sanitizeParam("Item: " + mask.TemplateEach.Item + ", Index: " + mask.TemplateEach.Index + ", Template: " + mask.TemplateEach.Template)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		maskSubgraph = unescapeTemplateValues(mask.TemplateEach.Template, "TemplateEach", masking.Selector.Jsonpath, variables, maskSubgraph)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Duration != "" {
		edgeToAdd.mask = "Duration"
		edgeToAdd.param = sanitizeParam(mask.Duration)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Remove {
		edgeToAdd.mask = "Remove"
		edgeToAdd.param = sanitizeParam(strconv.FormatBool(true))
		maskSubgraph.removed = true
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RangeMask != 0 {
		edgeToAdd.mask = "RangeMask"
		edgeToAdd.param = sanitizeParam(strconv.Itoa(mask.RangeMask))
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RandomDuration != (model.RandomDurationType{}) {
		edgeToAdd.mask = "RandomDuration"
		edgeToAdd.param = sanitizeParam("Min: " + mask.RandomDuration.Min + ", Max: " + mask.RandomDuration.Max)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.FluxURI != "" {
		edgeToAdd.mask = "FluxURI"
		edgeToAdd.param = sanitizeParam(mask.FluxURI)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.RandomDecimal != (model.RandomDecimalType{}) {
		edgeToAdd.mask = "RandomDecimal"
		min := strconv.FormatFloat(mask.RandomDecimal.Min, 'E', mask.RandomDecimal.Precision, 64)
		max := strconv.FormatFloat(mask.RandomDecimal.Max, 'E', mask.RandomDecimal.Precision, 64)
		precision := strconv.Itoa(mask.RandomDecimal.Precision)
		edgeToAdd.param = sanitizeParam("Min: " + min + ", Max: " + max + ", Precision: " + precision)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.DateParser != (model.DateParserType{}) {
		edgeToAdd.mask = "DateParser"
		edgeToAdd.param = sanitizeParam("InputFormat: " + mask.DateParser.InputFormat + ", OutputFormat: " + mask.DateParser.OutputFormat)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.FromCache != "" {
		edgeToAdd.mask = "FromCache"
		edgeToAdd.param = sanitizeParam(mask.FromCache)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.FF1 != (model.FF1Type{}) {
		edgeToAdd.mask = "FF1"
		edgeToAdd.param = sanitizeParam("KeyFromEnv: " + mask.FF1.KeyFromEnv + ", TweakField: " + mask.FF1.TweakField + ", Radix: " + strconv.FormatUint(uint64(mask.FF1.Radix), 10) + ", Decrypt: " + strconv.FormatBool(mask.FF1.Decrypt))
		maskSubgraph = exportFF1(maskSubgraph, edgeToAdd, mask, variables)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Pipe.Masking != nil {
		edgeToAdd.mask = "Pipe"
		edgeToAdd.param = sanitizeParam("DefinitionFile: " + mask.Pipe.DefinitionFile + ", InjectParent: " + mask.Pipe.InjectParent + ", InjectRoot: " + mask.Pipe.InjectRoot)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.FromJSON != "" {
		edgeToAdd.mask = "FromJSON"
		edgeToAdd.param = sanitizeParam(mask.FromJSON)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
	if mask.Luhn != nil {
		edgeToAdd.mask = "Luhn"
		edgeToAdd.param = sanitizeParam(mask.Luhn.Universe)
		maskSubgraph.masks = append(maskSubgraph.masks, edgeToAdd)
		variables[masking.Selector.Jsonpath] = maskSubgraph
	}
}

func checkSourceAndDestination(edgeToAdd edge, maskCount int, maskSubgraph subgraph, masking model.Masking) edge {
	if maskCount > 0 {
		edgeToAdd.source = maskSubgraph.masks[len(maskSubgraph.masks)-1].destination
		edgeToAdd.destination = masking.Selector.Jsonpath + "_" + strconv.Itoa(len(maskSubgraph.masks)+1)
	}
	return edgeToAdd
}

func exportAdd(maskSubgraph subgraph, addEdge edge, mask model.MaskType, masking model.Masking, variables map[string]subgraph) subgraph {
	if mask.Add.(string) != "" {
		maskSubgraph.masks = append(maskSubgraph.masks, addEdge)
	}
	if strings.Contains(mask.Add.(string), "{{") {
		maskSubgraph = unescapeTemplateValues(mask.Add.(string), "Add", masking.Selector.Jsonpath, variables, maskSubgraph)
	}
	maskSubgraph.added = true
	return maskSubgraph
}

func exportAddTransient(maskSubgraph subgraph, addEdge edge, mask model.MaskType, masking model.Masking, variables map[string]subgraph) subgraph {
	if mask.AddTransient.(string) != "" {
		maskSubgraph.masks = append(maskSubgraph.masks, addEdge)
	}
	if strings.Contains(mask.AddTransient.(string), "{{") {
		maskSubgraph = unescapeTemplateValues(mask.AddTransient.(string), "AddTransient", masking.Selector.Jsonpath, variables, maskSubgraph)
	}
	maskSubgraph.added = true
	return maskSubgraph
}

func exportFF1(maskSubgraph subgraph, addEdge edge, mask model.MaskType, variables map[string]subgraph) subgraph {
	edgeTweakField := edge{mask: "FF1", destination: addEdge.destination, param: addEdge.param}
	tweakFieldMaskCount := len(variables[mask.FF1.TweakField].masks)
	if tweakFieldMaskCount > 0 {
		edgeTweakField.source = mask.FF1.TweakField + "_" + strconv.Itoa(tweakFieldMaskCount)
	} else {
		edgeTweakField.source = mask.FF1.TweakField
	}
	maskSubgraph.masks = append(maskSubgraph.masks, addEdge, edgeTweakField)
	return maskSubgraph
}

func flattenChoices(mask model.MaskType) string {
	choices := make([]string, len(mask.RandomChoice))
	for i, v := range mask.RandomChoice {
		choices[i] = sanitizeParam(v.(string))
	}
	return strings.Join(choices, ",")
}

func flattenWeightedChoices(mask model.MaskType) string {
	choices := make([]string, len(mask.WeightedChoice))
	for i, v := range mask.WeightedChoice {
		weight := uint64(v.Weight)
		weightStr := strconv.FormatUint(weight, 10)
		choices[i] = v.Choice.(string) + " @ " + weightStr
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

func unescapeTemplateValues(templateValue, mask, jsonpath string, variables map[string]subgraph, maskSubgraph subgraph) subgraph {
	regex := regexp.MustCompile(`((?:({{)|\s)\.([0-z,\_,-]+))`)
	splittedTemplate := regex.FindAllSubmatch([]byte(templateValue), -1)
	edges := make([]edge, 0, 10)
	jsonpathMaskCount := len(variables[jsonpath].masks)
	for i := range splittedTemplate {
		templateEdge := edge{
			mask:  mask,
			param: sanitizeParam(templateValue),
		}

		value := string(splittedTemplate[i][1])
		if strings.Contains(value, "{{") {
			value = strings.TrimSpace(value[3:])
		} else if strings.Contains(value, " ") {
			value = strings.TrimSpace(value)
			value = value[1:]
		}
		// to avoid confusion with intermediate steps (i.e. "name_1")
		if strings.Contains(value, "#underscore;") {
			value = value[0:strings.LastIndex(value, "#underscore;")]
		}

		templateEdge.key = value

		maskNumber := len(variables[value].masks)

		if maskNumber == 0 {
			templateEdge.source = value
		} else {
			templateEdge.source = value + "#underscore;" + strconv.Itoa(maskNumber)
		}

		templateEdge.destination = jsonpath + "#underscore;" + strconv.Itoa(jsonpathMaskCount+1)
		edges = append(edges, templateEdge)
	}
	maskSubgraph.masks = append(maskSubgraph.masks, edges...)
	return maskSubgraph
}

func sanitizeParam(param string) string {
	return strings.ReplaceAll(param, `"`, "#quot;")
}

func printSubgraphs(variables map[string]subgraph, maskOrder []string) string {
	inputText := ""
	subgraphText := ""
	outputText := ""
	for _, key := range maskOrder {
		if variables[key].added {
			inputText += "!add[/Add/] --> " + key + "\n    "
		} else {
			inputText += "!input[(input)] --> " + key + "\n    "
		}

		count := len(variables[key].masks)
		if count > 0 {
			subgraphText += "subgraph " + variables[key].name
			for j := range variables[key].masks {
				subgraphText, inputText = printMask(subgraphText, inputText, variables[key].masks[j], variables)
			}
			subgraphText += "\n    end\n    "
			outputText += strings.Replace(variables[key].masks[count-1].destination, "#underscore;", "_", 1)
		} else {
			outputText += key
		}
		if variables[key].removed {
			outputText += " --> !remove[\\Remove\\]\n    "
		} else {
			outputText += " --> !output>Output]\n    "
		}
	}
	return strings.TrimSpace(inputText + subgraphText + outputText)
}

func printMask(subgraphText, inputText string, mask edge, variables map[string]subgraph) (string, string) {
	source := strings.Replace(mask.source, "#underscore;", "_", 1)
	lastUnderscore := strings.LastIndex(source, "_")
	var key string
	if lastUnderscore == -1 {
		key = source
	} else {
		if _, notNumeric := strconv.Atoi(source[lastUnderscore+1:]); notNumeric != nil {
			key = source
		} else {
			key = source[0:lastUnderscore]
		}
	}

	_, ok := variables[key]
	if !ok && strings.TrimSpace(key) != "" {
		inputText += "!input[(input)] --> " + key + "\n    "
	}

	destination := strings.Replace(mask.destination, "#underscore;", "_", 1)
	subgraphText += "\n        " + source + " -->|\"" + mask.mask + "(" + mask.param + ")\"| " + destination
	return subgraphText, inputText
}
