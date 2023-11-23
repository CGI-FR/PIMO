package configuration

import (
	"fmt"
	"io"
	"os"
	"time"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/rs/zerolog/log"
)

type DumpCacheType = func(name string, cache model.Cache, path string, reverse bool) error

type LoadCacheType = func(name string, cache model.Cache, path string, reverse bool) error

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

type UpdateContextHandler = func(counter int)

type InjectTemplateFuncsHandler = func()

type InjectMaskContextFactories = func() []model.MaskContextFactory

type InjectMaskFactories = func() []model.MaskFactory

func NewContext(pdef model.Definition) Context {
	return Context{pdef: pdef}
}

func (ctx *Context) Configure(cfg Config, updateContext UpdateContextHandler, injectTemplateFuncs InjectTemplateFuncsHandler, injectMaskFactories InjectMaskFactories, injectMaskContextFactories InjectMaskContextFactories) error {
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
		ctx.source = model.NewCallableMapSource()
	case cfg.EmptyInput:
		over.MDC().Set("context", "empty-input")
		ctx.source = model.NewSourceFromSlice([]model.Dictionary{model.NewPackedDictionary()})
	case cfg.SingleInput != nil:
		over.MDC().Set("context", "single-input")
		ctx.source = model.NewSourceFromSlice([]model.Dictionary{cfg.SingleInput.Pack()})
	default:
		over.MDC().Set("context", "stdin")
		ctx.source = jsonline.NewPackedSource(os.Stdin)
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

	ctx.pipeline = model.NewPipeline(ctx.source).
		Process(model.NewCounterProcessWithCallback("input-line", 0, updateContext)).
		Process(model.NewRepeaterProcess(cfg.Iteration))
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

func (ctx *Context) ExecuteMap(data map[string]string) (map[string]string, error) {
	input := model.NewDictionary()

	for k, v := range data {
		input = input.With(k, v)
	}
	source, ok := ctx.source.(*model.CallableMapSource)
	if !ok {
		return nil, fmt.Errorf("Source is not CallableMapSource")
	}
	source.SetValue(input)
	result := []model.Entry{}
	err := ctx.pipeline.AddSink(model.NewSinkToSlice(&result)).Run()
	if err != nil {
		return nil, err
	}

	newData := make(map[string]string)

	if len(result) > 0 {
		new_map, ok := result[0].(model.Dictionary)
		if !ok {
			return nil, fmt.Errorf("result is not Dictionary")
		}
		unordered := new_map.Unordered()
		for k, v := range unordered {
			stringValue, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("Result is not a string")
			}
			newData[k] = stringValue
		}
		return newData, nil
	}
	return nil, fmt.Errorf("Result is not a map[string]string")
}

func (ctx *Context) Execute(out io.Writer, loadCache LoadCacheType, dumpCache DumpCacheType) (statistics.ExecutionStats, error) {
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
