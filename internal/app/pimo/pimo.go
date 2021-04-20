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
	"fmt"
	"os"

	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
)

// YAMLStructure of the file
type YAMLStructure struct {
	Version string               `yaml:"version"`
	Seed    int64                `yaml:"seed"`
	Masking []model.Masking      `yaml:"masking"`
	Caches  map[string]YAMLCache `yaml:"caches"`
}

type MaskingV2 struct {
	Selector SelectorTypeV2         `yaml:"selector"`
	Mask     map[string]interface{} `yaml:"mask"`
	Cache    string                 `yaml:"cache"`
}

type SelectorTypeV2 struct {
	JsonPath string `yaml:"jsonpath"`
}

type YAMLStructureV2 struct {
	Version string               `yaml:"version"`
	Seed    int64                `yaml:"seed"`
	Masking []MaskingV2          `yaml:"masking"`
	Caches  map[string]YAMLCache `yaml:"caches"`
}

type YAMLCache struct {
	Unique bool `yaml:"unique"`
}

type CachedMaskEngineFactories func(model.MaskEngine) model.MaskEngine

func DumpCache(name string, cache model.Cache, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Cache %s not dump : %s", name, err.Error())
	}

	defer file.Close()
	err = model.NewPipeline(cache.Iterate()).AddSink(jsonline.NewSink(file)).Run()
	if err != nil {
		return fmt.Errorf("Cache %s not dump : %s", name, err.Error())
	}

	return nil
}

func LoadCache(name string, cache model.Cache, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	defer file.Close()
	err = model.NewPipeline(jsonline.NewSource(file)).AddSink(model.NewSinkToCache(cache)).Run()
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	return nil
}
