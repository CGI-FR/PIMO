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

	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

type Environment struct {
	env *env.Env
}

func NewEnvironment() Environment {
	e := Environment{env.NewEnv()}

	if err := e.init(); err != nil {
		return Environment{}
	}

	return e
}

func (e Environment) init() error {
	e.env = core.Import(e.env)

	return nil
}

// Compile the script and returns the environment needed to run it.
func (e Environment) Compile(script string) error {
	_, err := vm.Execute(e.env, nil, script)
	if err != nil {
		return fmt.Errorf("cannot compile environment: %w", err)
	}

	return nil
}

// Execute parses script and executes in the specified environment.
func (e Environment) Execute(script string) (interface{}, error) {
	output, err := vm.Execute(e.env, nil, script)
	if err != nil {
		return nil, fmt.Errorf("cannot execute environment: %w", err)
	}

	return output, nil
}
