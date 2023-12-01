package findincsv

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"html/template"
	"sort"
	"strconv"
	"strings"
	tmpl "text/template"
	"unicode/utf8"

	"github.com/cgi-fr/pimo/pkg/model"
	tmlmask "github.com/cgi-fr/pimo/pkg/template"
	"github.com/cgi-fr/pimo/pkg/uri"

	"github.com/rs/zerolog/log"
)

type CSVKey struct {
	filename string
	lineKey  string
}

type JaccardCSV struct {
	csvLine model.Entry
	lineKey string
}

type MaskEngine struct {
	seeder             model.Seeder
	templateURI        *template.Template
	temExactMatchCSV   *tmlmask.Engine // template to compute key for a csv entry
	temExactMatchEntry *tmlmask.Engine // template to compute key for json entry
	temJaccardCSV      *tmlmask.Engine // template to compute key for a csv entry
	temJaccardEntry    *tmlmask.Engine // template to compute key for json entry
	expected           string
	csvAllreadyRead    map[string][]model.Dictionary
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

	var temJaccardEntry *tmlmask.Engine
	var temJaccardCSV *tmlmask.Engine
	if len(conf.JaccardMatch.CSV) > 0 && len(conf.JaccardMatch.Entry) > 0 {
		temJaccardEntry, err = tmlmask.NewEngine(
			conf.JaccardMatch.Entry,
			tmpl.FuncMap{},
			seed,
			"", // use for seed functions ?
		)
		if err != nil {
			return MaskEngine{}, err
		}

		temJaccardCSV, err = tmlmask.NewEngine(
			conf.JaccardMatch.CSV,
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
		temJaccardCSV,
		temJaccardEntry,
		expected,
		map[string][]model.Dictionary{},
		map[CSVKey][]model.Entry{},
		conf.Header,
		sep, comment, conf.FieldsPerRecord, conf.TrimSpace,
	}, err
}

// Mask choose one or many line's value in csv which matched with given value from json entry.
func (me *MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask findInCSV")

	var filenameBuffer bytes.Buffer
	if len(context) == 0 {
		context = []model.Dictionary{model.NewPackedDictionary()}
	}
	if err := me.templateURI.Execute(&filenameBuffer, context[0].UnpackAsDict().Unordered()); err != nil {
		return nil, err
	}
	filename := filenameBuffer.String()
	var results []model.Entry
	var exactMatchFinded bool
	// Get ExactMatch results
	// If ExactMatch is nil, return a []Model.Entry{}
	if finded, exactMatchResult, err := me.ExactMatch(filename, context); err != nil {
		return nil, err
	} else {
		exactMatchFinded = finded
		results = exactMatchResult
	}
	// If exactMatch have no match, no need to pass jaccardMatch
	if me.temJaccardEntry != nil && me.temJaccardCSV != nil && exactMatchFinded {
		var jaccardEntryBuffer bytes.Buffer
		if err := me.temJaccardEntry.Execute(&jaccardEntryBuffer, context[0].UnpackAsDict().Unordered()); err != nil {
			return nil, err
		}
		jaccardEntryString := jaccardEntryBuffer.String()
		// If no exactMatch config
		if len(results) < 1 {
			var csvList []model.Dictionary
			if _, ok := me.csvAllreadyRead[filename]; !ok {
				var err error
				csvList, err = me.readCSV(filename)
				if err != nil {
					return nil, err
				}
			} else {
				csvList = me.csvAllreadyRead[filename]
			}
			var records []JaccardCSV
			for _, record := range csvList {
				lineKey, err := me.computeCSVLineKey(record, false)
				if err != nil {
					return nil, err
				}
				records = append(records, JaccardCSV{lineKey: lineKey, csvLine: record})
			}
			results = sortBySimilarity(jaccardEntryString, records)
		} else {
			var records []JaccardCSV
			for _, result := range results {
				lineKey, err := me.computeCSVLineKey(result.(model.Dictionary), false)
				if err != nil {
					return nil, err
				}
				records = append(records, JaccardCSV{lineKey: lineKey, csvLine: result})
			}
			results = sortBySimilarity(jaccardEntryString, records)
		}
	}
	return me.getExpectedResult(results)
}

func (me *MaskEngine) ExactMatch(filename string, context []model.Dictionary) (bool, []model.Entry, error) {
	if me.temExactMatchEntry != nil && me.temExactMatchCSV != nil {
		var csvList []model.Dictionary
		if _, ok := me.csvAllreadyRead[filename]; !ok {
			var err error
			csvList, err = me.readCSV(filename)
			if err != nil {
				return false, nil, err
			}
		}

		var exactEntryBuffer bytes.Buffer
		if err := me.temExactMatchEntry.Execute(&exactEntryBuffer, context[0].UnpackAsDict().Unordered()); err != nil {
			return false, nil, err
		}
		exactEntryString := exactEntryBuffer.String()
		err := me.getExactMatchCsvResult(filename, csvList)
		if err != nil {
			return false, []model.Entry{}, err
		}

		results := me.csvEntryByKey[CSVKey{
			filename: filename,
			lineKey:  exactEntryString,
		}]
		if len(results) < 1 {
			return false, results, nil
		}
		return true, results, nil
	}
	return true, []model.Entry{}, nil
}

func (me *MaskEngine) readCSV(filename string) ([]model.Dictionary, error) {
	recordsFromFile, err := uri.ReadCsv(filename, me.sep, me.comment, me.fieldsPerRecord, me.trimSpaces)
	if err != nil {
		return nil, err
	}
	csvList := me.createEntriesFromCSVLines(recordsFromFile)
	return csvList, nil
}

func (me *MaskEngine) computeCSVLineKey(record model.Dictionary, exactMatch bool) (string, error) {
	var output bytes.Buffer

	var template *tmlmask.Engine
	if exactMatch {
		template = me.temExactMatchCSV
	} else {
		template = me.temJaccardCSV
	}

	// Définissez un point de contrôle pour le débogage
	err := template.Execute(&output, record.Unordered())
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (me *MaskEngine) getExactMatchCsvResult(filename string, csvList []model.Dictionary) error {
	for _, record := range csvList {
		lineKey, err := me.computeCSVLineKey(record, true)
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

// Get numbers of result waited in expected config, by default return as at-least-one
func (me *MaskEngine) getExpectedResult(results []model.Entry) (model.Entry, error) {
	resultCount := len(results)

	switch me.expected {
	case "many":
		if resultCount < 1 {
			return []model.Entry{}, nil
		}
		return results, nil
	case "only-one":
		if resultCount != 1 {
			return nil, fmt.Errorf("Expected one result, but got %d", resultCount)
		}
	default:
		if resultCount == 0 {
			return nil, fmt.Errorf("Expected at least one result, but got none")
		}
		return results[0], nil
	}

	return results[0], nil
}

// JaccardSimilarity calculates the Jaccard similarity between two strings.
func JaccardSimilarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}

	set1 := convertStringToSet(s1)
	set2 := convertStringToSet(s2)

	intersection := setIntersection(set1, set2)
	union := setUnion(set1, set2)
	return float64(intersection) / float64(union)
}

// convertStringToSet converts a string into a set of tokens unordered.
// This doesn't measure similarity between texts, but if regarding a text as bag-of-letter
// Jaccard bigrammes
func convertStringToSet(s string) map[string]struct{} {
	set := make(map[string]struct{})
	runes := []rune(s)

	for i := 0; i < len(runes)-1; i++ {
		bigram := string(runes[i]) + string(runes[i+1])
		set[bigram] = struct{}{}
	}

	return set
}

// Get number of same bigrammes of 2 strings
func setIntersection(set1, set2 map[string]struct{}) int {
	intersection := 0
	for token := range set1 {
		if _, exists := set2[token]; exists {
			intersection++
		}
	}
	return intersection
}

// Get union of 2 strings's elements
func setUnion(set1, set2 map[string]struct{}) int {
	unionSet := make(map[string]struct{})
	for token := range set1 {
		unionSet[token] = struct{}{}
	}
	for token := range set2 {
		unionSet[token] = struct{}{}
	}
	return len(unionSet)
}

func sortBySimilarity(jaccardEntryString string, list []JaccardCSV) []model.Entry {
	type EntryWithSimilarity struct {
		Key        string
		Entry      model.Entry
		Similarity float64
	}

	var entriesWithSimilarity []EntryWithSimilarity

	for _, record := range list {
		similarity := JaccardSimilarity(jaccardEntryString, record.lineKey)
		entriesWithSimilarity = append(entriesWithSimilarity, EntryWithSimilarity{Key: record.lineKey, Entry: record.csvLine, Similarity: similarity})
	}

	// Function sort by similarity
	sort.Slice(entriesWithSimilarity, func(i, j int) bool {
		return entriesWithSimilarity[i].Similarity > entriesWithSimilarity[j].Similarity
	})

	// Extraire les entrées triées
	var sortedEntries []model.Entry
	for _, entryWithSimilarity := range entriesWithSimilarity {
		sortedEntries = append(sortedEntries, entryWithSimilarity.Entry)
	}

	return sortedEntries
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
	return func(uri string, exactMatchEntry string, exactMatchCsv string, jaccardMatchEntry string, jaccardMatchCsv string, expected string, header bool, sep string, comment string, fieldsPerRecord int, trimSpaces bool) (model.Entry, error) {
		exactMatch := model.ExactMatchType{
			CSV:   exactMatchCsv,
			Entry: exactMatchEntry,
		}

		jaccardMatch := model.ExactMatchType{
			CSV:   jaccardMatchCsv,
			Entry: jaccardMatchEntry,
		}
		mask, err := NewMask(model.FindInCSVType{URI: uri, Header: header, ExactMatch: exactMatch, JaccardMatch: jaccardMatch, Expected: expected, Separator: sep, Comment: comment, FieldsPerRecord: fieldsPerRecord, TrimSpace: trimSpaces}, seed+callnumber, model.NewSeeder(seedField, seed+callnumber))
		if err != nil {
			return "", err
		}
		callnumber++
		return mask.Mask(nil)
	}
}
