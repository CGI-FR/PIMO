package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/internal/app/pimo"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version   string
	commit    string
	buildDate string
	builtBy   string

	iteration int
)

var rootCmd = &cobra.Command{
	Use:     "pimo",
	Short:   "Command line to mask data from jsonlines",
	Long:    `Pimo is a tool to mask private data contained in jsonlines by using masking configurations`,
	Version: fmt.Sprintf("%v (commit=%v date=%v by=%v)", version, commit, buildDate, builtBy),
	Run: func(cmd *cobra.Command, args []string) {
		maskingfile := "masking.yml"
		config, err := pimo.YamlConfig(maskingfile)
		if err != nil {
			println("ERROR : masking.yml not working properly")
			os.Exit(1)
		}
		if iteration != 0 {
			maskingEngine := pimo.MaskingEngineFactory(config)
			reader := pimo.NewJSONLineIterator(os.Stdin)
			for {
				i := 0
				dic, err := reader.Next()
				for i != iteration {
					if err != nil {
						if (err == pimo.StopIteratorError{}) {
							os.Exit(0)
						} else {
							os.Stderr.WriteString(err.Error())
							os.Exit(1)
						}
					}
					masked := maskingEngine.Mask(dic).(map[string]pimo.Entry)
					jsonline, _ := pimo.DictionaryToJSON(masked)
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
}
