package cmd

import (
	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

var countryListISOCodesCmd = &cobra.Command{
	Use:   "list-iso-codes",
	Short: "Prints all countries and their two-letter ISO 3166 country code",
	Long:  `Prints country names in english along with their two-letter ISO 3166 country code`,
	Run: func(cmd *cobra.Command, args []string) {
		weatherman.ListISOCodes()
	},
}

func init() {
	countryCmd.AddCommand(countryListISOCodesCmd)
}
