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

type Applier func(rootContext Dictionary, parentContext Dictionary, key string, value Entry) (Action, Entry)

type Selector interface {
	Apply(Dictionary, ...Applier) bool
	ApplyContext(Dictionary, ...Applier) bool

	// old interface
	Delete(Dictionary) Dictionary
	ReadContext(Dictionary) (Dictionary, string, bool)
	WriteContext(Dictionary, Entry) Dictionary
	Read(Dictionary) (Entry, bool)
	Write(Dictionary, Entry) Dictionary
}

type selectorInternal interface {
	Selector
	// internal
	applySub(Dictionary, Dictionary, ...Applier) bool
	applySubContext(Dictionary, Dictionary, ...Applier) bool
}

type selector struct {
	path string
	sub  selectorInternal
}

func NewSelector(path string) Selector {
	paths := strings.SplitN(path, ".", 2)
	if len(paths) == 2 {
		return selector{paths[0], NewSelector(paths[1]).(selectorInternal)}
	}
	return selector{paths[0], nil}
}

func (s selector) Apply(root Dictionary, appliers ...Applier) bool {
	return s.applySub(root, root, appliers...)
}

func (s selector) applySub(root Dictionary, current Dictionary, appliers ...Applier) bool {
	entry, ok := current[s.path]
	if !ok {
		if s.sub == nil {
			// apply with nil value
			s.apply(root, current, nil, appliers)
		}
		return false
	}
	v := reflect.ValueOf(entry)
	kind := v.Kind()
	if s.sub != nil {
		switch kind {
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				s.sub.applySub(root, v.Index(i).Interface().(Dictionary), appliers...)
			}
			return true
		default:
			return s.sub.applySub(root, entry.(Dictionary), appliers...)
		}
	}
	switch kind {
	case reflect.Slice:
		actualSlice := []Entry{}
		for i := 0; i < v.Len(); i++ {
			actualSlice = append(actualSlice, v.Index(i).Interface())
		}
		s.applySlice(root, current, actualSlice, appliers)
	default:
		s.apply(root, current, entry, appliers)
	}
	return true
}

func (s selector) apply(root Dictionary, current Dictionary, entry Entry, appliers []Applier) {
	for _, applier := range appliers {
		action, entry := applier(root, current, s.path, entry)
		switch action {
		case WRITE:
			current[s.path] = entry
		case DELETE:
			delete(current, s.path)
		}
	}
}

func (s selector) applySlice(root Dictionary, current Dictionary, entries []Entry, appliers []Applier) {
	for _, applier := range appliers {
		result := []Entry{}
		for _, entry := range entries {
			action, modEntry := applier(root, current, s.path, entry)
			switch action {
			case WRITE:
				result = append(result, modEntry)
			case NOTHING:
				result = append(result, entry)
			}
		}
		current[s.path] = result
	}
}

func (s selector) ApplyContext(root Dictionary, appliers ...Applier) bool {
	return s.applySubContext(root, root, appliers...)
}

func (s selector) applySubContext(root Dictionary, current Dictionary, appliers ...Applier) bool {
	entry, ok := current[s.path]
	if !ok {
		if s.sub == nil {
			// apply with nil value
			s.applyContext(root, current, nil, appliers)
		}
		return false
	}
	v := reflect.ValueOf(entry)
	kind := v.Kind()
	if s.sub != nil {
		switch kind {
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				s.sub.applySubContext(root, v.Index(i).Interface().(Dictionary), appliers...)
			}
			return true
		default:
			return s.sub.applySubContext(root, entry.(Dictionary), appliers...)
		}
	}

	s.applyContext(root, current, entry, appliers)

	return true
}

func (s selector) applyContext(root Dictionary, current Dictionary, entry Entry, appliers []Applier) {
	for _, applier := range appliers {
		action, entry := applier(root, current, s.path, entry)
		switch action {
		case WRITE:
			current[s.path] = entry.(Dictionary)[s.path]
		case DELETE:
			delete(current, s.path)
		}
	}
}

func (s selector) Delete(dictionary Dictionary) Dictionary {
	s.Apply(dictionary, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		return DELETE, nil
	})
	return dictionary
}

func (s selector) ReadContext(dictionary Dictionary) (sub Dictionary, subkey string, found bool) {
	s.Apply(dictionary, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		sub = parentContext
		subkey = key
		found = true
		return NOTHING, nil
	})
	return
}

func (s selector) WriteContext(dictionary Dictionary, masked Entry) Dictionary {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	v := reflect.ValueOf(masked)
	s.ApplyContext(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		if key == "" || value == nil {
			return NOTHING, nil
		}
		if v.Kind() == reflect.Map && v.MapIndex(reflect.ValueOf(key)).IsValid() {
			masked = v.MapIndex(reflect.ValueOf(key)).Interface()
		}
		return WRITE, masked
	})
	return result
}

func (s selector) Read(dictionary Dictionary) (match Entry, found bool) {
	s.Apply(dictionary, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		match = value
		found = value != nil
		return NOTHING, nil
	})
	return
}

func (s selector) Write(dictionary Dictionary, masked Entry) Dictionary {
	result := Dictionary{}
	for k, v := range dictionary {
		result[k] = v
	}
	s.Apply(result, func(rootContext, parentContext Dictionary, key string, value Entry) (Action, Entry) {
		return WRITE, masked
	})
	return result
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
