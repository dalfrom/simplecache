package start

import (
	"github.com/spf13/cobra"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the TempoDB server",
	Run: func(cmd *cobra.Command, args []string) {
		println("Starting TempoDB server...")
	},
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the TempoDB server",
	Run: func(cmd *cobra.Command, args []string) {
		println("Checking TempoDB server status...")
	},
}
