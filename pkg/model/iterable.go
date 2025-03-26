package model

import "iter"

// Iterable is a resource that can be iterated over to produce values of type T. It extends the Resource interface.
type Iterable[T any] interface {
	Resource
	Values() iter.Seq2[T, error]
}
