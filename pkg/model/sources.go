package model

import (
	"iter"
)

// SliceSource is a source of dictionaries from a slice. It implements the Iterable and Resource interfaces.
type SliceSource struct {
	data []Dictionary
}

// NewSliceSource creates a new slice source from the given data.
func NewSliceSource(data []Dictionary) Iterable[Dictionary] {
	return &SliceSource{data}
}

// Open opens the source.
func (s *SliceSource) Open() error {
	return nil
}

// Close closes the source.
func (s *SliceSource) Close() error {
	return nil
}

// Values returns a sequence of dictionaries from the source.
func (s *SliceSource) Values() iter.Seq2[Dictionary, error] {
	return func(yield func(Dictionary, error) bool) {
		for _, entry := range s.data {
			if !yield(entry, nil) {
				return
			}
		}
	}
}
