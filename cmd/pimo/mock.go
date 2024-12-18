package main

import (
	"net/http"
	"net/url"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/internal/app/pimo/mock"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// pimo --serve ":8080" --mask 'name=[{add: true}, {randomChoiceInUri: "pimo://nameFR"}]'
// pimo -v5 mock -p :8081 http://localhost:8080

func setupMockCommand(rootCmd *cobra.Command) {
	mockAddr := ":8080"

	mockCmd := &cobra.Command{
		Use: "mock",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()
			if maxBufferCapacity > 0 {
				uri.MaxCapacityForEachLine = maxBufferCapacity * 1024
			}

			over.MDC().Set("config", mockConfigFile)
			backendURL := args[0]
			var globalSeed *int64
			if cmd.Flags().Changed("seed") {
				globalSeed = &seedValue
			}
			runMockCommand(backendURL, mockAddr, mockConfigFile, globalSeed)
		},
		Args: cobra.ExactArgs(1),
	}

	mockCmd.PersistentFlags().StringVarP(&mockAddr, "port", "p", mockAddr, "address and port number")
	addFlag(mockCmd, flagBufferSize)
	// addFlag(mockCmd, flagCatchErrors) // could use
	addFlag(mockCmd, flagConfigRoute)
	addFlag(mockCmd, flagCachesToDump)
	addFlag(mockCmd, flagCachesToLoad)
	// addFlag(mockCmd, flagProfiling) // could use
	addFlag(mockCmd, flagSeed)

	rootCmd.AddCommand(mockCmd)
}

func runMockCommand(backendURL, mockAddr, configFile string, globalSeed *int64) {
	pimo.InjectMasks()
	statistics.Reset()

	log.Info().Str("config", configFile).Msgf("Started mock for %s", backendURL)

	cfg, err := mock.LoadConfigFromFile(configFile)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load config from %s", configFile)
	}

	backend, err := url.Parse(backendURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse backend URL")
	}

	ctx, err := cfg.Build(backend, globalSeed)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to build routes from %s", configFile)
	}

	http.HandleFunc("/{path...}", func(w http.ResponseWriter, r *http.Request) {
		_, err := ctx.Process(w, r)
		if err != nil {
			log.Error().AnErr("error", err).Msg("")
		}
	})

	log.Err(http.ListenAndServe(mockAddr, nil)).Msgf("End mock for %s", backendURL) //nolint:gosec
}
