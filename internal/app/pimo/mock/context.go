package mock

import (
	"io"
	"net/http"
	"net/url"
	"strings"

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
	request  *Processor
	response *Processor
}

func (ctx Context) Process(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	requestPipeline, responsePipeline, captures := ctx.match(r)

	log.Warn().Interface("captures", captures).Msg("Captures not implemented yet")

	request, err := NewRequestDict(r)
	if err != nil {
		return nil, err
	}

	println(request.UnpackAsDict().String())

	if requestPipeline != nil {
		log.Info().
			Str("method", request.Method()).
			Str("path", request.URLPath()).
			Str("protocol", request.Protocol()).
			Msg("Request intercepted")

		request.Dictionary, err = requestPipeline.Process(request.Dictionary)
		if err != nil {
			return nil, err
		}

		println(request.String())
	}

	r, err = ToRequest(request.Dictionary)
	if err != nil {
		return nil, err
	}

	r.URL.Scheme = ctx.backend.Scheme
	r.URL.User = ctx.backend.User
	r.URL.Host = ctx.backend.Host

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
			Int("status", response.Status()).
			Str("protocol", response.Protocol()).
			Msg("Response intercepted")

		dict, err := responsePipeline.Process(response.Dictionary)
		if err != nil {
			return nil, err
		}

		println(dict.String())

		resp, err = ToResponse(dict)
		if err != nil {
			return resp, err
		}
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

	return resp, nil
}

func (ctx Context) match(r *http.Request) (*Processor, *Processor, model.Dictionary) {
	for _, route := range ctx.routes {
		if route.match(r) {
			return route.request, route.response, model.NewDictionary()
		}
	}

	return nil, nil, model.NewDictionary()
}

func (ctxr ContextRoute) match(r *http.Request) bool {
	if strings.TrimSpace(ctxr.route.Method) != "" && strings.TrimSpace(ctxr.route.Method) != r.Method {
		return false
	}

	return ctxr.route.Path == r.URL.Path
}
