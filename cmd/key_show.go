package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var keyShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show previously saved OpenWeather API key",
	Long:  `Read and display a previously saved OpenWeather API key from the configuration file located at $HOME/.weatherman/config.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.ShowAPIKey()
	},
}

func init() {
	keyCmd.AddCommand(keyShowCmd)
}
