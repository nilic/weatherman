package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var pollutionCmd = &cobra.Command{
	Use:   "pollution [location] [flags]",
	Short: "Retrieve pollution information for a given location",
	Long: `
Retrieve pollution information including Air Quality Index (AQI) and concentration of polluting gasses (CO, NO2, PM2.5 etc).

Location needs to be specified as command argument in format "<city name>,<ISO 3166 country code>", eg.

./weatherman pollution "Paris,FR"

If only city name is specified, Serbia (RS) is presumed.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.GetPollution(args)
	},
}

func init() {
	rootCmd.AddCommand(pollutionCmd)
}
