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

	"github.com/stretchr/testify/assert"
)

func TestPermuteMemory_add(t *testing.T) {
	mem := NewPermuteMemory(1)
	mem.add("value")
	assert.Equal(t, []Entry{"value"}, mem.memory)
	mem.add("value2")
	assert.Equal(t, []Entry{"value", "value2"}, mem.memory)
}

func TestPermuteMemory_pick(t *testing.T) {
	mem := NewPermuteMemory(1)
	mem.add("value")
	e, ok := mem.pick()
	assert.True(t, ok)
	assert.Equal(t, "value", e)
	e, ok = mem.pick()
	assert.False(t, ok)
	assert.Nil(t, e)
}
