package model

import "reflect"

func deepmask(me MaskEngine, entry interface{}, context ...Dictionary) (interface{}, error) {
	v := reflect.ValueOf(entry)
	switch v.Kind() {
	case reflect.Slice:
		masked := []Entry{}
		for i := 0; i < v.Len(); i++ {
			maskedEntry, err := deepmask(me, v.Index(i).Interface(), context...)
			if err != nil {
				return nil, err
			}
			masked = append(masked, maskedEntry)
		}
		return masked, nil
	default:
		return me.Mask(entry, context...)
	}
}
