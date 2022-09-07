package cmd

import (
	"os"

	"github.com/nilic/weatherman/weatherman"
	"github.com/spf13/cobra"
)

const logo = `
██╗    ██╗███████╗ █████╗ ████████╗██╗  ██╗███████╗██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
██║    ██║██╔════╝██╔══██╗╚══██╔══╝██║  ██║██╔════╝██╔══██╗████╗ ████║██╔══██╗████╗  ██║
██║ █╗ ██║█████╗  ███████║   ██║   ███████║█████╗  ██████╔╝██╔████╔██║███████║██╔██╗ ██║
██║███╗██║██╔══╝  ██╔══██║   ██║   ██╔══██║██╔══╝  ██╔══██╗██║╚██╔╝██║██╔══██║██║╚██╗██║
╚███╔███╔╝███████╗██║  ██║   ██║   ██║  ██║███████╗██║  ██║██║ ╚═╝ ██║██║  ██║██║ ╚████║
 ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`

var (
	rootCmd = &cobra.Command{
		Use:   "weatherman",
		Short: "Retrieve current weather, forecast and pollution information for any place in the world based on OpenWeather data",
		Long: logo + `

Retrieve current weather (weatherman current), forecast (weatherman forecast) and pollution (weatherman pollution) information for any place in the world.

Based on OpenWeather data.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			weatherman.InitConfig()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
