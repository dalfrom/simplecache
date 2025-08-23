package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dalfrom/tempodb/pkg/tcp"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the TempoDB server",
	Run: func(cmd *cobra.Command, args []string) {
		ServerCache = &tcp.ServerCache{Port: port}
		ServerCache.Start()
	},
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the TempoDB server",
	Run: func(cmd *cobra.Command, args []string) {
		println("Checking TempoDB server status...")
	},
}

func init() {
	// Add flags to the start command
	StartCmd.Flags().IntVar(&port, "p", 4000, "Server port [defaults to 4000]")
}
