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
	"log"
	"math"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

type Environment struct {
	Env *env.Env
}

// DefinePackage defines the packages that can be used.
func (env Environment) DefinePackage() {
	err := env.Env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	err = env.Env.Define("pow", math.Pow)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
}

// Compile returns the environment needed for a VM to run in after compiling the script.
func (env Environment) Compile(script string) {
	env.DefinePackage()

	_, err := vm.Execute(env.Env, nil, script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}
}

// Execute parses script and executes in the specified environment.
func (env Environment) Execute(script string) interface{} {
	output, err := vm.Execute(env.Env, nil, script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	return output
}

func (f Function) Build() string {
	script := ""
	params := ""
	for i, param := range f.Params {
		if i == 0 {
			params += param.Name
		} else {
			params += "," + param.Name
		}
	}
	script += "func " + f.Name + "(" + params + ") { " + f.Body + " }"

	return script
}
