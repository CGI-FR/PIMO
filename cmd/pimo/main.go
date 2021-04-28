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

package main

import (
	"fmt"
	"os"

	over "github.com/Trendyol/overlog"
	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/add"
	"github.com/cgi-fr/pimo/pkg/command"
	"github.com/cgi-fr/pimo/pkg/constant"
	"github.com/cgi-fr/pimo/pkg/dateparser"
	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/ff1"
	"github.com/cgi-fr/pimo/pkg/fluxuri"
	"github.com/cgi-fr/pimo/pkg/hash"
	"github.com/cgi-fr/pimo/pkg/increment"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/pipe"
	"github.com/cgi-fr/pimo/pkg/randdate"
	"github.com/cgi-fr/pimo/pkg/randdura"
	"github.com/cgi-fr/pimo/pkg/randomdecimal"
	"github.com/cgi-fr/pimo/pkg/randomint"
	"github.com/cgi-fr/pimo/pkg/randomlist"
	"github.com/cgi-fr/pimo/pkg/rangemask"
	"github.com/cgi-fr/pimo/pkg/regex"
	"github.com/cgi-fr/pimo/pkg/remove"
	"github.com/cgi-fr/pimo/pkg/replacement"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/cgi-fr/pimo/pkg/weightedchoice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version   string
	commit    string
	buildDate string
	builtBy   string

	verbosity    string
	jsonlog      bool
	iteration    int
	emptyInput   bool
	maskingFile  string
	cachesToDump map[string]string
	cachesToLoad map[string]string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "pimo",
		Short: "Command line to mask data from jsonlines",
		Long:  `Pimo is a tool to mask private data contained in jsonlines by using masking configurations`,
		Version: fmt.Sprintf(`%v (commit=%v date=%v by=%v)
Copyright (C) 2021 CGI France
License GPLv3: GNU GPL version 3 <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.`, version, commit, buildDate, builtBy),
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", "error", "set level of log verbosity, default is error : none (0), error (1), warn (2), info (3), debug (4), trace (5)")
	rootCmd.PersistentFlags().BoolVar(&jsonlog, "log-json", false, "output logs in JSON format")
	rootCmd.PersistentFlags().IntVarP(&iteration, "repeat", "r", 1, "number of iteration to mask each input")
	rootCmd.PersistentFlags().BoolVar(&emptyInput, "empty-input", false, "generate data without any input, to use with repeat flag")
	rootCmd.PersistentFlags().StringVarP(&maskingFile, "config", "c", "masking.yml", "name and location of the masking-config file")
	rootCmd.PersistentFlags().StringToStringVar(&cachesToDump, "dump-cache", map[string]string{}, "path for dumping cache into file")
	rootCmd.PersistentFlags().StringToStringVar(&cachesToLoad, "load-cache", map[string]string{}, "path for loading cache from file")

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}

func run() {
	initLog()

	log.Info().Msg("Start PIMO")

	var source model.Source
	if emptyInput {
		source = model.NewSourceFromSlice([]model.Dictionary{{}})
	} else {
		source = jsonline.NewSource(os.Stdin)
	}
	pipeline := model.NewPipeline(source).
		Process(model.NewRepeaterProcess(iteration))

	var (
		err    error
		caches map[string]model.Cache
	)

	model.InjectMaskContextFactories(injectMaskContextFactories())
	model.InjectMaskFactories(injectMaskFactories())

	pdef, err := model.LoadPipelineDefinitionFromYAML(maskingFile)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		log.Warn().Int("return", 1).Msg("End PIMO")
		os.Exit(1)
	}

	pipeline, caches, err = model.BuildPipeline(pipeline, pdef, nil)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		log.Error().Err(err).Msg("Cannot build pipeline")
		log.Warn().Int("return", 1).Msg("End PIMO")
		os.Exit(1)
	}

	for name, path := range cachesToLoad {
		cache, ok := caches[name]
		if !ok {
			log.Error().Str("cache-name", name).Msg("Cache not found")
			log.Warn().Int("return", 2).Msg("End PIMO")
			os.Exit(2)
		}
		err = pimo.LoadCache(name, cache, path)
		if err != nil {
			log.Err(err).Str("cache-name", name).Str("cache-path", path).Msg("Cannot load cache")
			log.Warn().Int("return", 3).Msg("End PIMO")
			os.Exit(3)
		}
	}

	err = pipeline.AddSink(jsonline.NewSink(os.Stdout)).Run()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		log.Err(err).Msg("Pipeline didn't complete run")
		log.Warn().Int("return", 4).Msg("End PIMO")
		os.Exit(4)
	}

	for name, path := range cachesToDump {
		cache, ok := caches[name]
		if !ok {
			log.Error().Str("cache-name", name).Msg("Cache not found")
			log.Warn().Int("return", 2).Msg("End PIMO")
			os.Exit(2)
		}
		err = pimo.DumpCache(name, cache, path)
		if err != nil {
			log.Err(err).Str("cache-name", name).Str("cache-path", path).Msg("Cannot dump cache")
			log.Warn().Int("return", 3).Msg("End PIMO")
			os.Exit(3)
		}
	}

	log.Info().Int("return", 0).Msg("End PIMO")
	os.Exit(0)
}

func injectMaskContextFactories() []model.MaskContextFactory {
	return []model.MaskContextFactory{
		fluxuri.Factory,
		add.Factory,
		remove.Factory,
		pipe.Factory,
	}
}

func injectMaskFactories() []model.MaskFactory {
	return []model.MaskFactory{
		constant.Factory,
		command.Factory,
		randomlist.Factory,
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
	}
}

func initLog() {
	var logger zerolog.Logger
	if jsonlog {
		logger = zerolog.New(os.Stderr) // .With().Timestamp().Caller().Logger()
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) // .With().Timestamp().Caller().Logger()
	}

	over.New(logger)
	over.MDC().Set("config", maskingFile)
	if iteration > 1 {
		over.MDC().Set("repeat", iteration)
	}
	if emptyInput {
		over.MDC().Set("empty-input", emptyInput)
	}
	if len(cachesToDump) > 0 {
		over.MDC().Set("dump-cache", cachesToDump)
	}
	if len(cachesToLoad) > 0 {
		over.MDC().Set("load-cache", cachesToLoad)
	}
	over.SetGlobalFields([]string{"config"})

	switch verbosity {
	case "trace", "5":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Info().Msg("Logger level set to trace")
	case "debug", "4":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("Logger level set to debug")
	case "info", "3":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Logger level set to info")
	case "warn", "2":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		log.Info().Msg("Logger level set to warn")
	case "error", "1":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		log.Info().Msg("Logger level set to error")
	default:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
}
