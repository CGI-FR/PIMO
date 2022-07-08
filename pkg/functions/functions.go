// Copyright (C) 2021 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package functions

import (
	"encoding/json"
	"reflect"
)

func getTypeFromString(t string) reflect.Type {
	switch t {
	case "int64":
		return reflect.TypeOf(int64(0))
	case "int32":
		return reflect.TypeOf(int32(0))
	case "int16":
		return reflect.TypeOf(int16(0))
	case "int8":
		return reflect.TypeOf(int8(0))
	case "uint":
		return reflect.TypeOf(uint(0))
	case "uint64":
		return reflect.TypeOf(uint64(0))
	case "uint32":
		return reflect.TypeOf(uint32(0))
	case "uint16":
		return reflect.TypeOf(uint16(0))
	case "uint8":
		return reflect.TypeOf(uint8(0))
	case "float64":
		return reflect.TypeOf(float64(0))
	case "float32":
		return reflect.TypeOf(float32(0))
	case "string":
		return reflect.TypeOf("0")
	case "json.Number":
		return reflect.TypeOf(json.Number("0"))
	case "bool":
		return reflect.TypeOf(true)
	case "":
		panic("")
	default:
		panic("")
	}
}

func GetFunction(params map[string]string, ret string, ankowrapper func(...interface{}) interface{}) reflect.Value {
	paramsType := []reflect.Type{}

	for _, t := range params {
		paramsType = append(paramsType, getTypeFromString(t))
	}

	typefn := reflect.FuncOf(paramsType, []reflect.Type{getTypeFromString(ret)}, false)

	wrapfn := func(args []reflect.Value) (results []reflect.Value) {
		argsInterface := []interface{}{}
		for i := range args {
			argsInterface = append(argsInterface, args[i].Interface())
		}

		result := ankowrapper(argsInterface...)

		return []reflect.Value{reflect.ValueOf(result)}
	}

	fn := reflect.MakeFunc(typefn, wrapfn)

	return fn
}
