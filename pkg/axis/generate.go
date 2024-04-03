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

package axis

import (
	"math/rand"
)

type Range struct {
	name, ref   string
	min, max    int64
	constraints []Constraint
}

type Generator struct {
	originName  string
	originValue int64
	points      []Range
}

func NewGenerator(originName string, originValue int64) *Generator {
	return &Generator{
		originName:  originName,
		originValue: originValue,
		points:      make([]Range, 0, 5),
	}
}

func (g *Generator) SetPoint(name string, reference string, min, max int64, constraints ...Constraint) {
	if reference == "" {
		reference = g.originName
	}

	g.points = append(g.points, Range{
		name:        name,
		ref:         reference,
		min:         min,
		max:         max,
		constraints: constraints,
	})
}

func (g *Generator) Generate(rng *rand.Rand) map[string]*int64 {
	loopCount := 0

Loop:
	for {
		loopCount++

		result := make(map[string]*int64, len(g.points)+1)

		result[g.originName] = &g.originValue

		for _, point := range g.points {
			var pointer *int64

			if ref := result[point.ref]; ref != nil {
				value := rng.Int63n(point.max-point.min) + point.min + *ref
				pointer = &(value)
			}

			if pointer != nil {
				for _, constraint := range point.constraints {
					if !constraint.Validate(*pointer, result) {
						switch constraint.Behavior() {
						case Nullify:
							pointer = nil
						case Reject:
							panic("rejected timeline generation, can't find valid value for " + point.name)
						case Retry:
							if loopCount > 200 {
								panic("rejected timeline generation, can't find valid value for " + point.name)
							}

							continue Loop
						}
					}
				}
			}

			result[point.name] = pointer
		}

		return result
	}
}

func (g *Generator) Origin() string {
	return g.originName
}
