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
	"regexp"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/add"
	"github.com/cgi-fr/pimo/pkg/addtransient"
	"github.com/cgi-fr/pimo/pkg/command"
	"github.com/cgi-fr/pimo/pkg/constant"
	"github.com/cgi-fr/pimo/pkg/dateparser"
	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/ff1"
	"github.com/cgi-fr/pimo/pkg/fluxuri"
	"github.com/cgi-fr/pimo/pkg/fromjson"
	"github.com/cgi-fr/pimo/pkg/hash"
	"github.com/cgi-fr/pimo/pkg/hashcsv"
	"github.com/cgi-fr/pimo/pkg/increment"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/luhn"
	"github.com/cgi-fr/pimo/pkg/markov"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/pipe"
	"github.com/cgi-fr/pimo/pkg/randdate"
	"github.com/cgi-fr/pimo/pkg/randdura"
	"github.com/cgi-fr/pimo/pkg/randomcsv"
	"github.com/cgi-fr/pimo/pkg/randomdecimal"
	"github.com/cgi-fr/pimo/pkg/randomint"
	"github.com/cgi-fr/pimo/pkg/randomlist"
	"github.com/cgi-fr/pimo/pkg/randomuri"
	"github.com/cgi-fr/pimo/pkg/rangemask"
	"github.com/cgi-fr/pimo/pkg/regex"
	"github.com/cgi-fr/pimo/pkg/remove"
	"github.com/cgi-fr/pimo/pkg/replacement"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/cgi-fr/pimo/pkg/templateeach"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/cgi-fr/pimo/pkg/transcode"
	"github.com/cgi-fr/pimo/pkg/weightedchoice"
	"github.com/invopop/jsonschema"
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

func InjectMaskContextFactories() []model.MaskContextFactory {
	return []model.MaskContextFactory{
		fluxuri.Factory,
		add.Factory,
		addtransient.Factory,
		remove.Factory,
		pipe.Factory,
		templateeach.Factory,
		fromjson.Factory,
	}
}

func InjectMaskFactories() []model.MaskFactory {
	return []model.MaskFactory{
		constant.Factory,
		command.Factory,
		randomlist.Factory,
		randomuri.Factory,
		randomint.Factory,
		weightedchoice.Factory,
		regex.Factory,
		hash.Factory,
		randdate.Factory,
		increment.Factory,
		replacement.Factory,
		duration.Factory,
		templatemask.Factory,
		rangemask.Factory,
		randdura.Factory,
		randomdecimal.Factory,
		dateparser.Factory,
		ff1.Factory,
		luhn.Factory,
		markov.Factory,
		transcode.Factory,
		randomcsv.Factory,
		hashcsv.Factory,
	}
}

func InjectTemplateFuncs() {
	template.InjectSeededFuncGenerator("MaskRegex", regex.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoice", randomlist.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoiceInUri", randomuri.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoiceInURI", randomuri.Func)
	template.InjectSeededFuncGenerator("MaskRandomInt", randomint.Func)
	template.InjectSeededFuncGenerator("MaskRandomDecimal", randomdecimal.Func)
	template.InjectSeededFuncGenerator("MaskCommand", command.Func)
	template.InjectSeededFuncGenerator("MaskWeightedChoice", weightedchoice.Func)
	template.InjectSeededFuncGenerator("MaskHash", hash.Func)
	template.InjectSeededFuncGenerator("MaskHashInUri", hash.FuncInUri)
	template.InjectSeededFuncGenerator("MaskHashInURI", hash.FuncInUri)
	template.InjectSeededFuncGenerator("MaskRandDate", randdate.Func)
	template.InjectSeededFuncGenerator("MaskDuration", duration.Func)
	template.InjectSeededFuncGenerator("MaskDateParser", dateparser.Func)
	template.InjectSeededFuncGenerator("MaskRandomDuration", randdura.Func)
	template.InjectSeededFuncGenerator("MaskFF1", ff1.Func)
	template.InjectSeededFuncGenerator("MaskFf1", ff1.Func)
	template.InjectSeededFuncGenerator("MaskFF1_v2", ff1.FuncV2)
	template.InjectSeededFuncGenerator("MaskFf1_v2", ff1.FuncV2)
	template.InjectSeededFuncGenerator("MaskRange", rangemask.Func)
	template.InjectSeededFuncGenerator("MaskLuhn", luhn.Func)
	template.InjectSeededFuncGenerator("MaskTranscode", transcode.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoiceInCsv", randomcsv.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoiceInCSV", randomcsv.Func)
	template.InjectSeededFuncGenerator("MaskHashInCsv", hashcsv.Func)
	template.InjectSeededFuncGenerator("MaskHashInCSV", hashcsv.Func)
}

var re = regexp.MustCompile(`(\[\d*\])?$`)

func UpdateContext(counter int) {
	context := over.MDC().GetString("context")
	over.MDC().Set("context", re.ReplaceAllString(context, fmt.Sprintf("[%d]", counter)))
}
