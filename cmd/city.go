package cmd

import (
	"github.com/spf13/cobra"
)

var cityCmd = &cobra.Command{
	Use:   "city [set|get]",
	Short: "Commands for managing default city",
	Long: `
Commands for managing default city, ie. setting the default city (city set) and retrieving the default city (city get).

Setting the default city allows for running current, forecast and pollution commands without needing to specify the city.

If no default city is set and commands are called without specifying a city, weather/pollution information will be shown for Belgrade,RS.`,
}

func init() {
	rootCmd.AddCommand(cityCmd)
}
