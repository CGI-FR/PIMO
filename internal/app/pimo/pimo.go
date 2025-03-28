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
	"io"
	"os"
	"regexp"
	"time"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/add"
	"github.com/cgi-fr/pimo/pkg/addtransient"
	"github.com/cgi-fr/pimo/pkg/apply"
	"github.com/cgi-fr/pimo/pkg/command"
	"github.com/cgi-fr/pimo/pkg/constant"
	"github.com/cgi-fr/pimo/pkg/dateparser"
	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/ff1"
	"github.com/cgi-fr/pimo/pkg/findincsv"
	"github.com/cgi-fr/pimo/pkg/fluxuri"
	"github.com/cgi-fr/pimo/pkg/fromjson"
	"github.com/cgi-fr/pimo/pkg/hash"
	"github.com/cgi-fr/pimo/pkg/hashcsv"
	"github.com/cgi-fr/pimo/pkg/increment"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/logmask"
	"github.com/cgi-fr/pimo/pkg/luhn"
	"github.com/cgi-fr/pimo/pkg/markov"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/parquet"
	"github.com/cgi-fr/pimo/pkg/partition"
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
	"github.com/cgi-fr/pimo/pkg/segment"
	"github.com/cgi-fr/pimo/pkg/sequence"
	"github.com/cgi-fr/pimo/pkg/sha3"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/cgi-fr/pimo/pkg/templateeach"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/cgi-fr/pimo/pkg/timeline"
	"github.com/cgi-fr/pimo/pkg/transcode"
	"github.com/cgi-fr/pimo/pkg/weightedchoice"
	"github.com/cgi-fr/pimo/pkg/xml"
	"github.com/invopop/jsonschema"
	"github.com/rs/zerolog/log"
)

type CachedMaskEngineFactories func(model.MaskEngine) model.MaskEngine

type Config struct {
	SingleInput      *model.Dictionary
	EmptyInput       bool
	RepeatUntil      string
	RepeatWhile      string
	Iteration        int
	SkipLineOnError  bool
	SkipFieldOnError bool
	SkipLogFile      string
	CachesToDump     map[string]string
	CachesToLoad     map[string]string
	XMLCallback      bool
	ParquetInput     string
	ParquetOutput    string
}

type Context struct {
	cfg                 Config
	pdef                model.Definition
	pipeline            model.Pipeline
	source              model.Source
	repeatCondition     string
	repeatConditionMode string
	caches              map[string]model.Cache
}

func NewContext(pdef model.Definition) Context {
	return Context{pdef: pdef}
}

func (ctx *Context) Configure(cfg Config) error {
	log.Info().
		Bool("skipLineOnError", cfg.SkipLineOnError).
		Bool("skipFieldOnError", cfg.SkipFieldOnError).
		Int("repeat", cfg.Iteration).
		Bool("empty-input", cfg.EmptyInput).
		Interface("dump-cache", cfg.CachesToDump).
		Interface("load-cache", cfg.CachesToLoad).
		Msg("Start PIMO")

	ctx.cfg = cfg

	over.AddGlobalFields("context")
	switch {
	case cfg.XMLCallback:
		over.MDC().Set("context", "callback-input")
		ctx.source = model.NewMutableSource()
	case cfg.EmptyInput:
		over.MDC().Set("context", "empty-input")
		ctx.source = model.NewSliceSource([]model.Dictionary{model.NewPackedDictionary()})
	case cfg.SingleInput != nil:
		over.MDC().Set("context", "single-input")
		ctx.source = model.NewSliceSource([]model.Dictionary{cfg.SingleInput.Pack()})
	case cfg.ParquetInput != "":
		over.MDC().Set("context", "parquet")
		source, err := parquet.NewPackedSource(cfg.ParquetInput)
		if err != nil {
			return err
		}
		ctx.source = source
	default:
		over.MDC().Set("context", "stdin")
		ctx.source = jsonline.NewPackedSource(os.Stdin)
	}

	if cfg.RepeatUntil != "" && cfg.RepeatWhile != "" {
		return fmt.Errorf("Cannot use repeatUntil and repeatWhile options together")
	}

	if cfg.Iteration > 1 {
		ctx.source = model.NewCountRepeater(ctx.source, cfg.Iteration)
	}

	ctx.repeatCondition = cfg.RepeatWhile
	ctx.repeatConditionMode = "while"
	if cfg.RepeatUntil != "" {
		ctx.repeatCondition = cfg.RepeatUntil
		ctx.repeatConditionMode = "until"
	}

	ctx.pipeline = model.NewPipeline(ctx.source).
		Process(model.NewCounterProcessWithCallback("input-line", 0, updateContext))

	over.AddGlobalFields("input-line")

	injectTemplateFuncs()
	model.InjectMaskContextFactories(injectMaskContextFactories())
	model.InjectMaskFactories(injectMaskFactories())
	model.InjectConfig(cfg.SkipLineOnError, cfg.SkipFieldOnError, cfg.SkipLogFile)

	var err error
	ctx.pipeline, ctx.caches, err = model.BuildPipeline(ctx.pipeline, ctx.pdef, nil, nil, ctx.repeatCondition, ctx.repeatConditionMode)
	if err != nil {
		return fmt.Errorf("Cannot build pipeline: %w", err)
	}

	return nil
}

func (ctx *Context) Execute(out io.Writer) (statistics.ExecutionStats, error) {
	for name, path := range ctx.cfg.CachesToLoad {
		cache, ok := ctx.caches[name]
		if !ok {
			return statistics.WithErrorCode(2), fmt.Errorf("Cache not found: %v", name)
		}
		if err := loadCache(name, cache, path, ctx.pdef.Caches[name].Reverse); err != nil {
			return statistics.WithErrorCode(3), fmt.Errorf("Cannot load cache: %v", err)
		}
	}

	// init stats and time measure to zero
	statistics.Reset()
	startTime := time.Now()

	over.AddGlobalFields("output-line")

	var sink model.SinkProcess

	switch {
	case ctx.cfg.ParquetOutput != "":
		parquetSource, ok := ctx.source.(*parquet.Source)
		if !ok {
			return statistics.WithErrorCode(5), fmt.Errorf("Source must be parquet")
		}
		schema, err := parquetSource.Schema()
		if err != nil {
			return statistics.WithErrorCode(6), fmt.Errorf("Can't read shema from source")
		}
		fileWriter, err := os.Create(ctx.cfg.ParquetOutput)
		if err != nil {
			return statistics.WithErrorCode(7), fmt.Errorf("Can't write file %s", ctx.cfg.ParquetOutput)
		}
		sink, err = parquet.NewSinkWithContext(fileWriter, schema, "output-line")
		if err != nil {
			return statistics.WithErrorCode(8), fmt.Errorf("Can't create parquet file %s : %v", ctx.cfg.ParquetOutput, err)
		}
	default:
		sink = jsonline.NewSinkWithContext(out, "output-line")
	}
	err := ctx.pipeline.AddSink(sink).Run()
	if ctx.cfg.ParquetOutput != "" {
		sink.(parquet.Sink).Close()
	}
	// include duration info and stats in log output
	duration := time.Since(startTime)
	over.MDC().Set("duration", duration)
	stats := statistics.Compute()

	over.SetGlobalFields([]string{"config", "output-line", "input-line", "duration"})
	if err != nil {
		return stats.WithErrorCode(4), fmt.Errorf("Pipeline didn't complete run: %v", err)
	}

	for name, path := range ctx.cfg.CachesToDump {
		cache, ok := ctx.caches[name]
		if !ok {
			return stats.WithErrorCode(2), fmt.Errorf("Cache not found: %v", name)
		}
		err = dumpCache(name, cache, path, ctx.pdef.Caches[name].Reverse)
		if err != nil {
			return stats.WithErrorCode(3), fmt.Errorf("Cannot dump cache %v: %w", name, err)
		}
	}

	return stats, nil
}

func dumpCache(name string, cache model.Cache, path string, reverse bool) error {
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

func loadCache(name string, cache model.Cache, path string, reverse bool) error {
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

func injectMaskContextFactories() []model.MaskContextFactory {
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

func injectMaskFactories() []model.MaskFactory {
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
		findincsv.Factory,
		xml.Factory,
		timeline.Factory,
		sequence.Factory,
		sha3.Factory,
		apply.Factory,
		partition.Factory,
		segment.Factory,
		logmask.Factory,
	}
}

func injectTemplateFuncs() {
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
	template.InjectSeededFuncGenerator("MaskFF1_v3", ff1.FuncV3)
	template.InjectSeededFuncGenerator("MaskFf1_v3", ff1.FuncV3)
	template.InjectSeededFuncGenerator("MaskRange", rangemask.Func)
	template.InjectSeededFuncGenerator("MaskLuhn", luhn.Func)
	template.InjectSeededFuncGenerator("MaskTranscode", transcode.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoiceInCsv", randomcsv.Func)
	template.InjectSeededFuncGenerator("MaskRandomChoiceInCSV", randomcsv.Func)
	template.InjectSeededFuncGenerator("MaskHashInCsv", hashcsv.Func)
	template.InjectSeededFuncGenerator("MaskHashInCSV", hashcsv.Func)
	template.InjectSeededFuncGenerator("MaskFindInCSV", findincsv.Func)
	template.InjectSeededFuncGenerator("MaskFindInCsv", findincsv.Func)
	template.InjectSeededFuncGenerator("MaskSequence", sequence.Func)
}

var re = regexp.MustCompile(`(\[\d*\])?$`)

func updateContext(counter int) {
	context := over.MDC().GetString("context")
	over.MDC().Set("context", re.ReplaceAllString(context, fmt.Sprintf("[%d]", counter)))
}

func (ctx *Context) ExecuteMap(data map[string]any) (map[string]any, error) {
	input := model.NewDictionary()

	for k, v := range data {
		input = input.With(k, v)
	}
	source, ok := ctx.source.(*model.MutableSource)
	if !ok {
		return nil, fmt.Errorf("Source is not MutableSource")
	}
	source.SetValues(input)
	result := []model.Entry{}
	err := ctx.pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	if err != nil {
		return nil, err
	}

	newData := make(map[string]any)

	if len(result) > 0 {
		newMap, ok := result[0].(model.Dictionary)
		if !ok {
			return nil, fmt.Errorf("result is not Dictionary")
		}
		unordered := newMap.Unordered()
		for k, v := range unordered {
			newData[k] = v
		}
		return newData, nil
	}
	return nil, fmt.Errorf("Result is not a map[string]string")
}

func XMLCallback(ctx Context, xmlList map[string]string) (map[string]string, error) {
	dictionary := make(map[string]any, len(xmlList))
	for k, v := range xmlList {
		dictionary[k] = v
	}
	transformedData, err := ctx.ExecuteMap(dictionary)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, len(xmlList))
	for k, v := range transformedData {
		stringValue, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("Result is not a string")
		}
		result[k] = stringValue
	}
	return result, nil
}

func InjectMasks() {
	injectTemplateFuncs()
	model.InjectMaskContextFactories(injectMaskContextFactories())
	model.InjectMaskFactories(injectMaskFactories())
}
