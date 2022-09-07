package cmd

import (
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Commands for managing OpenWeather API key",
	Long:  `Commands for managing OpenWeather API key, ie. saving it to the config file (key save), displaying it (key show) and removing it from the config file (key delete).`,
}

func init() {
	rootCmd.AddCommand(keyCmd)
}
