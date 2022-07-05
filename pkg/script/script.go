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

package script

import (
	"fmt"
	"log"
	"math"
	"regexp"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

type Engine struct {
	env   *env.Env
	names []string
}

// NewEngine create a template Engine
func NewEngine(script string) *Engine {
	return &Engine{env: Compile(env.NewEnv(), script), names: Names(script)}
}

// DefinePackage defines the packages that can be used.
func DefinePackage(e *env.Env) {
	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	err = e.Define("pow", math.Pow)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
}

// Compile returns the environment needed for a VM to run in after compiling the script.
func Compile(e *env.Env, script string) *env.Env {
	DefinePackage(e)

	_, err := vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	return e
}

// Execute parses script and executes in the specified environment.
func Execute(e *env.Env, script string) interface{} {
	output, err := vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	return output
}

// Names returns the list of function names found in script
func Names(script string) (names []string) {
	re := regexp.MustCompile(`func (.*)\(.*\)`)

	find := re.FindAllSubmatch([]byte(script), -1)

	for i := range find {
		names = append(names, string(find[i][1]))
	}

	return names
}
