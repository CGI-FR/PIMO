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
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/cgi-fr/pimo/pkg/maskingdata"
	"github.com/cgi-fr/pimo/pkg/model"
)

func ReadCsv(uri string, sep rune) ([][]string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "file":
		f, err := os.Open(uri)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		csvReader.Comma = sep

		records, err := csvReader.ReadAll()
		if err != nil {
			return nil, err
		}

		return records, nil
	case "http", "https":
		rep, err := http.Get(uri) //nolint:gosec
		if err != nil {
			return nil, err
		}
		defer rep.Body.Close()

		csvReader := csv.NewReader(rep.Body)
		csvReader.Comma = sep

		records, err := csvReader.ReadAll()
		if err != nil {
			return nil, err
		}

		return records, nil
	}

	return nil, nil
}

func Read(uri string) ([]model.Entry, error) {
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
		for scanner.Scan() {
			value := scanner.Text()
			intValue, err := strconv.Atoi(value)
			if err == nil {
				result = append(result, intValue)
			} else {
				result = append(result, scanner.Text())
			}
		}
		return result, nil
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		/* #nosec */
		rep, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		defer rep.Body.Close()
		scanner := bufio.NewScanner(rep.Body)
		for scanner.Scan() {
			value := scanner.Text()
			intValue, err := strconv.Atoi(value)
			if err == nil {
				result = append(result, intValue)
			} else {
				result = append(result, scanner.Text())
			}
		}
		return result, nil
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
		return result, nil
	}

	return nil, fmt.Errorf(u.Scheme + " is not a valid scheme")
}
