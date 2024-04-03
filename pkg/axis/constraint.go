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

type ConstraintBehavior int

const (
	Reject ConstraintBehavior = iota
	Retry
	Nullify
)

type Constraint interface {
	Validate(value int64, points map[string]*int64) bool
	Behavior() ConstraintBehavior
}

type constraint struct {
	behavior  ConstraintBehavior
	reference string
}

func (c constraint) Behavior() ConstraintBehavior {
	return c.behavior
}

func LowerThan(reference string, behavior ConstraintBehavior) Constraint {
	return lowerThan{
		constraint: constraint{
			behavior:  behavior,
			reference: reference,
		},
	}
}

type lowerThan struct {
	constraint
}

func (lt lowerThan) Validate(value int64, points map[string]*int64) bool {
	ref := points[lt.reference]

	if ref == nil {
		return true
	}

	return value < *ref
}

func GreaterThan(reference string, behavior ConstraintBehavior) Constraint {
	return greaterThan{
		constraint: constraint{
			behavior:  behavior,
			reference: reference,
		},
	}
}

type greaterThan struct {
	constraint
}

func (gt greaterThan) Validate(value int64, points map[string]*int64) bool {
	ref := points[gt.reference]

	if ref == nil {
		return true
	}

	return value > *ref
}
