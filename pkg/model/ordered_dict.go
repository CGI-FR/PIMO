package model

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gitlab.com/c0b/go-ordered-json"
)

type Dictionary struct {
	*ordered.OrderedMap
}

func NewDictionary() Dictionary {
	return Dictionary{ordered.NewOrderedMap()}
}

// CopyDictionary clone deeply dictionary
func CopyDictionary(other Dictionary) Dictionary {
	if other.OrderedMap == nil {
		return NewDictionary()
	}

	om := ordered.NewOrderedMap()
	iter := other.EntriesIter()

	for {
		pair, ok := iter()
		if !ok {
			break
		}
		switch typedVal := pair.Value.(type) {
		case Dictionary:
			om.Set(pair.Key, CopyDictionary(typedVal))
		default:
			om.Set(pair.Key, pair.Value)
		}
	}
	return Dictionary{om}
}

// CleanTypes always return either primitive types, or a composition of Dictionaries
func CleanTypes(inter interface{}) interface{} {
	switch typedInter := inter.(type) {
	case map[string]Entry:
		log.Warn().Msg("conversion from map[string]Entry will lose order of keys")
		dict := NewDictionary()
		for k, v := range typedInter {
			dict.Set(k, CleanTypes(v))
		}
		return dict
	case *Dictionary:
		iter := typedInter.EntriesIter()
		dict := NewDictionary()
		for pair, ok := iter(); ok; pair, ok = iter() {
			dict.Set(pair.Key, CleanTypes(pair.Value))
		}
		return dict
	case Dictionary:
		iter := typedInter.EntriesIter()
		dict := NewDictionary()
		for pair, ok := iter(); ok; pair, ok = iter() {
			dict.Set(pair.Key, CleanTypes(pair.Value))
		}
		return dict
	case *ordered.OrderedMap:
		iter := typedInter.EntriesIter()
		dict := NewDictionary()
		for pair, ok := iter(); ok; pair, ok = iter() {
			dict.Set(pair.Key, CleanTypes(pair.Value))
		}
		return dict
	case ordered.OrderedMap:
		iter := typedInter.EntriesIter()
		dict := NewDictionary()
		for pair, ok := iter(); ok; pair, ok = iter() {
			dict.Set(pair.Key, CleanTypes(pair.Value))
		}
		return dict
	case []interface{}:
		tab := []Entry{}

		for _, item := range typedInter {
			tab = append(tab, CleanTypes(item))
		}

		return tab
	default:
		return inter
	}
}

// CleanDictionary fixes all types in the structure
func CleanDictionary(dict interface{}) Dictionary {
	return CleanTypes(dict).(Dictionary)
}

func CleanDictionarySlice(dictSlice interface{}) []Dictionary {
	result := []Dictionary{}

	switch typedInter := dictSlice.(type) {
	case []interface{}:
		for _, d := range typedInter {
			result = append(result, CleanDictionary(d))
		}

	case []Dictionary:
		for _, d := range typedInter {
			result = append(result, CleanDictionary(d))
		}

	case []Entry:
		for _, d := range typedInter {
			result = append(result, CleanDictionary(d))
		}
	}

	return result
}

// UnorderedTypes a composition of map[string]Entry
func UnorderedTypes(inter interface{}) interface{} {
	switch typedInter := inter.(type) {
	case map[string]Entry:
		cleanmap := map[string]Entry{}
		for k, v := range typedInter {
			cleanmap[k] = UnorderedTypes(v)
		}
		return cleanmap
	case *Dictionary:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]Entry{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = UnorderedTypes(pair.Value)
		}
		return cleanmap
	case Dictionary:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]Entry{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = UnorderedTypes(pair.Value)
		}
		return cleanmap
	case *ordered.OrderedMap:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]Entry{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = UnorderedTypes(pair.Value)
		}
		return cleanmap
	case ordered.OrderedMap:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]Entry{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = UnorderedTypes(pair.Value)
		}
		return cleanmap
	case []interface{}:
		tab := []Entry{}

		for _, item := range typedInter {
			tab = append(tab, UnorderedTypes(item))
		}

		return tab
	default:
		return inter
	}
}

func (d Dictionary) Copy() Dictionary {
	return CopyDictionary(d)
}

func (d Dictionary) Unordered() map[string]Entry {
	return UnorderedTypes(d).(map[string]Entry)
}

func (d Dictionary) String() string {
	b, err := d.MarshalJSON()
	if err != nil {
		return fmt.Sprint(d.OrderedMap)
	}
	return string(b)
}

func (d Dictionary) With(key string, value interface{}) Dictionary {
	result := CleanDictionary(d)
	result.Set(key, CleanTypes(value))
	return result
}
