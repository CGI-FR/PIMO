// Copyright (C) 2024 CGI France
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

package main

import "github.com/spf13/cobra"

type flag[T any] struct {
	name      string // Name of the flag
	shorthand string // Optional short name
	variable  *T     // Pointer to the variable
	usage     string // Description of the flag
}

// flags with default values
//
//nolint:gochecknoglobals
var (
	maxBufferCapacity = 64
	catchErrors       = ""
	maskingFile       = "masking.yml"
	mockConfigFile    = "routes.yaml"
	cachesToDump      = map[string]string{}
	cachesToLoad      = map[string]string{}
	emptyInput        = false
)

var (
	flagBufferSize    = flag[int]{name: "buffer-size", variable: &maxBufferCapacity, usage: "buffer size in kB to load data from uri for each line"}
	flagCatchErrors   = flag[string]{name: "catch-errors", shorthand: "e", variable: &catchErrors, usage: "catch errors and write line in file, same as using skip-field-on-error + skip-log-file"}
	flagConfigMasking = flag[string]{name: "config", shorthand: "c", variable: &maskingFile, usage: "name and location of the masking config file"}
	flagConfigRoute   = flag[string]{name: "config", shorthand: "c", variable: &mockConfigFile, usage: "name and location of the routes config file"}
	flagCachesToDump  = flag[map[string]string]{name: "dump-cache", variable: &cachesToDump, usage: "path for dumping cache into file"}
	flagCachesToLoad  = flag[map[string]string]{name: "load-cache", variable: &cachesToLoad, usage: "path for loading cache from file"}
	flagEmptyInput    = flag[bool]{name: "empty-input", variable: &emptyInput, usage: "generate data without any input, to use with repeat flag"}
)

func addFlag[T any](cmd *cobra.Command, flag flag[T]) {
	switch typedVar := any(*flag.variable).(type) {
	case int:
		if len(flag.shorthand) > 0 {
			cmd.Flags().IntVarP(&typedVar, flag.name, flag.shorthand, typedVar, flag.usage)
		} else {
			cmd.Flags().IntVar(&typedVar, flag.name, typedVar, flag.usage)
		}
	case bool:
		if len(flag.shorthand) > 0 {
			cmd.Flags().BoolVarP(&typedVar, flag.name, flag.shorthand, typedVar, flag.usage)
		} else {
			cmd.Flags().BoolVar(&typedVar, flag.name, typedVar, flag.usage)
		}
	case string:
		if len(flag.shorthand) > 0 {
			cmd.Flags().StringVarP(&typedVar, flag.name, flag.shorthand, typedVar, flag.usage)
		} else {
			cmd.Flags().StringVar(&typedVar, flag.name, typedVar, flag.usage)
		}
	case int64:
		if len(flag.shorthand) > 0 {
			cmd.Flags().Int64VarP(&typedVar, flag.name, flag.shorthand, typedVar, flag.usage)
		} else {
			cmd.Flags().Int64Var(&typedVar, flag.name, typedVar, flag.usage)
		}
	case map[string]string:
		if len(flag.shorthand) > 0 {
			cmd.Flags().StringToStringVarP(&typedVar, flag.name, flag.shorthand, typedVar, flag.usage)
		} else {
			cmd.Flags().StringToStringVar(&typedVar, flag.name, typedVar, flag.usage)
		}
	case []string:
		if len(flag.shorthand) > 0 {
			cmd.Flags().StringArrayVarP(&typedVar, flag.name, flag.shorthand, typedVar, flag.usage)
		} else {
			cmd.Flags().StringArrayVar(&typedVar, flag.name, typedVar, flag.usage)
		}
	}
}
