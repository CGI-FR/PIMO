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

package hashcsv

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a list of masking value and a rand init to mask
type MaskEngine struct {
	rand            *rand.Rand
	seeder          model.Seeder
	template        *template.Template
	cache           map[string][][]string
	header          bool
	sep             rune
	comment         rune
	fieldsPerRecord int
	trimSpaces      bool
}

// NewMask create a MaskRandomChoiceInCSV with a seed
func NewMask(conf model.ChoiceInCSVType, seed int64, seeder model.Seeder) (MaskEngine, error) {
	template, err := template.New("template-randomInCsv").Parse(conf.URI)
	sep := ','
	if len(conf.Separator) > 0 {
		sep, _ = utf8.DecodeRune([]byte(conf.Separator))
	}
	var comment rune
	if len(conf.Comment) > 0 {
		comment, _ = utf8.DecodeRune([]byte(conf.Comment))
	}
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, template, map[string][][]string{}, conf.Header, sep, comment, conf.FieldsPerRecord, conf.TrimSpace}, err
}

// Mask choose a mask value randomly
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask hashInCSV")

	if len(context) > 0 {
		seed, ok, err := mrl.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			mrl.rand.Seed(seed)
		}
	}

	var output bytes.Buffer
	if len(context) == 0 {
		context = []model.Dictionary{model.NewDictionary()}
	}
	if err := mrl.template.Execute(&output, context[0].UnpackAsDict().Unordered()); err != nil {
		return nil, err
	}
	filename := output.String()

	var records [][]string
	if recordsFromCache, ok := mrl.cache[filename]; ok {
		records = recordsFromCache
	} else {
		recordsFromFile, err := uri.ReadCsv(filename, mrl.sep, mrl.comment, mrl.fieldsPerRecord, mrl.trimSpaces)
		if err != nil {
			return nil, err
		}
		mrl.cache[filename] = recordsFromFile
		records = recordsFromFile
	}

	if mrl.header {
		h := fnv.New32a()
		if _, err := h.Write([]byte(fmt.Sprintf("%v", e))); err != nil {
			return nil, err
		}
		record := records[1:][int(h.Sum32())%(len(records)-1)]
		obj := model.NewDictionary()
		headers := records[0]
		for i, header := range headers {
			if mrl.trimSpaces {
				obj.Set(strings.TrimSpace(header), strings.TrimSpace(record[i]))
			} else {
				obj.Set(header, record[i])
			}
		}
		return obj, nil
	} else {
		h := fnv.New32a()
		if _, err := h.Write([]byte(fmt.Sprintf("%v", e))); err != nil {
			return nil, err
		}
		record := records[int(h.Sum32())%len(records)]
		obj := model.NewDictionary()
		for i, value := range record {
			if mrl.trimSpaces {
				obj.Set(strconv.Itoa(i), strings.TrimSpace(value))
			} else {
				obj.Set(strconv.Itoa(i), value)
			}
		}
		return obj, nil
	}
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

	// set differents seeds for differents jsonpath
	h := fnv.New64a()
	h.Write([]byte(conf.Masking.Selector.Jsonpath))
	conf.Seed += int64(h.Sum64())

	if len(conf.Masking.Mask.HashInCSV.URI) != 0 {
		mask, err := NewMask(conf.Masking.Mask.HashInCSV, conf.Seed, seeder)
		return mask, true, err
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(uri string, header bool, sep string, comment string, fieldsPerRecord int, trimSpaces bool) (model.Entry, error) {
		mask, err := NewMask(model.ChoiceInCSVType{URI: uri, Header: header, Separator: sep, Comment: comment, FieldsPerRecord: fieldsPerRecord, TrimSpace: trimSpaces}, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		if err != nil {
			return "", err
		}
		callnumber++
		return mask.Mask(nil)
	}
}
