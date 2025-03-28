package mock

import (
	"net/http"
	"net/url"
	"os"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/goccy/go-yaml"
)

type Config struct {
	Routes []Route `yaml:"routes,omitempty"`
}

type Route struct {
	Method  string  `yaml:"method,omitempty"`
	Path    string  `yaml:"path"`
	Masking Masking `yaml:"masking,omitempty"`
}

type Masking struct {
	Request  string `yaml:"request,omitempty"`
	Response string `yaml:"response,omitempty"`
}

func LoadConfigFromFile(filename string) (*Config, error) {
	source, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := yaml.Unmarshal(source, config); err != nil {
		return config, err
	}

	return config, nil
}

func (cfg *Config) Build(backend *url.URL, globalSeed *int64, cachesToLoad map[string]string, client *http.Client) (Context, error) {
	ctx := Context{
		client:  client,
		backend: backend,
		routes:  []ContextRoute{},
		caches:  map[string]model.Cache{},
	}

	for _, route := range cfg.Routes {
		var request, response *Processor
		var err error

		if route.Masking.Request != "" {
			request, ctx.caches, err = NewProcessor(route.Masking.Request, globalSeed, ctx.caches, cachesToLoad)
			if err != nil {
				return ctx, err
			}
		}

		if route.Masking.Response != "" {
			response, ctx.caches, err = NewProcessor(route.Masking.Response, globalSeed, ctx.caches, cachesToLoad)
			if err != nil {
				return ctx, err
			}
		}

		ctx.routes = append(ctx.routes, NewContextRoute(route, request, response))
	}

	return ctx, nil
}
