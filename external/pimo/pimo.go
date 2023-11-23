package pimo

import (
	"errors"
	"strings"

	internal "github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/configuration"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

func Play(yaml string, data string) (string, error) {
	config := configuration.Config{
		EmptyInput:       false,
		RepeatUntil:      "",
		RepeatWhile:      "",
		Iteration:        1,
		SkipLineOnError:  false,
		SkipFieldOnError: false,
		CachesToDump:     map[string]string{},
		CachesToLoad:     map[string]string{},
	}

	pdef, pipe_err := model.LoadPipelineDefinitionFromYAML([]byte(yaml))
	if pipe_err != nil {
		log.Err(pipe_err).Msg("Cannot load pipeline definition")
		return "", errors.New("Cannot load pipeline definition")
	}

	input, json_err := jsonline.JSONToDictionary([]byte(data))
	if json_err != nil {
		log.Err(json_err).Msg("Cannot load pipeline definition")
		return "", json_err
	}

	config.SingleInput = &input
	context := configuration.NewContext(pdef)

	if ctx_err := context.Configure(config, internal.UpdateContext, internal.InjectTemplateFuncs, internal.InjectMaskFactories, internal.InjectMaskContextFactories); ctx_err != nil {
		log.Err(ctx_err).Msg("Cannot configure pipeline")
		return "", ctx_err
	}

	output := &strings.Builder{}

	stats, err := context.Execute(output, internal.LoadCache, internal.DumpCache)
	if err != nil {
		log.Err(err).Msg("Cannot execute pipeline")
	}

	log.Info().Interface("stats", stats).Msg("Input masked")
	return output.String(), err
}
