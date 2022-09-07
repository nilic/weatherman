package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current [location] [flags]",
	Short: "Retrieve current weather conditions for a given location",
	Long: `
Retrieve current weather conditions including temperature, wind speed, pressure, humidity, sunrise and sunset time.

Location needs to be specified as command argument in format "<city name>,<ISO 3166 country code>", eg.

./weatherman current "Paris,FR"

If only city name is specified, Serbia (RS) is presumed.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.GetWeatherCurrent(args)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
