package model

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"gitlab.com/c0b/go-ordered-json"
)

type Dictionary struct {
	*ordered.OrderedMap
}

func NewPackedDictionary() Dictionary {
	return NewDictionary().Pack()
}

func NewDictionary() Dictionary {
	return Dictionary{ordered.NewOrderedMap()}
}

func Copy(other Entry) Entry {
	switch d := other.(type) {
	case Dictionary:
		return d.Copy()
	case Entry:
		return NewDictionary().With("entry", d).Copy().Get("entry")
	default:
		panic("")
	}
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
	case map[string]interface{}:
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

	case json.Number:

		resFloat64, err := typedInter.Float64()
		if err == nil {
			return resFloat64
		}

		return typedInter.String()

	case uint64:
		res := float64(typedInter)
		return res

	case int64:
		res := float64(typedInter)
		return res
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

func Untyped(inter interface{}) interface{} {
	switch typedInter := inter.(type) {
	case map[string]Entry:
		cleanmap := map[string]interface{}{}
		for k, v := range typedInter {
			cleanmap[k] = Untyped(v)
		}
		return cleanmap
	case *Dictionary:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]interface{}{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = Untyped(pair.Value)
		}
		return cleanmap
	case Dictionary:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]interface{}{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = Untyped(pair.Value)
		}
		return cleanmap
	case *ordered.OrderedMap:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]interface{}{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = Untyped(pair.Value)
		}
		return cleanmap
	case ordered.OrderedMap:
		iter := typedInter.EntriesIter()
		cleanmap := map[string]interface{}{}
		for pair, ok := iter(); ok; pair, ok = iter() {
			cleanmap[pair.Key] = Untyped(pair.Value)
		}
		return cleanmap
	case []interface{}:
		tab := []Entry{}

		for _, item := range typedInter {
			tab = append(tab, Untyped(item))
		}

		return tab

	case []Dictionary:
		tab := []interface{}{}

		for _, item := range typedInter {
			tab = append(tab, Untyped(item))
		}

		return tab

	case []Entry:
		tab := []interface{}{}

		for _, item := range typedInter {
			tab = append(tab, Untyped(item))
		}

		return tab

	default:
		return inter
	}
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

	case []Dictionary:
		tab := []Entry{}

		for _, item := range typedInter {
			tab = append(tab, UnorderedTypes(item))
		}

		return tab

	case []Entry:
		tab := []Entry{}

		for _, item := range typedInter {
			tab = append(tab, UnorderedTypes(item))
		}

		return tab

	default:
		return inter
	}
}

func (d Dictionary) IsPacked() bool {
	_, packed := d.GetValue(".")
	return packed
}

func (d Dictionary) Pack() Dictionary {
	return NewDictionary().With(".", d)
}

func (d Dictionary) Unpack() Entry {
	return d.Get(".")
}

func (d Dictionary) UnpackUnordered() Entry {
	return d.Unordered()["."]
}

func (d Dictionary) UnpackUntyped() interface{} {
	return d.Untyped()["."]
}

func (d Dictionary) UnpackAsDict() Dictionary {
	return d.Get(".").(Dictionary)
}

func (d Dictionary) TryUnpackAsDict() (Dictionary, bool) {
	dict, ok := d.Get(".").(Dictionary)
	return dict, ok
}

func (d Dictionary) CanUnpackAsDict() bool {
	_, ok := d.Get(".").(Dictionary)
	return ok
}

func (d Dictionary) Copy() Dictionary {
	return CopyDictionary(d)
}

func (d Dictionary) Unordered() map[string]Entry {
	return UnorderedTypes(d).(map[string]Entry)
}

func (d Dictionary) Untyped() map[string]interface{} {
	return Untyped(d).(map[string]interface{})
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
