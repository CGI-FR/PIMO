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

package apply

import (
	"hash/fnv"
	"strings"
	"text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a value that always mask the same way
type MaskEngine struct {
	pipeline model.Pipeline
	lib      string
}

// NewMask return a MaskEngine from a value
func NewMask(seed int64, caches map[string]model.Cache, fns template.FuncMap, uri string) (MaskEngine, error) {
	libname := "default"
	libmasking := uri

	libparts := strings.SplitN(uri, ":", 2)
	if len(libparts) == 2 {
		libname = libparts[0]
		libmasking = libparts[1]
		olddefault := model.SetDefault(libname)
		defer model.SetDefault(olddefault)
	}

	pipeline, err := model.Load(libname, libmasking, seed, caches, fns)

	return MaskEngine{pipeline, libname}, err
}

// func (me MaskEngine) MaskContext(e model.Dictionary, key string, context ...model.Dictionary) (model.Dictionary, error) {
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask apply")

	var result []model.Entry

	err := me.pipeline.
		WithSource(model.NewSourceFromSlice([]model.Dictionary{model.NewDictionary().With(".", e)})).
		// Process(model.NewCounterProcessWithCallback("internal", 1, updateContext)).
		AddSink(model.NewSinkToSlice(&result)).
		Run()
	if err != nil {
		return nil, err
	}

	return result[0], nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Apply.URI) > 0 {
		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		mask, err := NewMask(conf.Seed, conf.Cache, conf.Functions, conf.Masking.Mask.Apply.URI)
		if err != nil {
			return mask, true, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
