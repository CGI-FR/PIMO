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
	"runtime"
	"strings"
	"time"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/flow"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/mattn/go-isatty"
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

	verbosity        string
	debug            bool
	jsonlog          bool
	colormode        string
	iteration        int
	emptyInput       bool
	maskingFile      string
	cachesToDump     map[string]string
	cachesToLoad     map[string]string
	skipLineOnError  bool
	skipFieldOnError bool
	seedValue        int64
	maskingOneLiner  []string
	repeatUntil      string
	repeatWhile      string
)

func main() {
	rootCmd := &cobra.Command{
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

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", "error", "set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "add debug information to logs (very slow)")
	rootCmd.PersistentFlags().BoolVar(&jsonlog, "log-json", false, "output logs in JSON format")
	rootCmd.PersistentFlags().StringVar(&colormode, "color", "auto", "use colors in log outputs : yes, no or auto")
	rootCmd.PersistentFlags().IntVarP(&iteration, "repeat", "r", 1, "number of iteration to mask each input")
	rootCmd.PersistentFlags().BoolVar(&emptyInput, "empty-input", false, "generate data without any input, to use with repeat flag")
	rootCmd.PersistentFlags().StringVarP(&maskingFile, "config", "c", "masking.yml", "name and location of the masking-config file")
	rootCmd.PersistentFlags().StringToStringVar(&cachesToDump, "dump-cache", map[string]string{}, "path for dumping cache into file")
	rootCmd.PersistentFlags().StringToStringVar(&cachesToLoad, "load-cache", map[string]string{}, "path for loading cache from file")
	rootCmd.PersistentFlags().BoolVar(&skipLineOnError, "skip-line-on-error", false, "skip a line if an error occurs while masking a field")
	rootCmd.PersistentFlags().BoolVar(&skipFieldOnError, "skip-field-on-error", false, "remove a field if an error occurs while masking this field")
	rootCmd.PersistentFlags().Int64VarP(&seedValue, "seed", "s", -1, "set seed")
	rootCmd.PersistentFlags().StringArrayVarP(&maskingOneLiner, "mask", "m", []string{}, "one liner masking")
	rootCmd.PersistentFlags().StringVar(&repeatUntil, "repeat-until", "", "mask each input repeatedly until the given condition is met")
	rootCmd.PersistentFlags().StringVar(&repeatWhile, "repeat-while", "", "mask each input repeatedly while the given condition is met")

	rootCmd.AddCommand(&cobra.Command{
		Use: "jsonschema",
		Run: func(cmd *cobra.Command, args []string) {
			jsonschema, err := pimo.GetJsonSchema()
			if err != nil {
				os.Exit(8)
			}
			fmt.Println(jsonschema)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "flow",
		Run: func(cmd *cobra.Command, args []string) {
			pdef, err := model.LoadPipelineDefinitionFromFile(maskingFile)
			if err != nil {
				log.Err(err).Msg("Cannot load pipeline definition from file")
				log.Warn().Int("return", 1).Msg("End PIMO")
				os.Exit(1)
			}
			flow, err := flow.Export(pdef)
			if err != nil {
				os.Exit(9)
			}
			fmt.Println(flow)
		},
	})

	playPort := 3010
	playSecure := false
	playCmd := &cobra.Command{
		Use: "play",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()

			router := pimo.Play(playSecure)
			port := fmt.Sprintf("0.0.0.0:%d", playPort)

			if err := router.Start(port); err != nil {
				os.Exit(8)
			}
		},
	}
	playCmd.PersistentFlags().IntVarP(&playPort, "port", "p", 3010, "port number")
	playCmd.PersistentFlags().BoolVarP(&playSecure, "secure", "s", false, "enable security features (use this flag if PIMO Play is publicly exposed)")
	rootCmd.AddCommand(playCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}

func run() {
	initLog()

	config := pimo.Config{
		EmptyInput:       emptyInput,
		RepeatUntil:      repeatUntil,
		RepeatWhile:      repeatWhile,
		Iteration:        iteration,
		SkipLineOnError:  skipLineOnError,
		SkipFieldOnError: skipFieldOnError,
		CachesToDump:     cachesToDump,
		CachesToLoad:     cachesToLoad,
	}

	var pdef model.Definition
	var err error
	if len(maskingOneLiner) > 0 {
		pdef, err = model.LoadPipelineDefintionFromOneLiner(maskingOneLiner)
	} else {
		pdef, err = model.LoadPipelineDefinitionFromFile(maskingFile)
	}

	switch {
	case seedValue != -1:
		pdef.Seed = seedValue
	case seedValue == 0:
		pdef.Seed = time.Now().UnixNano()
	}

	if err != nil {
		log.Err(err).Msg("Cannot load pipeline definition from file")
		log.Warn().Int("return", 1).Msg("End PIMO")
		os.Exit(1)
	}

	ctx := pimo.NewContext(pdef)

	if err := ctx.Configure(config); err != nil {
		log.Err(err).Msg("Cannot configure pipeline")
		log.Warn().Int("return", 1).Msg("End PIMO")
		os.Exit(1)
	}

	stats, err := ctx.Execute(os.Stdout)
	if err != nil {
		log.Err(err).Msg("Cannot execute pipeline")
		log.Warn().Int("return", stats.GetErrorCode()).Msg("End PIMO")
		os.Exit(stats.GetErrorCode())
	}

	log.Info().RawJSON("stats", stats.ToJSON()).Int("return", 0).Msg("End PIMO")
	os.Exit(0)
}

func initLog() {
	color := false
	switch strings.ToLower(colormode) {
	case "auto":
		if isatty.IsTerminal(os.Stdout.Fd()) && runtime.GOOS != "windows" {
			color = true
		}
	case "yes", "true", "1", "on", "enable":
		color = true
	}

	var logger zerolog.Logger
	if jsonlog {
		logger = zerolog.New(os.Stderr) // .With().Caller().Logger()
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: !color}) // .With().Caller().Logger()
	}
	if debug {
		logger = logger.With().Caller().Logger()
	}

	over.New(logger)

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
	case "error", "1":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
	over.MDC().Set("config", maskingFile)
	over.SetGlobalFields([]string{"config"})
}
