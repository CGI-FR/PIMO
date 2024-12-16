package mock

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
)

type ResponseDict struct {
	model.Dictionary
}

func (r ResponseDict) ToResponse() (*http.Response, error) {
	dict := r.Dictionary
	if d, ok := dict.TryUnpackAsDict(); ok {
		dict = d
	}

	var status int
	if m, ok := dict.GetValue(keyStatus); ok {
		if s, ok := m.(int); ok {
			status = s
		}
		if s, ok := m.(float64); ok {
			status = int(s)
		}
	}

	var protocol string
	if p, ok := dict.GetValue(keyProtocol); ok {
		if s, ok := p.(string); ok {
			protocol = s
		}
	}

	var body string
	if b, ok := dict.GetValue(keyBody); ok {
		if s, ok := b.(string); ok {
			body = s
		} else if bytes, err := json.Marshal(b); err == nil {
			body = string(bytes) + "\n"
		} else {
			log.Err(err).Msg("Failed to read body")
		}
	}

	response := &http.Response{
		Status:           fmt.Sprintf("%d %s", status, http.StatusText(status)),
		StatusCode:       status,
		Proto:            protocol,
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           http.Header{},
		Body:             io.NopCloser(strings.NewReader(body)),
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     false,
		Trailer:          http.Header{},
		Request:          nil,
		TLS:              nil,
	}

	var headers model.Dictionary
	if h, ok := dict.GetValue(keyHeaders); ok {
		if d, ok := h.(model.Dictionary); ok {
			headers = d
		}
	}

	for _, key := range headers.Keys() {
		if values, ok := headers.Get(key).([]model.Entry); ok {
			for _, value := range values {
				if s, ok := value.(string); ok {
					if key == "Content-Length" {
						s = strconv.Itoa(len(body))
					}
					response.Header.Add(key, s)
				}
			}
		}
	}

	var trailers model.Dictionary
	if h, ok := dict.GetValue(keyTrailers); ok {
		if d, ok := h.(model.Dictionary); ok {
			trailers = d
		}
	}

	for _, key := range trailers.Keys() {
		if values, ok := trailers.Get(key).([]model.Entry); ok {
			for _, value := range values {
				if s, ok := value.(string); ok {
					response.Trailer.Add(key, s)
				}
			}
		}
	}

	return response, nil
}

func NewResponseDict(response *http.Response) (ResponseDict, error) {
	dict := model.NewDictionary()

	dict.Set(keyStatus, response.StatusCode)
	dict.Set(keyProtocol, response.Proto)

	headers := model.NewDictionary()
	for key, values := range response.Header {
		entries := make([]model.Entry, len(values))
		for i, val := range values {
			entries[i] = val
		}
		headers.Set(key, entries)
	}
	dict.Set(keyHeaders, headers)

	if response.Body != nil {
		b, err := io.ReadAll(response.Body)
		if err != nil {
			return ResponseDict{dict.Pack()}, err
		}

		var bodydict interface{}
		if err := json.Unmarshal(b, &bodydict); err != nil {
			dict.Set(keyBody, string(b))
		} else {
			dict.Set(keyBody, model.CleanTypes(bodydict))
		}
	}

	trailers := model.NewDictionary()
	for key, values := range response.Trailer {
		entries := make([]model.Entry, len(values))
		for i, val := range values {
			entries[i] = val
		}
		trailers.Set(key, entries)
	}
	dict.Set(keyTrailers, trailers)

	return ResponseDict{dict.Pack()}, nil
}

func (r ResponseDict) Status() int {
	return r.UnpackAsDict().Get(keyStatus).(int)
}

func (r ResponseDict) Protocol() string {
	return r.UnpackAsDict().Get(keyProtocol).(string)
}
