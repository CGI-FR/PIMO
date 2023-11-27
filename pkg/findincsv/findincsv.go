package findincsv

import (
	"bytes"
	"hash/fnv"
	"html/template"
	"strconv"
	"strings"
	"unicode/utf8"

	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	tmlmask "github.com/cgi-fr/pimo/pkg/template"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
)

type CSVKey struct {
	filename string
	lineKey  string
}

type MaskEngine struct {
	seeder             model.Seeder
	templateURI        *template.Template
	temExactMatchCSV   *tmlmask.Engine // template to compute key for a csv entry
	temExactMatchEntry *tmlmask.Engine // template to compute key for json entry
	expected           string
	csvAllreadyRead    map[string]struct{}
	csvEntryByKey      map[CSVKey][]model.Entry
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
	var tempExactMatchEntry *tmlmask.Engine
	var tempExactMatchCSV *tmlmask.Engine
	if len(conf.ExactMatch.CSV) > 0 && len(conf.ExactMatch.Entry) > 0 {
		tempExactMatchEntry, err = tmlmask.NewEngine(
			conf.ExactMatch.Entry,
			tmpl.FuncMap{},
			seed,
			"", // use for seed functions ?
		)
		if err != nil {
			return MaskEngine{}, err
		}

		tempExactMatchCSV, err = tmlmask.NewEngine(
			conf.ExactMatch.CSV,
			tmpl.FuncMap{},
			seed,
			"")

		if err != nil {
			return MaskEngine{}, err
		}
	}

	var expected string
	if len(conf.Expected) > 0 {
		expected = conf.Expected
	}

	return MaskEngine{
		seeder,
		template,
		tempExactMatchCSV,
		tempExactMatchEntry,
		expected,
		map[string]struct{}{},
		map[CSVKey][]model.Entry{},
		conf.Header,
		sep, comment, conf.FieldsPerRecord, conf.TrimSpace,
	}, err
}

// Mask choose one or many line's value in csv which matched with given value from json entry.
func (me *MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask findInCSV")

	var filename bytes.Buffer
	if len(context) == 0 {
		context = []model.Dictionary{model.NewPackedDictionary()}
	}
	if err := me.templateURI.Execute(&filename, context[0].UnpackAsDict().Unordered()); err != nil {
		return nil, err
	}

	if _, ok := me.csvAllreadyRead[filename.String()]; !ok {
		err := me.readCSV(filename.String())
		if err != nil {
			return nil, err
		}
	}

	// Get template Entry value
	var exactEntryBuffer bytes.Buffer
	if err := me.temExactMatchEntry.Execute(&exactEntryBuffer, context[0].UnpackAsDict().Unordered()); err != nil {
		return nil, err
	}
	exactEntryString := exactEntryBuffer.String()

	results, ok := me.csvEntryByKey[CSVKey{
		filename: filename.String(),
		lineKey:  exactEntryString,
	}]

	if !ok {
		return []model.Entry{}, nil
	}
	if len(results) > 0 {
		return results[0], nil
	}

	return []model.Entry{}, nil
}

func (me *MaskEngine) readCSV(filename string) error {
	recordsFromFile, err := uri.ReadCsv(filename, me.sep, me.comment, me.fieldsPerRecord, me.trimSpaces)
	if err != nil {
		return err
	}

	for _, record := range me.createEntriesFromCSVLines(recordsFromFile) {
		lineKey, err := me.computeCSVLineKey(record)
		if err != nil {
			return err
		}

		key := CSVKey{
			filename: filename,
			lineKey:  lineKey,
		}

		if records, ok := me.csvEntryByKey[key]; ok {
			records = append(records, record)
			me.csvEntryByKey[key] = records
		} else {
			me.csvEntryByKey[key] = []model.Entry{record}
		}
	}

	return nil
}

func (me *MaskEngine) computeCSVLineKey(record model.Dictionary) (string, error) {
	var output bytes.Buffer

	err := me.temExactMatchCSV.Execute(&output, record.Unordered())
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (me *MaskEngine) createEntriesFromCSVLines(records [][]string) []model.Dictionary {
	results := []model.Dictionary{}

	for _, record := range records {
		if me.header {
			obj := model.NewDictionary()
			headers := records[0]
			for i, header := range headers {
				if me.trimSpaces {
					obj.Set(strings.TrimSpace(header), strings.TrimSpace(record[i]))
				} else {
					obj.Set(header, record[i])
				}
			}
			results = append(results, obj)
		} else {
			obj := model.NewDictionary()
			for i, value := range record {
				if me.trimSpaces {
					obj.Set(strconv.Itoa(i), strings.TrimSpace(value))
				} else {
					obj.Set(strconv.Itoa(i), value)
				}
			}
			results = append(results, obj)
		}
	}
	return results
}

// Factory create a mask from a yaml config
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

	h := fnv.New64a()
	h.Write([]byte(conf.Masking.Selector.Jsonpath))
	conf.Seed += int64(h.Sum64())

	if len(conf.Masking.Mask.FindInCSV.URI) != 0 {
		mask, err := NewMask(conf.Masking.Mask.FindInCSV, conf.Seed, seeder)
		return &mask, true, err
	}
	return nil, false, nil
}

func Func(seed int64, seedField string) interface{} {
	var callnumber int64
	return func(uri string, exactMatchEntry string, exactMatchCsv string, expected string, header bool, sep string, comment string, fieldsPerRecord int, trimSpaces bool) (model.Entry, error) {
		exactMatch := model.ExactMatchType{
			CSV:   exactMatchCsv,
			Entry: exactMatchEntry,
		}
		mask, err := NewMask(model.FindInCSVType{URI: uri, Header: header, ExactMatch: exactMatch, Expected: expected, Separator: sep, Comment: comment, FieldsPerRecord: fieldsPerRecord, TrimSpace: trimSpaces}, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		if err != nil {
			return "", err
		}
		callnumber++
		return mask.Mask(nil)
	}
}
