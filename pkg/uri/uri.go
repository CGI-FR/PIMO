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

package uri

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cgi-fr/pimo/pkg/maskingdata"
	"github.com/cgi-fr/pimo/pkg/model"
)

var MaxCapacityForEachLine = bufio.MaxScanTokenSize

type Records[T model.Entry | []string] interface {
	Len() int
	Get(idx int) T
	Collect() []T
}

type (
	CSVRecords   Records[[]string]
	EntryRecords Records[model.Entry]
	DictRecords  Records[model.Dictionary]
)

type records[T model.Entry | []string] struct {
	cache []T
}

func (r records[T]) Len() int {
	return len(r.cache)
}

func (r records[T]) Get(idx int) T {
	return r.cache[idx]
}

func (r records[T]) Collect() []T {
	result := make([]T, r.Len())
	copy(result, r.cache)
	return result
}

var (
	cacheCSV   map[string]records[[]string]         = map[string]records[[]string]{}
	cacheEntry map[string]records[model.Entry]      = map[string]records[model.Entry]{}
	cacheDict  map[string]records[model.Dictionary] = map[string]records[model.Dictionary]{}
)

func ReadCsvAsDicts(uri string, sep rune, comment rune, fieldsPerRecord int, trimLeadingSpaces bool, hasHeaders bool) (DictRecords, error) {
	if records, present := cacheDict[uri]; present {
		return records, nil
	}

	csvRecords, err := ReadCsv(uri, sep, comment, fieldsPerRecord, trimLeadingSpaces)
	if err != nil {
		return nil, err
	}

	headers := []string{}
	if hasHeaders {
		headerRecord := csvRecords.Get(0)
		for _, header := range headerRecord {
			if trimLeadingSpaces {
				headers = append(headers, strings.TrimSpace(header))
			} else {
				headers = append(headers, header)
			}
		}
	}

	i := 0
	if hasHeaders {
		i = 1
	}

	results := make([]model.Dictionary, csvRecords.Len())
	for ; i < csvRecords.Len(); i++ {
		record := csvRecords.Get(i)
		obj := model.NewDictionary()
		if hasHeaders {
			for i, header := range headers {
				if trimLeadingSpaces {
					obj.Set(header, strings.TrimSpace(record[i]))
				} else {
					obj.Set(header, record[i])
				}
			}
		} else {
			for i, value := range record {
				if trimLeadingSpaces {
					obj.Set(strconv.Itoa(i), strings.TrimSpace(value))
				} else {
					obj.Set(strconv.Itoa(i), value)
				}
			}
		}
		results[i] = obj
	}

	if hasHeaders {
		results = results[1:]
	}

	return records[model.Dictionary]{results}, nil
}

func ReadCsv(uri string, sep rune, comment rune, fieldsPerRecord int, trimLeadingSpaces bool) (CSVRecords, error) {
	if records, present := cacheCSV[uri]; present {
		return records, nil
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	var result records[[]string]

	switch u.Scheme {
	case "file":
		f, err := os.Open(u.Host + u.Path)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		csvReader.Comma = sep
		csvReader.Comment = comment
		csvReader.FieldsPerRecord = fieldsPerRecord
		csvReader.TrimLeadingSpace = trimLeadingSpaces

		r, err := csvReader.ReadAll()
		if err != nil {
			return nil, err
		}

		result = records[[]string]{r}
	case "http", "https":
		rep, err := http.Get(uri) //nolint:gosec
		if err != nil {
			return nil, err
		}
		defer rep.Body.Close()

		csvReader := csv.NewReader(rep.Body)
		csvReader.Comma = sep
		csvReader.Comment = comment
		csvReader.FieldsPerRecord = fieldsPerRecord
		csvReader.TrimLeadingSpace = trimLeadingSpaces

		r, err := csvReader.ReadAll()
		if err != nil {
			return nil, err
		}

		result = records[[]string]{r}
	}

	cacheCSV[uri] = result

	return result, nil
}

func Read(uri string) (EntryRecords, error) {
	if records, present := cacheEntry[uri]; present {
		return records, nil
	}

	var result []model.Entry

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "file" {
		file, err := os.Open(u.Host + u.Path)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		scanner.Buffer(make([]byte, MaxCapacityForEachLine), MaxCapacityForEachLine)
		for scanner.Scan() {
			value := scanner.Text()
			intValue, err := strconv.Atoi(value)
			if err == nil {
				result = append(result, intValue)
			} else {
				result = append(result, scanner.Text())
			}
		}
		if errors.Is(scanner.Err(), bufio.ErrTooLong) {
			return records[model.Entry]{result}, errors.New("error reading " + uri + ": line is too long, use --buffer-size parameter to increase buffer size")
		}
		cacheEntry[uri] = records[model.Entry]{result}
		return cacheEntry[uri], scanner.Err()
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		/* #nosec */
		rep, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		defer rep.Body.Close()
		scanner := bufio.NewScanner(rep.Body)
		scanner.Buffer(make([]byte, MaxCapacityForEachLine), MaxCapacityForEachLine)
		for scanner.Scan() {
			value := scanner.Text()
			intValue, err := strconv.Atoi(value)
			if err == nil {
				result = append(result, intValue)
			} else {
				result = append(result, scanner.Text())
			}
		}
		if errors.Is(scanner.Err(), bufio.ErrTooLong) {
			return records[model.Entry]{result}, errors.New("error reading " + uri + ": line is too long, use --buffer-size parameter to increase buffer size")
		}
		cacheEntry[uri] = records[model.Entry]{result}
		return cacheEntry[uri], scanner.Err()
	}
	if u.Scheme == "pimo" {
		list, ok := maskingdata.MapData[u.Host]
		if !ok {
			return nil, fmt.Errorf("Not a Pimo inside file")
		}
		result := make([]model.Entry, len(list))
		for i, v := range list {
			result[i] = v
		}
		cacheEntry[uri] = records[model.Entry]{result}
		return cacheEntry[uri], nil
	}

	return nil, fmt.Errorf("%s is not a valid scheme", u.Scheme)
}
