package pimo

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"strings"

	"github.com/cgi-fr/pimo/pkg/flow"
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
	router.POST("/flow", flowchart)

	return router
}

// play binds client interface data entry to be processed and returns the transformed json to client
func play(ctx echo.Context) error {
	// dont't panic
	defer func() error {
		if r := recover(); r != nil {
			log.Error().AnErr("panic", r.(error)).Msg("Recovering from panic in play.")
			return ctx.String(http.StatusInternalServerError, r.(error).Error())
		}
		return nil
	}() //nolint:errcheck

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

	pdef, err := model.LoadPipelineDefinitionFromYAML([]byte(yaml))
	if err != nil {
		log.Err(err).Msg("Cannot load pipeline definition")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

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

	config.SingleInput = &input
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

func flowchart(ctx echo.Context) error {
	// dont't panic
	defer func() error {
		if r := recover(); r != nil {
			log.Error().AnErr("panic", r.(error)).Msg("Recovering from panic in play.")
			return ctx.String(http.StatusInternalServerError, r.(error).Error())
		}
		return nil
	}() //nolint:errcheck

	var dataInput map[string]interface{}

	err := ctx.Bind(&dataInput)
	if err != nil {
		log.Err(err).Msg("Failed to bind client data")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	yaml := fmt.Sprintf("%v", dataInput["masking"])

	pdef, err := model.LoadPipelineDefinitionFromYAML([]byte(yaml))
	if err != nil {
		log.Err(err).Msg("Cannot load pipeline definition")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	flowchart, err := flow.Export(pdef)
	if err != nil {
		log.Err(err).Msg("Cannot generate flowchart")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	log.Info().Msg("Mask flowchart generated")
	return ctx.JSONBlob(http.StatusOK, []byte(flowchart))
}

func checkSecurityRequirements(pdef model.Definition) error {
	for _, mask := range pdef.Masking {
		if err := checkMask(mask.Mask); err != nil {
			return err
		}

		for _, m := range mask.Masks {
			if err := checkMask(m); err != nil {
				return err
			}
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

func checkMask(mask model.MaskType) error {
	// usage of command is not allowed with pimo play
	if len(mask.Command) > 0 {
		return fmt.Errorf("Usage of `command` mask is forbidden")
	}

	// usage of MaskCommand in templates is not allowed
	if len(mask.Template) > 0 {
		if strings.Contains(mask.Template, "MaskCommand") {
			return fmt.Errorf("Usage of `command` mask is forbidden")
		}
	}

	// usage of MaskCommand in templates is not allowed
	if tmplstr, ok := mask.Add.(string); ok {
		if strings.Contains(tmplstr, "MaskCommand") {
			return fmt.Errorf("Usage of `command` mask is forbidden")
		}
	}

	// usage of MaskCommand in templates is not allowed
	if tmplstr, ok := mask.AddTransient.(string); ok {
		if strings.Contains(tmplstr, "MaskCommand") {
			return fmt.Errorf("Usage of `command` mask is forbidden")
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
