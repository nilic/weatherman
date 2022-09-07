package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var citySetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set default city",
	Long: `
Save default city to the configuration file located at $HOME/.weatherman/config, eg.
		
./weatherman city set Paris,FR

Setting the default city allows for running current, forecast and pollution commands without needing to specify the city.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.SetCity(args[0])
	},
}

func init() {
	cityCmd.AddCommand(citySetCmd)
}
