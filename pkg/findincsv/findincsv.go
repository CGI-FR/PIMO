package findincsv

import (
	"bytes"
	"hash/fnv"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"unicode/utf8"

	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	tmlmask "github.com/cgi-fr/pimo/pkg/template"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	rand               *rand.Rand
	seeder             model.Seeder
	template           *template.Template
	temExactMatchCSV   *tmlmask.Engine
	temExactMatchEntry *tmlmask.Engine
	expected           string
	cache              map[string][][]string
	findCache          map[string][]string
	header             bool
	sep                rune
	comment            rune
	fieldsPerRecord    int
	trimSpaces         bool
}

func NewMask(conf model.FindInCSVType, seed int64, seeder model.Seeder) (MaskEngine, error) {
	template, err := template.New("template-findInCsv").Parse(conf.URI)
	sep := ','
	if len(conf.Separator) > 0 {
		sep, _ = utf8.DecodeRune([]byte(conf.Separator))
	}
	var comment rune
	if len(conf.Comment) > 0 {
		comment, _ = utf8.DecodeRune([]byte(conf.Comment))
	}
	var tempExactMatchEntry, tempExactMatchCSV *tmlmask.Engine
	if len(conf.ExactMatch.CSV) > 0 && len(conf.ExactMatch.Entry) > 0 {
		tempExactMatchEntry, err = tmlmask.NewEngine(conf.ExactMatch.Entry, tmpl.FuncMap{}, seed, "")
		if err != nil {
			return MaskEngine{}, err
		}
		tempExactMatchCSV, err = tmlmask.NewEngine(conf.ExactMatch.CSV, tmpl.FuncMap{}, seed, "")
		if err != nil {
			return MaskEngine{}, err
		}
	}

	var expected string
	if len(conf.Expected) > 0 {
		expected = conf.Expected
	}
	// nolint: gosec
	return MaskEngine{rand.New(rand.NewSource(seed)), seeder, template, tempExactMatchCSV, tempExactMatchEntry, expected, map[string][][]string{}, map[string][]string{}, conf.Header, sep, comment, conf.FieldsPerRecord, conf.TrimSpace}, err
}

// Mask choose one or many line's value in csv which matched with given value from json entry.
func (mrl MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask findInCSV")

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
		context = []model.Dictionary{model.NewPackedDictionary()}
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

	// Get template Entry value
	var exactEntryBuffer bytes.Buffer
	if err := mrl.temExactMatchEntry.Execute(&exactEntryBuffer, context[0].UnpackAsDict().Unordered()); err != nil {
		return nil, err
	}
	exactEntryString := exactEntryBuffer.String()
	headers := records[0]
	// Get template CSV value
	for _, record := range records {
		dictionary := model.NewDictionary()
		for i, header := range headers {
			if mrl.trimSpaces {
				dictionary.Set(strings.TrimSpace(header), strings.TrimSpace(record[i]))
			} else {
				dictionary.Copy().Set(header, record[i])
			}
		}
		var exactCSVBuffer bytes.Buffer
		if err := mrl.temExactMatchCSV.Execute(&exactCSVBuffer, context[0].UnpackAsDict().Unordered()); err != nil {
			return nil, err
		}
		exactEntryCSV := exactCSVBuffer.String()
		if exactEntryString == exactEntryCSV {
			if mrl.header {
				record := records[1:][mrl.rand.Intn(len(records)-1)]
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
				record := records[mrl.rand.Intn(len(records))]
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
	}

	if mrl.header {
		record := records[1:][mrl.rand.Intn(len(records)-1)]
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
		record := records[mrl.rand.Intn(len(records))]
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

	h := fnv.New64a()
	h.Write([]byte(conf.Masking.Selector.Jsonpath))
	conf.Seed += int64(h.Sum64())

	if len(conf.Masking.Mask.FindInCSV.URI) != 0 {
		mask, err := NewMask(conf.Masking.Mask.FindInCSV, conf.Seed, seeder)
		return mask, true, err
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(uri string, header bool, sep string, comment string, fieldsPerRecord int, trimSpaces bool) (model.Entry, error) {
		mask, err := NewMask(model.FindInCSVType{URI: uri, Header: header, Separator: sep, Comment: comment, FieldsPerRecord: fieldsPerRecord, TrimSpace: trimSpaces}, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		if err != nil {
			return "", err
		}
		callnumber++
		return mask.Mask(nil)
	}
}
