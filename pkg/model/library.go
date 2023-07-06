package model

import (
	"fmt"
	"html/template"
	"net/url"
	"os"
	"path"
	"strings"
)

var libraries = map[string]Loader{}

type Loader interface {
	Load(name string) ([]byte, error)
}

func Declare(libname string, uri string) error {
	switch {
	case strings.HasPrefix(uri, "file://"):
		libraries[libname] = localDirLoader{uri}
	default:
		return fmt.Errorf("invalid uri scheme for library : %s", uri)
	}

	return nil
}

func Load(libname string, masking string, globalSeed int64, globalCaches map[string]Cache, globalFunctions template.FuncMap) (Pipeline, error) {
	loader, ok := libraries[libname]
	if !ok {
		return nil, fmt.Errorf("library not found : %s", libname)
	}

	yaml, err := loader.Load(masking)
	if err != nil {
		return nil, err
	}

	definition, err := LoadPipelineDefinitionFromYAML(yaml)
	if err != nil {
		return nil, err
	}

	definition.Seed += globalSeed

	pipeline := NewPipeline(nil)
	pipeline, _, err = BuildPipeline(pipeline, definition, globalCaches, globalFunctions, "", "")
	if err != nil {
		return nil, err
	}

	return pipeline, nil
}

type localDirLoader struct {
	uri string
}

func (l localDirLoader) Load(name string) ([]byte, error) {
	u, err := url.Parse(l.uri)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(u.Host, u.Path, name))
}
