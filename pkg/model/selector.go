package model

import (
	"log"
	"reflect"
	"strings"

	"github.com/cgi-fr/pimo/pkg/selector"
)

// This file contains the old implementation of selectors.
// To activate the old version,
//   - rename function NewPathSelectorV1 to NewPathSelector (and fix compile error in model.go)
//   - rename interface SelectorV1 to Selector (and fix compile error in model.go)
//   - change the following lines in model.go :
//       type Dictionary = selector.Dictionary
//       type Entry = selector.Entry
//       type Selector = selector.Selector
//     to:
//       type Dictionary = map[string]Entry
//       type Entry = interface{}
// To activate the new version,
//   - rename function NewPathSelector to NewPathSelectorV1 (and fix compile error in model.go)
//   - rename interface Selector to SelectorV1 (and fix compile error in model.go)
//   - change the following lines in model.go :
//       type Dictionary = map[string]Entry
//       type Entry = interface{}
//     to:
//       type Dictionary = selector.Dictionary
//       type Entry = selector.Entry
//       type Selector = selector.Selector

func NewPathSelector(path string) selector.Selector {
	return selector.NewSelector(path)
}

// Selector acces to a specific data into a dictonary
type SelectorV1 interface {
	Read(Dictionary) (Entry, bool)
	Write(Dictionary, Entry) Dictionary
	Delete(Dictionary) Dictionary
	ReadContext(Dictionary) (Dictionary, string, bool)
	WriteContext(Dictionary, Entry) Dictionary
}

func NewPathSelectorV1(path string) Selector {
	mainEntry := strings.SplitN(path, ".", 2)
	if len(mainEntry) == 2 {
		return ComplexePathSelector{mainEntry[0], NewPathSelector(mainEntry[1])}
	} else {
		return NewSimplePathSelector(mainEntry[0])
	}
}

type ComplexePathSelector struct {
	path        string
	subSelector Selector
}

func (s ComplexePathSelector) ApplyWithContext(Dictionary, Dictionary, ...selector.Applier) bool {
	log.Fatal("Not implemented")
	return false
}

func (s ComplexePathSelector) Apply(Dictionary, ...selector.Applier) bool {
	log.Fatal("Not implemented")
	return false
}

func (s ComplexePathSelector) Read(dictionary Dictionary) (entry Entry, ok bool) {
	entry, ok = dictionary[s.path]
	if !ok {
		return
	}

	switch typedMatch := entry.(type) {
	case []Entry:
		entry = []Entry{}
		for _, subEntry := range typedMatch {
			var subResult Entry
			subResult, ok = s.subSelector.Read(subEntry.(Dictionary))
			if !ok {
				return
			}

			entry = append(entry.([]Entry), subResult.(Entry))
		}
		return
	case Dictionary:
		return s.subSelector.Read(typedMatch)
	default:
		return nil, false
	}
}

func (s ComplexePathSelector) ReadContext(dictionary Dictionary) (Dictionary, string, bool) {
	entry, ok := dictionary[s.path]
	if !ok {
		return dictionary, "", false
	}
	subEntry, ok := entry.(Dictionary)
	if !ok {
		return dictionary, "", false
	}
	return s.subSelector.ReadContext(subEntry)
}

func (s ComplexePathSelector) Write(dictionary Dictionary, entry Entry) Dictionary {
	result := Dictionary{}

	for k, v := range dictionary {
		result[k] = v
	}

	v := reflect.ValueOf(entry)
	switch v.Kind() {
	case reflect.Slice:
		subTargetArray, isArray := result[s.path].([]Entry)
		if !isArray {
			result[s.path] = s.subSelector.Write(result[s.path].(Dictionary), entry)
		} else {
			subArray := []Entry{}
			for i := 0; i < v.Len(); i++ {
				subArray = append(subArray, s.subSelector.Write(subTargetArray[i].(Dictionary), v.Index(i).Interface()))
			}
			result[s.path] = subArray
		}

	default:
		result[s.path] = s.subSelector.Write(result[s.path].(Dictionary), entry)
	}

	return result
}

func (s ComplexePathSelector) WriteContext(dictionary Dictionary, entry Entry) Dictionary {
	return dictionary
}

func (s ComplexePathSelector) Delete(dictionary Dictionary) Dictionary {
	result := Dictionary{}

	for k, v := range dictionary {
		if k == s.path {
			result[k] = s.subSelector.Delete(v.(Dictionary))
		}
	}
	return result
}

func NewSimplePathSelector(path string) Selector {
	return SimplePathSelector{path}
}

type SimplePathSelector struct {
	Path string
}

func (s SimplePathSelector) ApplyWithContext(Dictionary, Dictionary, ...selector.Applier) bool {
	log.Fatal("Not implemented")
	return false
}

func (s SimplePathSelector) Apply(Dictionary, ...selector.Applier) bool {
	log.Fatal("Not implemented")
	return false
}

func (s SimplePathSelector) Read(dictionary Dictionary) (entry Entry, ok bool) {
	entry, ok = dictionary[s.Path]
	return
}

func (s SimplePathSelector) ReadContext(dictionary Dictionary) (Dictionary, string, bool) {
	return dictionary, s.Path, true
}

func (s SimplePathSelector) Write(dictionary Dictionary, entry Entry) Dictionary {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	result[s.Path] = entry
	return result
}

func (s SimplePathSelector) WriteContext(dictionary Dictionary, entry Entry) Dictionary {
	return dictionary
}

func (s SimplePathSelector) Delete(dictionary Dictionary) Dictionary {
	result := Dictionary{}

	for k, v := range dictionary {
		if k != s.Path {
			result[k] = v
		}
	}
	return result
}
