package mock

import (
	"io"
	"net/http"
	"net/url"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

type Context struct {
	client  *http.Client
	backend *url.URL
	routes  []ContextRoute
	caches  map[string]model.Cache
}

type ContextRoute struct {
	route    Route
	request  model.SinkedPipeline
	response model.SinkedPipeline
}

func (ctx Context) Process(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	requestPipeline, responsePipeline, captures := ctx.match(r)

	log.Warn().Interface("captures", captures).Msg("Captures not implemented yet")

	if requestPipeline != nil {
		request, err := NewRequestDict(r)
		if err != nil {
			return nil, err
		}

		println(request.UnpackAsDict().String())

		log.Info().
			Str("method", request.Method()).
			Str("path", request.URLPath()).
			Str("protocol", request.Protocol()).
			Msg("Request intercepted")

		// TODO : call pipeline

		r, err = ToRequest(request.Dictionary)
		if err != nil {
			return nil, err
		}

		r.URL.Scheme = ctx.backend.Scheme
		r.URL.User = ctx.backend.User
		r.URL.Host = ctx.backend.Host
	}

	resp, err := ctx.client.Do(r)
	if err != nil {
		return resp, err
	}

	if responsePipeline != nil {
		response, err := NewResponseDict(resp)
		if err != nil {
			return resp, err
		}

		println(response.UnpackAsDict().String())

		log.Info().
			Str("status", response.Status()).
			Str("protocol", response.Protocol()).
			Msg("Response intercepted")

		// TODO : call pipeline

		resp, err = ToResponse(response.Dictionary)
		if err != nil {
			return resp, err
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

		// trailers ?
	}

	return resp, nil
}

func (ctx Context) match(r *http.Request) (model.Pipeline, model.Pipeline, model.Dictionary) {
	return nil, nil, model.NewDictionary()
}
