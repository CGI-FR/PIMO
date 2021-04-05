package model

import "github.com/cgi-fr/pimo/pkg/selector"

// Dictionary is a Map with string as key and Entry as value
type Dictionary = selector.Dictionary

// Entry is a dictionary value
type Entry = selector.Entry

type Selector = selector.Selector

func NewPathSelector(path string) Selector {
	return selector.NewSelector(path)
}
