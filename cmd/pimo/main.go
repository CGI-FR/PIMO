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
			println("ERROR : Masking.yml not working properly")
			os.Exit(-1)
		}
		maskingEngine := pimo.MaskingEngineFactory(config)
		reader := pimo.NewJSONLineIterator(os.Stdin)
		for {
			dic, err := reader.Next()
			if err != nil {
				println("Enable to  transform this line")
				break
			}
			masked := maskingEngine.Mask(dic).(map[string]pimo.Entry)
			jsonline, _ := pimo.DictionaryToJSON(masked)
			os.Stdout.Write(jsonline)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
