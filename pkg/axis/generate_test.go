// Copyright (C) 2024 CGI France
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

package axis_test

import (
	"math/rand"
	"testing"

	"github.com/cgi-fr/pimo/pkg/axis"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	generator := axis.NewGenerator("start", 0)

	generator.SetPoint("birth", "start", -80, -18)
	generator.SetPoint("contract", "birth", 18, 40, axis.LowerThan("start", axis.Retry))
	generator.SetPoint("promotion", "contract", 0, 5)

	result := generator.Generate(rand.New(rand.NewSource(11)))

	for key, value := range result {
		if value != nil {
			println(key, *value)
		} else {
			println(key, "nil")
		}
	}

	assert.Len(t, result, 4)
}
