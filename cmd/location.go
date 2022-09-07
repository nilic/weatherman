package cmd

import (
	"github.com/spf13/cobra"
)

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Commands for managing default location configuration",
	Long: `
Commands for managing default location configuration, ie. setting the default location (location set) and retrieving the default location (location get).

Setting the default location allows for running current, forecast and pollution commands without needing to specify the location.

If no default location is set and commands are called without specifying a location, weather/pollution information will be shown for Belgrade,RS.`,
}

func init() {
	rootCmd.AddCommand(locationCmd)
}
