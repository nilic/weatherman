package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var locationShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show default location",
	Long:  `Display previously defined default location.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.ShowLocation()
	},
}

func init() {
	locationCmd.AddCommand(locationShowCmd)
}
