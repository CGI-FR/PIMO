package model

import (
	"fmt"
	"html/template"
	"net/url"
	"os"
	"path"
	"strings"
)

var (
	libdefault = "default"
	libraries  = map[string]LibraryLoader{}
)

type LibraryLoader interface {
	Load(name string) ([]byte, error)
}

func DeclareLibrary(libname string, uri string) error {
	switch {
	case strings.HasPrefix(uri, "file://"):
		libraries[libname] = localDirLibraryLoader{uri}
	default:
		return fmt.Errorf("invalid uri scheme for library : %s", uri)
	}

	return nil
}

func SetDefaultLibrary(libname string) string {
	olddefault := libdefault
	libdefault = libname
	return olddefault
}

func LoadDefaultLibrary(masking string, globalSeed int64, globalCaches map[string]Cache, globalFunctions template.FuncMap) (Pipeline, error) {
	return LoadLibrary(libdefault, masking, globalSeed, globalCaches, globalFunctions)
}

func LoadLibrary(libname string, masking string, globalSeed int64, globalCaches map[string]Cache, globalFunctions template.FuncMap) (Pipeline, error) {
	if libname == "default" {
		libname = libdefault
	}

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

type localDirLibraryLoader struct {
	uri string
}

func (l localDirLibraryLoader) Load(name string) ([]byte, error) {
	u, err := url.Parse(l.uri)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(u.Host, u.Path, name))
}
