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

package main

import (
	"fmt"
	"os"

	"github.com/cgi-fr/pimo/internal/app/pimo"
	"github.com/cgi-fr/pimo/pkg/add"
	"github.com/cgi-fr/pimo/pkg/command"
	"github.com/cgi-fr/pimo/pkg/constant"
	"github.com/cgi-fr/pimo/pkg/dateparser"
	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/ff1"
	"github.com/cgi-fr/pimo/pkg/fluxuri"
	"github.com/cgi-fr/pimo/pkg/hash"
	"github.com/cgi-fr/pimo/pkg/increment"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/pipe"
	"github.com/cgi-fr/pimo/pkg/randdate"
	"github.com/cgi-fr/pimo/pkg/randdura"
	"github.com/cgi-fr/pimo/pkg/randomdecimal"
	"github.com/cgi-fr/pimo/pkg/randomint"
	"github.com/cgi-fr/pimo/pkg/randomlist"
	"github.com/cgi-fr/pimo/pkg/rangemask"
	"github.com/cgi-fr/pimo/pkg/regex"
	"github.com/cgi-fr/pimo/pkg/remove"
	"github.com/cgi-fr/pimo/pkg/replacement"
	"github.com/cgi-fr/pimo/pkg/templatemask"
	"github.com/cgi-fr/pimo/pkg/weightedchoice"
	"github.com/spf13/cobra"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version   string
	commit    string
	buildDate string
	builtBy   string

	iteration    int
	emptyInput   bool
	maskingFile  string
	cachesToDump map[string]string
	cachesToLoad map[string]string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "pimo",
		Short: "Command line to mask data from jsonlines",
		Long:  `Pimo is a tool to mask private data contained in jsonlines by using masking configurations`,
		Version: fmt.Sprintf(`%v (commit=%v date=%v by=%v)
Copyright (C) 2021 CGI France
License GPLv3: GNU GPL version 3 <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.`, version, commit, buildDate, builtBy),
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	rootCmd.PersistentFlags().IntVarP(&iteration, "repeat", "r", 1, "number of iteration to mask each input")
	rootCmd.PersistentFlags().BoolVar(&emptyInput, "empty-input", false, "generate data without any input, to use with repeat flag")
	rootCmd.PersistentFlags().StringVarP(&maskingFile, "config", "c", "masking.yml", "name and location of the masking-config file")
	rootCmd.PersistentFlags().StringToStringVar(&cachesToDump, "dump-cache", map[string]string{}, "path for dumping cache into file")
	rootCmd.PersistentFlags().StringToStringVar(&cachesToLoad, "load-cache", map[string]string{}, "path for loading cache from file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() {
	var source model.Source
	if emptyInput {
		source = model.NewSourceFromSlice([]model.Dictionary{{}})
	} else {
		source = jsonline.NewSource(os.Stdin)
	}
	pipeline := model.NewPipeline(source).
		Process(model.NewRepeaterProcess(iteration))

	var (
		err    error
		caches map[string]model.Cache
	)

	model.InjectMaskContextFactories(injectMaskContextFactories())
	model.InjectMaskFactories(injectMaskFactories())
	pdef, err := model.LoadPipelineDefinitionFromYAML(maskingFile)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	pipeline, caches, err = model.BuildPipeline(pipeline, pdef)

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	for name, path := range cachesToLoad {
		cache, ok := caches[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "Cache %s not found", name)
			os.Exit(2)
		}
		err = pimo.LoadCache(name, cache, path)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(3)
		}
	}

	err = pipeline.AddSink(jsonline.NewSink(os.Stdout)).Run()

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(4)
	}

	for name, path := range cachesToDump {
		cache, ok := caches[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "Cache %s not found", name)
			os.Exit(2)
		}
		err = pimo.DumpCache(name, cache, path)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(3)
		}
	}

	os.Exit(0)
}

func injectMaskContextFactories() []model.MaskContextFactory {
	return []model.MaskContextFactory{
		fluxuri.Factory,
		add.Factory,
		remove.Factory,
	}
}

func injectMaskFactories() []model.MaskFactory {
	return []model.MaskFactory{

		constant.Factory,
		command.Factory,
		randomlist.Factory,
		randomint.Factory,
		weightedchoice.Factory,
		regex.Factory,
		hash.Factory,
		randdate.Factory,
		increment.Factory,
		replacement.Factory,
		duration.Factory,
		templatemask.Factory,

		rangemask.Factory,
		randdura.Factory,
		randomdecimal.Factory,
		dateparser.Factory,
		ff1.Factory,
		pipe.Factory,
	}
}
