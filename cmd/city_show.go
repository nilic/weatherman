package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var cityShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show default city",
	Long:  `Display previously defined default city.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.ShowCity()
	},
}

func init() {
	cityCmd.AddCommand(cityShowCmd)
}
