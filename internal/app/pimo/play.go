package pimo

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

//go:embed client
var content embed.FS

func Play(enableSecurity bool) *echo.Echo {
	router := echo.New()

	router.Use(middleware.CORS())
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} : method=${method}, uri=${uri}, status=${status}, host=${host}, origin=${user_agent}\n" +
			" err=${error}\n",
	}))

	if enableSecurity {
		router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Set("enableSecurity", true)
				return next(c)
			}
		})
	}

	router.GET("/*", echo.WrapHandler(handleClient()))
	router.POST("/play", play)

	return router
}

// play binds client interface data entry to be processed and returns the transformed json to client
func play(ctx echo.Context) error {
	config := Config{
		EmptyInput:       false,
		RepeatUntil:      "",
		RepeatWhile:      "",
		Iteration:        1,
		SkipLineOnError:  false,
		SkipFieldOnError: false,
		CachesToDump:     map[string]string{},
		CachesToLoad:     map[string]string{},
	}

	var dataInput map[string]interface{}

	err := ctx.Bind(&dataInput)
	if err != nil {
		log.Err(err).Msg("Failed to bind client data")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	// yaml := ctx.FormValue("masking")
	// data := ctx.FormValue("data")
	yaml := fmt.Sprintf("%v", dataInput["masking"])
	data := fmt.Sprintf("%v", dataInput["data"])
	jsonl, _ := strconv.ParseBool(fmt.Sprintf("%v", dataInput["jsonl"]))

	pdef, err := model.LoadPipelineDefinitionFromYAML([]byte(yaml))
	if err != nil {
		log.Err(err).Msg("Cannot load pipeline definition")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	inputs := []model.Dictionary{}
	if !jsonl {
		input, err := jsonline.JSONToDictionary([]byte(data))
		if err != nil {
			log.Err(err).Msg("Cannot load pipeline definition")
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		if ctx.Get("enableSecurity") == true {
			if err := checkSecurityRequirements(pdef); err != nil {
				log.Err(err).Msg("Forbidden request")
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
		}
		inputs = append(inputs, input)

	} else {
		for _, line := range strings.Split(strings.TrimRight(data, "\n"), "\n") {
			input, err := jsonline.JSONToDictionary([]byte(line))
			if err != nil {
				log.Err(err).Msg("Cannot load pipeline definition")
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			if ctx.Get("enableSecurity") == true {
				if err := checkSecurityRequirements(pdef); err != nil {
					log.Err(err).Msg("Forbidden request")
					return ctx.String(http.StatusInternalServerError, err.Error())
				}
			}

			inputs = append(inputs, input)
		}
	}

	config.MultipleInput = inputs
	context := NewContext(pdef)

	if err := context.Configure(config); err != nil {
		log.Err(err).Msg("Cannot configure pipeline")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	result := &strings.Builder{}

	stats, err := context.Execute(result)
	if err != nil {
		log.Err(err).Msg("Cannot execute pipeline")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	log.Info().Interface("stats", stats).Msg("Input masked")
	return ctx.JSONBlob(http.StatusOK, []byte(result.String()))
}

func checkSecurityRequirements(pdef model.Definition) error {
	for _, mask := range pdef.Masking {
		// usage of command is not allowed with pimo play
		if len(mask.Mask.Command) > 0 {
			return fmt.Errorf("Usage of `command` mask is forbidden")
		}

		// usage of file scheme is not allowed
		if err := checkUriScheme(mask.Mask.FluxURI); err != nil {
			return err
		}

		if err := checkUriScheme(mask.Mask.HashInURI); err != nil {
			return err
		}

		if err := checkUriScheme(mask.Mask.Markov.Sample); err != nil {
			return err
		}

		if err := checkUriScheme(mask.Mask.RandomChoiceInURI); err != nil {
			return err
		}
	}

	return nil
}

func checkUriScheme(uri string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return nil
	}

	if u.Scheme == "file" {
		return fmt.Errorf("Usage of `file` scheme is forbidden")
	}

	return nil
}

func handleClient() http.Handler {
	fSys, err := fs.Sub(content, "client")
	if err != nil {
		log.Err(err).Msg("cannot find subtree")
	}
	return http.FileServer(http.FS(fSys))
}
