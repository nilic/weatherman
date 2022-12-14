package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var keyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a previously saved OpenWeather API key",
	Long:  `Remove a previously saved OpenWeather API key from the configuration file.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.DeleteAPIKey()
	},
}

func init() {
	keyCmd.AddCommand(keyDeleteCmd)
}
