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
func NewEngine(text string, funcMap template.FuncMap) (*Engine, error) {
	funcMap["ToUpper"] = strings.ToUpper
	funcMap["ToLower"] = strings.ToLower
	funcMap["NoAccent"] = rmAcc
	temp, err := template.New("template").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(text)
	return &Engine{temp}, err
}
