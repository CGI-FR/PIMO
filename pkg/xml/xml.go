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

package xml

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"html/template"
	"strings"

	"github.com/CGI-FR/xixo/pkg/xixo"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a struct to mask XML content within JSON values
type MaskEngine struct {
	source       string
	xPath        string
	injectParent string
	pipeline     model.Pipeline
}

// NewMask return a MaskEngine from xPath name, injectParent and Masking config
func NewMask(xPath, injectParent string, caches map[string]model.Cache, fns template.FuncMap, filename string, seed int64, subMasking ...model.Masking) (MaskEngine, error) {
	var definition model.Definition
	var err error
	if len(filename) > 0 {
		definition, err = model.LoadPipelineDefinitionFromFile(filename)
		if err != nil {
			return MaskEngine{filename, xPath, injectParent, nil}, err
		}
		// merge the current seed with the seed provided by configuration on the pipe
		definition.Seed += seed
	} else {
		definition = model.Definition{Seed: seed + 1, Masking: subMasking}
	}
	pipeline := model.NewPipeline(nil)
	pipeline, _, err = model.BuildPipeline(pipeline, definition, caches, fns, "", "")
	return MaskEngine{"", xPath, injectParent, pipeline}, err
}

// Mask choose the target attribute or tag value and apply masking configuration
func (engine MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask XML")

	// Prepare date
	stringXML, ok := e.(string)
	if !ok {
		return nil, fmt.Errorf("jsonpath content is not a string")
	}
	contentReader := strings.NewReader(stringXML)
	var resultBuffer bytes.Buffer
	// Create xml parser
	parser := xixo.NewXMLParser(contentReader, &resultBuffer).EnableXpath()
	// Apply masking
	parser.RegisterMapCallback(engine.xPath, func(m map[string]string) (map[string]string, error) {
		source := model.NewCallableMapSource()
		input := model.NewDictionary()
		for k, v := range m {
			input = input.With(k, v)
		}
		source.SetValue(input)
		result := []model.Entry{}
		err := engine.pipeline.WithSource(source).AddSink(model.NewSinkToSlice(&result)).Run()
		if err != nil {
			return nil, err
		}

		if len(result) > 0 {
			newMap, ok := result[0].(model.Dictionary)
			if !ok {
				return nil, fmt.Errorf("Result is not Dictionary")
			}
			unordered := newMap.Unordered()
			result := make(map[string]string, len(unordered))
			for k, v := range unordered {
				stringValue, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("Result is not a string")
				}
				result[k] = stringValue
			}
			return result, nil
		} else {
			return nil, fmt.Errorf("Result is not a map[string]string")
		}
	})

	err := parser.Stream()
	if err != nil {
		log.Err(err).Msg("Error during parsing XML document")
	}
	// Return masked xml value in dictionary
	result := resultBuffer.String()
	return result, nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.XML.XPath) > 0 && len(conf.Masking.Mask.XML.Masking) > 0 {
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		mask, err := NewMask(conf.Masking.Mask.XML.XPath, conf.Masking.Mask.XML.InjectParent, conf.Cache, conf.Functions, conf.Masking.Mask.XML.DefinitionFile, conf.Seed, conf.Masking.Mask.XML.Masking...)
		if err != nil {
			log.Err(err).Msg("Error during factoring XML mask")
			return mask, true, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
