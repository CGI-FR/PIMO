package template

import (
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//nolint:checknoglobals
var seededFuncs = map[string]interface{}{}

func InjectSeededFuncGenerator(name string, f interface{}) {
	seededFuncs[name] = f
}

// Engine is a template engine
type Engine struct {
	*template.Template
}

// rmAcc removes accents from string
// Function derived from: http://blog.golang.org/normalization
func rmAcc(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// NewEngine create a template Engine
func NewEngine(text string, funcMap template.FuncMap, seed int64) (*Engine, error) {
	funcMap["ToUpper"] = strings.ToUpper
	funcMap["ToLower"] = strings.ToLower
	funcMap["NoAccent"] = rmAcc
	for k, v := range seededFuncs {
		if f, ok := v.(func(int64) interface{}); ok {
			funcMap[k] = f(seed)
		}
	}
	temp, err := template.New("template").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(text)
	return &Engine{temp}, err
}
