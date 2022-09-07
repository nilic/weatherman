package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var countryShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show default country",
	Long:  `Display previously defined default country.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.ShowCountry()
	},
}

func init() {
	countryCmd.AddCommand(countryShowCmd)
}
