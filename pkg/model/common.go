package model

type Cloneable[T any] interface {
	Copy() T
}

type Resource interface {
	Open() error
	Close() error
}
