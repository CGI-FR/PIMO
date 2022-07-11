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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTypeFromString(t *testing.T) {
	t.Parallel()

	type1, err1 := getTypeFromString("int64")
	assert.Nil(t, err1)
	assert.Equal(t, reflect.TypeOf(int64(25)), type1)

	type2, err2 := getTypeFromString("string")
	assert.Nil(t, err2)
	assert.Equal(t, reflect.TypeOf("hello"), type2)
}
