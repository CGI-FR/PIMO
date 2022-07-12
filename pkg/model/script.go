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

package model

import (
	"fmt"
	"math"

	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

type Environment struct {
	Env *env.Env
}

func NewEnvironment() Environment {
	e := env.NewEnv()

	if err := DefinePackage(e); err != nil {
		return Environment{}
	}

	return Environment{e}
}

// DefinePackage defines the packages that can be used.
func DefinePackage(e *env.Env) error {
	e = core.Import(e)

	if err := e.Define("pow", math.Pow); err != nil {
		return fmt.Errorf("cannot define package: %w", err)
	}

	return nil
}

// Compile returns the environment needed for a VM to run in after compiling the script.
func (env Environment) Compile(script string) error {
	// env.DefinePackage()

	_, err := vm.Execute(env.Env, nil, script)
	if err != nil {
		return fmt.Errorf("cannot compile environment: %w", err)
	}

	return nil
}

// Execute parses script and executes in the specified environment.
func (env Environment) Execute(script string) (interface{}, error) {
	output, err := vm.Execute(env.Env, nil, script)
	if err != nil {
		return nil, fmt.Errorf("cannot execute environment: %w", err)
	}

	return output, nil
}

func (f Function) Build(name string) string {
	script := ""
	params := ""
	i := 0
	for _, param := range f.Params {
		if i == 0 {
			params += param.Name
		} else {
			params += "," + param.Name
		}
		i++
	}
	script += "func " + name + "(" + params + ") { " + f.Body + " }"

	return script
}
