package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func setupMockCommand(rootCmd *cobra.Command) {
	mockAddr := ":8080"
	mockConfigFile := "routes.yaml"

	mockCmd := &cobra.Command{
		Use: "mock",
		Run: func(cmd *cobra.Command, args []string) {
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
	log.Info().Msgf("started mock for %s", backendURL)

	client := http.DefaultClient

	http.HandleFunc("/{path...}", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(r.Method, fmt.Sprintf("%s%s", backendURL, r.URL.Path), r.Body)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
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

	log.Err(http.ListenAndServe(mockAddr, nil)).Msgf("end mock for %s", backendURL)
}
