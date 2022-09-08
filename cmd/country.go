package cmd

import (
	"github.com/spf13/cobra"
)

var countryCmd = &cobra.Command{
	Use:   "country",
	Short: "Commands for managing default country configuration",
	Long: `
Commands for managing default country configuration, ie. setting the default country (country set) and retrieving the default country (country get).

Setting the default coutry allows for running current, forecast and pollution commands without needing to specify the city.

If no default country is set and commands are called without specifying a country, Serbia (RS) is presumed.`,
}

func init() {
	rootCmd.AddCommand(countryCmd)
}
