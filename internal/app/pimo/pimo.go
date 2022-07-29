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

	over "github.com/Trendyol/overlog"
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
	"github.com/cgi-fr/pimo/pkg/increment"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/luhn"
	"github.com/cgi-fr/pimo/pkg/markov"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/pipe"
	"github.com/cgi-fr/pimo/pkg/randdate"
	"github.com/cgi-fr/pimo/pkg/randdura"
	"github.com/cgi-fr/pimo/pkg/randomdecimal"
	"github.com/cgi-fr/pimo/pkg/randomint"
	"github.com/cgi-fr/pimo/pkg/randomlist"
	"github.com/cgi-fr/pimo/pkg/randomuri"
	"github.com/cgi-fr/pimo/pkg/rangemask"
	"github.com/cgi-fr/pimo/pkg/regex"
	"github.com/cgi-fr/pimo/pkg/remove"
	"github.com/cgi-fr/pimo/pkg/replacement"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/templateeach"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/cgi-fr/pimo/pkg/transcode"
	"github.com/cgi-fr/pimo/pkg/weightedchoice"
	"github.com/invopop/jsonschema"
	"github.com/rs/zerolog/log"
)

type CachedMaskEngineFactories func(model.MaskEngine) model.MaskEngine

type Config struct {
	MultipleInput    []model.Dictionary
	SingleInput      *model.Dictionary
	EmptyInput       bool
	RepeatUntil      string
	RepeatWhile      string
	Iteration        int
	SkipLineOnError  bool
	SkipFieldOnError bool
	CachesToDump     map[string]string
	CachesToLoad     map[string]string
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
	case cfg.EmptyInput:
		over.MDC().Set("context", "empty-input")
		ctx.source = model.NewSourceFromSlice([]model.Dictionary{{}})
	case cfg.SingleInput != nil:
		over.MDC().Set("context", "single-input")
		ctx.source = model.NewSourceFromSlice([]model.Dictionary{*cfg.SingleInput})
	case cfg.MultipleInput != nil:
		over.MDC().Set("context", "multiple-input")
		ctx.source = model.NewSourceFromSlice(cfg.MultipleInput)
	default:
		over.MDC().Set("context", "stdin")
		ctx.source = jsonline.NewSource(os.Stdin)
	}

	if cfg.RepeatUntil != "" && cfg.RepeatWhile != "" {
		return fmt.Errorf("Cannot use repeatUntil and repeatWhile options together")
	}

	ctx.repeatCondition = cfg.RepeatWhile
	ctx.repeatConditionMode = "while"
	if cfg.RepeatUntil != "" {
		ctx.repeatCondition = cfg.RepeatUntil
		ctx.repeatConditionMode = "until"
	}

	if ctx.repeatCondition != "" {
		ctx.source = model.NewTempSource(ctx.source)
	}

	ctx.pipeline = model.NewPipeline(ctx.source).
		Process(model.NewCounterProcessWithCallback("input-line", 0, updateContext)).
		Process(model.NewRepeaterProcess(cfg.Iteration))
	over.AddGlobalFields("input-line")

	model.InjectMaskContextFactories(injectMaskContextFactories())
	model.InjectMaskFactories(injectMaskFactories())
	model.InjectConfig(cfg.SkipLineOnError, cfg.SkipFieldOnError)

	var err error
	ctx.pipeline, ctx.caches, err = model.BuildPipeline(ctx.pipeline, ctx.pdef, nil)
	if err != nil {
		return fmt.Errorf("Cannot build pipeline: %w", err)
	}

	if ctx.repeatCondition != "" {
		processor, err := model.NewRepeaterUntilProcess(ctx.source.(*model.TempSource), ctx.repeatCondition, ctx.repeatConditionMode)
		if err != nil {
			return fmt.Errorf("Cannot build pipeline: %w", err)
		}
		ctx.pipeline = ctx.pipeline.Process(processor)
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
	err := ctx.pipeline.AddSink(jsonline.NewSinkWithContext(out, "output-line")).Run()

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
	}
}

var re = regexp.MustCompile(`(\[\d*\])?$`)

func updateContext(counter int) {
	context := over.MDC().GetString("context")
	over.MDC().Set("context", re.ReplaceAllString(context, fmt.Sprintf("[%d]", counter)))
}
