package mock

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
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
	regex    *regexp.Regexp
	request  *Processor
	response *Processor
}

func NewContextRoute(route Route, request, response *Processor) ContextRoute {
	pre := regexp.MustCompile(`\{([^\}]+)\}`)

	return ContextRoute{
		route:    route,
		regex:    regexp.MustCompile("^" + pre.ReplaceAllString(route.Path, `(?P<$1>[^/]*)`) + "$"),
		request:  request,
		response: response,
	}
}

func (ctx Context) Process(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	requestPipeline, responsePipeline, captures := ctx.match(r)

	request, err := NewRequestDict(r)
	if err != nil {
		return nil, err
	}

	request.SetCaptures(captures)

	if requestPipeline != nil {
		log.Info().
			Str("method", request.Method()).
			Str("path", request.URLPath()).
			Str("protocol", request.Protocol()).
			Msg("Request intercepted")

		var origin model.Dictionary
		if log.Trace().Enabled() {
			origin = request.Copy()
		}

		if err := requestPipeline.Process(request.Dictionary); err != nil {
			return nil, err
		}

		log.Trace().RawJSON("request", []byte(origin.UnpackAsDict().String())).Msg("Origin")
		log.Trace().RawJSON("request", []byte(request.UnpackAsDict().String())).Msg("Masked")
	}

	r, err = request.ToRequest()
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

		log.Info().
			Int("status", response.Status()).
			Str("protocol", response.Protocol()).
			Msg("Response intercepted")

		var origin model.Dictionary
		if log.Trace().Enabled() {
			origin = response.Copy()
		}

		if err := responsePipeline.Process(response.Dictionary); err != nil {
			return nil, err
		}

		log.Trace().RawJSON("response", []byte(origin.UnpackAsDict().String())).Msg("Origin")
		log.Trace().RawJSON("response", []byte(response.UnpackAsDict().String())).Msg("Masked")

		resp, err = response.ToResponse()
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
		if match, captures := route.match(r); match {
			return route.request, route.response, captures
		}
	}

	return nil, nil, model.NewDictionary()
}

func (ctxr ContextRoute) match(r *http.Request) (bool, model.Dictionary) {
	if strings.TrimSpace(ctxr.route.Method) != "" && strings.TrimSpace(ctxr.route.Method) != r.Method {
		return false, model.NewDictionary()
	}

	if submatches := ctxr.regex.FindStringSubmatch(r.URL.Path); submatches != nil {
		captures := model.NewDictionary()
		names := ctxr.regex.SubexpNames()
		for i, submatch := range submatches {
			captures.Set(names[i], submatch)
		}
		captures.Delete("")
		return true, captures
	}

	return false, model.NewDictionary()
}
