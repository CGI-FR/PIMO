package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/internal/app/pimo"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/add"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/command"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/constant"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/dateparser"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/duration"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/fluxuri"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/hash"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/increment"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/jsonline"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randdate"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randdura"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomdecimal"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomint"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomlist"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/rangemask"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/regex"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/remove"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/replacement"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/templatemask"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/weightedchoice"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version   string
	commit    string
	buildDate string
	builtBy   string

	iteration   int
	emptyInput  bool
	maskingFile string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "pimo",
		Short:   "Command line to mask data from jsonlines",
		Long:    `Pimo is a tool to mask private data contained in jsonlines by using masking configurations`,
		Version: fmt.Sprintf("%v (commit=%v date=%v by=%v)\n© CGI Inc 2020 All rights reserved", version, commit, buildDate, builtBy),
		Run: func(cmd *cobra.Command, args []string) {
			/*
				config, err := pimo.YamlConfig(maskingFile, injectMaskFactories(), injectMaskContextFactories())
				if err != nil {
					os.Stderr.WriteString("ERROR : masking.yml not working properly, " + err.Error())
					os.Exit(1)
				}

				run(config)
			*/
			run2()
		},
	}

	rootCmd.PersistentFlags().IntVarP(&iteration, "repeat", "r", 1, "number of iteration to mask each input")
	rootCmd.PersistentFlags().BoolVar(&emptyInput, "empty-input", false, "generate datas without any input, to use with repeat flag")
	rootCmd.PersistentFlags().StringVarP(&maskingFile, "config", "c", "masking.yml", "name and location of the masking-config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run2() {
	var source model.ISource
	if emptyInput {
		source = model.NewSourceFromSlice([]model.Dictionary{{}})
	} else {
		source = jsonline.NewSource(os.Stdin)
	}
	pipeline := model.NewPipeline(source).
		Process(model.NewRepeaterProcess(iteration))

	var err error

	pipeline, err = pimo.YamlPipeline(pipeline, maskingFile, injectMaskFactories(), injectMaskContextFactories())

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	err = pipeline.AddSink(jsonline.NewSink(os.Stdout)).Run()

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(4)
	}
	os.Exit(0)
}

/*
func run(config model.MaskConfiguration) {
	if iteration == 0 {
		return
	}
	maskingEngine := model.MaskingEngineFactory(config, true)
	reader := pimo.NewJSONLineIterator(os.Stdin)
	for {
		i := 0
		var dic model.Dictionary
		var err error
		if emptyInput {
			// TODO replace that if statement by an EmptyDictionaryIterator
			dic = model.Dictionary{}
		} else {
			dic, err = reader.Next()
		}
		for i < iteration {
			if (err == pimo.StopIteratorError{}) {
				os.Exit(0)
			}
			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(2)
			}
			masked, err := maskingEngine.Mask(dic)
			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				if skipLine {
					break
				}
				if !skipField {
					os.Exit(4)
				}
			}
			maskedmap := masked.(map[string]model.Entry)
			jsonline, err := pimo.DictionaryToJSON(maskedmap)
			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(3)
			}
			os.Stdout.Write(jsonline)
			i++
		}
		if emptyInput {
			os.Exit(0)
		}
	}
}
*/

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
	}
}
