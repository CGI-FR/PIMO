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
	"errors"
	"fmt"
	"reflect"
)

func getTypeFromString(t string) (reflect.Type, error) {
	switch t {
	case "int64":
		return reflect.TypeOf(int64(0)), nil
	case "int32":
		return reflect.TypeOf(int32(0)), nil
	case "int16":
		return reflect.TypeOf(int16(0)), nil
	case "int8":
		return reflect.TypeOf(int8(0)), nil
	case "uint":
		return reflect.TypeOf(uint(0)), nil
	case "uint64":
		return reflect.TypeOf(uint64(0)), nil
	case "uint32":
		return reflect.TypeOf(uint32(0)), nil
	case "uint16":
		return reflect.TypeOf(uint16(0)), nil
	case "uint8":
		return reflect.TypeOf(uint8(0)), nil
	case "float64":
		return reflect.TypeOf(float64(0)), nil
	case "float32":
		return reflect.TypeOf(float32(0)), nil
	case "string":
		return reflect.TypeOf("0"), nil
	case "json.Number":
		return reflect.TypeOf(json.Number("0")), nil
	case "bool":
		return reflect.TypeOf(true), nil
	default:
		return nil, errors.New("Invalid type declared in the yaml configuration file")
	}
}

func BuildFunction(types []string, ret string,
	ankowrapper func(args ...interface{}) (interface{}, error),
) (reflect.Value, error) {
	paramsType := []reflect.Type{}

	for _, t := range types {
		typeString, err := getTypeFromString(t)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("cannot build function: %w", err)
		}

		paramsType = append(paramsType, typeString)
	}

	typeStringReturn, err := getTypeFromString(ret)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("cannot build function: %w", err)
	}

	typefn := reflect.FuncOf(paramsType, []reflect.Type{typeStringReturn}, false)

	wrapfn := func(args []reflect.Value) []reflect.Value {
		argsInterface := []interface{}{}
		for i := range args {
			argsInterface = append(argsInterface, args[i].Interface())
		}

		result, err := ankowrapper(argsInterface...)
		if err != nil {
			panic(err)
		}

		return []reflect.Value{reflect.ValueOf(result)}
	}

	fn := reflect.MakeFunc(typefn, wrapfn)

	return fn, nil
}
