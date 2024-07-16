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

package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCounters(t *testing.T) {
	tests := map[string]struct {
		count    int
		length   int
		base     int
		expected []int
	}{
		"0000": {
			count:    0,
			length:   4,
			base:     10,
			expected: []int{0, 0, 0, 0},
		},
		"0001": {
			count:    1,
			length:   4,
			base:     10,
			expected: []int{1, 0, 0, 0},
		},
		"0010": {
			count:    10,
			length:   4,
			base:     10,
			expected: []int{0, 1, 0, 0},
		},
		"0100": {
			count:    100,
			length:   4,
			base:     10,
			expected: []int{0, 0, 1, 0},
		},
		"1000": {
			count:    1000,
			length:   4,
			base:     10,
			expected: []int{0, 0, 0, 1},
		},
		"9999": {
			count:    9999,
			length:   4,
			base:     10,
			expected: []int{9, 9, 9, 9},
		},
		"99": {
			count:    99,
			length:   4,
			base:     10,
			expected: []int{9, 9, 0, 0},
		},
		"033": {
			count:    33,
			length:   3,
			base:     10,
			expected: []int{3, 3, 0},
		},
		"base2": {
			count:    54,
			length:   8,
			base:     2,
			expected: []int{0, 1, 1, 0, 1, 1, 0, 0},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			state := MaskEngineState{test.count}
			assert.Equal(t, test.expected, state.toCounters(test.length, test.base))
		})
	}
}
