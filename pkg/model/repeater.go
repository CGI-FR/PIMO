package model

import "iter"

type Predicate[T any] interface {
	Reset()
	Test(value T) bool
}

type TruePredicate[T any] struct{}

func (TruePredicate[T]) Test(T) bool {
	return true
}

func (TruePredicate[T]) Reset() {}

type FalsePredicate[T any] struct{}

func (FalsePredicate[T]) Test(T) bool {
	return false
}

func (FalsePredicate[T]) Reset() {}

type CountPredicate[T any] struct {
	max   int
	count int
}

func (c *CountPredicate[T]) Test(T) bool {
	c.count++
	return c.count <= c.max
}

func (c *CountPredicate[T]) Reset() {
	c.count = 0
}

type Repeater[T Cloneable[T]] struct {
	input Iterable[T]
	while Predicate[T]
	until Predicate[T]
}

func (r *Repeater[T]) Open() error {
	return r.input.Open()
}

func (r *Repeater[T]) Close() error {
	return r.input.Close()
}

func NewRepeater[T Cloneable[T]](input Iterable[T], while Predicate[T], until Predicate[T]) Iterable[T] {
	return &Repeater[T]{input, while, until}
}

func NewCountRepeater[T Cloneable[T]](input Iterable[T], count int) Iterable[T] {
	return NewRepeater(input, &CountPredicate[T]{count, 0}, &FalsePredicate[T]{})
}

func (r *Repeater[T]) Values() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for value, err := range r.input.Values() {
			if err != nil {
				yield(value, err)
				return
			}

			r.while.Reset()
			r.until.Reset()

			vcopy := value.Copy()

			for r.while.Test(vcopy) {
				if !yield(vcopy, nil) {
					return
				}

				if r.until.Test(vcopy) {
					break
				}

				vcopy = value.Copy()
			}
		}
	}
}
