//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"os"
	"strings"
	"syscall/js"

	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/flow"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

func play(yaml string, data string) (result interface{}, err error) {
	config := pimo.Config{
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
		return
	}

	input, json_err := jsonline.JSONToDictionary([]byte(data))
	if json_err != nil {
		log.Err(json_err).Msg("Cannot load pipeline definition")
		return nil, json_err
	}

	config.SingleInput = &input
	context := pimo.NewContext(pdef)

	if ctx_err := context.Configure(config); ctx_err != nil {
		log.Err(ctx_err).Msg("Cannot configure pipeline")
		return nil, ctx_err
	}

	output := &strings.Builder{}

	stats, err := context.Execute(output)
	if err != nil {
		log.Err(err).Msg("Cannot execute pipeline")
	}

	log.Info().Interface("stats", stats).Msg("Input masked")
	return output.String(), err
}

func main() {
	// Environement variable for FF1 encryption
	os.Setenv("FF1_ENCRYPTION_KEY", "LW4gInNlY3JldDEyc2VjcmV0MTIiIA0K")

	pimo := map[string]interface{}{}

	// JS Constructor
	errorConstructor := js.Global().Get("Error")
	promiseConstructor := js.Global().Get("Promise")

	pimo["play"] = js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		yaml := fmt.Sprintf("%v", inputs[0]) // masking
		data := fmt.Sprintf("%v", inputs[1]) // data
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				data, err := play(yaml, data)
				if err != nil {
					// err should be an instance of `error`, eg `errors.New("some error")`
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
				} else {
					resolve.Invoke(js.ValueOf(data))
				}
			}()

			return nil
		})

		return promiseConstructor.New(handler)
	})

	pimo["flow"] = js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		yaml := fmt.Sprintf("%v", inputs[0]) // masking
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				pdef, err := model.LoadPipelineDefinitionFromYAML([]byte(yaml))
				if err != nil {
					log.Err(err).Msg("Cannot load pipeline definition")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				data, err := flow.Export(pdef)
				if err != nil {
					// err should be an instance of `error`, eg `errors.New("some error")`
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
				} else {
					resolve.Invoke(js.ValueOf(data))
				}
			}()

			return nil
		})

		return promiseConstructor.New(handler)
	})

	// Expose pimo object
	js.Global().Set("pimo", pimo)

	select {} // sleep forever
}
