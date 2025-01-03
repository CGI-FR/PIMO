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

import (
	"os"

	"github.com/spf13/cobra"
)

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
	maxBufferCapacity     = 64
	catchErrors           = ""
	maskingFile           = "masking.yml"
	mockConfigFile        = "routes.yaml"
	cachesToDump          = map[string]string{}
	cachesToLoad          = map[string]string{}
	emptyInput            = false
	maskingOneLiner       = []string{}
	profiling             = ""
	iteration             = 1
	repeatUntil           = ""
	repeatWhile           = ""
	seedValue             = int64(0)
	serve                 = ""
	skipLineOnError       = false
	skipFieldOnError      = false
	skipLogFile           = ""
	statisticsDestination = os.Getenv("PIMO_STATS_URL")
	statsTemplate         = os.Getenv("PIMO_STATS_TEMPLATE")
	xmlSubscriberName     = map[string]string{}
)

//nolint:gochecknoglobals
var (
	flagBufferSize        = flag[int]{name: "buffer-size", variable: &maxBufferCapacity, usage: "buffer size in kB to load data from uri for each line"}
	flagCatchErrors       = flag[string]{name: "catch-errors", shorthand: "e", variable: &catchErrors, usage: "catch errors and write line in file, same as using skip-field-on-error + skip-log-file"}
	flagConfigMasking     = flag[string]{name: "config", shorthand: "c", variable: &maskingFile, usage: "name and location of the masking config file"}
	flagConfigRoute       = flag[string]{name: "config", shorthand: "c", variable: &mockConfigFile, usage: "name and location of the routes config file"}
	flagCachesToDump      = flag[map[string]string]{name: "dump-cache", variable: &cachesToDump, usage: "path for dumping cache into file"}
	flagCachesToLoad      = flag[map[string]string]{name: "load-cache", variable: &cachesToLoad, usage: "path for loading cache from file"}
	flagEmptyInput        = flag[bool]{name: "empty-input", variable: &emptyInput, usage: "generate data without any input, to use with repeat flag"}
	flagMaskOneLiner      = flag[[]string]{name: "mask", shorthand: "m", variable: &maskingOneLiner, usage: "one liner masking"}
	flagProfiling         = flag[string]{name: "pprof", variable: &profiling, usage: "create a pprof file - use 'cpu' to create a CPU pprof file or 'mem' to create an memory pprof file"}
	flagRepeat            = flag[int]{name: "repeat", shorthand: "r", variable: &iteration, usage: "number of iteration to mask each input"}
	flagRepeatUntil       = flag[string]{name: "repeat-until", variable: &repeatUntil, usage: "mask each input repeatedly until the given condition is met"}
	flagRepeatWhile       = flag[string]{name: "repeat-while", variable: &repeatWhile, usage: "mask each input repeatedly while the given condition is met"}
	flagSeed              = flag[int64]{name: "seed", shorthand: "s", variable: &seedValue, usage: "set global seed"}
	flagServe             = flag[string]{name: "serve", variable: &serve, usage: "listen/respond to HTTP interface and port instead of stdin/stdout, <ip>:<port> or :<port> to listen to all local networks"}
	flagSkipLineOnError   = flag[bool]{name: "skip-line-on-error", variable: &skipLineOnError, usage: "skip a line if an error occurs while masking a field"}
	flagSkipFieldOnError  = flag[bool]{name: "skip-field-on-error", variable: &skipFieldOnError, usage: "remove a field if an error occurs while masking this field"}
	flagSkipLogFile       = flag[string]{name: "skip-log-file", variable: &skipLogFile, usage: "skipped lines will be written to this log file"}
	flagStatsDestination  = flag[string]{name: "stats", variable: &statisticsDestination, usage: "generate execution statistics in the specified dump file"}
	flagStatsTemplate     = flag[string]{name: "statsTemplate", variable: &statsTemplate, usage: "template string to format stats (to include them you have to specify them as `{{ .Stats }}` like `{\"software\":\"PIMO\",\"stats\":{{ .Stats }}}`)"}
	flagXMLSubscriberName = flag[map[string]string]{name: "subscriber", variable: &xmlSubscriberName, usage: "name of element to mask"}
)

func addFlag[T any](cmd *cobra.Command, flag flag[T]) {
	switch variable := any(flag.variable).(type) {
	case *int:
		if len(flag.shorthand) > 0 {
			cmd.Flags().IntVarP(variable, flag.name, flag.shorthand, *variable, flag.usage)
		} else {
			cmd.Flags().IntVar(variable, flag.name, *variable, flag.usage)
		}
	case *bool:
		if len(flag.shorthand) > 0 {
			cmd.Flags().BoolVarP(variable, flag.name, flag.shorthand, *variable, flag.usage)
		} else {
			cmd.Flags().BoolVar(variable, flag.name, *variable, flag.usage)
		}
	case *string:
		if len(flag.shorthand) > 0 {
			cmd.Flags().StringVarP(variable, flag.name, flag.shorthand, *variable, flag.usage)
		} else {
			cmd.Flags().StringVar(variable, flag.name, *variable, flag.usage)
		}
	case *int64:
		if len(flag.shorthand) > 0 {
			cmd.Flags().Int64VarP(variable, flag.name, flag.shorthand, *variable, flag.usage)
		} else {
			cmd.Flags().Int64Var(variable, flag.name, *variable, flag.usage)
		}
	case *map[string]string:
		if len(flag.shorthand) > 0 {
			cmd.Flags().StringToStringVarP(variable, flag.name, flag.shorthand, *variable, flag.usage)
		} else {
			cmd.Flags().StringToStringVar(variable, flag.name, *variable, flag.usage)
		}
	case *[]string:
		if len(flag.shorthand) > 0 {
			cmd.Flags().StringArrayVarP(variable, flag.name, flag.shorthand, *variable, flag.usage)
		} else {
			cmd.Flags().StringArrayVar(variable, flag.name, *variable, flag.usage)
		}
	}
}
