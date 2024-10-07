package pimo

import (
	"io"
	"net/http"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
)

type RequestDict struct {
	model.Dictionary
}

const (
	KeyMethod   = "method"
	KeyPath     = "path"
	KeyProtocol = "protocol"
	KeyHeaders  = "headers"
	KeyBody     = "body"
)

func (r RequestDict) Method() string {
	return r.UnpackAsDict().Get(KeyMethod).(string)
}

func (r RequestDict) Path() string {
	return r.UnpackAsDict().Get(KeyPath).(string)
}

func (r RequestDict) Protocol() string {
	return r.UnpackAsDict().Get(KeyProtocol).(string)
}

func (r RequestDict) Body() string {
	return r.UnpackAsDict().Get(KeyBody).(string)
}

func NewRequestDict(request *http.Request) (RequestDict, error) {
	dict := model.NewDictionary()
	dict.Set(KeyMethod, request.Method)
	dict.Set(KeyPath, request.URL.String())
	dict.Set(KeyProtocol, request.Proto)
	headers := model.NewDictionary()
	for key, values := range request.Header {
		entries := make([]model.Entry, len(values))
		for i, val := range values {
			entries[i] = val
		}
		headers.Set(key, entries)
	}
	dict.Set(KeyHeaders, headers)
	b, err := io.ReadAll(request.Body)
	dict.Set(KeyBody, string(b))
	return RequestDict{dict.Pack()}, err
}

func ToRequest(dict model.Dictionary) (*http.Request, error) {
	if d, ok := dict.TryUnpackAsDict(); ok {
		dict = d
	}

	var method string
	if m, ok := dict.GetValue(KeyMethod); ok {
		if s, ok := m.(string); ok {
			method = s
		}
	}

	var path string
	if p, ok := dict.GetValue(KeyPath); ok {
		if s, ok := p.(string); ok {
			path = s
		}
	}

	var body string
	if b, ok := dict.GetValue(KeyBody); ok {
		if s, ok := b.(string); ok {
			body = s
		}
	}

	r, err := http.NewRequest(method, path, strings.NewReader(body))

	var headers model.Dictionary
	if h, ok := dict.GetValue(KeyHeaders); ok {
		if d, ok := h.(model.Dictionary); ok {
			headers = d
		}
	}

	for _, key := range headers.Keys() {
		if values, ok := headers.Get(key).([]model.Entry); ok {
			for _, value := range values {
				if s, ok := value.(string); ok {
					r.Header.Add(key, s)
				}
			}
		}
	}

	if p, ok := dict.GetValue(KeyProtocol); ok {
		if s, ok := p.(string); ok {
			r.Proto = s
		}
	}

	return r, err
}
