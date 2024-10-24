package mock

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
)

type RequestDict struct {
	model.Dictionary
}

func (r RequestDict) ToRequest() (*http.Request, error) {
	dict := r.Dictionary
	if d, ok := dict.TryUnpackAsDict(); ok {
		dict = d
	}

	var method string
	if m, ok := dict.GetValue(keyMethod); ok {
		if s, ok := m.(string); ok {
			method = s
		}
	}

	var url string
	if u, ok := dict.Get(keyURL).(model.Dictionary); ok {
		var err error
		url, err = dictToURL(u)
		if err != nil {
			return nil, err
		}
	}

	var body string
	if b, ok := dict.GetValue(keyBody); ok {
		if s, ok := b.(string); ok {
			body = s
		} else if d, ok := b.(model.Dictionary); ok {
			bytes, err := d.MarshalJSON()
			if err != nil {
				return nil, err
			}
			body = string(bytes)
		}
	}

	req, err := http.NewRequest(method, url, strings.NewReader(body))

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
					req.Header.Add(key, s)
				}
			}
		}
	}

	if p, ok := dict.GetValue(keyProtocol); ok {
		if s, ok := p.(string); ok {
			req.Proto = s
		}
	}

	return req, err
}

func NewRequestDict(request *http.Request) (RequestDict, error) {
	dict := model.NewDictionary()
	dict.Set(keyMethod, request.Method)
	dict.Set(keyURL, urlToDict(request.URL))
	dict.Set(keyProtocol, request.Proto)
	headers := model.NewDictionary()

	for key, values := range request.Header {
		entries := make([]model.Entry, len(values))
		for i, val := range values {
			entries[i] = val
		}
		headers.Set(key, entries)
	}
	dict.Set(keyHeaders, headers)

	if request.Body != nil {
		b, err := io.ReadAll(request.Body)
		if err != nil {
			return RequestDict{dict.Pack()}, err
		}
		if bodydict, err := jsonline.JSONToDictionary(b); err != nil {
			dict.Set(keyBody, string(b))
		} else {
			dict.Set(keyBody, bodydict)
		}
	}

	return RequestDict{dict.Pack()}, nil
}

func (r RequestDict) Method() string {
	return r.UnpackAsDict().Get(keyMethod).(string)
}

func (r RequestDict) URLPath() string {
	return r.UnpackAsDict().Get(keyURL).(model.Dictionary).Get(keyURLPath).(string)
}

func (r RequestDict) URLFragment() (string, bool) {
	s, ok := r.UnpackAsDict().Get(keyURL).(model.Dictionary).Get(keyURLFragment).(string)
	return s, ok
}

func (r RequestDict) Protocol() string {
	return r.UnpackAsDict().Get(keyProtocol).(string)
}

func (r RequestDict) Body() string {
	return r.UnpackAsDict().Get(keyBody).(string)
}

func (r RequestDict) SetCaptures(captures model.Dictionary) {
	r.UnpackAsDict().Set("captures", captures)
}

func urlToDict(url *url.URL) model.Dictionary {
	dict := model.NewDictionary()
	if url == nil {
		return dict
	}

	// dict.Set(keyURLScheme, url.Scheme)

	// if url.User != nil {
	// 	user := model.NewDictionary().With(keyURLUserName, url.User.Username())
	// 	if password, set := url.User.Password(); set {
	// 		user.Set(keyURLUserPassword, password)
	// 	}
	// 	dict.Set(keyURLUser, user)
	// }

	// host := model.NewDictionary().With(keyURLHostName, url.Hostname())
	// if url.Port() != "" {
	// 	host.Set(keyURLHostPort, url.Port())
	// }
	// dict.Set(keyURLHost, host)

	dict.Set(keyURLPath, url.EscapedPath())
	if url.RawQuery != "" {
		dict.Set(keyURLQuery, queryToDict(url.RawQuery))
	}

	if url.RawFragment != "" {
		dict.Set(keyURLFragment, url.EscapedFragment())
	}

	return dict
}

func queryToDict(rawquery string) model.Dictionary {
	query := model.NewDictionary()
	for _, pair := range strings.Split(rawquery, "&") {
		if key, value, found := strings.Cut(pair, "="); found {
			if values, exists := query.GetValue(key); exists {
				query.Set(key, append(values.([]model.Entry), value))
			} else {
				query.Set(key, []model.Entry{value})
			}
		} else {
			if values, exists := query.GetValue(key); exists {
				query.Set(key, append(values.([]model.Entry), nil))
			} else {
				query.Set(key, []model.Entry{nil})
			}
		}
	}
	return query
}

func dictToURL(dict model.Dictionary) (string, error) {
	// var user *url.Userinfo
	// if userinfo, ok := dict.Get(keyURLUser).(model.Dictionary); ok {
	// 	if pass, ok := userinfo.Get(keyURLUserPassword).(string); ok {
	// 		user = url.UserPassword(userinfo.Get(keyURLUserName).(string), pass)
	// 	}
	// }

	// hostport := dict.Get(keyURLHost).(model.Dictionary).Get(keyURLHostName).(string)
	// if port, ok := dict.Get(keyURLHost).(model.Dictionary).Get(keyURLHostPort).(string); ok {
	// 	hostport = hostport + ":" + port
	// }

	rawpath := dict.Get(keyURLPath).(string)
	path, err := url.PathUnescape(rawpath)
	if err != nil {
		return "", err
	}

	queryDict, ok := dict.GetValue(keyURLQuery)
	queryRaw := ""
	if ok {
		queryRaw = dictToRawQuery(queryDict.(model.Dictionary))
	}

	fragmentDict, ok := dict.GetValue(keyURLFragment)
	fragmentRaw := ""
	fragment := ""
	if ok {
		fragmentRaw = fragmentDict.(string)
		fragment, err = url.PathUnescape(fragmentRaw)
		if err != nil {
			return "", err
		}
	}

	url := &url.URL{
		Scheme:      "", // dict.Get(keyURLScheme).(string),
		Opaque:      "",
		User:        nil, // user,
		Host:        "",  // hostport,
		Path:        path,
		RawPath:     rawpath,
		OmitHost:    false,
		ForceQuery:  false,
		RawQuery:    queryRaw,
		Fragment:    fragment,
		RawFragment: fragmentRaw,
	}

	return url.String(), nil
}

func dictToRawQuery(dict model.Dictionary) string {
	query := url.Values{}

	for _, key := range dict.Keys() {
		values := dict.Get(key).([]model.Entry)
		for _, value := range values {
			query.Add(key, value.(string))
		}
	}

	return query.Encode()
}
