package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var keySaveCmd = &cobra.Command{
	Use:   "save [apikey] [flags]",
	Short: "Save OpenWeather API key",
	Long: `
Save OpenWeather API key to the configuration file, eg.
	
./weatherman key save 123456qwerty789`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.SaveAPIKey(args[0])
	},
}

func init() {
	keyCmd.AddCommand(keySaveCmd)
}
