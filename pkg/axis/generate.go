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
	"errors"
	"fmt"
	"math/rand"
)

var ErrRejectedGeneration = errors.New("rejected timeline generation, can't find a value that satisfy constraint")

type Range struct {
	name, ref          string
	min, max           int64
	constraints        []Constraint
	onErrorReplaceWith string
}

type Generator struct {
	originName  string
	originValue int64
	points      []Range
	maxRetry    int
}

const (
	DefaultPointsSize = 5
	DefaultMaxRetry   = 200
)

func NewGenerator(originName string, originValue int64) *Generator {
	return &Generator{
		originName:  originName,
		originValue: originValue,
		points:      make([]Range, 0, DefaultPointsSize),
		maxRetry:    DefaultMaxRetry,
	}
}

func (g *Generator) SetMaxRetry(maxRetry int) {
	g.maxRetry = maxRetry
}

func (g *Generator) SetPoint(name string, reference string, min, max int64, defaultValueFrom string, constraints ...Constraint) {
	if reference == "" {
		reference = g.originName
	}

	g.points = append(g.points, Range{
		name:               name,
		ref:                reference,
		min:                min,
		max:                max,
		constraints:        constraints,
		onErrorReplaceWith: defaultValueFrom,
	})
}

func (g *Generator) Generate(rng *rand.Rand, initialStates ...map[string]*int64) (map[string]*int64, error) {
	loopCount := 0

Loop:
	for {
		loopCount++

		result := make(map[string]*int64, len(g.points)*2+1)

		result[g.originName] = &g.originValue

		loadStates(result, initialStates)

		for _, point := range g.points {
			var pointer *int64

			if ref := result[point.ref]; ref != nil {
				switch {
				case point.max > point.min:
					value := rng.Int63n(point.max-point.min) + point.min + *ref
					pointer = &(value)
				case point.max == point.min:
					value := point.max + *ref
					pointer = &(value)
				case point.max < point.min:
					value := rng.Int63n(point.min-point.max) + point.max + *ref
					pointer = &(value)
				}
			}

			if pointer != nil {
				for _, constraint := range point.constraints {
					if !constraint.Validate(*pointer, result) {
						if loopCount <= g.maxRetry {
							continue Loop // next try, better luck ?
						}

						switch constraint.Behavior() {
						case Nullify:
							pointer = nil
						case Replace:
							pointer = result[point.onErrorReplaceWith]
						// case Default:
						// 	pointer = &point.defaultValue
						case Reject:
							return nil, fmt.Errorf("%w: %s", ErrRejectedGeneration, point.name)
						}

						break // do not check next constraint, not necessary because pointer is nil
					}
				}
			}

			result[point.name] = pointer
		}

		return result, nil
	}
}

func (g *Generator) Origin() string {
	return g.originName
}

func loadStates(target map[string]*int64, states []map[string]*int64) {
	for _, state := range states {
		for name, value := range state {
			target[name] = value
		}
	}
}
