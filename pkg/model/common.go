package model

type Resource interface {
	Open() error
	Close() error
}
