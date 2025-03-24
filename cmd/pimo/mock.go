package main

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	over "github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/internal/app/pimo/mock"
	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/uri"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// pimo --serve ":8080" --mask 'name=[{add: true}, {randomChoiceInUri: "pimo://nameFR"}]'
// pimo -v5 mock -p :8081 http://localhost:8080

func askPassword(prompt string) string {
	if term.IsTerminal(int(os.Stdin.Fd())) {
		os.Stdout.Write([]byte(prompt))
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		os.Stdout.Write([]byte("\n"))
		if err != nil {
			os.Exit(1)
		}
		return string(bytePassword)
	}
	return ""
}

func setupMockCommand(rootCmd *cobra.Command) {
	mockAddr := ":8080"
	proxyURL := ""
	proxyAuth := ""
	serverAuth := ""

	mockCmd := &cobra.Command{
		Use:   "mock",
		Short: "Use masking configurations to proxyfy remote HTTP service",
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

			client := http.DefaultClient
			if proxyURL != "" {
				proxyURLParsed, err := url.Parse(proxyURL)
				if err != nil {
					log.Fatal().Err(err).Msg("Failed to parse proxy URL")
				}
				if proxyAuth != "" {
					if !strings.Contains(proxyAuth, ":") {
						password := askPassword("enter password for proxy authentication: ")
						proxyAuth = proxyAuth + ":" + password
					}
					proxyURLParsed.User = url.UserPassword(strings.Split(proxyAuth, ":")[0], strings.Split(proxyAuth, ":")[1])
				}
				client.Transport = &http.Transport{
					Proxy: http.ProxyURL(proxyURLParsed),
				}
			}

			if serverAuth != "" {
				if !strings.Contains(serverAuth, ":") {
					password := askPassword("enter password for server authentication: ")
					proxyAuth = proxyAuth + ":" + password
				}
				client.Transport = &TransportWithAuth{
					wrapped: client.Transport,
					user:    strings.Split(serverAuth, ":")[0],
					pass:    strings.Split(serverAuth, ":")[1],
				}
			}

			runMockCommand(backendURL, mockAddr, mockConfigFile, globalSeed, client)
		},
		Args: cobra.ExactArgs(1),
	}

	mockCmd.PersistentFlags().StringVarP(&mockAddr, "port", "p", mockAddr, "address and port number")
	mockCmd.PersistentFlags().StringVarP(&proxyURL, "proxy", "x", proxyURL, "use the specified proxy to perform backend request")
	mockCmd.PersistentFlags().StringVarP(&proxyAuth, "proxy-user", "U", proxyAuth, "specify user:password to use for proxy authentication")
	mockCmd.PersistentFlags().StringVarP(&serverAuth, "user", "u", serverAuth, "specify user:password to use for server authentication")
	addFlag(mockCmd, flagBufferSize)
	// addFlag(mockCmd, flagCatchErrors) // could use
	addFlag(mockCmd, flagConfigRoute)
	// addFlag(mockCmd, flagCachesToDump) // could use
	addFlag(mockCmd, flagCachesToLoad)
	// addFlag(mockCmd, flagProfiling) // could use
	addFlag(mockCmd, flagSeed)
	// addFlag(mockCmd, flagSkipFieldOnError) // could use
	// addFlag(mockCmd, flagSkipLineOnError)  // could use
	// addFlag(mockCmd, flagSkipLogFile)      // could use
	// addFlag(mockCmd, flagStatsDestination) // could use
	// addFlag(mockCmd, flagStatsTemplate)    // could use

	rootCmd.AddCommand(mockCmd)
}

type TransportWithAuth struct {
	wrapped http.RoundTripper
	user    string
	pass    string
}

func (t *TransportWithAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.user, t.pass)
	if t.wrapped == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.wrapped.RoundTrip(req)
}

func runMockCommand(backendURL, mockAddr, configFile string, globalSeed *int64, client *http.Client) {
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

	ctx, err := cfg.Build(backend, globalSeed, cachesToLoad, client)
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
