package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [apikey] [flags]",
	Short: "Save OpenWeather API key",
	Long: `
Save OpenWeather API key to the configuration file located at $HOME/.weatherman/config, eg.
	
./weatherman key save 123456qwerty789`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.SaveAPIKey(args)
	},
}

func init() {
	keyCmd.AddCommand(saveCmd)
}
