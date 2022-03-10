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
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/jsonschema"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
)

type CachedMaskEngineFactories func(model.MaskEngine) model.MaskEngine

func DumpCache(name string, cache model.Cache, path string, reverse bool) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Cache %s not dump : %s", name, err.Error())
	}

	defer file.Close()

	pipe := model.NewPipeline(cache.Iterate())

	if reverse {
		reverseFunc := func(d model.Dictionary) (model.Dictionary, error) {
			reverse := model.NewDictionary()
			reverse.Set("key", d.Get("value"))
			reverse.Set("value", d.Get("key"))
			return reverse, nil
		}

		pipe = pipe.Process(model.NewMapProcess(reverseFunc))
	}

	err = pipe.AddSink(jsonline.NewSink(file)).Run()

	if err != nil {
		return fmt.Errorf("Cache %s not dump : %s", name, err.Error())
	}

	return nil
}

func LoadCache(name string, cache model.Cache, path string, reverse bool) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	defer file.Close()

	pipe := model.NewPipeline(jsonline.NewSource(file))

	if reverse {
		reverseFunc := func(d model.Dictionary) (model.Dictionary, error) {
			reverse := model.NewDictionary()
			reverse.Set("key", d.Get("value"))
			reverse.Set("value", d.Get("key"))
			return reverse, nil
		}

		pipe = pipe.Process(model.NewMapProcess(reverseFunc))
	}

	err = pipe.AddSink(model.NewSinkToCache(cache)).Run()

	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	return nil
}

func GetJsonSchema() (string, error) {
	resBytes, err := json.MarshalIndent(jsonschema.Reflect(&model.Definition{}), "", "  ")
	if err != nil {
		return "", err
	}
	return string(resBytes), nil
}
