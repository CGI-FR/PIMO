package main

import (
	"io"
	"net/http"
	"net/url"

	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func setupMockCommand(rootCmd *cobra.Command) {
	mockAddr := ":8080"
	mockConfigFile := "routes.yaml"

	mockCmd := &cobra.Command{
		Use: "mock",
		Run: func(cmd *cobra.Command, args []string) {
			initLog()
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
	log.Info().Str("config", configFile).Msgf("Started mock for %s", backendURL)

	client := http.DefaultClient

	http.HandleFunc("/{path...}", func(w http.ResponseWriter, r *http.Request) {
		request, err := pimo.NewRequestDict(r)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		log.Info().
			Str("method", request.Method()).
			Str("path", request.URLPath()).
			Str("protocol", request.Protocol()).
			Msg("Request intercepted")

		req, err := pimo.ToRequest(request.Dictionary)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		req.URL, err = url.Parse(backendURL)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		req.URL.Path = request.URLPath()

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		// copy headers
		for name, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}

		w.WriteHeader(resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		if _, err := w.Write(body); err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	log.Err(http.ListenAndServe(mockAddr, nil)).Msgf("End mock for %s", backendURL) //nolint:gosec
}
