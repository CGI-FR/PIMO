package selector

import (
	"reflect"
	"strings"
)

type Entry interface{}
type Dictionary map[string]Entry

type Action uint

const (
	NOTHING Action = iota
	WRITE          = iota
	DELETE         = iota
)

type Applier func(rootContext Dictionary, parentContext Dictionary, value Entry) (Action, Entry)

type Selector interface {
	Apply(Dictionary, ...Applier) bool
	ApplyWithContext(Dictionary, Dictionary, ...Applier) bool
}

type selector struct {
	path string
	sub  Selector
}

func NewSelector(path string) Selector {
	paths := strings.SplitN(path, ".", 2)
	if len(paths) == 2 {
		return selector{paths[0], NewSelector(paths[1])}
	} else {
		return selector{paths[0], nil}
	}
}

func (s selector) Apply(root Dictionary, appliers ...Applier) bool {
	return s.ApplyWithContext(root, root, appliers...)
}

func (s selector) ApplyWithContext(root Dictionary, current Dictionary, appliers ...Applier) bool {
	entry, ok := current[s.path]
	if !ok {
		return false
	}
	v := reflect.ValueOf(entry)
	kind := v.Kind()
	if s.sub != nil {
		switch kind {
		case reflect.Slice:
		default:
			return s.sub.ApplyWithContext(root, entry.(Dictionary), appliers...)
		}
	}
	switch kind {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			s.apply(root, current, v.Index(i).Interface(), appliers)
		}
	default:
		s.apply(root, current, entry, appliers)
	}
	return true
}

func (s selector) apply(root Dictionary, current Dictionary, entry Entry, appliers []Applier) {
	for _, applier := range appliers {
		action, entry := applier(root, current, entry)
		switch action {
		case WRITE:
			current[s.path] = entry
		case DELETE:
			delete(current, s.path)
		}
	}
}

// InterfaceToMap returns a Dictionary from an interface
func InterfaceToDictionary(inter interface{}) Dictionary {
	dic := make(map[string]Entry)
	mapint := inter.(map[string]interface{})

	for k, v := range mapint {
		switch typedValue := v.(type) {
		case map[string]interface{}:
			dic[k] = InterfaceToDictionary(v)
		case []interface{}:
			tab := []Entry{}
			for _, item := range typedValue {
				_, dico := item.(map[string]interface{})

				if dico {
					tab = append(tab, InterfaceToDictionary(item))
				} else {
					tab = append(tab, item)
				}
			}
			dic[k] = tab
		default:
			dic[k] = v
		}
	}

	return dic
}
