package main

import (
	"net/http"
	"net/url"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/internal/app/pimo/mock"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// pimo --serve ":8080" --mask 'name=[{add: true}, {randomChoiceInUri: "pimo://nameFR"}]'
// pimo -v5 mock -p :8081 http://localhost:8080

func setupMockCommand(rootCmd *cobra.Command) {
	mockAddr := ":8080"
	mockConfigFile := "routes.yaml"

	mockCmd := &cobra.Command{
		Use: "mock",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()
			over.MDC().Set("config", mockConfigFile)
			backendURL := args[0]
			runMockCommand(backendURL, mockAddr, mockConfigFile)
		},
		Args: cobra.ExactArgs(1),
	}

	mockCmd.PersistentFlags().StringVarP(&mockAddr, "port", "p", mockAddr, "address and port number")
	mockCmd.PersistentFlags().StringVarP(&mockConfigFile, "config", "c", mockConfigFile, "name and location of the routes config file")

	rootCmd.AddCommand(mockCmd)
}

func runMockCommand(backendURL, mockAddr, configFile string) {
	pimo.InjectMasks()

	log.Info().Str("config", configFile).Msgf("Started mock for %s", backendURL)

	cfg, err := mock.LoadConfigFromFile(configFile)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load config from %s", configFile)
	}

	backend, err := url.Parse(backendURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse backend URL")
	}

	ctx, err := cfg.Build(backend)
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
