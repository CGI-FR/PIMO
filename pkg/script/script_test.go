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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	t.Parallel()

	function := `func add(x,y) {
					return x+y
			    }`

	env := NewEngine(function)

	template := `add(4,6)`

	result := Execute(env.env, template)

	assert.Equal(t, result, int64(10))
}

func TestName(t *testing.T) {
	t.Parallel()

	function1 := `Hello
	func add(x,y) {
		return x+y
	}`

	result1 := Names(function1)
	assert.Equal(t, []string{"add"}, result1)

	function2 := `Hello
	func add(x,y) {
		return x+y
	}
	func sub(x,y) {
		return x-y
	}`

	result2 := Names(function2)
	assert.Equal(t, []string{"add", "sub"}, result2)
}
