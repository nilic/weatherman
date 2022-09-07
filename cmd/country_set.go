package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var countrySetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set default country",
	Long: `
Save default country to the configuration file, eg.
		
./weatherman country set FR
	
Setting the default country allows for running current, forecast and pollution commands without needing to specify the country, eg.

./weatherman pollution Paris`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.SetCountry(args[0])
	},
}

func init() {
	countryCmd.AddCommand(countrySetCmd)
}
