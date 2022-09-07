package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var locationSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set default location",
	Long: `
Save default location to the configuration file, eg.
		
./weatherman location set Paris,FR

Setting the default location allows for running current, forecast and pollution commands without needing to specify the location.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.SetLocation(args[0])
	},
}

func init() {
	locationCmd.AddCommand(locationSetCmd)
}
