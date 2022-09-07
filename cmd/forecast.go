package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var forecastCmd = &cobra.Command{
	Use:   "forecast [location] [flags]",
	Short: "Retrieve 5-day weather forecast for a given location",
	Long: `
Retrieve 5-day weather forecast including temperature, wind speed and chance of rain.

Location needs to be specified as command argument in format "<city name>,<ISO 3166 country code>", eg.

./weatherman forecast "Paris,FR"

If only city name is specified, Serbia (RS) is presumed.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.GetWeatherForecast(args)
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}
