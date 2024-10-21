package mock

import (
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
	Request  model.Masking `yaml:"request,omitempty"`
	Response model.Masking `yaml:"response,omitempty"`
}

func LoadConfigFromFile(filename string) (*Config, error) {
	source, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config *Config

	if err := yaml.Unmarshal(source, config); err != nil {
		return config, err
	}

	return config, nil
}

func (cfg *Config) Build() {
	
}
