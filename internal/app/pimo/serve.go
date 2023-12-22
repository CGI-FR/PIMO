package pimo

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func ServeCommand() *cobra.Command {
	port := 8080

	cmd := &cobra.Command{
		Use:        "serve",
		Aliases:    []string{},
		SuggestFor: []string{},
		Short:      "Start a HTTP Rest service",
		Long:       "TODO",
		Example:    "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			Serve(port)
		},
	}

	cmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "port number")

	return cmd
}

var pimo Context

func Serve(port int) {
	router := echo.New()

	router.Use(middleware.CORS())
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} : method=${method}, uri=${uri}, status=${status}, host=${host}, origin=${user_agent}\n" +
			" err=${error}\n",
	}))

	router.GET("/", emptyInputHandler)

	config := Config{
		EmptyInput:       false,
		RepeatUntil:      "",
		RepeatWhile:      "",
		Iteration:        1,
		SkipLineOnError:  false,
		SkipFieldOnError: false,
		CachesToDump:     map[string]string{},
		CachesToLoad:     map[string]string{},
		XMLCallback:      true,
	}

	pdef, err := model.LoadPipelineDefinitionFromFile("masking.yml")
	if err != nil {
		log.Err(err).Msg("Cannot load pipeline definition from file")
		log.Warn().Int("return", 1).Msg("End PIMO")
		os.Exit(1)
	}

	pimo = NewContext(pdef)

	if err := pimo.Configure(config); err != nil {
		log.Err(err).Msg("Cannot configure pipeline")
		log.Warn().Int("return", 1).Msg("End PIMO")
		os.Exit(1)
	}

	if err := router.Start(fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
		os.Exit(8)
	}
}

func emptyInputHandler(ctx echo.Context) error {
	// dont't panic
	defer func() error {
		if r := recover(); r != nil {
			log.Error().AnErr("panic", r.(error)).Msg("Recovering from panic in rest server.")
			return ctx.String(http.StatusInternalServerError, r.(error).Error())
		}
		return nil
	}() //nolint:errcheck

	result, err := pimo.ExecuteMap(map[string]string{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func payloadInputHandler(ctx echo.Context) error {
	// dont't panic
	defer func() error {
		if r := recover(); r != nil {
			log.Error().AnErr("panic", r.(error)).Msg("Recovering from panic in rest server.")
			return ctx.String(http.StatusInternalServerError, r.(error).Error())
		}
		return nil
	}() //nolint:errcheck

	var payload map[string]string

	if err := ctx.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result, err := pimo.ExecuteMap(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, result)
}
