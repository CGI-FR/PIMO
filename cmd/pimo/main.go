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
	"bytes"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/flow"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/labstack/echo/v4"
	"github.com/mattn/go-isatty"
	"github.com/pkg/profile"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Provisioned by ldflags
//
//nolint:gochecknoglobals
var (
	version   string
	commit    string
	buildDate string
	builtBy   string

	verbosity string
	debug     bool
	jsonlog   bool
	colormode string

	parquetInput  string
	parquetOutput string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "pimo",
		Short: "Command line to mask data from jsonlines",
		Long: `Pimo is a tool to mask private data contained in jsonlines by using masking configurations

Environment Variables:
  PIMO_STATS_URL      The URL where statistics will be sent
  PIMO_STATS_TEMPLATE The template string to format statistics`,

		Version: fmt.Sprintf(`%v (commit=%v date=%v by=%v)
Copyright (C) 2021 CGI France
License GPLv3: GNU GPL version 3 <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.`, version, commit, buildDate, builtBy),
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", "error", "set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "add debug information to logs (very slow)")
	rootCmd.PersistentFlags().BoolVar(&jsonlog, "log-json", false, "output logs in JSON format")
	rootCmd.PersistentFlags().StringVar(&colormode, "color", "auto", "use colors in log outputs : yes, no or auto")

	addFlag(rootCmd, flagBufferSize)
	addFlag(rootCmd, flagCatchErrors)
	addFlag(rootCmd, flagConfigMasking)
	addFlag(rootCmd, flagCachesToDump)
	addFlag(rootCmd, flagCachesToLoad)
	addFlag(rootCmd, flagEmptyInput)
	addFlag(rootCmd, flagMaskOneLiner)
	addFlag(rootCmd, flagProfiling)
	addFlag(rootCmd, flagRepeat)
	addFlag(rootCmd, flagRepeatUntil)
	addFlag(rootCmd, flagRepeatWhile)
	addFlag(rootCmd, flagSeed)
	addFlag(rootCmd, flagServe)
	addFlag(rootCmd, flagSkipFieldOnError)
	addFlag(rootCmd, flagSkipLineOnError)
	addFlag(rootCmd, flagSkipLogFile)
	addFlag(rootCmd, flagStatsDestination)
	addFlag(rootCmd, flagStatsTemplate)

	rootCmd.AddCommand(&cobra.Command{
		Use: "jsonschema",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()
			jsonschema, err := pimo.GetJsonSchema()
			if err != nil {
				os.Exit(8)
			}
			fmt.Println(jsonschema)
		},
	})
	// Add command for XML transformer
	xmlCmd := &cobra.Command{
		Use:   "xml",
		Short: "Parsing and masking XML file",
		Run: func(cmd *cobra.Command, _ []string) {
			initLog()
			if maxBufferCapacity > 0 {
				uri.MaxCapacityForEachLine = maxBufferCapacity * 1024
			}
			if len(catchErrors) > 0 {
				skipLineOnError = true
				skipLogFile = catchErrors
			}
			config := pimo.Config{
				EmptyInput:       emptyInput,
				RepeatUntil:      repeatUntil,
				RepeatWhile:      repeatWhile,
				Iteration:        iteration,
				SkipLineOnError:  skipLineOnError,
				SkipFieldOnError: skipFieldOnError,
				SkipLogFile:      skipLogFile,
				CachesToDump:     cachesToDump,
				CachesToLoad:     cachesToLoad,
				XMLCallback:      true,
			}

			parser := pimo.ParseXML(cmd.InOrStdin(), cmd.OutOrStdout())
			// Map the command line balise name to fit the masking configuration
			for elementName, mask := range xmlSubscriberName {
				pdef, err := model.LoadPipelineDefinitionFromFile(mask)
				if err != nil {
					fmt.Printf("Error when charging pipeline for %s : %v\n", elementName, err)
					return
				}

				if cmd.Flags().Changed("seed") {
					(&pdef).SetSeed(seedValue)
				}

				ctx := pimo.NewContext(pdef)
				if err := ctx.Configure(config); err != nil {
					log.Err(err).Msg("Cannot configure pipeline")
					log.Warn().Int("return", 1).Msg("End PIMO")
					os.Exit(1)
				}

				parser.RegisterMapCallback(elementName, func(m map[string]string) (map[string]string, error) {
					newList, _ := pimo.XMLCallback(ctx, m)
					return newList, nil
				})
			}
			err := parser.Stream()
			if err != nil {
				log.Err(err).Msg("Error during parsing XML document")
			}
		},
	}
	addFlag(xmlCmd, flagBufferSize)
	addFlag(xmlCmd, flagCatchErrors)
	addFlag(xmlCmd, flagCachesToDump)
	addFlag(xmlCmd, flagCachesToLoad)
	// addFlag(xmlCmd, flagProfiling) //could use
	addFlag(xmlCmd, flagSeed)
	addFlag(xmlCmd, flagSkipFieldOnError)
	addFlag(xmlCmd, flagSkipLineOnError)
	addFlag(xmlCmd, flagSkipLogFile)
	// addFlag(xmlCmd, flagStatsDestination) // could use
	// addFlag(xmlCmd, flagStatsTemplate)    // could use
	addFlag(xmlCmd, flagXMLSubscriberName)
	rootCmd.AddCommand(xmlCmd)

	// Add command for parquet transformer
	parquetCmd := &cobra.Command{
		Use:   "parquet input_parquet_file output_parquet_file",
		Short: "Parsing and masking a parquet file",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			initLog()
			if len(catchErrors) > 0 {
				skipLineOnError = true
				skipLogFile = catchErrors
			}
			parquetInput = args[0]
			parquetOutput = args[1]

			run(cmd)
		},
	}
	addFlag(parquetCmd, flagBufferSize)
	addFlag(parquetCmd, flagCatchErrors)
	addFlag(parquetCmd, flagConfigMasking)
	addFlag(parquetCmd, flagCachesToDump)
	addFlag(parquetCmd, flagCachesToLoad)
	addFlag(parquetCmd, flagMaskOneLiner)
	addFlag(parquetCmd, flagProfiling)
	addFlag(parquetCmd, flagSeed)
	addFlag(parquetCmd, flagSkipFieldOnError)
	addFlag(parquetCmd, flagSkipLineOnError)
	addFlag(parquetCmd, flagSkipLogFile)
	addFlag(parquetCmd, flagStatsDestination)
	addFlag(parquetCmd, flagStatsTemplate)
	rootCmd.AddCommand(parquetCmd)

	flowCmd := &cobra.Command{
		Use: "flow",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()
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
	}
	rootCmd.AddCommand(flowCmd)

	playPort := 3010
	playSecure := false
	playCmd := &cobra.Command{
		Use: "play",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()

			if maxBufferCapacity > 0 {
				uri.MaxCapacityForEachLine = maxBufferCapacity * 1024
			}

			router := pimo.Play(playSecure)
			port := fmt.Sprintf("0.0.0.0:%d", playPort)

			if err := router.Start(port); err != nil {
				os.Exit(8)
			}
		},
	}
	playCmd.PersistentFlags().IntVarP(&playPort, "port", "p", 3010, "port number")
	playCmd.PersistentFlags().BoolVarP(&playSecure, "secure", "s", false, "enable security features (use this flag if PIMO Play is publicly exposed)")
	addFlag(playCmd, flagBufferSize)
	rootCmd.AddCommand(playCmd)

	setupMockCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}

func run(cmd *cobra.Command) {
	initLog()

	if len(catchErrors) > 0 {
		skipLineOnError = true
		skipLogFile = catchErrors
	}

	if maxBufferCapacity > 0 {
		uri.MaxCapacityForEachLine = maxBufferCapacity * 1024
	}

	config := pimo.Config{
		EmptyInput:       emptyInput,
		RepeatUntil:      repeatUntil,
		RepeatWhile:      repeatWhile,
		Iteration:        iteration,
		SkipLineOnError:  skipLineOnError,
		SkipFieldOnError: skipFieldOnError,
		SkipLogFile:      skipLogFile,
		CachesToDump:     cachesToDump,
		CachesToLoad:     cachesToLoad,
		XMLCallback:      len(serve) > 0,
		ParquetInput:     parquetInput,
		ParquetOutput:    parquetOutput,
	}

	var pdef model.Definition
	var err error
	if len(maskingOneLiner) > 0 {
		pdef, err = model.LoadPipelineDefintionFromOneLiner(maskingOneLiner)
	} else {
		pdef, err = model.LoadPipelineDefinitionFromFile(maskingFile)
	}

	if cmd.Flags().Changed("seed") {
		(&pdef).SetSeed(seedValue)
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

	var profiler interface{ Stop() }

	if profiling == "cpu" {
		profiler = profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.Quiet)
	} else if profiling == "memory" {
		profiler = profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.Quiet)
	}

	if len(serve) > 0 {
		router := echo.New()
		router.HideBanner = true
		router.GET("/", httpHandler(ctx))
		router.POST("/", httpHandler(ctx))
		if err := router.Start(serve); err != nil {
			log.Err(err).Msg("Failed to start server")
			os.Exit(8)
		}
	} else {
		startTime := time.Now()
		stats, err := ctx.Execute(os.Stdout)
		if err != nil {
			log.Err(err).Msg("Cannot execute pipeline")
			log.Warn().Int("return", stats.GetErrorCode()).Msg("End PIMO")
			os.Exit(stats.GetErrorCode())
		}

		duration := time.Since(startTime)
		statistics.SetDuration(duration)
		dumpStats(stats)
	}

	if profiling == "cpu" || profiling == "memory" {
		profiler.Stop()
	}

	os.Exit(0)
}

func httpHandler(pimo pimo.Context) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		// dont't panic
		defer func() error {
			if r := recover(); r != nil {
				log.Error().AnErr("panic", r.(error)).Msg("Recovering from panic in rest server.")
				return ctx.String(http.StatusInternalServerError, r.(error).Error())
			}
			return nil
		}() //nolint:errcheck

		payload := map[string]any{}

		if err := ctx.Bind(&payload); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		result, err := pimo.ExecuteMap(payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, result)
	}
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
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger() // .With().Caller().Logger()
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

func dumpStats(stats statistics.ExecutionStats) {
	statsToWrite := stats.ToJSON()
	if statsTemplate != "" {
		tmpl, err := template.New("statsTemplate").Parse(statsTemplate)
		if err != nil {
			log.Error().Err(err).Msg(("Error parsing statistics template"))
			os.Exit(1)
		}
		var output bytes.Buffer
		err = tmpl.ExecuteTemplate(&output, "statsTemplate", Stats{Stats: string(stats.ToJSON())})
		if err != nil {
			log.Error().Err(err).Msg("Error adding stats to template")
			os.Exit(1)
		}
		statsToWrite = output.Bytes()
	}
	if statisticsDestination != "" {
		if strings.HasPrefix(statisticsDestination, "http") {
			sendMetrics(statisticsDestination, statsToWrite)
		} else {
			writeMetricsToFile(statisticsDestination, statsToWrite)
		}
	}

	log.Info().RawJSON("stats", stats.ToJSON()).Int("return", 0).Msg("End PIMO")
}

func writeMetricsToFile(statsFile string, statsByte []byte) {
	file, err := os.Create(statsFile)
	if err != nil {
		log.Error().Err(err).Msg("Error generating statistics dump file")
	}
	defer file.Close()

	_, err = file.Write(statsByte)
	if err != nil {
		log.Error().Err(err).Msg("Error writing statistics to dump file")
	}
	log.Info().Msgf("Statistics exported to file %s", file.Name())
}

func sendMetrics(statsDestination string, statsByte []byte) {
	requestBody := bytes.NewBuffer(statsByte)
	// nolint: gosec
	_, err := http.Post(statsDestination, "application/json", requestBody)
	if err != nil {
		log.Error().Err(err).Msgf("An error occurred trying to send metrics to %s", statsDestination)
	}
	log.Info().Msgf("Statistics sent to %s", statsDestination)
}

type Stats struct {
	Stats interface{} `json:"stats"`
}
