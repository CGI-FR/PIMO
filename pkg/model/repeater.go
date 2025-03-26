package model

import (
	"iter"
)

type Predicate interface {
	Reset() error
	Test(value Entry) bool
	Error() error
}

type TruePredicate struct{}

func (TruePredicate) Test(Entry) bool {
	return true
}

func (TruePredicate) Reset() error {
	return nil
}

func (TruePredicate) Error() error {
	return nil
}

type FalsePredicate struct{}

func (FalsePredicate) Test(Entry) bool {
	return false
}

func (FalsePredicate) Reset() error {
	return nil
}

func (FalsePredicate) Error() error {
	return nil
}

type CountPredicate struct {
	max   int
	count int
}

func (c *CountPredicate) Test(Entry) bool {
	c.count++
	return c.count <= c.max
}

func (c *CountPredicate) Reset() error {
	c.count = 0
	return nil
}

func (c *CountPredicate) Error() error {
	return nil
}

type Repeater struct {
	input Source
	while Predicate
	until Predicate
}

func (r *Repeater) Open() error {
	return r.input.Open()
}

func (r *Repeater) Close() error {
	return r.input.Close()
}

func NewRepeater(input Source, while Predicate, until Predicate) *Repeater {
	return &Repeater{input, while, until}
}

func NewCountRepeater(input Source, count int) *Repeater {
	return NewRepeater(input, &CountPredicate{count, 0}, &FalsePredicate{})
}

func (r *Repeater) Values() iter.Seq2[Entry, error] {
	return func(yield func(Entry, error) bool) {
		for value, err := range r.input.Values() {
			if err != nil {
				yield(value, err)
				return
			}

			if err := r.while.Reset(); err != nil {
				yield(value, err)
				return
			}

			if err := r.until.Reset(); err != nil {
				yield(value, err)
				return
			}

			vcopy := Copy(value)

			for r.while.Test(vcopy) {
				if !yield(vcopy, nil) {
					return
				}

				if r.until.Test(vcopy) || r.until.Error() != nil {
					break
				}

				vcopy = Copy(value)
			}

			if r.while.Error() != nil {
				yield(vcopy, r.while.Error())
				return
			}

			if r.until.Error() != nil {
				yield(vcopy, r.until.Error())
				return
			}
		}
	}
}
