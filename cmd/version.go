package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the current version of weatherman",
	Long:  `Returns the current version of weatherman`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("weatherman v%s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
