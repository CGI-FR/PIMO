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
	"testing"

	"github.com/mattn/anko/env"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	t.Parallel()

	function := `func add(x,y) {
					return x+y
			    }`

	env := Environment{Env: env.NewEnv()}
	env.Compile(function)

	template := `add(4,6)`

	result := env.Execute(template)

	assert.Equal(t, result, int64(10))
}

func TestBuildScriptFunction(t *testing.T) {
	t.Parallel()

	Param := make(map[string]string)
	Param["i"] = "int64"

	function := Function{Params: Param, Body: "return i + 1"}

	res := function.Build("add")

	assert.Equal(t, "func add(i) { return i + 1 }", res)
}
