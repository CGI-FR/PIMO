package model

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"

	goccy "github.com/goccy/go-json"
	ordered "github.com/iancoleman/orderedmap"
)

type Dictionary struct {
	*ordered.OrderedMap
}

func NewPackedDictionary() Dictionary {
	return NewDictionary().Pack()
}

func NewDictionary() Dictionary {
	return Dictionary{ordered.New()}
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

	om := ordered.New()

	for _, key := range other.Keys() {
		value := other.Get(key)
		switch typedVal := value.(type) {
		case Dictionary:
			om.Set(key, CopyDictionary(typedVal))
		default:
			om.Set(key, value)
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
		dict := NewDictionary()
		for _, key := range typedInter.Keys() {
			dict.Set(key, CleanTypes(typedInter.Get(key)))
		}
		return dict
	case Dictionary:
		dict := NewDictionary()
		for _, key := range typedInter.Keys() {
			dict.Set(key, CleanTypes(typedInter.Get(key)))
		}
		return dict
	case *ordered.OrderedMap:
		dict := NewDictionary()
		for _, key := range typedInter.Keys() {
			value, _ := typedInter.Get(key)
			dict.Set(key, CleanTypes(value))
		}
		return dict
	case ordered.OrderedMap:
		dict := NewDictionary()
		for _, key := range typedInter.Keys() {
			value, _ := typedInter.Get(key)
			dict.Set(key, CleanTypes(value))
		}
		return dict
	case []Entry:
		tab := make([]Entry, 0, len(typedInter))
		for _, item := range typedInter {
			tab = append(tab, CleanTypes(item))
		}
		return tab
	case []interface{}:
		tab := make([]Entry, 0, len(typedInter))
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
		cleanmap := map[string]interface{}{}
		for _, key := range typedInter.Keys() {
			cleanmap[key] = Untyped(typedInter.Get(key))
		}
		return cleanmap
	case Dictionary:
		cleanmap := map[string]interface{}{}
		for _, key := range typedInter.Keys() {
			cleanmap[key] = Untyped(typedInter.Get(key))
		}
		return cleanmap
	case *ordered.OrderedMap:
		cleanmap := map[string]interface{}{}
		for _, key := range typedInter.Keys() {
			value, _ := typedInter.Get(key)
			cleanmap[key] = Untyped(value)
		}
		return cleanmap
	case ordered.OrderedMap:
		cleanmap := map[string]interface{}{}
		for _, key := range typedInter.Keys() {
			value, _ := typedInter.Get(key)
			cleanmap[key] = Untyped(value)
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
		cleanmap := map[string]Entry{}
		for _, key := range typedInter.Keys() {
			cleanmap[key] = UnorderedTypes(typedInter.Get(key))
		}
		return cleanmap
	case Dictionary:
		cleanmap := map[string]Entry{}
		for _, key := range typedInter.Keys() {
			cleanmap[key] = UnorderedTypes(typedInter.Get(key))
		}
		return cleanmap
	case *ordered.OrderedMap:
		cleanmap := map[string]Entry{}
		for _, key := range typedInter.Keys() {
			value, _ := typedInter.Get(key)
			cleanmap[key] = UnorderedTypes(value)
		}
		return cleanmap
	case ordered.OrderedMap:
		cleanmap := map[string]Entry{}
		for _, key := range typedInter.Keys() {
			value, _ := typedInter.Get(key)
			cleanmap[key] = UnorderedTypes(value)
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
	_, packed := d.OrderedMap.Get(".")
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

func (d Dictionary) GetValue(key string) (Entry, bool) {
	return d.OrderedMap.Get(key)
}

func (d Dictionary) Get(key string) Entry {
	entry, _ := d.OrderedMap.Get(key)
	return entry
}

func (d Dictionary) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	encoder := goccy.NewEncoder(&buf)
	for i, k := range d.Keys() {
		if i > 0 {
			buf.WriteByte(',')
		}
		// add key
		if err := encoder.Encode(k); err != nil {
			return nil, err
		}
		buf.Truncate(buf.Len() - 1) // remove last new line
		buf.WriteByte(':')
		// add value
		if err := encoder.Encode(d.Get(k)); err != nil {
			return nil, err
		}
		buf.Truncate(buf.Len() - 1) // remove last new line
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
