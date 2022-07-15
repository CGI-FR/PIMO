// Copyright (C) 2021 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package functions

import (
	"fmt"
	"strings"
	"text/template"
)

type ParamDefinition struct {
	Name string
	Type string
}

type Definition struct {
	Params  []ParamDefinition
	Returns string
	Body    string
}

type Builder struct {
	Definitions map[string]Definition
}

// AsScript returns the function as a compilable script.
func (d Definition) AsScript(name string) string {
	script := &strings.Builder{}

	script.WriteString("func ")
	script.WriteString(name)
	script.WriteByte('(')

	for i, param := range d.Params {
		if i > 0 {
			script.WriteByte(',')
		}
		script.WriteString(param.Name)
	}

	script.WriteString(") { ")
	script.WriteString(d.Body)
	script.WriteString(" }")

	return script.String()
}

// AsCall returns a call instruction of the function with the given param values.
func (d Definition) AsCall(name string, args ...interface{}) string {
	var joined string
	for _, arg := range args {
		switch arg.(type) {
		case string:
			joined = fmt.Sprintf("\"%s\"", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args...)), ","), "[]"))
		default:
			joined = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args...)), ","), "[]")
		}
	}
	return fmt.Sprintf("%s(%s)", name, joined)
}

// Build returns a FuncMap (map[string]interface{}) that can be used in Go Template API, with all functions.
func (b Builder) Build() (template.FuncMap, error) {
	funcMap := make(template.FuncMap, len(b.Definitions))

	if len(b.Definitions) == 0 {
		return funcMap, nil
	}

	env := NewEnvironment()

	for name, def := range b.Definitions {
		name := name // https://stackoverflow.com/a/26694016/2531684

		err := env.Compile(def.AsScript(name))
		if err != nil {
			return nil, fmt.Errorf("cannot compile function: %w", err)
		}

		wrapper := func(args ...interface{}) (interface{}, error) {
			output, err := env.Execute(def.AsCall(name, args))
			if err != nil {
				return nil, fmt.Errorf("cannot execute function: %w", err)
			}
			return output, nil
		}

		funcMap[name] = wrapper
	}

	return funcMap, nil
}
