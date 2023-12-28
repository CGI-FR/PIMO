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

package pimo

import (
	"bytes"
	"hash/fnv"
	"io"
	"math/rand"
	"os"
	"strings"

	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a struct to mask XML content within JSON values
type MaskEngine struct {
	rand         *rand.Rand
	seeder       model.Seeder
	xPath        string
	injectParent string
	masking      []byte
}

// NewMask return a MaskEngine from xPath name, injectParent and Masking config
func NewMask(xPath, injectParent, subMasking string, seed int64, seeder model.Seeder) MaskEngine {
	prefix := `version: "1"
seed: 42
masking:`

	subMasking = prefix + subMasking
	subMasking = strings.ReplaceAll(subMasking, "\t", "  ")

	maskingConfig := []byte(subMasking)
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, xPath, injectParent, maskingConfig}
}

// Mask choose the target attribute or tag value and apply masking configuration
func (engine MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask XML")
	// Get masking configuration
	var conf model.Definition
	err := yaml.Unmarshal(engine.masking, &conf)
	if err != nil {
		return conf, err
	}

	(&conf).SetSeed(engine.rand.Int63())
	ctx := pimo.NewContext(conf)
	cfg := pimo.Config{
		Iteration:   1,
		XMLCallback: true,
	}

	if err := ctx.Configure(cfg); err != nil {
		log.Err(err).Msg("Cannot configure pipeline")
		log.Warn().Int("return", 1).Msg("End PIMO")
	}
	// Get xml value
	jsonDict := context[0].UnpackAsDict().Unordered()
	xmlValue := jsonDict[e.(string)]

	// Create xml parser
	contentReader := strings.NewReader(xmlValue.(string))
	var resultBuffer bytes.Buffer
	outputWriter := io.MultiWriter(&resultBuffer, os.Stdout)
	parser := pimo.ParseXML(contentReader, outputWriter)
	// Apply masking
	parser.RegisterMapCallback(engine.xPath, func(m map[string]string) (map[string]string, error) {
		return pimo.XMLCallback(ctx, m)
	})
	// Return masked xml value in json
	result := resultBuffer.String()
	return result, nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.XML.XPath) != 0 || len(conf.Masking.Mask.XML.Masking) != 0 {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)
		if len(conf.Masking.Mask.XML.InjectParent) == 0 {
			conf.Masking.Mask.XML.InjectParent = ""
		}
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		return NewMask(conf.Masking.Mask.XML.XPath,
			conf.Masking.Mask.XML.InjectParent,
			conf.Masking.Mask.XML.Masking,
			conf.Seed, seeder), true, nil
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(xPath string, injectParent string, masking string) (model.Entry, error) {
		mask := NewMask(xPath, injectParent, masking, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		callnumber++
		return mask.Mask(nil)
	}
}
