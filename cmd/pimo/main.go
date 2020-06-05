package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/internal/app/pimo"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/command"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/constant"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/duration"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/hash"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/increment"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randdate"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomint"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomlist"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/regex"
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

	iteration int
	skipLine  bool
	skipField bool
)

var rootCmd = &cobra.Command{
	Use:     "pimo",
	Short:   "Command line to mask data from jsonlines",
	Long:    `Pimo is a tool to mask private data contained in jsonlines by using masking configurations`,
	Version: fmt.Sprintf("%v (commit=%v date=%v by=%v)", version, commit, buildDate, builtBy),
	Run: func(cmd *cobra.Command, args []string) {
		if skipField && skipLine {
			os.Stderr.WriteString("Can't use both flags \n")
			os.Exit(5)
		}
		maskingfile := "masking.yml"
		config, err := pimo.YamlConfig(maskingfile, injectMaskFactories())
		if err != nil {
			println("ERROR : masking.yml not working properly, ", err.Error())
			os.Exit(1)
		}
		if iteration != 0 {
			maskingEngine := model.MaskingEngineFactory(config)
			reader := pimo.NewJSONLineIterator(os.Stdin)
			for {
				i := 0
				dic, err := reader.Next()
				for i != iteration {
					if err != nil {
						if (err == pimo.StopIteratorError{}) {
							os.Exit(0)
						} else {
							os.Stderr.WriteString(err.Error() + "\n")
							os.Exit(2)
						}
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
			}
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&iteration, "repeat", "r", 1, "number of iteration to mask each input")
	rootCmd.PersistentFlags().BoolVar(&skipLine, "skip-line-on-error", false, "if an error occurs, skip the line")
	rootCmd.PersistentFlags().BoolVar(&skipField, "skip-field-on-error", false, "if an error occurs, skip the field")
}

func injectMaskFactories() []func(model.Masking, int64) (model.MaskEngine, bool, error) {
	return []func(model.Masking, int64) (model.MaskEngine, bool, error){
		constant.NewMaskFromConfig,
		command.NewMaskFromConfig,
		randomlist.NewMaskFromConfig,
		randomint.NewMaskFromConfig,
		weightedchoice.NewMaskFromConfig,
		regex.NewMaskFromConfig,
		hash.NewMaskFromConfig,
		randdate.NewMaskFromConfig,
		increment.NewMaskFromConfig,
		replacement.NewMaskFromConfig,
		duration.NewMaskFromConfig,
		templatemask.NewMaskFromConfig,
	}
}
